package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/webui-dev/go-webui/v2"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"git-issues-chat/src/handlers"
)

func initAppEnv(dataPath string) (*sql.DB, error) {
	// Database file path
	dbPath := filepath.Join(dataPath, "chat.db")

	// Open SQLite database (this will create the file if it doesn't exist)
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return db, nil
}

func main() {
	// Define command line flags
	versionFlag := flag.Bool("version", false, "Print version information")
	dataPath := flag.String("data-path", "", "Path to data directory (required)")
	uiPath := flag.String("ui-path", "", "Path to UI directory (required)")
	browser := flag.String("browser", "", "Browser to use: firefox, chrome, edge (required)")

	flag.Parse()

	// Handle version flag
	if *versionFlag {
		fmt.Println("Git Issues Chat v0.1.0")
		return
	}

	// Check if required flags are provided
	if *dataPath == "" || *uiPath == "" || *browser == "" {
		fmt.Println("Error: --data-path, --ui-path, and --browser are required")
		flag.Usage()
		os.Exit(1)
	}

	// Initialize application environment
	db, err := initAppEnv(*dataPath)
	if err != nil {
		log.Fatal("Failed to initialize app environment:", err)
	}
	defer db.Close()

	// Create new webui window
	w := webui.NewWindow()

	// Bind handler functions to the window
	// App handlers
	w.Bind("getAppInfo", handlers.GetAppInfo)
	w.Bind("verifyAppUser", handlers.VerifyAppUser)
	w.Bind("saveAppUserInfo", handlers.SaveAppUserInfo)
	w.Bind("updateAppUserPassword", handlers.UpdateAppUserPassword)

	// Platform handlers
	w.Bind("getGitPlatformList", handlers.GetGitPlatformList)
	w.Bind("saveGitPlatform", handlers.SaveGitPlatform)
	w.Bind("getGitRepoList", handlers.GetGitRepoList)
	w.Bind("getGitRepoInfo", handlers.GetGitRepoInfo)
	w.Bind("saveGitRepo", handlers.SaveGitRepo)

	// Repo handlers
	w.Bind("getGitIssueList", handlers.GetGitIssueList)
	w.Bind("saveGitIssue", handlers.SaveGitIssue)

	// Issue handlers
	w.Bind("getGitIssueCommentList", handlers.GetGitIssueCommentList)
	w.Bind("getGitIssueParticipantList", handlers.GetGitIssueParticipantList)
	w.Bind("saveGitIssueComment", handlers.SaveGitIssueComment)

	// Set the root folder for UI files
	w.SetRootFolder(*uiPath)

	// Determine browser type
	var browserType webui.Browser
	switch *browser {
	case "firefox":
		browserType = webui.Firefox
	case "chrome":
		browserType = webui.Chrome
	case "edge":
		browserType = webui.Edge
	default:
		log.Fatal("Invalid browser type. Use firefox, chrome, or edge")
	}

	// Show the UI in the specified browser
	err = w.ShowBrowser("index.html", browserType)
	if err != nil {
		log.Fatal("Failed to show browser window:", err)
	}

	// Wait for the window to be closed
	webui.Wait()
}