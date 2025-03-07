package core

// Version is the current version of the CronScribe core package
const Version = "1.0.0"

// CronScribe is the main entry point for using the core functionality
type CronScribe struct {
	mapper *HumanCronMapper
}

// New creates a new CronScribe instance
func New(rulesDir string) (*CronScribe, error) {
	mapper, err := NewHumanCronMapper(rulesDir)
	if err != nil {
		return nil, err
	}

	return &CronScribe{
		mapper: mapper,
	}, nil
}

// Convert transforms a human-readable scheduling expression to a cron expression
func (c *CronScribe) Convert(expression string) (string, error) {
	return c.mapper.ToCron(expression)
}

// AutoDetect tries to automatically detect the language and convert the expression
func (c *CronScribe) AutoDetect(expression string) (string, error) {
	return c.mapper.AutoDetectAndConvert(expression)
}

// SetLanguage sets the language for processing expressions
func (c *CronScribe) SetLanguage(lang string) error {
	return c.mapper.SetLanguage(lang)
}

// GetSupportedLanguages returns a list of supported languages
func (c *CronScribe) GetSupportedLanguages() []string {
	return c.mapper.GetSupportedLanguages()
}

// AddRulesFromFile adds rules from a file
func (c *CronScribe) AddRulesFromFile(filePath string) error {
	return c.mapper.AddRulesFromFile(filePath)
}