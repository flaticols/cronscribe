package cronscribe

import (
	"testing"
)

// TestEnglishBasicExpressions tests basic time expressions
func TestEnglishBasicExpressions(t *testing.T) {
	configDir := "../config"

	m, err := NewHumanCronMapper(configDir)
	if err != nil {
		t.Fatalf("Failed to create mapper: %v", err)
	}

	// Set English language
	if err := m.SetLanguage("en"); err != nil {
		t.Fatalf("Failed to set language: %v", err)
	}

	tests := []struct {
		expression string
		expected   string
	}{
		// Time-based expressions
		{"every hour", "0 * * * *"},
		{"every day at 12 pm", "0 12 * * *"},
		{"every day at midnight", "0 0 * * *"},
		{"every day at noon", "0 12 * * *"},
		{"every 5 minutes", "*/5 * * * *"},
		{"every 2 hours", "0 */2 * * *"},
	}

	for _, test := range tests {
		result, err := m.ToCron(test.expression)
		if err != nil {
			t.Errorf("Error parsing '%s': %v", test.expression, err)
			continue
		}

		if result != test.expected {
			t.Errorf("For expression '%s' expected result '%s', got '%s'",
				test.expression, test.expected, result)
		}
	}
}

// TestEnglishDayTimeExpressions tests expressions with AM/PM time specifications
func TestEnglishDayTimeExpressions(t *testing.T) {
	configDir := "../config"

	m, err := NewHumanCronMapper(configDir)
	if err != nil {
		t.Fatalf("Failed to create mapper: %v", err)
	}

	// Set English language
	if err := m.SetLanguage("en"); err != nil {
		t.Fatalf("Failed to set language: %v", err)
	}

	tests := []struct {
		expression string
		expected   string
	}{
		{"every day at 3 am", "0 3 * * *"},
		{"every day at 3:30 am", "30 3 * * *"},
		{"every day at 3 pm", "0 15 * * *"},
		{"every day at 3:45 pm", "45 15 * * *"},
		{"every day at 12 am", "0 0 * * *"},  // Midnight
		{"every day at 12 pm", "0 12 * * *"}, // Noon
	}

	for _, test := range tests {
		result, err := m.ToCron(test.expression)
		if err != nil {
			t.Errorf("Error parsing '%s': %v", test.expression, err)
			continue
		}

		if result != test.expected {
			t.Errorf("For expression '%s' expected result '%s', got '%s'",
				test.expression, test.expected, result)
		}
	}
}

// TestEnglishWeekdayExpressions tests weekday expressions
func TestEnglishWeekdayExpressions(t *testing.T) {
	configDir := "../config"

	m, err := NewHumanCronMapper(configDir)
	if err != nil {
		t.Fatalf("Failed to create mapper: %v", err)
	}

	// Set English language
	if err := m.SetLanguage("en"); err != nil {
		t.Fatalf("Failed to set language: %v", err)
	}

	tests := []struct {
		expression string
		expected   string
	}{
		// Weekday expressions
		{"every Monday", "0 0 * * 1"},
		{"every Tuesday", "0 0 * * 2"},
		{"every Wednesday", "0 0 * * 3"},
		{"every Thursday", "0 0 * * 4"},
		{"every Friday", "0 0 * * 5"},
		{"every Saturday", "0 0 * * 6"},
		{"every Sunday", "0 0 * * 0"},

		// Weekdays with time
		{"every Monday at 9 am", "0 9 * * 1"},
		{"every Friday at 5:30 pm", "30 17 * * 5"},

		// Weekday ranges and groups
		{"every weekday", "0 0 * * 1-5"},
		{"every weekday at 9 am", "0 9 * * 1-5"},
		{"every weekend", "0 0 * * 0,6"},
		{"every weekend at 10 am", "0 10 * * 0,6"},
	}

	for _, test := range tests {
		result, err := m.ToCron(test.expression)
		if err != nil {
			t.Errorf("Error parsing '%s': %v", test.expression, err)
			continue
		}

		if result != test.expected {
			t.Errorf("For expression '%s' expected result '%s', got '%s'",
				test.expression, test.expected, result)
		}
	}
}

// TestEnglishMonthlyExpressions tests monthly recurring expressions
func TestEnglishMonthlyExpressions(t *testing.T) {
	configDir := "../config"

	m, err := NewHumanCronMapper(configDir)
	if err != nil {
		t.Fatalf("Failed to create mapper: %v", err)
	}

	// Set English language
	if err := m.SetLanguage("en"); err != nil {
		t.Fatalf("Failed to set language: %v", err)
	}

	tests := []struct {
		expression string
		expected   string
	}{
		// Nth weekday of month
		{"each first Monday of month", "0 0 * * 1#1"},
		{"each second Tuesday of month", "0 0 * * 2#2"},
		{"each third Wednesday of month", "0 0 * * 3#3"},
		{"each fourth Thursday of month", "0 0 * * 4#4"},
		{"each last Friday of month", "0 0 * * 5L"},

		// Specific day of month
		{"every 1st of the month", "0 0 1 * *"},
		{"every 15th of the month", "0 0 15 * *"},
		{"every 1st of the month at 3 pm", "0 15 1 * *"},

		// Last day of month
		{"the last day of the month", "0 0 L * *"},
		{"the last day of the month at 11:30 pm", "30 23 L * *"},
	}

	for _, test := range tests {
		result, err := m.ToCron(test.expression)
		if err != nil {
			t.Errorf("Error parsing '%s': %v", test.expression, err)
			continue
		}

		if result != test.expected {
			t.Errorf("For expression '%s' expected result '%s', got '%s'",
				test.expression, test.expected, result)
		}
	}
}

