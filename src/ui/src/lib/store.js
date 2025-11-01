// store.js - Application state management
import { writable } from 'svelte/store';

// App user state
export const appUser = writable(null);

// Git platform states
export const gitPlatformList = writable([]);
export const selectedGitPlatform = writable(null);

// Git repository states
export const gitRepoList = writable([]);
export const selectedGitRepo = writable(null);

// Git issue states
export const gitIssueList = writable([]);
export const selectedGitIssue = writable(null);

// Git comment states
export const gitIssueCommentList = writable([]);
export const gitIssueParticipantList = writable([]);

// Action log states
export const appUserActionLogList = writable([]);

// UI states
export const showAboutDialog = writable(false);
export const showUserProfileDialog = writable(false);
export const showUserPasswordDialog = writable(false);
export const showAddPlatformDialog = writable(false);
export const showAddRepoDialog = writable(false);
export const showAddIssueDialog = writable(false);
export const notification = writable({ type: '', message: '' });

// Update functions
export const updateAppUser = (user) => {
    appUser.set(user);
};

export const updateGitPlatformList = (platforms) => {
    gitPlatformList.set(platforms);
};

export const updateSelectedGitPlatform = (platform) => {
    selectedGitPlatform.set(platform);
};

export const updateGitRepoList = (repos) => {
    gitRepoList.set(repos);
};

export const updateSelectedGitRepo = (repo) => {
    selectedGitRepo.set(repo);
};

export const updateGitIssueList = (issues) => {
    gitIssueList.set(issues);
};

export const updateSelectedGitIssue = (issue) => {
    selectedGitIssue.set(issue);
};

export const updateGitIssueCommentList = (comments) => {
    gitIssueCommentList.set(comments);
};

export const updateGitIssueParticipantList = (participants) => {
    gitIssueParticipantList.set(participants);
};

export const updateAppUserActionLogList = (logs) => {
    appUserActionLogList.set(logs);
};

export const showNotification = (type, message) => {
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
};