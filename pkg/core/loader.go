package core

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// LoadRulesFromFile loads rules from a YAML file
func LoadRulesFromFile(filePath string) (*Rules, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading rules file: %w", err)
	}

	var rules Rules
	if err := yaml.Unmarshal(data, &rules); err != nil {
		return nil, fmt.Errorf("error parsing YAML: %w", err)
	}

	// Compile regular expressions for all rules
	for i := range rules.Rules {
		if err := rules.Rules[i].CompilePattern(); err != nil {
			return nil, fmt.Errorf("error compiling regex for rule %s: %w", rules.Rules[i].Name, err)
		}
	}

	return &rules, nil
}

// LoadAllRules loads rules for all languages from a directory
func LoadAllRules(directory string) (map[string]*Rules, error) {
	files, err := filepath.Glob(filepath.Join(directory, "*.yaml"))
	if err != nil {
		return nil, fmt.Errorf("error finding rule files: %w", err)
	}

	allRules := make(map[string]*Rules)
	for _, file := range files {
		rules, err := LoadRulesFromFile(file)
		if err != nil {
			return nil, fmt.Errorf("error loading rules from %s: %w", file, err)
		}

		allRules[rules.Language] = rules
	}

	return allRules, nil
}