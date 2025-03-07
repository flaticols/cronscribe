# CronScribe AI Package

This package extends the core CronScribe functionality with AI-powered features that allow for more flexible parsing of human-readable schedule descriptions.

## Features

- AI provider interface for integrating any LLM service
- Fallback mechanism between rule-based and AI-based conversion
- Customizable options for prioritizing AI or rules
- Example implementations for common AI providers

## Installation

```bash
go get github.com/flaticols/cronscribe/pkg/ai
```

## Dependencies

- `github.com/flaticols/cronscribe/pkg/core`: Core functionality
- Any AI provider dependencies based on your implementation

## Usage

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/flaticols/cronscribe/pkg/ai"
    "github.com/sashabaranov/go-openai" // Example using OpenAI
)

// Implement the AIProvider interface
type OpenAIProvider struct {
    client *openai.Client
}

func NewOpenAIProvider(apiKey string) *OpenAIProvider {
    client := openai.NewClient(apiKey)
    return &OpenAIProvider{client: client}
}

func (p *OpenAIProvider) GenerateCron(ctx context.Context, input string) (string, error) {
    // Use recommended prompts from the ai package
    systemPrompt := ai.RecommendedSystemPrompt()
    userPrompt := ai.RecommendedUserPrompt(input)
    
    // Call the OpenAI API
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
    
    return resp.Choices[0].Message.Content, nil
}

func main() {
    // Create your AI provider
    provider := NewOpenAIProvider(os.Getenv("OPENAI_API_KEY"))
    
    // Create a CronScribeAI instance
    cronscribeAI, err := ai.New("./rules", provider,
        ai.WithAIFirst(false), // Try rules first, then AI
    )
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    // Convert a human-readable expression
    cronExpr, err := cronscribeAI.ToCron("run every 3 hours starting at 9am")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("Cron expression: %s\n", cronExpr)
}
```

## Configuration Options

- `WithAIFirst(bool)`: Set to true to try AI conversion before rule-based conversion
- `WithAIProvider(provider)`: Set a custom AI provider implementation