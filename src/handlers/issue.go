package handlers

import (
	"fmt"
	ui "github.com/webui-dev/go-webui/v2"
	"git-issues-chat/src/types"
)

// GetGitIssueList returns list of issues for a Git repository
func GetGitIssueList(e ui.Event) any {
	// Parse input parameters
	params, err := ui.GetArg[ParamsOfGetGitIssueList](e)
	if err != nil {
		result := types.Response{
			Success: false,
			Msg:     "Invalid parameters",
		}
		return result
	}
	
	// Use the params to avoid "declared and not used" error
	_ = params
	
	// For simulation, return 50 sample issues
	issues := make([]types.GitIssue, 50)
	
	for i := 0; i < 50; i++ {
		issues[i] = types.GitIssue{
			ID:    fmt.Sprintf("0192d4a4-75d4-7f8a-8c7d-2f8b0d4f1a%02d", i),
			Title: fmt.Sprintf("Sample Issue #%d", i+1),
			Code:  fmt.Sprintf("%d", i+1),
			URL:   fmt.Sprintf("https://github.com/example/repo/issues/%d", i+1),
			CreatedBy: types.GitUser{
				Username:    fmt.Sprintf("dev%d", (i%5)+1),
				DisplayName: fmt.Sprintf("Developer %d", (i%5)+1),
				URL:         fmt.Sprintf("https://github.com/dev%d", (i%5)+1),
			},
			CreatedAt:   int64(1735689600000 + i*3600000), // Sample epoch time in milliseconds
			Description: fmt.Sprintf("This is a sample issue #%d for demonstration purposes", i+1),
		}
	}
	
	result := types.Response{
		Success: true,
		Data:    issues,
	}
	
	return result
}

// SaveGitIssue saves Git repository issue information
func SaveGitIssue(e ui.Event) any {
	// Parse input parameters
	params, err := ui.GetArg[ParamsOfSaveGitIssue](e)
	if err != nil {
		result := types.Response{
			Success: false,
			Msg:     "Invalid parameters",
		}
		return result
	}
	
	// For simulation, we'll always succeed
	fmt.Printf("Saving Git repository issue: %+v\n", params)
	
	result := types.Response{
		Success: true,
	}
	
	return result
}

// GetGitIssueCommentList returns list of comments for a Git issue
func GetGitIssueCommentList(e ui.Event) any {
	// Parse input parameters
	params, err := ui.GetArg[ParamsOfGetGitIssueCommentList](e)
	if err != nil {
		result := types.Response{
			Success: false,
			Msg:     "Invalid parameters",
		}
		return result
	}
	
	// Use the params to avoid "declared and not used" error
	_ = params
	
	// For simulation, return 30 sample comments with 10 from current user
	comments := make([]types.GitIssueComment, 30)
	
	for i := 0; i < 30; i++ {
		// Make 10 comments from the current user (dev1) and 20 from other users
		var username, displayName string
		if i < 10 {
			// Comments from current user
			username = "dev1"
			displayName = "Developer One"
		} else {
			// Comments from other users
			userNum := (i % 5) + 1
			username = fmt.Sprintf("dev%d", userNum)
			displayName = fmt.Sprintf("Developer %d", userNum)
		}
		
		// Create markdown content for comments
		var content string
		switch i % 5 {
		case 0:
			content = fmt.Sprintf("# Comment #%d\n\nThis is a sample comment from %s.\n\n## Key Points\n\n- Point one\n- Point two\n- Point three\n\n[Link to documentation](https://example.com)", i+1, displayName)
		case 1:
			content = fmt.Sprintf("## Update #%d\n\nHello team, I've reviewed the code and here are my thoughts:\n\n```go\nfunc example() {\n    fmt.Println(\"Hello, World!\")\n}\n```\n\nPlease review this implementation.", i+1)
		case 2:
			content = fmt.Sprintf("**Comment #%d** from *%s*\n\n**Laplace** transform:\n$$\n  \\mathcal{L}\\{f\\}(s) = \\int_0^{\\infty} {f(t)e^{-st}dt}\n$$\n\nI think we should consider the following:\n\n1. First item\n2. Second item\n3. Third item\n\n> This is an important quote from the documentation.\n\n---\n\nLet me know what you think.", i+1, displayName)
		case 3:
			content = fmt.Sprintf("# Technical Review #%d\n\n## Summary\n\nThis is a comprehensive review of the changes.\n\n## Code Changes\n\nThe following files were modified:\n\n- `src/main.go`\n- `src/handlers/issue.go`\n\n## Issues Found\n\n- [ ] Memory leak in function `GetIssues()`\n- [x] Fixed typo in variable name\n- [ ] Need to add more test cases\n\n@%s Please address these issues.", i+1, username)
		case 4:
			content = fmt.Sprintf("Comment #%d\n\nThis is a regular comment with some **bold text** and some *italic text*.\n\nHere's a list of items:\n\n* Item one\n* Item two\n* Item three\n\nAnd here's a numbered list:\n\n1. First step\n2. Second step\n3. Third step\n\n---\n\nThanks for your work on this!", i+1)
		}
		
		comments[i] = types.GitIssueComment{
			ID: fmt.Sprintf("0192d4a4-75d4-7f8a-8c7d-2f8b0d4f1a%02d", i),
			CreatedBy: types.GitUser{
				Username:    username,
				DisplayName: displayName,
				URL:         fmt.Sprintf("https://github.com/%s", username),
			},
			CreatedAt: int64(1735689600000 + i*3600000), // Sample epoch time in milliseconds
			Content:   content,
		}
	}
	
	result := types.Response{
		Success: true,
		Data:    comments,
	}
	
	return result
}

// GetGitIssueParticipantList returns list of participants for a Git issue
func GetGitIssueParticipantList(e ui.Event) any {
	// Parse input parameters
	params, err := ui.GetArg[ParamsOfGetGitIssueParticipantList](e)
	if err != nil {
		result := types.Response{
			Success: false,
			Msg:     "Invalid parameters",
		}
		return result
	}
	
	// Use the params to avoid "declared and not used" error
	_ = params
	
	// For simulation, return 20 sample participants
	participants := make([]types.GitUser, 20)
	
	for i := 0; i < 20; i++ {
		userNum := (i % 10) + 1
		participants[i] = types.GitUser{
			ID:          fmt.Sprintf("0192d4a4-75d4-7f8a-8c7d-2f8b0d4f1a%02d", userNum),
			Username:    fmt.Sprintf("dev%d", userNum),
			DisplayName: fmt.Sprintf("Developer %d", userNum),
			URL:         fmt.Sprintf("https://github.com/dev%d", userNum),
		}
	}
	
	result := types.Response{
		Success: true,
		Data:    participants,
	}
	
	return result
}

// SaveGitIssueComment creates a comment for a Git issue
func SaveGitIssueComment(e ui.Event) any {
	// Parse input parameters
	params, err := ui.GetArg[ParamsOfSaveGitIssueComment](e)
	if err != nil {
		result := types.Response{
			Success: false,
			Msg:     "Invalid parameters",
		}
		return result
	}
	
	// For simulation, we'll always succeed
	fmt.Printf("Saving Git issue comment: %+v\n", params)
	
	result := types.Response{
		Success: true,
	}
	
	return result
}
