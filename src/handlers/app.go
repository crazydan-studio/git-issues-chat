package handlers

import (
	"fmt"
	ui "github.com/webui-dev/go-webui/v2"
)

// GetAppInfo returns application information
func GetAppInfo(e ui.Event) any {
	appInfo := map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"name":    "Git Issues Chat",
			"version": "0.1.0",
			"author":  "flytreeleft@crazydan.org",
			"source":  "https://github.com/yourusername/git-issues-chat",
			"license": map[string]string{
				"name": "Apache 2.0",
				"url":  "https://www.apache.org/licenses/LICENSE-2.0.txt",
			},
			"description": "A desktop chat application that transforms GitHub repositories into chat rooms using Git Issues as topics and Comments as messages.",
		},
	}
	
	return appInfo
}

// VerifyAppUser verifies user password
func VerifyAppUser(e ui.Event) any {
	// Parse input parameters
	params, err := ui.GetArg[map[string]interface{}](e)
	if err != nil {
		result := map[string]interface{}{
			"success": false,
			"msg":     "Invalid parameters",
		}
		return result
	}
	
	// For simulation, we'll accept any non-empty password
	password, _ := params["password"].(string)
	userId, _ := params["userId"].(string)
	
	if (userId == "" && password != "") || (userId != "" && password != "") {
		user := map[string]interface{}{
			"id":          "0192d4a4-75d4-7f8a-8c7d-2f8b0d4f1a7f",
			"username":    "dev1",
			"displayName": "Git User",
			"avatar":      "avatar-hash-123",
		}
		
		result := map[string]interface{}{
			"success": true,
			"data":    user,
		}
		
		return result
	}
	
	result := map[string]interface{}{
		"success": false,
		"msg":     "username or password is wrong",
	}
	
	return result
}

// SaveAppUserInfo saves user information
func SaveAppUserInfo(e ui.Event) any {
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
	fmt.Printf("Saving user info: %+v\n", params)
	
	result := map[string]interface{}{
		"success": true,
	}
	
	return result
}

// UpdateAppUserPassword updates user password
func UpdateAppUserPassword(e ui.Event) any {
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
	fmt.Printf("Updating user password: %+v\n", params)
	
	result := map[string]interface{}{
		"success": true,
	}
	
	return result
}

// GetAppUserActionLogList gets user action log list
func GetAppUserActionLogList(e ui.Event) any {
	// Parse input parameters
	params, err := ui.GetArg[map[string]interface{}](e)
	if err != nil {
		result := map[string]interface{}{
			"success": false,
			"msg":     "Invalid parameters",
		}
		return result
	}

	userId, _ := params["userId"].(string)
	if userId == "" {
		result := map[string]interface{}{
			"success": false,
			"msg":     "User ID is required",
		}
		return result
	}

	// For simulation, return 20 sample action logs
	logs := make([]map[string]interface{}, 20)

	for i := 0; i < 20; i++ {
		// Alternate between success and error status
		status := "success"
		if i%3 == 0 {
			status = "error"
		}

		logs[i] = map[string]interface{}{
			"id":        fmt.Sprintf("0192d4a4-75d4-7f8a-8c7d-2f8b0d4f1a%02d", 20-i),
			"status":    status,
			"createdAt": fmt.Sprintf("2025-01-%02d %02d:%02d:%02d", (i%28)+1, (i%24), (i%60), (i%60)),
			"content":   fmt.Sprintf("Sample action log entry #%d for demonstration purposes", 20-i),
		}
	}

	result := map[string]interface{}{
		"success": true,
		"data":    logs,
	}

	return result
}