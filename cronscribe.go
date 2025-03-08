package cronscribe

import (
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
