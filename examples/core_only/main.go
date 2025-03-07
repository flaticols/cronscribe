package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/flaticols/cronscribe/pkg/core"
)

func main() {
	// Get absolute path to the examples directory
	exDir, err := filepath.Abs(".")
	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err)
	}
	fmt.Printf("Current directory: %s\n", exDir)
	
	// Try different paths for rules directory
	var rulesPath string
	possiblePaths := []string{
		filepath.Join(exDir, "pkg", "core", "rules"),                   // If running from project root
		filepath.Join(exDir, "..", "..", "pkg", "core", "rules"),       // If running from examples/core_only
	}
	
	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			rulesPath = path
			break
		}
	}
	
	if rulesPath == "" {
		log.Fatalf("Unable to find rules directory in any of: %v", possiblePaths)
	}
	
	fmt.Printf("Found rules directory at: %s\n", rulesPath)
	
	// List files in the directory
	files, err := filepath.Glob(filepath.Join(rulesPath, "*.yaml"))
	if err != nil {
		log.Fatalf("Failed to glob files: %v", err)
	}
	fmt.Printf("Found rule files: %v\n", files)
	
	cs, err := core.New(rulesPath)
	if err != nil {
		log.Fatalf("Failed to create CronScribe: %v", err)
	}

	// Convert a human-readable expression to cron
	cronExpr, err := cs.Convert("every day at 10:30")
	if err != nil {
		log.Fatalf("Failed to convert: %v", err)
	}

	fmt.Printf("Cron expression: %s\n", cronExpr)

	// List supported languages
	languages := cs.GetSupportedLanguages()
	fmt.Printf("Supported languages: %v\n", languages)

	// Try with another language
	err = cs.SetLanguage("ru")
	if err != nil {
		log.Fatalf("Failed to set language: %v", err)
	}

	cronExpr, err = cs.Convert("каждый день в 10:30")
	if err != nil {
		log.Fatalf("Failed to convert: %v", err)
	}

	fmt.Printf("Russian cron expression: %s\n", cronExpr)

	// Auto-detect language
	cronExpr, err = cs.AutoDetect("elke dag om 10:30")
	if err != nil {
		log.Fatalf("Failed to auto-detect: %v", err)
	}

	fmt.Printf("Auto-detected cron expression: %s\n", cronExpr)
}