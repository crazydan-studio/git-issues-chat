package handlers

import (
	"fmt"
	ui "github.com/webui-dev/go-webui/v2"
	"git-issues-chat/src/types"
)

// GetGitIssueList returns list of issues for a Git repository
func GetGitIssueList(e ui.Event) any {
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
	params, err := ui.GetArg[map[string]interface{}](e)
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
		
		comments[i] = types.GitIssueComment{
			ID: fmt.Sprintf("0192d4a4-75d4-7f8a-8c7d-2f8b0d4f1a%02d", i),
			CreatedBy: types.GitUser{
				Username:    username,
				DisplayName: displayName,
				URL:         fmt.Sprintf("https://github.com/%s", username),
			},
			CreatedAt: int64(1735689600000 + i*3600000), // Sample epoch time in milliseconds
			Content:   fmt.Sprintf("This is sample comment #%d for the issue. This comment is from %s.", i+1, displayName),
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
	params, err := ui.GetArg[map[string]interface{}](e)
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