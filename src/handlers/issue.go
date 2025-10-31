package handlers

import (
	"fmt"
	ui "github.com/webui-dev/go-webui/v2"
)

// GetGitIssueCommentList returns list of comments for a Git issue
func GetGitIssueCommentList(e ui.Event) any {
	// For simulation, return some sample comments
	comments := []map[string]interface{}{
		{
			"id": "0192d4a4-75d4-7f8a-8c7d-2f8b0d4f1a86",
			"author": map[string]string{
				"username":    "dev1",
				"displayName": "Developer One",
				"url":         "https://github.com/dev1",
			},
			"createdAt": "2025-01-10 23:15:32",
			"content":   "I'm working on this issue and should have a fix ready by tomorrow.",
		},
		{
			"id": "0192d4a4-75d4-7f8a-8c7d-2f8b0d4f1a87",
			"author": map[string]string{
				"username":    "dev2",
				"displayName": "Developer Two",
				"url":         "https://github.com/dev2",
			},
			"createdAt": "2025-01-11 09:45:22",
			"content":   "Thanks for the update. Please make sure to add unit tests for the authentication module.",
		},
	}
	
	result := map[string]interface{}{
		"success": true,
		"data":    comments,
	}
	
	return result
}

// GetGitIssueParticipantList returns list of participants for a Git issue
func GetGitIssueParticipantList(e ui.Event) any {
	// For simulation, return some sample participants
	participants := []map[string]interface{}{
		{
			"username":    "dev1",
			"displayName": "Developer One",
			"url":         "https://github.com/dev1",
		},
		{
			"username":    "dev2",
			"displayName": "Developer Two",
			"url":         "https://github.com/dev2",
		},
		{
			"username":    "dev3",
			"displayName": "Developer Three",
			"url":         "https://github.com/dev3",
		},
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