// TestEnglishYearlyExpressions tests expressions that occur on specific days of the year
func TestEnglishYearlyExpressions(t *testing.T) {
	configDir := "../config"

	m, err := NewHumanCronMapper(configDir)
	if err != nil {
		t.Fatalf("Failed to create mapper: %v", err)
	}

	// Set English language
	if err := m.SetLanguage("en"); err != nil {
		t.Fatalf("Failed to set language: %v", err)
	}

	tests := []struct {
		expression string
		expected   string
	}{
		// Month and day combinations
		{"every January 1st", "0 0 1 1 *"},
		{"every December 31st", "0 0 31 12 *"},
		{"every April 15th at 9 am", "0 9 15 4 *"},
	}

	for _, test := range tests {
		result, err := m.ToCron(test.expression)
		if err != nil {
			t.Errorf("Error parsing '%s': %v", test.expression, err)
			continue
		}

		if result != test.expected {
			t.Errorf("For expression '%s' expected result '%s', got '%s'",
				test.expression, test.expected, result)
		}
	}
}

// TestEnglishCaseInsensitivity verifies that expressions are case insensitive
func TestEnglishCaseInsensitivity(t *testing.T) {
	configDir := "../config"

	m, err := NewHumanCronMapper(configDir)
	if err != nil {
		t.Fatalf("Failed to create mapper: %v", err)
	}

	// Set English language
	if err := m.SetLanguage("en"); err != nil {
		t.Fatalf("Failed to set language: %v", err)
	}

	tests := []struct {
		lowerCase string
		mixedCase string
		expected  string
	}{
		{"every day at 3 pm", "EVERY Day At 3 PM", "0 15 * * *"},
		{"every monday", "Every MonDay", "0 0 * * 1"},
		{"each first monday of month", "Each FIRST Monday Of Month", "0 0 * * 1#1"},
	}

	for _, test := range tests {
		// Test lower case
		resultLower, err := m.ToCron(test.lowerCase)
		if err != nil {
			t.Errorf("Error parsing '%s': %v", test.lowerCase, err)
			continue
		}

		// Test mixed case
		resultMixed, err := m.ToCron(test.mixedCase)
		if err != nil {
			t.Errorf("Error parsing '%s': %v", test.mixedCase, err)
			continue
		}

		// Check that results match
		if resultLower != resultMixed {
			t.Errorf("Case sensitivity issue: '%s' -> '%s', '%s' -> '%s'",
				test.lowerCase, resultLower, test.mixedCase, resultMixed)
		}

		// Check that results match expected output
		if resultLower != test.expected {
			t.Errorf("For expression '%s' expected result '%s', got '%s'",
				test.lowerCase, test.expected, resultLower)
		}
	}
}

// TestEnglishAlternativePhrases tests equivalent expressions with different wording
func TestEnglishAlternativePhrases(t *testing.T) {
	configDir := "../config"

	m, err := NewHumanCronMapper(configDir)
	if err != nil {
		t.Fatalf("Failed to create mapper: %v", err)
	}

	// Set English language
	if err := m.SetLanguage("en"); err != nil {
		t.Fatalf("Failed to set language: %v", err)
	}

	tests := []struct {
		phrase1  string
		phrase2  string
		expected string
	}{
		{"every day at 3 pm", "each day at 3 pm", "0 15 * * *"},
		{"every monday", "each monday", "0 0 * * 1"},
		{"every first monday of month", "each first monday of the month", "0 0 * * 1#1"},
		{"every hour", "each hour", "0 * * * *"},
	}

	for _, test := range tests {
		// Test first phrase
		result1, err := m.ToCron(test.phrase1)
		if err != nil {
			t.Errorf("Error parsing '%s': %v", test.phrase1, err)
			continue
		}

		// Test second phrase
		result2, err := m.ToCron(test.phrase2)
		if err != nil {
			t.Errorf("Error parsing '%s': %v", test.phrase2, err)
			continue
		}

		// Check that results match
		if result1 != result2 {
			t.Errorf("Different phrases produce different results: '%s' -> '%s', '%s' -> '%s'",
				test.phrase1, result1, test.phrase2, result2)
		}

		// Check that results match expected output
		if result1 != test.expected {
			t.Errorf("For expression '%s' expected result '%s', got '%s'",
				test.phrase1, test.expected, result1)
		}
	}
}

// TestEnglishSpecialExpressions tests special cron expressions
func TestEnglishSpecialExpressions(t *testing.T) {
	configDir := "../rules"

	m, err := NewHumanCronMapper(configDir)
	if err != nil {
		t.Fatalf("Failed to create mapper: %v", err)
	}

	// Set English language
	if err := m.SetLanguage("en"); err != nil {
		t.Fatalf("Failed to set language: %v", err)
	}

	tests := []struct {
		expression string
		expected   string
	}{
		{"every weekday nearest 15th", "0 0 15W * *"},
		{"every business day", "0 0 * * 1-5"},
		{"at startup", "@reboot"},
		{"every year", "@yearly"},
		{"every month", "@monthly"},
		{"every week", "@weekly"},
		{"every day", "@daily"},
		{"every hour", "@hourly"},
	}

	for _, test := range tests {
		result, err := m.ToCron(test.expression)
		if err != nil {
			t.Errorf("Error parsing '%s': %v", test.expression, err)
			continue
		}

		if result != test.expected {
			t.Errorf("For expression '%s' expected result '%s', got '%s'",
				test.expression, test.expected, result)
		}
	}
}
