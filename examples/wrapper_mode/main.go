package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/flaticols/cronscribe"
)

func main() {
	// Find the rules directory
	exDir, err := filepath.Abs(".")
	if err != nil {
		fmt.Printf("Failed to get current directory: %v\n", err)
		return
	}
	
	// Try different paths for rules directory
	var rulesPath string
	possiblePaths := []string{
		filepath.Join(exDir, "pkg", "core", "rules"),                   // If running from project root
		filepath.Join(exDir, "..", "..", "pkg", "core", "rules"),       // If running from examples/wrapper_mode
	}
	
	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			rulesPath = path
			break
		}
	}
	
	if rulesPath == "" {
		fmt.Printf("Unable to find rules directory\n")
		return
	}
	
	fmt.Printf("Using rules from: %s\n", rulesPath)
	
	// Create a new CronScribe instance with default rules directory
	mapper, err := cronscribe.New(rulesPath)
	if err != nil {
		fmt.Printf("Error creating mapper: %v\n", err)
		return
	}

	// Display supported languages
	fmt.Println("Supported languages:", mapper.GetSupportedLanguages())

	// Example expressions
	expressions := []string{
		"every day at noon",
		"every Monday at 9:15 AM",
		"first day of every month at 3 PM",
		"every 15 minutes from 9am to 5pm on weekdays",
		"every four hours starting at midnight",
	}

	// Using English rules (default)
	fmt.Println("\n== English expressions ==")
	for _, expr := range expressions {
		fmt.Printf("Human: %s\n", expr)
		
		cronExpr, err := mapper.Convert(expr)
		if err != nil {
			fmt.Printf("Error: %s\n\n", err)
			continue
		}
		
		fmt.Printf("Cron: %s\n\n", cronExpr)
	}

	// Try with another language if available
	if len(mapper.GetSupportedLanguages()) > 1 {
		// Use first non-English language
		for _, lang := range mapper.GetSupportedLanguages() {
			if lang != "en" {
				fmt.Printf("\n== %s language ==\n", lang)
				mapper.SetLanguage(lang)
				
				// Try a simple expression
				expr := "every day at noon"
				fmt.Printf("Human: %s\n", expr)
				
				cronExpr, err := mapper.Convert(expr)
				if err != nil {
					fmt.Printf("Error: %s\n\n", err)
				} else {
					fmt.Printf("Cron: %s\n\n", cronExpr)
				}
				
				break
			}
		}
	}

	// Try auto-detection
	fmt.Println("\n== Auto-detection ==")
	expr := "every day at noon"
	fmt.Printf("Human: %s\n", expr)
	
	cronExpr, err := mapper.AutoDetect(expr)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Cron: %s\n", cronExpr)
	}
}