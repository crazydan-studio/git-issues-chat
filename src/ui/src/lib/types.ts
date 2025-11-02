// Frontend type definitions corresponding to backend Go structs

// License represents a license agreement
export interface License {
  name: string;
  url: string;
}

// App represents an application
export interface App {
  name: string;
  version: string;
  author: string;
  source: string;
  description: string;
  license: License;
}

// AppUser represents an application user
export interface AppUser {
  id: string;
  username: string;
  displayName: string;
  avatar: string;
}

// AppUserActionLog represents an application user action log
export interface AppUserActionLog {
  id: string;
  status: string;
  content: string;
  createdAt: number;
}

// GitPlatform represents a Git platform
export interface GitPlatform {
  id: string;
  name: string;
  url: string;
  icon: string;
  description?: string;
}

// GitRepo represents a Git repository
export interface GitRepo {
  id: string;
  name: string;
  platform: string;
  url: string;
  icon?: string;
  description?: string;
}

// GitUser represents a Git user
export interface GitUser {
  id: string;
  username: string;
  displayName: string;
  url: string;
}

// GitIssue represents a Git issue
export interface GitIssue {
  id: string;
  title: string;
  code: string;
  url: string;
  description: string;
  createdBy: GitUser;
  createdAt: number;
}

// GitIssueComment represents a Git issue comment
export interface GitIssueComment {
  id: string;
  content: string;
  createdBy: GitUser;
  createdAt: number;
}