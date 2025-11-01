package handlers

import (
	"fmt"
	ui "github.com/webui-dev/go-webui/v2"
)

// GetGitPlatformList returns list of Git platforms
func GetGitPlatformList(e ui.Event) any {
	platforms := []map[string]interface{}{
		{
			"id":          "0192d4a4-75d4-7f8a-8c7d-2f8b0d4f1a80",
			"name":        "GitHub",
			"url":         "https://github.com",
			"icon":        "github-icon-hash",
			"description": "GitHub is a development platform inspired by the way you work.",
		},
		{
			"id":          "0192d4a4-75d4-7f8a-8c7d-2f8b0d4f1a81",
			"name":        "GitLab",
			"url":         "https://gitlab.com",
			"icon":        "gitlab-icon-hash",
			"description": "GitLab is a complete DevOps platform, delivered as a single application.",
		},
	}
	
	result := map[string]interface{}{
		"success": true,
		"data":    platforms,
	}
	
	return result
}

// SaveGitPlatform saves Git platform information
func SaveGitPlatform(e ui.Event) any {
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
	fmt.Printf("Saving Git platform: %+v\n", params)
	
	result := map[string]interface{}{
		"success": true,
	}
	
	return result
}