# CronScribe

Convert human-readable text into a cron expression. It supports Dutch, English, and Russian by default and can be extended with custom rules in YAML.

## Features

- Convert natural language schedule descriptions to cron expressions
- Support for multiple languages (English, Russian, Dutch)
- Extensible rule-based system with YAML configuration
- Optional AI-powered mode with pluggable AI provider interface

## Installation

```bash
go get github.com/flaticols/cronscribe
```

## Usage

### Basic (recommended) usage with Rule-Based Conversion

```go
package main

import (
    "fmt"

    "github.com/flaticols/cronscribe"
)

func main() {
    // Create a new mapper with rules from the "rules" directory
    mapper, err := cronscribe.NewHumanCronMapper("./rules")
    if err != nil {
        panic(err)
    }

    // Convert a human-readable expression to cron
    cronExpr, err := mapper.ToCron("every day at noon")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Cron expression:", cronExpr) // Output: 0 12 * * *

    // Use a specific language
    err = mapper.SetLanguage("ru")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    cronExpr, err = mapper.ToCron("каждый день в полдень")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Cron expression:", cronExpr) // Output: 0 12 * * *
}
```

### Using AI-Powered "Brave Mode"

This library provides a flexible AIProvider interface that you can implement with any LLM provider.

#### Built-in LangChain Provider

CronScribe includes a built-in LangChain provider for easy integration with various LLMs:

```go
package main

import (
    "fmt"
    "os"
    "log"

    "github.com/flaticols/cronscribe"
    "github.com/tmc/langchaingo/llms/openai"
)

func main() {
    // Make sure to set OPENAI_API_KEY in your environment
    if os.Getenv("OPENAI_API_KEY") == "" {
        log.Fatal("OPENAI_API_KEY environment variable must be set")
    }

    // Create a LangChain provider with default settings (GPT-3.5-turbo)
    provider, err := cronscribe.NewLangChainProvider()
    if err != nil {
        log.Fatalf("Failed to create LangChain provider: %v", err)
    }

    // Optionally, specify a different model
    // provider, err := cronscribe.NewLangChainProvider(
    //     cronscribe.WithLangChainModel(openai.GPT4),
    // )

    // Create a brave mapper with LangChain provider
    mapper, err := cronscribe.NewBraveHumanCronMapper(
        "./rules",
        provider,
        cronscribe.WithAIFirst(true), // Try AI first, fallback to local rules
    )
    if err != nil {
        log.Fatalf("Failed to create mapper: %v", err)
    }

    cronExpr, err := mapper.ToCron("run every Monday at 9:15 AM")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Cron expression:", cronExpr)
}
```

#### Custom AI Provider Implementation

You can also implement your own AI provider:

```go
package main

import (
    "context"
    "fmt"

    "github.com/flaticols/cronscribe"
    "github.com/sashabaranov/go-openai" // For OpenAI implementation example
)

// Implement the AIProvider interface with your preferred AI provider
type MyOpenAIProvider struct {
    client *openai.Client
}

func NewMyOpenAIProvider(apiKey string) *MyOpenAIProvider {
    client := openai.NewClient(apiKey)
    return &MyOpenAIProvider{client: client}
}

// GenerateCron implements the AIProvider interface
func (p *MyOpenAIProvider) GenerateCron(ctx context.Context, input string) (string, error) {
    // Use the recommended prompts from cronscribe
    systemPrompt := cronscribe.RecommendedSystemPrompt()
    userPrompt := cronscribe.RecommendedUserPrompt(input)

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
    openaiProvider := NewMyOpenAIProvider("your-openai-api-key")

    // Create a brave mapper that can use both local rules and AI
    mapper, err := cronscribe.NewBraveHumanCronMapper(
        "./rules",
        openaiProvider,
        cronscribe.WithAIFirst(false), // Try local rules first, fallback to AI
    )
    if err != nil {
        panic(err)
    }

    // Try to convert using local rules first, then AI if needed
    cronExpr, err := mapper.ToCron("run every Monday at 9:15 AM")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Cron expression:", cronExpr)
}
```

## Custom Rules

You can create your own rules by adding YAML files to the rules directory. See the existing files in the `rules/` directory for examples.
