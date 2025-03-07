package cronscribe

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// LoadRulesFromFile загружает правила из YAML-файла
func LoadRulesFromFile(filePath string) (*Rules, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла правил: %w", err)
	}

	var rules Rules
	if err := yaml.Unmarshal(data, &rules); err != nil {
		return nil, fmt.Errorf("ошибка разбора YAML: %w", err)
	}

	// Компилируем регулярные выражения для всех правил
	for i := range rules.Rules {
		if err := rules.Rules[i].CompilePattern(); err != nil {
			return nil, fmt.Errorf("ошибка компиляции regex для правила %s: %w", rules.Rules[i].Name, err)
		}
	}

	return &rules, nil
}

// LoadAllRules загружает правила для всех языков из директории
func LoadAllRules(directory string) (map[string]*Rules, error) {
	files, err := filepath.Glob(filepath.Join(directory, "*.yaml"))
	if err != nil {
		return nil, fmt.Errorf("ошибка поиска файлов правил: %w", err)
	}

	allRules := make(map[string]*Rules)
	for _, file := range files {
		rules, err := LoadRulesFromFile(file)
		if err != nil {
			return nil, fmt.Errorf("ошибка загрузки правил из %s: %w", file, err)
		}

		allRules[rules.Language] = rules
	}

	return allRules, nil
}
