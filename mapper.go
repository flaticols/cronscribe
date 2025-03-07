package cronscribe

import (
	"fmt"
	"strings"
)

// HumanCronMapper преобразует человекочитаемые выражения планирования в cron-формат
type HumanCronMapper struct {
	allRules     map[string]*Rules
	currentRules *Rules
}

// NewHumanCronMapper создает новый экземпляр маппера
func NewHumanCronMapper(rulesDir string) (*HumanCronMapper, error) {
	allRules, err := LoadAllRules(rulesDir)
	if err != nil {
		return nil, err
	}

	mapper := &HumanCronMapper{
		allRules: allRules,
	}

	// По умолчанию используем английские правила, если они есть
	if rules, ok := allRules["en"]; ok {
		mapper.currentRules = rules
	} else {
		// Иначе используем первые доступные правила
		for _, rules := range allRules {
			mapper.currentRules = rules
			break
		}
	}

	return mapper, nil
}

// SetLanguage устанавливает язык для маппера
func (m *HumanCronMapper) SetLanguage(lang string) error {
	rules, ok := m.allRules[lang]
	if !ok {
		return fmt.Errorf("неподдерживаемый язык: %s", lang)
	}

	m.currentRules = rules
	return nil
}

// ToCron преобразует человекочитаемое выражение в cron-формат
func (m *HumanCronMapper) ToCron(expression string) (string, error) {
	if m.currentRules == nil {
		return "", fmt.Errorf("правила не загружены")
	}

	// Приводим выражение к нижнему регистру для унификации
	expr := strings.ToLower(strings.TrimSpace(expression))

	// Проходим по всем правилам и пытаемся найти соответствие
	for _, rule := range m.currentRules.Rules {
		if match := rule.Match(expr); match != nil {
			return TranslateRule(&rule, match, m.currentRules.Dictionaries)
		}
	}

	return "", fmt.Errorf("неподдерживаемый формат выражения: %s", expression)
}

// AutoDetectAndConvert пытается автоматически определить язык и преобразовать выражение
func (m *HumanCronMapper) AutoDetectAndConvert(expression string) (string, error) {
	expr := strings.ToLower(strings.TrimSpace(expression))

	// Проходим по всем языкам
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

	return "", fmt.Errorf("неподдерживаемый формат выражения: %s", expression)
}

// GetSupportedLanguages возвращает список поддерживаемых языков
func (m *HumanCronMapper) GetSupportedLanguages() []string {
	languages := make([]string, 0, len(m.allRules))
	for lang := range m.allRules {
		languages = append(languages, lang)
	}
	return languages
}

// AddRulesFromFile добавляет правила из файла
func (m *HumanCronMapper) AddRulesFromFile(filePath string) error {
	rules, err := LoadRulesFromFile(filePath)
	if err != nil {
		return err
	}

	m.allRules[rules.Language] = rules
	return nil
}
