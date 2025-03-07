package cronscribe

import (
	"fmt"
	"strconv"
	"strings"
)

// TranslateRule преобразует совпадение в cron-выражение по правилу
func TranslateRule(rule *Rule, match []string, dictionaries map[string]map[string]string) (string, error) {
	// Извлекаем переменные из совпадения
	variables := make(map[string]string)
	for name, index := range rule.Variables {
		if index < len(match) {
			variables[name] = match[index]
		}
	}

	// Применяем значения по умолчанию для отсутствующих переменных
	for name, value := range rule.DefaultValues {
		if _, exists := variables[name]; !exists || variables[name] == "" {
			variables[name] = value
		}
	}

	// Преобразуем строковые переменные в числовые, если необходимо
	for name, value := range variables {
		if name == "hour" || name == "minute" || name == "day" {
			if i, err := strconv.Atoi(value); err == nil {
				variables[name] = strconv.Itoa(i)
			}
		}
	}

	// Применяем преобразования для переменных
	if err := rule.ApplyTransformations(variables, dictionaries); err != nil {
		return "", err
	}

	// Проверяем специальные случаи
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

	// Используем стандартный формат
	return applyFormatWithDictionaries(rule.Format, variables, dictionaries, rule.Dictionaries)
}

// applyFormatWithDictionaries применяет формат с заменой переменных и значений словарей
func applyFormatWithDictionaries(format string, variables map[string]string, dictionaries map[string]map[string]string, dictionaryMap map[string]string) (string, error) {
	result := format

	// Заменяем переменные в формате
	for name, value := range variables {
		// Проверяем, нужно ли использовать словарь для этой переменной
		if dictName, ok := dictionaryMap[name]; ok {
			dict, dictExists := dictionaries[dictName]
			if dictExists {
				// Ищем значение в словаре
				if dictValue, valueExists := dict[value]; valueExists {
					result = strings.ReplaceAll(result, "%"+name, dictValue)
				} else {
					return "", fmt.Errorf("значение '%s' не найдено в словаре '%s'", value, dictName)
				}
			} else {
				return "", fmt.Errorf("словарь '%s' не найден", dictName)
			}
		} else {
			// Прямая замена значения
			result = strings.ReplaceAll(result, "%"+name, value)
		}
	}

	return result, nil
}
