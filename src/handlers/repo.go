package handlers

import (
	"fmt"
	ui "github.com/webui-dev/go-webui/v2"
	"git-issues-chat/src/types"
)

// GetGitRepoList returns list of repositories for a Git platform
func GetGitRepoList(e ui.Event) any {
	// Parse input parameters
	params, err := ui.GetArg[map[string]interface{}](e)
	if err != nil {
		result := types.Response{
			Success: false,
			Msg:     "Invalid parameters",
		}
		return result
	}
	
	platformId, _ := params["platformId"].(string)
	
	// For simulation, return 20 sample repositories
	repos := make([]types.GitRepo, 20)
	
	for i := 0; i < 20; i++ {
		repos[i] = types.GitRepo{
			ID:          fmt.Sprintf("0192d4a4-75d4-7f8a-8c7d-2f8b0d4f1a%02d", i+2),
			Platform:    platformId,
			Name:        fmt.Sprintf("user%d/repository%d", (i%5)+1, i+1),
			URL:         fmt.Sprintf("https://github.com/user%d/repository%d", (i%5)+1, i+1),
			Icon:        fmt.Sprintf("repo-icon-hash-%d", i+1),
			Description: fmt.Sprintf("This is sample repository #%d for demonstration purposes on platform %s.", i+1, platformId),
		}
	}
	
	result := types.Response{
		Success: true,
		Data:    repos,
	}
	
	return result
}

// GetGitRepoInfo returns information about a specific repository
func GetGitRepoInfo(e ui.Event) any {
	// Parse input parameters
	params, err := ui.GetArg[map[string]interface{}](e)
	if err != nil {
		result := types.Response{
			Success: false,
			Msg:     "Invalid parameters",
		}
		return result
	}
	
	platformId, _ := params["platformId"].(string)
	repoName, _ := params["repoName"].(string)
	
	// For simulation, return sample repository info
	repo := types.GitRepo{
		Name:        repoName,
		URL:         fmt.Sprintf("https://github.com/%s", repoName),
		Icon:        "repo-icon-hash-new",
		Description: fmt.Sprintf("Repository %s on platform %s", repoName, platformId),
	}
	
	result := types.Response{
		Success: true,
		Data:    repo,
	}
	
	return result
}

// SaveGitRepo saves repository information to a Git platform
func SaveGitRepo(e ui.Event) any {
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
	fmt.Printf("Saving Git platform repository: %+v\n", params)
	
	result := types.Response{
		Success: true,
	}
	
	return result
}