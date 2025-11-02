// store.ts - Application state management
import { writable } from 'svelte/store';
import type { 
  AppUser, 
  AppUserActionLog,
  GitPlatform, 
  GitRepo, 
  GitIssue, 
  GitIssueComment,
  GitUser
} from './types';

// App user state
export const appUser = writable<AppUser | null>(null);

// Git platform states
export const gitPlatformList = writable<GitPlatform[]>([]);
export const selectedGitPlatform = writable<GitPlatform | null>(null);

// Git repository states
export const gitRepoList = writable<GitRepo[]>([]);
export const selectedGitRepo = writable<GitRepo | null>(null);

// Git issue states
export const gitIssueList = writable<GitIssue[]>([]);
export const selectedGitIssue = writable<GitIssue | null>(null);

// Git comment states
export const gitIssueCommentList = writable<GitIssueComment[]>([]);
export const gitIssueParticipantList = writable<GitUser[]>([]);

// Action log states
export const appUserActionLogList = writable<AppUserActionLog[]>([]);

// UI states
export const showAboutDialog = writable<boolean>(false);
export const showUserProfileDialog = writable<boolean>(false);
export const showUserPasswordDialog = writable<boolean>(false);
export const showAddPlatformDialog = writable<boolean>(false);
export const showAddRepoDialog = writable<boolean>(false);
export const showAddIssueDialog = writable<boolean>(false);
export const notification = writable<{ type: string; message: string }>({ type: '', message: '' });

// Update functions
export const updateAppUser = (user: AppUser | null) => {
    appUser.set(user);
};

export const updateGitPlatformList = (platforms: GitPlatform[]) => {
    gitPlatformList.set(platforms);
};

export const updateSelectedGitPlatform = (platform: GitPlatform | null) => {
    selectedGitPlatform.set(platform);
};

export const updateGitRepoList = (repos: GitRepo[]) => {
    gitRepoList.set(repos);
};

export const updateSelectedGitRepo = (repo: GitRepo | null) => {
    selectedGitRepo.set(repo);
};

export const updateGitIssueList = (issues: GitIssue[]) => {
    gitIssueList.set(issues);
};

export const updateSelectedGitIssue = (issue: GitIssue | null) => {
    selectedGitIssue.set(issue);
};

export const updateGitIssueCommentList = (comments: GitIssueComment[]) => {
    gitIssueCommentList.set(comments);
};

export const updateGitIssueParticipantList = (participants: GitUser[]) => {
    gitIssueParticipantList.set(participants);
};

export const updateAppUserActionLogList = (logs: AppUserActionLog[]) => {
    appUserActionLogList.set(logs);
};

export const showNotification = (type: string, message: string) => {
    notification.set({ type, message });
    // Auto-hide notification after 3 seconds
    setTimeout(() => {
        notification.set({ type: '', message: '' });
    }, 3000);
};

export const clearAllStates = () => {
    appUser.set(null);
    gitPlatformList.set([]);
    selectedGitPlatform.set(null);
    gitRepoList.set([]);
    selectedGitRepo.set(null);
    gitIssueList.set([]);
    selectedGitIssue.set(null);
    gitIssueCommentList.set([]);
    gitIssueParticipantList.set([]);
    appUserActionLogList.set([]);
};