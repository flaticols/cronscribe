package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/flaticols/cronscribe/pkg/ai"
	"github.com/flaticols/cronscribe/pkg/core"
	"github.com/sashabaranov/go-openai"
)

// OpenAIProvider implements ai.AIProvider using OpenAI API
type OpenAIProvider struct {
	client *openai.Client
}

// NewOpenAIProvider creates a new OpenAI provider
func NewOpenAIProvider(apiKey string) *OpenAIProvider {
	client := openai.NewClient(apiKey)
	return &OpenAIProvider{client: client}
}

// GenerateCron generates a cron expression from human text using OpenAI
func (p *OpenAIProvider) GenerateCron(ctx context.Context, input string) (string, error) {
	// Use recommended prompts from ai package
	systemPrompt := ai.RecommendedSystemPrompt()
	userPrompt := ai.RecommendedUserPrompt(input)

	// Call OpenAI API
	resp, err := p.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemPrompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: userPrompt,
				},
			},
			Temperature: 0.0,
		},
	)

	if err != nil {
		return "", fmt.Errorf("OpenAI API error: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI API")
	}

	return resp.Choices[0].Message.Content, nil
}

func main() {
	// Get API key from environment variable
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("Please set the OPENAI_API_KEY environment variable")
		return
	}

	// Find the rules directory
	exDir, err := filepath.Abs(".")
	if err != nil {
		fmt.Println("Failed to get current directory:", err)
		return
	}
	
	// Try different paths for rules directory
	var rulesPath string
	possiblePaths := []string{
		filepath.Join(exDir, "pkg", "core", "rules"),                   // If running from project root
		filepath.Join(exDir, "..", "..", "pkg", "core", "rules"),       // If running from examples/brave_mode
	}
	
	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			rulesPath = path
			break
		}
	}
	
	if rulesPath == "" {
		fmt.Println("Unable to find rules directory")
		return
	}
	
	fmt.Println("Using rules from:", rulesPath)
	
	// OPTION 1: Create standalone CronScribeAI instance
	fmt.Println("=== Using AI with fallback to local rules ===")
	openaiProvider := NewOpenAIProvider(apiKey)
	cronscribeAI, err := ai.New(
		rulesPath,
		openaiProvider,
		ai.WithAIFirst(false), // Try local rules first, fallback to AI
	)
	if err != nil {
		fmt.Println("Error creating CronScribeAI:", err)
		return
	}

	// OPTION 2: Create core instance first, then integrate with AI
	fmt.Println("\n=== Using core instance with AI extension ===")
	coreInstance, err := core.New(rulesPath)
	if err != nil {
		fmt.Println("Error creating CronScribe core:", err)
		return
	}

	// Now create AI version using the existing core instance
	cronscribeWithCore, err := ai.WithCore(
		coreInstance,
		openaiProvider,
		ai.WithAIFirst(true), // Try AI first this time
	)
	if err != nil {
		fmt.Println("Error creating AI with core:", err)
		return
	}

	// Example expressions
	expressions := []string{
		"every day at noon",
		"every Monday at 9:15 AM",
		"first day of every month at 3 PM",
		"every 15 minutes from 9am to 5pm on weekdays",
		"every four hours starting at midnight",
	}

	// Test with first instance (local rules first)
	fmt.Println("\n--- Using local rules first, AI as fallback ---")
	for _, expr := range expressions {
		fmt.Printf("Human: %s\n", expr)
		
		cronExpr, err := cronscribeAI.ToCron(expr)
		if err != nil {
			fmt.Printf("Error: %s\n\n", err)
			continue
		}
		
		fmt.Printf("Cron: %s\n\n", cronExpr)
	}

	// Test with second instance (AI first)
	fmt.Println("\n--- Using AI first, local rules as fallback ---")
	for _, expr := range expressions {
		fmt.Printf("Human: %s\n", expr)
		
		cronExpr, err := cronscribeWithCore.ToCron(expr)
		if err != nil {
			fmt.Printf("Error: %s\n\n", err)
			continue
		}
		
		fmt.Printf("Cron: %s\n\n", cronExpr)
	}

	// Demonstrate language support using core functionality
	fmt.Println("\n--- Demonstrating language features (from core) ---")
	fmt.Printf("Supported languages: %v\n", cronscribeAI.GetSupportedLanguages())
	
	// Try auto-detection
	russianExpr := "каждый понедельник в 9 утра"
	fmt.Printf("Russian: %s\n", russianExpr)
	cronExpr, err := cronscribeAI.AutoDetect(russianExpr)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Cron: %s\n", cronExpr)
	}
}