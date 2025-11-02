package types

// GitPlatform represents a Git platform
type GitPlatform struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Icon        string `json:"icon"`
	Description string `json:"description,omitempty"`
}

// GitRepo represents a Git repository
type GitRepo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Platform    string `json:"platform"`
	URL         string `json:"url"`
	Icon        string `json:"icon,omitempty"`
	Description string `json:"description,omitempty"`
}

// GitUser represents a Git user
type GitUser struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	URL         string `json:"url"`
}

// GitIssue represents a Git issue
type GitIssue struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Code        string   `json:"code"`
	URL         string   `json:"url"`
	Description string   `json:"description"`
	CreatedBy   GitUser  `json:"createdBy"`
	CreatedAt   int64    `json:"createdAt"`
}

// GitIssueComment represents a Git issue comment
type GitIssueComment struct {
	ID        string  `json:"id"`
	Content   string  `json:"content"`
	CreatedBy GitUser `json:"createdBy"`
	CreatedAt int64   `json:"createdAt"`
}