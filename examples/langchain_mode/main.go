package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/flaticols/cronscribe/pkg/ai"
	"github.com/flaticols/cronscribe/pkg/core"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

// LangChainProvider implements ai.AIProvider using the LangChain library
type LangChainProvider struct {
	model llms.Model
}

// NewLangChainProvider creates a new LangChain provider
func NewLangChainProvider(opts ...LangChainOption) (*LangChainProvider, error) {
	// Default options
	options := langChainOptions{
		model: "gpt-3.5-turbo",
	}

	// Apply options
	for _, opt := range opts {
		opt(&options)
	}

	// Initialize OpenAI LLM
	llm, err := openai.New(openai.WithModel(options.model))
	if err != nil {
		return nil, fmt.Errorf("failed to create LangChain model: %w", err)
	}

	return &LangChainProvider{
		model: llm,
	}, nil
}

// GenerateCron generates a cron expression from a human text using LangChain
func (p *LangChainProvider) GenerateCron(ctx context.Context, input string) (string, error) {
	// Use recommended prompts from the AI package
	systemPrompt := ai.RecommendedSystemPrompt()
	userPrompt := ai.RecommendedUserPrompt(input)

	// Create message content
	content := []llms.MessageContent{
		{

			Role:  "system",
			Parts: []llms.ContentPart{llms.TextContent{Text: systemPrompt}},
		},
		{
			Role:  "user",
			Parts: []llms.ContentPart{llms.TextContent{Text: userPrompt}},
		},
	}

	// Generate completion using LangChain
	resp, err := p.model.GenerateContent(ctx, content, llms.WithTemperature(0.0))
	if err != nil {
		return "", fmt.Errorf("LangChain error: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from LangChain")
	}

	return resp.Choices[0].Content, nil
}

// Options for configuring the LangChain provider
type langChainOptions struct {
	model string
}

// LangChainOption defines a function to configure langChainOptions
type LangChainOption func(*langChainOptions)

// WithLangChainModel sets the model to use with LangChain
func WithLangChainModel(model string) LangChainOption {
	return func(opts *langChainOptions) {
		opts.model = model
	}
}

func main() {
	// Check if OpenAI API key is set
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY environment variable must be set")
	}

	// Create a LangChain provider with default model (gpt-3.5-turbo)
	provider, err := NewLangChainProvider()
	if err != nil {
		log.Fatalf("Failed to create LangChain provider: %v", err)
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
		filepath.Join(exDir, "pkg", "core", "rules"),             // If running from project root
		filepath.Join(exDir, "..", "..", "pkg", "core", "rules"), // If running from examples/langchain_mode
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

	// OPTION 1: Create standalone CronScribeAI instance with LangChain provider
	fmt.Println("=== Using LangChain with fallback to local rules ===")
	cronscribeAI, err := ai.New(
		rulesPath,
		provider,
		ai.WithAIFirst(true), // Try AI first, fallback to local rules
	)
	if err != nil {
		fmt.Println("Error creating CronScribeAI:", err)
		return
	}

	// Example expressions
	expressions := []string{
		"every day at 3 PM",
		"every weekday at 9 AM",
		"at 10:30 PM on the 1st and 15th of every month",
		"every hour from 9 AM to 5 PM on weekdays",
	}

	fmt.Println("\n--- Converting human-readable schedules using LangChain ---")
	for _, expr := range expressions {
		fmt.Printf("Human: %s\n", expr)

		cronExpr, err := cronscribeAI.ToCron(expr)
		if err != nil {
			fmt.Printf("Error: %s\n\n", err)
			continue
		}

		fmt.Printf("Cron: %s\n\n", cronExpr)
	}

	// Example using a different model
	fmt.Println("\n=== Using a more capable model ===")

	// Create a provider with a more capable model
	gpt4Provider, err := NewLangChainProvider(
		WithLangChainModel("gpt-4"),
	)
	if err != nil {
		log.Fatalf("Failed to create GPT-4 provider: %v", err)
	}

	// OPTION 2: Create core instance first, then add AI with GPT-4
	fmt.Println("\n--- Using core with GPT-4 provider ---")
	coreInstance, err := core.New(rulesPath)
	if err != nil {
		fmt.Println("Error creating CronScribe core:", err)
		return
	}

	// Create AI version using the existing core instance
	gpt4Mapper, err := ai.WithCore(
		coreInstance,
		gpt4Provider,
		ai.WithAIFirst(true), // Try AI first
	)
	if err != nil {
		fmt.Println("Error creating AI with core:", err)
		return
	}

	// Test with a more complex expression
	complexExpr := "run at 3:45 AM on the second Tuesday of each month and on the last day of each quarter"
	fmt.Printf("Complex: %s\n", complexExpr)

	cronExpr, err := gpt4Mapper.ToCron(complexExpr)
	if err != nil {
		fmt.Printf("Error: %s\n\n", err)
	} else {
		fmt.Printf("Cron: %s\n\n", cronExpr)
	}

	// Example of direct use of the provider
	fmt.Println("\n=== Direct use of the LangChain provider ===")
	directInput := "every Monday at 7 AM"
	fmt.Printf("Direct: %s\n", directInput)

	directResult, err := provider.GenerateCron(context.Background(), directInput)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Result: %s\n", directResult)
	}
}
