package cronscribe

import (
	"slices"
	"testing"
)

// TestMapperInitialization verifies the mapper initializes correctly
func TestMapperInitialization(t *testing.T) {
	configDir := "../rules" // Path to the rules directory

	m, err := NewHumanCronMapper(configDir)
	if err != nil {
		t.Fatalf("Failed to create mapper: %v", err)
	}

	// Check that English language is available
	langs := m.GetSupportedLanguages()

	englishFound := slices.Contains(langs, "en")

	if !englishFound {
		t.Error("English language not found in supported languages list")
	}
}

// TestLanguageSelection verifies language selection functionality
func TestLanguageSelection(t *testing.T) {
	configDir := "../rules"

	m, err := NewHumanCronMapper(configDir)
	if err != nil {
		t.Fatalf("Failed to create mapper: %v", err)
	}

	// Set English language explicitly
	err = m.SetLanguage("en")
	if err != nil {
		t.Errorf("Failed to set English language: %v", err)
	}

	// Try to set a non-existent language
	err = m.SetLanguage("nonexistent")
	if err == nil {
		t.Error("Expected error when setting non-existent language, but got none")
	}
}

// TestInvalidExpressions verifies handling of invalid expressions
func TestInvalidExpressions(t *testing.T) {
	configDir := "../rules"

	m, err := NewHumanCronMapper(configDir)
	if err != nil {
		t.Fatalf("Failed to create mapper: %v", err)
	}

	invalidExpressions := []string{
		"nonsense string",
		"run at some point",
		"123456",
		"",
	}

	for _, expr := range invalidExpressions {
		_, err := m.ToCron(expr)
		if err == nil {
			t.Errorf("Expected error for invalid expression '%s', but got none", expr)
		}
	}
}

// TestAutoDetection verifies language auto-detection functionality
func TestAutoDetection(t *testing.T) {
	configDir := "../rules"

	m, err := NewHumanCronMapper(configDir)
	if err != nil {
		t.Fatalf("Failed to create mapper: %v", err)
	}

	// Test auto-detection with an English expression
	result, err := m.AutoDetectAndConvert("every day at 12 pm")
	if err != nil {
		t.Errorf("Error in auto-detection for 'every day at 12 pm': %v", err)
	}

	expected := "0 12 * * *"
	if result != expected {
		t.Errorf("Auto-detection gave incorrect result: expected '%s', got '%s'", expected, result)
	}
}
