package handlers

import (
	"fmt"
	ui "github.com/webui-dev/go-webui/v2"
	"git-issues-chat/src/types"
)

// GetAppInfo returns application information
func GetAppInfo(e ui.Event) any {
	appInfo := types.Response{
		Success: true,
		Data: types.App{
			Name:        "Git Issues Chat",
			Version:     "0.1.0",
			Author:      "flytreeleft@crazydan.org",
			Source:      "https://github.com/yourusername/git-issues-chat",
			Description: "A desktop chat application that transforms GitHub repositories into chat rooms using Git Issues as topics and Comments as messages.",
			License: types.License{
				Name: "Apache 2.0",
				URL:  "https://www.apache.org/licenses/LICENSE-2.0.txt",
			},
		},
	}
	
	return appInfo
}

// VerifyAppUser verifies user password
func VerifyAppUser(e ui.Event) any {
	// Parse input parameters
	params, err := ui.GetArg[map[string]interface{}](e)
	if err != nil {
		result := types.Response{
			Success: false,
			Msg:     "Invalid parameters",
		}
		return result
	}
	
	// For simulation, we'll accept any non-empty password
	password, _ := params["password"].(string)
	userId, _ := params["userId"].(string)
	
	if (userId == "" && password != "") || (userId != "" && password != "") {
		user := types.AppUser{
			ID:          "0192d4a4-75d4-7f8a-8c7d-2f8b0d4f1a7f",
			Username:    "dev1",
			DisplayName: "Git User",
			Avatar:      "avatar-hash-123",
		}
		
		result := types.Response{
			Success: true,
			Data:    user,
		}
		
		return result
	}
	
	result := types.Response{
		Success: false,
		Msg:     "username or password is wrong",
	}
	
	return result
}

// SaveAppUserInfo saves user information
func SaveAppUserInfo(e ui.Event) any {
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
	fmt.Printf("Saving user info: %+v\n", params)
	
	result := types.Response{
		Success: true,
	}
	
	return result
}

// UpdateAppUserPassword updates user password
func UpdateAppUserPassword(e ui.Event) any {
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
	fmt.Printf("Updating user password: %+v\n", params)
	
	result := types.Response{
		Success: true,
	}
	
	return result
}

// GetAppUserActionLogList gets user action log list
func GetAppUserActionLogList(e ui.Event) any {
	// Parse input parameters
	params, err := ui.GetArg[map[string]interface{}](e)
	if err != nil {
		result := types.Response{
			Success: false,
			Msg:     "Invalid parameters",
		}
		return result
	}

	userId, _ := params["userId"].(string)
	if userId == "" {
		result := types.Response{
			Success: false,
			Msg:     "User ID is required",
		}
		return result
	}

	// For simulation, return 20 sample action logs
	logs := make([]types.AppUserActionLog, 20)

	for i := 0; i < 20; i++ {
		// Alternate between success and error status
		status := "success"
		if i%3 == 0 {
			status = "error"
		}

		logs[i] = types.AppUserActionLog{
			ID:        fmt.Sprintf("0192d4a4-75d4-7f8a-8c7d-2f8b0d4f1a%02d", 20-i),
			Status:    status,
			CreatedAt: int64(1735689600000 + i*3600000), // Sample epoch time in milliseconds
			Content:   fmt.Sprintf("Sample action log entry #%d for demonstration purposes", 20-i),
		}
	}

	result := types.Response{
		Success: true,
		Data:    logs,
	}

	return result
}