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

// GetGitPlatformRepoList returns list of repositories for a Git platform
func GetGitPlatformRepoList(e ui.Event) any {
	// Parse input parameters
	params, err := ui.GetArg[map[string]interface{}](e)
	if err != nil {
		result := map[string]interface{}{
			"success": false,
			"msg":     "Invalid parameters",
		}
		return result
	}
	
	platformId, _ := params["platformId"].(string)
	
	// For simulation, return 20 sample repositories
	repos := make([]map[string]interface{}, 20)
	
	for i := 0; i < 20; i++ {
		repos[i] = map[string]interface{}{
			"id":          fmt.Sprintf("0192d4a4-75d4-7f8a-8c7d-2f8b0d4f1a%02d", i+2),
			"platform":    platformId,
			"name":        fmt.Sprintf("user%d/repository%d", (i%5)+1, i+1),
			"url":         fmt.Sprintf("https://github.com/user%d/repository%d", (i%5)+1, i+1),
			"icon":        fmt.Sprintf("repo-icon-hash-%d", i+1),
			"description": fmt.Sprintf("This is sample repository #%d for demonstration purposes on platform %s.", i+1, platformId),
		}
	}
	
	result := map[string]interface{}{
		"success": true,
		"data":    repos,
	}
	
	return result
}

// GetGitPlatformRepoInfo returns information about a specific repository
func GetGitPlatformRepoInfo(e ui.Event) any {
	// Parse input parameters
	params, err := ui.GetArg[map[string]interface{}](e)
	if err != nil {
		result := map[string]interface{}{
			"success": false,
			"msg":     "Invalid parameters",
		}
		return result
	}
	
	platformId, _ := params["platformId"].(string)
	repoName, _ := params["repoName"].(string)
	
	// For simulation, return sample repository info
	repo := map[string]interface{}{
		"name":        repoName,
		"url":         fmt.Sprintf("https://github.com/%s", repoName),
		"icon":        "repo-icon-hash-new",
		"description": fmt.Sprintf("Repository %s on platform %s", repoName, platformId),
	}
	
	result := map[string]interface{}{
		"success": true,
		"data":    repo,
	}
	
	return result
}

// SaveGitPlatformRepo saves repository information to a Git platform
func SaveGitPlatformRepo(e ui.Event) any {
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
	fmt.Printf("Saving Git platform repository: %+v\n", params)
	
	result := map[string]interface{}{
		"success": true,
	}
	
	return result
}