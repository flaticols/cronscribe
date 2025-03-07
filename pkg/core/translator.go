package core

import (
	"fmt"
	"strconv"
	"strings"
)

// TranslateRule converts a match to a cron expression according to the rule
func TranslateRule(rule *Rule, match []string, dictionaries map[string]map[string]string) (string, error) {
	// Extract variables from the match
	variables := make(map[string]string)
	for name, index := range rule.Variables {
		if index < len(match) {
			variables[name] = match[index]
		}
	}

	// Apply default values for missing variables
	for name, value := range rule.DefaultValues {
		if _, exists := variables[name]; !exists || variables[name] == "" {
			variables[name] = value
		}
	}

	// Convert string variables to numeric if needed
	for name, value := range variables {
		if name == "hour" || name == "minute" || name == "day" {
			if i, err := strconv.Atoi(value); err == nil {
				variables[name] = strconv.Itoa(i)
			}
		}
	}

	// Apply transformations to variables
	if err := rule.ApplyTransformations(variables, dictionaries); err != nil {
		return "", err
	}

	// Check special cases
	for _, specialCase := range rule.SpecialCases {
		condition := specialCase.Condition
		for k, v := range variables {
			condition = strings.ReplaceAll(condition, k, fmt.Sprintf("\"%s\"", v))
		}

		if evalCondition(condition) {
			format := specialCase.Format
			return applyFormatWithDictionaries(format, variables, dictionaries, rule.Dictionaries)
		}
	}

	// Use standard format
	return applyFormatWithDictionaries(rule.Format, variables, dictionaries, rule.Dictionaries)
}

// applyFormatWithDictionaries applies format with variable and dictionary value substitution
func applyFormatWithDictionaries(format string, variables map[string]string, dictionaries map[string]map[string]string, dictionaryMap map[string]string) (string, error) {
	result := format

	// Replace variables in the format
	for name, value := range variables {
		// Check if we need to use a dictionary for this variable
		if dictName, ok := dictionaryMap[name]; ok {
			dict, dictExists := dictionaries[dictName]
			if dictExists {
				// Look up the value in the dictionary
				if dictValue, valueExists := dict[value]; valueExists {
					result = strings.ReplaceAll(result, "%"+name, dictValue)
				} else {
					return "", fmt.Errorf("value '%s' not found in dictionary '%s'", value, dictName)
				}
			} else {
				return "", fmt.Errorf("dictionary '%s' not found", dictName)
			}
		} else {
			// Direct value replacement
			result = strings.ReplaceAll(result, "%"+name, value)
		}
	}

	return result, nil
}