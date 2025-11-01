package handlers

import (
	"fmt"
	ui "github.com/webui-dev/go-webui/v2"
)

// GetGitIssueList returns list of issues for a Git repository
func GetGitIssueList(e ui.Event) any {
	// For simulation, return 50 sample issues
	issues := make([]map[string]interface{}, 50)
	
	for i := 0; i < 50; i++ {
		issues[i] = map[string]interface{}{
			"id":    fmt.Sprintf("0192d4a4-75d4-7f8a-8c7d-2f8b0d4f1a%02d", i),
			"title": fmt.Sprintf("Sample Issue #%d", i+1),
			"code":  fmt.Sprintf("%d", i+1),
			"url":   fmt.Sprintf("https://github.com/example/repo/issues/%d", i+1),
			"author": map[string]string{
				"username":    fmt.Sprintf("dev%d", (i%5)+1),
				"displayName": fmt.Sprintf("Developer %d", (i%5)+1),
				"url":         fmt.Sprintf("https://github.com/dev%d", (i%5)+1),
			},
			"createdAt":   fmt.Sprintf("2025-01-%02d %02d:%02d:%02d", (i%28)+1, (i%24), (i%60), (i%60)),
			"description": fmt.Sprintf("This is a sample issue #%d for demonstration purposes", i+1),
		}
	}
	
	result := map[string]interface{}{
		"success": true,
		"data":    issues,
	}
	
	return result
}

// SaveGitIssue saves Git repository issue information
func SaveGitIssue(e ui.Event) any {
	// Parse input parameters
	params, err := ui.GetArg[map[string]interface{}](e)
	if err != nil {
		result := map[string]interface{}{
			"success": false,
			"msg":     "Invalid parameters",
		}
		return result
	}
	
	// For simulation, we'll always succeed
	fmt.Printf("Saving Git repository issue: %+v\n", params)
	
	result := map[string]interface{}{
		"success": true,
	}
	
	return result
}

// GetGitIssueCommentList returns list of comments for a Git issue
func GetGitIssueCommentList(e ui.Event) any {
	// For simulation, return 30 sample comments with 10 from current user
	comments := make([]map[string]interface{}, 30)
	
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
		
		comments[i] = map[string]interface{}{
			"id": fmt.Sprintf("0192d4a4-75d4-7f8a-8c7d-2f8b0d4f1a%02d", i),
			"author": map[string]interface{}{
				"username":    username,
				"displayName": displayName,
				"url":         fmt.Sprintf("https://github.com/%s", username),
			},
			"createdAt": fmt.Sprintf("2025-01-%02d %02d:%02d:%02d", (i%28)+1, (i%24), (i%60), (i%60)),
			"content":   fmt.Sprintf("This is sample comment #%d for the issue. This comment is from %s.", i+1, displayName),
		}
	}
	
	result := map[string]interface{}{
		"success": true,
		"data":    comments,
	}
	
	return result
}

// GetGitIssueParticipantList returns list of participants for a Git issue
func GetGitIssueParticipantList(e ui.Event) any {
	// For simulation, return 20 sample participants
	participants := make([]map[string]interface{}, 20)
	
	for i := 0; i < 20; i++ {
		userNum := (i % 10) + 1
		participants[i] = map[string]interface{}{
			"username":    fmt.Sprintf("dev%d", userNum),
			"displayName": fmt.Sprintf("Developer %d", userNum),
			"url":         fmt.Sprintf("https://github.com/dev%d", userNum),
		}
	}
	
	result := map[string]interface{}{
		"success": true,
		"data":    participants,
	}
	
	return result
}

// SaveGitIssueComment creates a comment for a Git issue
func SaveGitIssueComment(e ui.Event) any {
	// Parse input parameters
	params, err := ui.GetArg[map[string]interface{}](e)
	if err != nil {
		result := map[string]interface{}{
			"success": false,
			"msg":     "Invalid parameters",
		}
		return result
	}
	
	// For simulation, we'll always succeed
	fmt.Printf("Saving Git issue comment: %+v\n", params)
	
	result := map[string]interface{}{
		"success": true,
	}
	
	return result
}