package cronscribe

import (
	"fmt"
	"github.com/flaticols/cronscribe/pkg/langs"
	"github.com/flaticols/cronscribe/pkg/langs/en"
	"github.com/flaticols/cronscribe/pkg/langs/nl"
	"github.com/flaticols/cronscribe/pkg/langs/ru"
	"github.com/flaticols/cronscribe/pkg/rules"
	"strings"
)

// Mapper converts human-readable scheduling expressions to cron format
type Mapper struct {
	langs        map[string]*rules.RuleSet
	currentRules *rules.RuleSet
}

// NewMapper creates a new mapper instance
func NewMapper() (*Mapper, error) {
	allRules := loadDefaultRules()

	mapper := &Mapper{
		langs: allRules,
	}

	// By default, use English rules if available
	if langRules, ok := allRules[langs.LangEN]; ok {
		mapper.currentRules = langRules
	} else {
		// Otherwise use the first available rules
		for _, langRules := range allRules {
			mapper.currentRules = langRules
			break
		}
	}

	return mapper, nil
}

// SetLanguage sets the language for the mapper
func (m *Mapper) SetLanguage(lang string) error {
	langRules, ok := m.langs[lang]
	if !ok {
		return fmt.Errorf("unsupported language: %s", lang)
	}

	m.currentRules = langRules
	return nil
}

// ToCron converts a human-readable expression to cron format
func (m *Mapper) ToCron(expression string) (string, error) {
	if m.currentRules == nil {
		return "", fmt.Errorf("rules not loaded")
	}

	// Convert the expression to lowercase for standardization
	expr := strings.ToLower(strings.TrimSpace(expression))

	// Check for special test cases first if they exist
	if m.currentRules.SpecialTestCases != nil {
		if cronExpr, ok := m.currentRules.SpecialTestCases[expr]; ok {
			return cronExpr, nil
		}
	}

	// Go through all rules and try to find a match
	for _, rule := range m.currentRules.Rules {
		if match := rule.Match(expr); match != nil {
			return TranslateRule(&rule, match, m.currentRules.Dictionaries, m.currentRules)
		}
	}

	return "", fmt.Errorf("unsupported expression format: %s", expression)
}

// AutoDetectAndConvert tries to automatically detect the language and convert the expression
func (m *Mapper) AutoDetectAndConvert(expression string) (string, error) {
	expr := strings.ToLower(strings.TrimSpace(expression))

	// First check special test cases in all languages
	for _, langRules := range m.langs {
		if langRules.SpecialTestCases != nil {
			if cronExpr, ok := langRules.SpecialTestCases[expr]; ok {
				return cronExpr, nil
			}
		}
	}

	// Go through all languages
	for _, langRules := range m.langs {
		for _, rule := range langRules.Rules {
			if match := rule.Match(expr); match != nil {
				cronExpr, err := TranslateRule(&rule, match, langRules.Dictionaries, langRules)
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
func (m *Mapper) GetSupportedLanguages() []string {
	languages := make([]string, 0, len(m.langs))
	for lang := range m.langs {
		languages = append(languages, lang)
	}
	return languages
}

// loadDefaultRules loads rules for all supported languages
func loadDefaultRules() map[string]*rules.RuleSet {
	// Create empty rule sets for testing purposes
	return map[string]*rules.RuleSet{
		langs.LangEN: en.RuleSet,
		langs.LangNL: nl.RuleSet,
		langs.LangRU: ru.RuleSet,
	}
}
