package handlers

// ParamsOfVerifyAppUser represents the parameters for VerifyAppUser function
type ParamsOfVerifyAppUser struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
}

// ParamsOfSaveAppUserInfo represents the parameters for SaveAppUserInfo function
type ParamsOfSaveAppUserInfo struct {
	// Add fields as needed based on what the frontend sends
}

// ParamsOfUpdateAppUserPassword represents the parameters for UpdateAppUserPassword function
type ParamsOfUpdateAppUserPassword struct {
	// Add fields as needed based on what the frontend sends
}

// ParamsOfGetGitRepoList represents the parameters for GetGitRepoList function
type ParamsOfGetGitRepoList struct {
	PlatformId string `json:"platformId"`
}

// ParamsOfGetGitRepoInfo represents the parameters for GetGitRepoInfo function
type ParamsOfGetGitRepoInfo struct {
	PlatformId string `json:"platformId"`
	RepoName   string `json:"repoName"`
}

// ParamsOfSaveGitRepo represents the parameters for SaveGitRepo function
type ParamsOfSaveGitRepo struct {
	// Add fields as needed based on what the frontend sends
}

// ParamsOfGetGitIssueList represents the parameters for GetGitIssueList function
type ParamsOfGetGitIssueList struct {
	RepoId string `json:"repoId"`
}

// ParamsOfSaveGitIssue represents the parameters for SaveGitIssue function
type ParamsOfSaveGitIssue struct {
	// Add fields as needed based on what the frontend sends
}

// ParamsOfGetGitIssueCommentList represents the parameters for GetGitIssueCommentList function
type ParamsOfGetGitIssueCommentList struct {
	IssueId string `json:"issueId"`
}

// ParamsOfGetGitIssueParticipantList represents the parameters for GetGitIssueParticipantList function
type ParamsOfGetGitIssueParticipantList struct {
	IssueId string `json:"issueId"`
}

// ParamsOfSaveGitIssueComment represents the parameters for SaveGitIssueComment function
type ParamsOfSaveGitIssueComment struct {
	IssueId string `json:"issueId"`
	Content string `json:"content"`
}

// ParamsOfGetAppUserActionLogList represents the parameters for GetAppUserActionLogList function
type ParamsOfGetAppUserActionLogList struct {
	UserId string `json:"userId"`
}

// ParamsOfSaveGitPlatform represents the parameters for SaveGitPlatform function
type ParamsOfSaveGitPlatform struct {
	// Add fields as needed based on what the frontend sends
}