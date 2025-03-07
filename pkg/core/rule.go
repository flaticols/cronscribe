package core

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Rule represents a rule for converting human-readable expression to cron
type Rule struct {
	Name            string                      `yaml:"name"`
	Pattern         string                      `yaml:"pattern"`
	Variables       map[string]int              `yaml:"variables"`
	Dictionaries    map[string]string           `yaml:"dictionaries"`
	Format          string                      `yaml:"format"`
	DefaultValues   map[string]string           `yaml:"default_values"`
	SpecialCases    []SpecialCase               `yaml:"special_cases"`
	Transformations map[string][]Transformation `yaml:"transformations"`

	compiledPattern *regexp.Regexp
}

// SpecialCase represents a special case for conversion
type SpecialCase struct {
	Condition string `yaml:"condition"`
	Format    string `yaml:"format"`
}

// Transformation represents a variable transformation
type Transformation struct {
	Condition string `yaml:"condition"`
	Operation string `yaml:"operation"`
}

// Rules contains all rules for a language
type Rules struct {
	Language     string                       `yaml:"language"`
	Rules        []Rule                       `yaml:"rules"`
	Dictionaries map[string]map[string]string `yaml:"dictionaries"`
}

// CompilePattern compiles the regular expression for the rule
func (r *Rule) CompilePattern() error {
	var err error
	r.compiledPattern, err = regexp.Compile(r.Pattern)
	return err
}

// Match checks if the expression matches this rule
func (r *Rule) Match(expression string) []string {
	if r.compiledPattern == nil {
		if err := r.CompilePattern(); err != nil {
			return nil
		}
	}
	return r.compiledPattern.FindStringSubmatch(expression)
}

// ApplyTransformations applies transformations to variables
func (r *Rule) ApplyTransformations(variables map[string]string, dictionaries map[string]map[string]string) error {
	for varName, transformations := range r.Transformations {
		value, exists := variables[varName]
		if !exists {
			continue
		}

		for _, t := range transformations {
			// Replace variables in the condition
			condition := t.Condition
			for k, v := range variables {
				condition = strings.ReplaceAll(condition, k, fmt.Sprintf("\"%s\"", v))
			}

			// Evaluate the condition
			// Note: For simplicity, we use basic condition evaluation.
			// In a real application, it's better to use a library for expressions
			if evalCondition(condition) {
				// Replace variables in the operation
				operation := t.Operation
				for k, v := range variables {
					operation = strings.ReplaceAll(operation, k, fmt.Sprintf("\"%s\"", v))
				}

				// Perform the operation
				result, err := evalOperation(operation, value)
				if err != nil {
					return err
				}

				variables[varName] = result
				break
			}
		}
	}

	return nil
}

// evalCondition evaluates a simple condition
// Simplified version for example
func evalCondition(condition string) bool {
	if strings.Contains(condition, "==") {
		parts := strings.Split(condition, "==")
		left := strings.TrimSpace(parts[0])
		right := strings.TrimSpace(parts[1])
		return left == right
	}

	if strings.Contains(condition, "<") {
		parts := strings.Split(condition, "<")
		left, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		right, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		return left < right
	}

	if strings.Contains(condition, ">") {
		parts := strings.Split(condition, ">")
		left, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		right, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		return left > right
	}

	return false
}

// evalOperation evaluates a simple operation
// Simplified version for example
func evalOperation(operation, currentValue string) (string, error) {
	if strings.Contains(operation, "+") {
		parts := strings.Split(operation, "+")
		left, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		right, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		return strconv.Itoa(left + right), nil
	}

	// If the operation is just a string (e.g., 'first')
	if strings.HasPrefix(operation, "'") && strings.HasSuffix(operation, "'") {
		return operation[1 : len(operation)-1], nil
	}

	return operation, nil
}