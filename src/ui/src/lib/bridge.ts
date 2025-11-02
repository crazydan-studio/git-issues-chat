// bridge.ts - Bridge between UI and Go backend
import type { 
  App, 
  AppUser, 
  GitPlatform, 
  GitRepo, 
  GitIssue, 
  GitIssueComment,
  AppUserActionLog,
  GitUser
} from './types';

// Response interface for typed API responses
export interface Response<D> {
  success: boolean;
  msg?: string;
  data?: D;
}

// Declare the webui global object
declare global {
  interface Window {
    webui: any;
  }
}

interface WebUI {
  isConnected(): boolean;
  call(functionName: string, params?: string): Promise<string>;
  setEventCallback(callback: (e: any) => void): void;
}

// Function to get webui instance with proper checking
async function getWebUI(): Promise<WebUI> {
    // Return a Promise that resolves with the connected webui object
    return new Promise((resolve, reject) => {
        function checkWebUI(webui: WebUI) {
            if (webui.isConnected()) {
                // Connection to the backend is established
                console.log('Connected.');
                // Resolve with the webui object when connected
                resolve(webui);
            } else {
                // Connection to the backend is lost
                console.log('Disconnected.');
                // Reject with an error when disconnected
                reject(new Error('WebUI connection to backend was lost.'));
            }
        }

        if (typeof window.webui !== 'undefined') {
            checkWebUI(window.webui);
        } else {
            document.addEventListener('DOMContentLoaded', function() {
                // DOM is loaded. Check if `webui` object is available
                if (typeof window.webui !== 'undefined') {
                    // Set events callback
                    window.webui.setEventCallback((e: any) => {
                        checkWebUI(window.webui);
                    });
                } else {
                    // The virtual file `webui.js` is not included
                    reject(new Error('Please add webui.js to your HTML.'));
                }
            });
        }
    });
}

// Common function to handle Go function calls with parameter serialization and return value deserialization
async function callGoFunction(functionName: string, params: any = null): Promise<any> {
    try {
        const webui = await getWebUI();
        
        let result: string;
        if (params !== null) {
            result = await webui.call(functionName, JSON.stringify(params));
        } else {
            result = await webui.call(functionName);
        }
        // Parse the JSON response from the Go backend
        return JSON.parse(result);
    } catch (error: any) {
        console.error(`Error calling ${functionName}:`, error);
        // Return error information in the msg property instead of throwing
        return { success: false, msg: `${functionName}: ${error.message || 'Unknown error occurred'}` };
    }
}

export async function getAppInfo(): Promise<Response<App>> {
    const result = await callGoFunction('getAppInfo');
    return result;
}

export async function verifyAppUser(params: {userId: string, password: string}): Promise<Response<AppUser>> {
    const result = await callGoFunction('verifyAppUser', params);
    return result;
}

export async function saveAppUserInfo(params: any): Promise<Response<null>> {
    const result = await callGoFunction('saveAppUserInfo', params);
    return result;
}

export async function updateAppUserPassword(params: any): Promise<Response<null>> {
    const result = await callGoFunction('updateAppUserPassword', params);
    return result;
}

export async function getGitPlatformList(): Promise<Response<GitPlatform[]>> {
    const result = await callGoFunction('getGitPlatformList');
    return result;
}

export async function saveGitPlatform(params: any): Promise<Response<null>> {
    const result = await callGoFunction('saveGitPlatform', params);
    return result;
}

export async function getGitRepoList(params: {platformId: string}): Promise<Response<GitRepo[]>> {
    const result = await callGoFunction('getGitRepoList', params);
    return result;
}

export async function getGitRepoInfo(params: {platformId: string, repoName: string}): Promise<Response<GitRepo>> {
    const result = await callGoFunction('getGitRepoInfo', params);
    return result;
}

export async function saveGitRepo(params: any): Promise<Response<null>> {
    const result = await callGoFunction('saveGitRepo', params);
    return result;
}

export async function getGitIssueList(params: {repoId: string}): Promise<Response<GitIssue[]>> {
    const result = await callGoFunction('getGitIssueList', params);
    return result;
}

export async function saveGitIssue(params: any): Promise<Response<null>> {
    const result = await callGoFunction('saveGitIssue', params);
    return result;
}

export async function getGitIssueCommentList(params: {issueId: string}): Promise<Response<GitIssueComment[]>> {
    const result = await callGoFunction('getGitIssueCommentList', params);
    return result;
}

export async function getGitIssueParticipantList(params: {issueId: string}): Promise<Response<GitUser[]>> {
    const result = await callGoFunction('getGitIssueParticipantList', params);
    return result;
}

export async function saveGitIssueComment(params: any): Promise<Response<null>> {
    const result = await callGoFunction('saveGitIssueComment', params);
    return result;
}

export async function getAppUserActionLogList(params: {userId: string}): Promise<Response<AppUserActionLog[]>> {
    const result = await callGoFunction('getAppUserActionLogList', params);
    return result;
}