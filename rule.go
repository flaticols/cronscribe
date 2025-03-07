package cronscribe

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Rule представляет правило для преобразования человекочитаемого выражения в cron
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

// SpecialCase представляет особый случай для преобразования
type SpecialCase struct {
	Condition string `yaml:"condition"`
	Format    string `yaml:"format"`
}

// Transformation представляет трансформацию переменной
type Transformation struct {
	Condition string `yaml:"condition"`
	Operation string `yaml:"operation"`
}

// Rules содержит все правила для языка
type Rules struct {
	Language     string                       `yaml:"language"`
	Rules        []Rule                       `yaml:"rules"`
	Dictionaries map[string]map[string]string `yaml:"dictionaries"`
}

// CompilePattern компилирует регулярное выражение для правила
func (r *Rule) CompilePattern() error {
	var err error
	r.compiledPattern, err = regexp.Compile(r.Pattern)
	return err
}

// Match проверяет, соответствует ли выражение этому правилу
func (r *Rule) Match(expression string) []string {
	if r.compiledPattern == nil {
		if err := r.CompilePattern(); err != nil {
			return nil
		}
	}
	return r.compiledPattern.FindStringSubmatch(expression)
}

// ApplyTransformations применяет трансформации к переменным
func (r *Rule) ApplyTransformations(variables map[string]string, dictionaries map[string]map[string]string) error {
	for varName, transformations := range r.Transformations {
		value, exists := variables[varName]
		if !exists {
			continue
		}

		for _, t := range transformations {
			// Заменяем переменные в условии
			condition := t.Condition
			for k, v := range variables {
				condition = strings.ReplaceAll(condition, k, fmt.Sprintf("\"%s\"", v))
			}

			// Вычисляем условие
			// Примечание: Для простоты используем базовую оценку условий.
			// В реальном приложении лучше использовать библиотеку для выражений
			if evalCondition(condition) {
				// Заменяем переменные в операции
				operation := t.Operation
				for k, v := range variables {
					operation = strings.ReplaceAll(operation, k, fmt.Sprintf("\"%s\"", v))
				}

				// Выполняем операцию
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

// evalCondition оценивает простое условие
// Упрощенная версия для примера
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

// evalOperation оценивает простую операцию
// Упрощенная версия для примера
func evalOperation(operation, currentValue string) (string, error) {
	if strings.Contains(operation, "+") {
		parts := strings.Split(operation, "+")
		left, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		right, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		return strconv.Itoa(left + right), nil
	}

	// Если операция это просто строка (например, 'первый')
	if strings.HasPrefix(operation, "'") && strings.HasSuffix(operation, "'") {
		return operation[1 : len(operation)-1], nil
	}

	return operation, nil
}
