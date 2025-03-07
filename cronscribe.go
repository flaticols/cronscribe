package cronscribe

import (
	"github.com/flaticols/cronscribe/pkg/ai"
	"github.com/flaticols/cronscribe/pkg/core"
)

// CronScribe provides the main functionality for cron expression conversion
type CronScribe struct {
	*core.CronScribe
}

// New creates a new instance of CronScribe with the default rule-based mapper
func New(rulesDir string) (*CronScribe, error) {
	c, err := core.New(rulesDir)
	if err != nil {
		return nil, err
	}
	return &CronScribe{CronScribe: c}, nil
}

// NewWithAI creates a new instance of CronScribe with AI-powered features
// This is a convenience function that creates a new CronScribeAI instance
func NewWithAI(rulesDir string, provider ai.AIProvider, opts ...ai.BraveOption) (*ai.CronScribeAI, error) {
	return ai.New(rulesDir, provider, opts...)
}

// WithCore extends an existing CronScribe instance with AI capabilities
// This is a convenience function that wraps core.CronScribe with AI features
func WithCore(cs *CronScribe, provider ai.AIProvider, opts ...ai.BraveOption) (*ai.CronScribeAI, error) {
	return ai.WithCore(cs.CronScribe, provider, opts...)
}

// All core methods are automatically available:
// - Convert(expression string) (string, error)
// - AutoDetect(expression string) (string, error)
// - SetLanguage(lang string) error
// - GetSupportedLanguages() []string
// - AddRulesFromFile(filePath string) error