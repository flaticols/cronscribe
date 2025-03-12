package main

import (
	"fmt"
	"os"

	"github.com/flaticols/cronscribe"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go \"expression\"")
		os.Exit(1)
	}

	expression := os.Args[1]
	
	// Create a new CronScribe instance
	cs, err := cronscribe.New()
	if err != nil {
		fmt.Printf("Error creating CronScribe: %v\n", err)
		os.Exit(1)
	}

	// Try to auto-detect the language and convert
	result, err := cs.AutoDetect(expression)
	if err != nil {
		fmt.Printf("Error auto-detecting: %v\n", err)
		
		// Try Russian specifically
		if err := cs.SetLanguage("ru"); err != nil {
			fmt.Printf("Error setting language to Russian: %v\n", err)
			os.Exit(1)
		}
		
		result, err = cs.Convert(expression)
		if err != nil {
			fmt.Printf("Error converting with Russian: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Printf("Expression: %s\nCron: %s\n", expression, result)
}