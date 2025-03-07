package cronscribe

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

// DefaultModel is the default OpenAI model to use
const DefaultModel = "gpt-3.5-turbo"

// LangChainProvider implements AIProvider interface using LangChain
type LangChainProvider struct {
	model string
	llm   llms.Model
}

// LangChainOption represents a functional option for configuring LangChainProvider
type LangChainOption func(*LangChainProvider)

// WithLangChainModel sets the model to use
func WithLangChainModel(model string) LangChainOption {
	return func(p *LangChainProvider) {
		p.model = model
	}
}

// WithCustomLLM sets a custom LLM implementation
func WithCustomLLM(llm llms.Model) LangChainOption {
	return func(p *LangChainProvider) {
		p.llm = llm
	}
}

// NewLangChainProvider creates a new LangChainProvider
func NewLangChainProvider(options ...LangChainOption) (*LangChainProvider, error) {
	// Default provider with OpenAI's GPT-3.5-turbo
	provider := &LangChainProvider{
		model: DefaultModel,
	}

	// Apply all options
	for _, option := range options {
		option(provider)
	}

	// Initialize LLM if not provided
	if provider.llm == nil {
		llm, err := openai.New(openai.WithModel(provider.model))
		if err != nil {
			return nil, fmt.Errorf("failed to initialize OpenAI LLM: %w", err)
		}
		provider.llm = llm
	}

	return provider, nil
}

// GenerateCron generates a cron expression from a human-readable string
func (p *LangChainProvider) GenerateCron(ctx context.Context, input string) (string, error) {
	// Prepare the prompt with system and user messages
	systemMessage := llms.TextParts(llms.ChatMessageTypeSystem, RecommendedSystemPrompt())
	userMessage := llms.TextParts(llms.ChatMessageTypeHuman, RecommendedUserPrompt(input))
	
	messages := []llms.MessageContent{
		systemMessage,
		userMessage,
	}

	// Set up the options
	opts := []llms.CallOption{
		llms.WithModel(p.model),
		llms.WithTemperature(0.1), // Low temperature for deterministic output
	}

	// Generate response
	response, err := p.llm.GenerateContent(ctx, messages, opts...)
	if err != nil {
		return "", fmt.Errorf("LangChain LLM call failed: %w", err)
	}

	// Extract content from response using reflection since the ContentResponse type 
	// might change in future versions
	content, err := extractContentFromResponse(response)
	if err != nil {
		return "", err
	}

	// Process the response
	cronExpr := strings.TrimSpace(content)
	if !isValidCronExpression(cronExpr) {
		return "", errors.New("invalid cron expression returned by LLM")
	}

	return cronExpr, nil
}

// extractContentFromResponse uses reflection to safely extract content from 
// a ContentResponse which may vary between LangChain versions
func extractContentFromResponse(response interface{}) (string, error) {
	// Try to access Choices field using reflection
	responseVal := reflect.ValueOf(response)
	if responseVal.Kind() == reflect.Ptr {
		responseVal = responseVal.Elem()
	}

	// Find the Choices field
	choicesField := responseVal.FieldByName("Choices")
	if !choicesField.IsValid() {
		return "", errors.New("no Choices field found in response")
	}

	// Check if there are any choices
	if choicesField.Len() == 0 {
		return "", errors.New("no content choices returned from LLM")
	}

	// Get the first choice
	firstChoice := choicesField.Index(0)
	
	// Get the Content field from the choice
	contentField := firstChoice.FieldByName("Content")
	if !contentField.IsValid() {
		return "", errors.New("no Content field found in choice")
	}

	// Convert to string and return
	return contentField.String(), nil
}