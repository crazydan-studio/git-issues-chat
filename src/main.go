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
)

func initAppEnv(dataPath string) (*sql.DB, error) {
	// Create data directory if it doesn't exist
	err := os.MkdirAll(dataPath, 0755)
	if err != nil {
		return nil, fmt.Errorf("failed to create data directory: %v", err)
	}

	// Database file path
	dbPath := filepath.Join(dataPath, "chat.db")

	// Open SQLite database
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
	w.ShowBrowser("index.html", browserType)

	// Wait for the window to be closed
	webui.Wait()
}