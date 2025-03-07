package ai

import (
	"github.com/flaticols/cronscribe/pkg/core"
)

// CronScribeAI extends core functionality with AI capabilities
type CronScribeAI struct {
	*BraveHumanCronMapper
}

// New creates a new CronScribeAI instance
func New(rulesDir string, provider AIProvider, options ...BraveOption) (*CronScribeAI, error) {
	braveCronMapper, err := NewBraveHumanCronMapper(rulesDir, provider, options...)
	if err != nil {
		return nil, err
	}

	return &CronScribeAI{
		BraveHumanCronMapper: braveCronMapper,
	}, nil
}

// ToCron converts a human-readable expression to a cron expression
// using AI if local rules fail or if UseAIFirst option was provided
func (c *CronScribeAI) ToCron(expression string) (string, error) {
	return c.BraveHumanCronMapper.ToCron(expression)
}

// WithCore creates a new CronScribeAI instance using an existing core instance
func WithCore(coreInstance *core.CronScribe, provider AIProvider, options ...BraveOption) (*CronScribeAI, error) {
	if provider == nil {
		return nil, nil
	}

	mapper := &BraveHumanCronMapper{
		coreMapper: coreInstance,
		aiProvider: provider,
		useAIFirst: false, // Default to using local rules first
	}

	// Apply all options
	for _, option := range options {
		option(mapper)
	}

	return &CronScribeAI{
		BraveHumanCronMapper: mapper,
	}, nil
}
