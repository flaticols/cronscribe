package core

import (
	"fmt"
	"strings"
)

// HumanCronMapper converts human-readable scheduling expressions to cron format
type HumanCronMapper struct {
	allRules     map[string]*Rules
	currentRules *Rules
}

// NewHumanCronMapper creates a new mapper instance
func NewHumanCronMapper(rulesDir string) (*HumanCronMapper, error) {
	allRules, err := LoadAllRules(rulesDir)
	if err != nil {
		return nil, err
	}

	mapper := &HumanCronMapper{
		allRules: allRules,
	}

	// By default, use English rules if available
	if rules, ok := allRules["en"]; ok {
		mapper.currentRules = rules
	} else {
		// Otherwise use the first available rules
		for _, rules := range allRules {
			mapper.currentRules = rules
			break
		}
	}

	return mapper, nil
}

// SetLanguage sets the language for the mapper
func (m *HumanCronMapper) SetLanguage(lang string) error {
	rules, ok := m.allRules[lang]
	if !ok {
		return fmt.Errorf("unsupported language: %s", lang)
	}

	m.currentRules = rules
	return nil
}

// ToCron converts a human-readable expression to cron format
func (m *HumanCronMapper) ToCron(expression string) (string, error) {
	if m.currentRules == nil {
		return "", fmt.Errorf("rules not loaded")
	}

	// Convert the expression to lowercase for standardization
	expr := strings.ToLower(strings.TrimSpace(expression))

	// Go through all rules and try to find a match
	for _, rule := range m.currentRules.Rules {
		if match := rule.Match(expr); match != nil {
			return TranslateRule(&rule, match, m.currentRules.Dictionaries)
		}
	}

	return "", fmt.Errorf("unsupported expression format: %s", expression)
}

// AutoDetectAndConvert tries to automatically detect the language and convert the expression
func (m *HumanCronMapper) AutoDetectAndConvert(expression string) (string, error) {
	expr := strings.ToLower(strings.TrimSpace(expression))

	// Go through all languages
	for _, rules := range m.allRules {
		for _, rule := range rules.Rules {
			if match := rule.Match(expr); match != nil {
				cronExpr, err := TranslateRule(&rule, match, rules.Dictionaries)
				if err != nil {
					continue
				}
				return cronExpr, nil
			}
		}
	}

	return "", fmt.Errorf("unsupported expression format: %s", expression)
}

// GetSupportedLanguages returns a list of supported languages
func (m *HumanCronMapper) GetSupportedLanguages() []string {
	languages := make([]string, 0, len(m.allRules))
	for lang := range m.allRules {
		languages = append(languages, lang)
	}
	return languages
}

// AddRulesFromFile adds rules from a file
func (m *HumanCronMapper) AddRulesFromFile(filePath string) error {
	rules, err := LoadRulesFromFile(filePath)
	if err != nil {
		return err
	}

	m.allRules[rules.Language] = rules
	return nil
}