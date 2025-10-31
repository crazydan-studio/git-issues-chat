// bridge.js - Bridge between UI and Go backend

// Common function to handle Go function calls with parameter serialization and return value deserialization
async function callGoFunction(functionName, params = null) {
    try {
        let result;
        if (params !== null) {
            result = await webui.call(functionName, JSON.stringify(params));
        } else {
            result = await webui.call(functionName);
        }
        // Parse the JSON response from the Go backend
        return JSON.parse(result);
    } catch (error) {
        console.error(`Error calling ${functionName}:`, error);
        // Return error information in the msg property instead of throwing
        return { success: false, msg: `${functionName}: ${error.message || 'Unknown error occurred'}` };
    }
}

export async function getAppInfo() {
    const result = await callGoFunction('getAppInfo');
    return result;
}

export async function verifyAppUser(params) {
    const result = await callGoFunction('verifyAppUser', params);
    return result;
}

export async function saveAppUserInfo(params) {
    const result = await callGoFunction('saveAppUserInfo', params);
    return result;
}

export async function updateAppUserPassword(params) {
    const result = await callGoFunction('updateAppUserPassword', params);
    return result;
}

export async function getGitPlatformList() {
    const result = await callGoFunction('getGitPlatformList');
    return result;
}

export async function saveGitPlatform(params) {
    const result = await callGoFunction('saveGitPlatform', params);
    return result;
}

export async function getGitPlatformRepoList(params) {
    const result = await callGoFunction('getGitPlatformRepoList', params);
    return result;
}

export async function getGitPlatformRepoInfo(params) {
    const result = await callGoFunction('getGitPlatformRepoInfo', params);
    return result;
}

export async function saveGitPlatformRepo(params) {
    const result = await callGoFunction('saveGitPlatformRepo', params);
    return result;
}

export async function getGitRepoIssueList(params) {
    const result = await callGoFunction('getGitRepoIssueList', params);
    return result;
}

export async function saveGitRepoIssue(params) {
    const result = await callGoFunction('saveGitRepoIssue', params);
    return result;
}

export async function getGitIssueCommentList(params) {
    const result = await callGoFunction('getGitIssueCommentList', params);
    return result;
}

export async function getGitIssueParticipantList(params) {
    const result = await callGoFunction('getGitIssueParticipantList', params);
    return result;
}

export async function saveGitIssueComment(params) {
    const result = await callGoFunction('saveGitIssueComment', params);
    return result;
}