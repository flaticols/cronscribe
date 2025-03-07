package main

import (
	"context"
	"fmt"
	"os"

	"github.com/flaticols/cronscribe"
	"github.com/sashabaranov/go-openai"
)

// OpenAIProvider implements cronscribe.AIProvider using OpenAI API
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
	// Use recommended prompts from cronscribe
	systemPrompt := cronscribe.RecommendedSystemPrompt()
	userPrompt := cronscribe.RecommendedUserPrompt(input)

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

	// Create OpenAI provider
	openaiProvider := NewOpenAIProvider(apiKey)

	// Create a brave mapper that can use both local rules and AI
	mapper, err := cronscribe.NewBraveHumanCronMapper(
		"../../rules", 
		openaiProvider,
		cronscribe.WithAIFirst(false), // Try local rules first, fallback to AI
	)
	if err != nil {
		fmt.Println("Error creating mapper:", err)
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

	for _, expr := range expressions {
		fmt.Printf("Human: %s\n", expr)
		
		cronExpr, err := mapper.ToCron(expr)
		if err != nil {
			fmt.Printf("Error: %s\n\n", err)
			continue
		}
		
		fmt.Printf("Cron: %s\n\n", cronExpr)
	}
}