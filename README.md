# cronscribe

> [!WARNING]
> Library in active development

Convert human-readable text into a cron expression. It supports Dutch, English, and Russian by default and can be extended with custom rules in YAML.

## Features

- Convert natural language schedule descriptions to cron expressions
- Support for multiple languages (English, Russian, Dutch)
- Extensible rule-based system with YAML configuration
- Optional AI-powered mode with pluggable AI provider interface
- Modular design: use only what you need

## Installation

### Core Package Only (No AI Dependencies)

```bash
go get github.com/flaticols/cronscribe
```

### With AI Support

```bash
go get github.com/flaticols/cronscribe/pkg/ai
```

## Usage

### Basic Usage with Rule-Based Conversion

```go
package main

import (
    "fmt"

    "github.com/flaticols/cronscribe"
)

func main() {
    // Create a new CronScribe instance with rules from the "pkg/core/rules" directory
    cs, err := cronscribe.New("./pkg/core/rules")
    if err != nil {
        panic(err)
    }

    // Convert a human-readable expression to cron
    cronExpr, err := cs.Convert("every day at noon")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Cron expression:", cronExpr) // Output: 0 12 * * *

    // Use a specific language
    err = cs.SetLanguage("ru")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    cronExpr, err = cs.Convert("каждый день в полдень")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Cron expression:", cronExpr) // Output: 0 12 * * *
}
```

### Using AI-Powered "Brave Mode"

For AI-powered functionality, import the AI package:

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/flaticols/cronscribe/pkg/ai"
    "github.com/sashabaranov/go-openai" // For OpenAI implementation example
)

// Implement the AIProvider interface with your preferred AI provider
type OpenAIProvider struct {
    client *openai.Client
}

func NewOpenAIProvider(apiKey string) *OpenAIProvider {
    client := openai.NewClient(apiKey)
    return &OpenAIProvider{client: client}
}

// GenerateCron implements the AIProvider interface
func (p *OpenAIProvider) GenerateCron(ctx context.Context, input string) (string, error) {
    // Use the recommended prompts from the ai package
    systemPrompt := ai.RecommendedSystemPrompt()
    userPrompt := ai.RecommendedUserPrompt(input)

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
    // Get API key from environment variable
    apiKey := os.Getenv("OPENAI_API_KEY")
    if apiKey == "" {
        fmt.Println("Please set the OPENAI_API_KEY environment variable")
        return
    }

    // Create OpenAI provider
    openaiProvider := NewOpenAIProvider(apiKey)

    // Create a CronScribeAI instance with the OpenAI provider
    cronscribeAI, err := ai.New(
        "./pkg/core/rules", 
        openaiProvider,
        ai.WithAIFirst(false), // Try local rules first, fallback to AI
    )
    if err != nil {
        fmt.Println("Error creating CronScribeAI:", err)
        return
    }

    // Convert using local rules first, then AI if needed
    cronExpr, err := cronscribeAI.ToCron("run every Monday at 9:15 AM")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Cron expression:", cronExpr)
}
```

## Custom Rules

You can create your own rules by adding YAML files to the rules directory. See the existing files in the `pkg/core/rules/` directory for examples.

## Module Structure

```
cronscribe/
├── pkg/
│   ├── core/                # Core package - YAML rule based conversion
│   │   ├── mapper.go        # Rule-based mapping implementation
│   │   ├── rule.go          # Rule definitions and logic
│   │   ├── loader.go        # YAML rules loader
│   │   ├── translator.go    # Expression translator
│   │   ├── cronscribe.go    # Core package entrypoint
│   │   └── rules/           # YAML rule definitions
│   │       ├── en.yaml      # English rules
│   │       ├── ru.yaml      # Russian rules
│   │       └── nl.yaml      # Dutch rules
│   │
│   └── ai/                  # AI package - AI-powered conversion
│       ├── ai_provider.go   # AI provider interface
│       ├── brave_mapper.go  # AI-powered mapper implementation
│       ├── cronscribe_ai.go # AI package entrypoint
│       └── rules/           # Copy of core rules
│
├── cronscribe.go            # Main package entrypoint (wrapper)
│
└── examples/
    ├── core_only/           # Example using only core features
    ├── brave_mode/          # Example using AI features with OpenAI
    ├── langchain_mode/      # Example using AI features with LangChain
    └── wrapper_mode/        # Example using the main package wrapper
```

This modular design allows you to only import what you need, keeping dependencies minimal when you only need the core functionality.

## Dependencies

- Core package:
  - `gopkg.in/yaml.v3`: YAML parsing for rule files

- AI package:
  - Core package dependencies
  - AI provider specific dependencies (based on your implementation)

## License

See the LICENSE file for details.
