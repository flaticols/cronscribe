package ai

import (
	"context"
	"fmt"
	"strings"

	"github.com/flaticols/cronscribe/pkg/core"
)

// BraveOption represents a functional option for configuring BraveHumanCronMapper
type BraveOption func(*BraveHumanCronMapper)

// WithAIFirst configures whether to try AI first before local rules
func WithAIFirst(useAIFirst bool) BraveOption {
	return func(m *BraveHumanCronMapper) {
		m.useAIFirst = useAIFirst
	}
}

// WithAIProvider sets a custom AI provider implementation
func WithAIProvider(provider AIProvider) BraveOption {
	return func(m *BraveHumanCronMapper) {
		m.aiProvider = provider
	}
}

// BraveHumanCronMapper extends core.CronScribe with AI API capabilities
type BraveHumanCronMapper struct {
	coreMapper *core.CronScribe
	aiProvider AIProvider
	useAIFirst bool
}

// NewBraveHumanCronMapper creates a new brave mapper that can use both local rules and AI
func NewBraveHumanCronMapper(rulesDir string, provider AIProvider, options ...BraveOption) (*BraveHumanCronMapper, error) {
	// Create the base mapper using local rules
	coreMapper, err := core.New(rulesDir)
	if err != nil {
		return nil, fmt.Errorf("failed to create base mapper: %w", err)
	}

	// Create the mapper with provided AI provider
	if provider == nil {
		return nil, fmt.Errorf("AI provider cannot be nil")
	}

	mapper := &BraveHumanCronMapper{
		coreMapper: coreMapper,
		aiProvider: provider,
		useAIFirst: false, // Default to using local rules first
	}

	// Apply all options
	for _, option := range options {
		option(mapper)
	}

	return mapper, nil
}

// ToCron converts a human-readable expression to a cron expression
// In brave mode, it can use AI if local rules fail or if useAIFirst is true
func (m *BraveHumanCronMapper) ToCron(expression string) (string, error) {
	if m.useAIFirst {
		// Try AI first
		ctx := context.Background()
		cronExpr, err := m.aiProvider.GenerateCron(ctx, expression)
		if err == nil && isValidCronExpression(cronExpr) {
			return cronExpr, nil
		}
		// If AI fails, fall back to local rules
	}

	// Try local rules
	cronExpr, err := m.coreMapper.Convert(expression)
	if err == nil {
		return cronExpr, nil
	}

	// If local rules fail and we didn't try AI yet, use AI as fallback
	if !m.useAIFirst {
		ctx := context.Background()
		cronExpr, err := m.aiProvider.GenerateCron(ctx, expression)
		if err == nil && isValidCronExpression(cronExpr) {
			return cronExpr, nil
		}
		return "", fmt.Errorf("unable to convert expression with local rules or AI: %s", expression)
	}

	return "", err
}

// SetLanguage sets the language for the underlying mapper
func (m *BraveHumanCronMapper) SetLanguage(lang string) error {
	return m.coreMapper.SetLanguage(lang)
}

// GetSupportedLanguages returns the list of supported languages
func (m *BraveHumanCronMapper) GetSupportedLanguages() []string {
	return m.coreMapper.GetSupportedLanguages()
}

// AutoDetect tries to automatically detect the language and convert the expression
func (m *BraveHumanCronMapper) AutoDetect(expression string) (string, error) {
	// Try with local rules first
	cronExpr, err := m.coreMapper.AutoDetect(expression)
	if err == nil {
		return cronExpr, nil
	}

	// Fall back to AI
	ctx := context.Background()
	cronExpr, err = m.aiProvider.GenerateCron(ctx, expression)
	if err == nil && isValidCronExpression(cronExpr) {
		return cronExpr, nil
	}

	return "", fmt.Errorf("unable to convert expression with local rules or AI: %s", expression)
}

// isValidCronExpression performs basic validation on a cron expression
func isValidCronExpression(expr string) bool {
	// Trim the expression and remove any quotation marks that might be in the AI response
	expr = strings.TrimSpace(expr)
	expr = strings.Trim(expr, `"'`)

	// Basic validation: check if it has 5 fields separated by whitespace
	fields := strings.Fields(expr)
	return len(fields) == 5
}
