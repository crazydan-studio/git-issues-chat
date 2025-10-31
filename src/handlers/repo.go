package handlers

import (
	"fmt"
	ui "github.com/webui-dev/go-webui/v2"
)

// GetGitRepoIssueList returns list of issues for a Git repository
func GetGitRepoIssueList(e ui.Event) any {
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

// SaveGitRepoIssue saves Git repository issue information
func SaveGitRepoIssue(e ui.Event) any {
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