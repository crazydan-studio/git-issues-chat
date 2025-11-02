package handlers

import (
	"fmt"
	ui "github.com/webui-dev/go-webui/v2"
	"git-issues-chat/src/types"
)

// GetGitPlatformList returns list of Git platforms
func GetGitPlatformList(e ui.Event) any {
	platforms := []types.GitPlatform{
		{
			ID:          "0192d4a4-75d4-7f8a-8c7d-2f8b0d4f1a80",
			Name:        "GitHub",
			URL:         "https://github.com",
			Icon:        "github-icon-hash",
			Description: "GitHub is a development platform inspired by the way you work.",
		},
		{
			ID:          "0192d4a4-75d4-7f8a-8c7d-2f8b0d4f1a81",
			Name:        "GitLab",
			URL:         "https://gitlab.com",
			Icon:        "gitlab-icon-hash",
			Description: "GitLab is a complete DevOps platform, delivered as a single application.",
		},
	}
	
	result := types.Response{
		Success: true,
		Data:    platforms,
	}
	
	return result
}

// SaveGitPlatform saves Git platform information
func SaveGitPlatform(e ui.Event) any {
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
	fmt.Printf("Saving Git platform: %+v\n", params)
	
	result := types.Response{
		Success: true,
	}
	
	return result
}