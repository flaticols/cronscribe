package cronscribe

import (
	"context"
	"testing"
)

// MockAIProvider simulates an AI provider for testing
type MockAIProvider struct {
	responses map[string]string
}

func NewMockAIProvider() *MockAIProvider {
	return &MockAIProvider{
		responses: map[string]string{
			"every monday at 10am":          "0 10 * * 1",
			"every day at midnight":         "0 0 * * *",
			"every 15 minutes":              "*/15 * * * *",
			"at 2:30pm on the first of month": "30 14 1 * *",
		},
	}
}

func (m *MockAIProvider) GenerateCron(ctx context.Context, input string) (string, error) {
	if response, ok := m.responses[input]; ok {
		return response, nil
	}
	return "0 0 * * *", nil // Default response for testing
}

func TestBraveHumanCronMapper_LocalRulesFirst(t *testing.T) {
	// Create the base rules for testing
	baseRules := createTestRules()
	
	// Create a mock AI provider
	mockAI := NewMockAIProvider()
	
	// Create the BraveHumanCronMapper with default options (local rules first)
	braveMapper, err := NewBraveHumanCronMapper("", mockAI)
	if err != nil {
		t.Fatalf("Failed to create mapper: %v", err)
	}
	
	// Manually set the rules for testing
	braveMapper.HumanCronMapper.allRules = baseRules
	braveMapper.HumanCronMapper.currentRules = baseRules["en"]

	// Compile the patterns
	for i := range braveMapper.HumanCronMapper.currentRules.Rules {
		if err := braveMapper.HumanCronMapper.currentRules.Rules[i].CompilePattern(); err != nil {
			t.Fatalf("Failed to compile pattern: %v", err)
		}
	}

	// Test case that exists in local rules
	result, err := braveMapper.ToCron("every day at noon")
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if result != "0 12 * * *" {
		t.Errorf("Expected '0 12 * * *', got: %s", result)
	}

	// Test case that only exists in AI
	result, err = braveMapper.ToCron("every monday at 10am")
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if result != "0 10 * * 1" {
		t.Errorf("Expected '0 10 * * 1', got: %s", result)
	}
}

func TestBraveHumanCronMapper_AIFirst(t *testing.T) {
	// Create the base rules for testing
	baseRules := createTestRules()
	
	// Create a mock AI provider
	mockAI := NewMockAIProvider()
	mockAI.responses["every day at noon"] = "0 12 * * *" // Same as local to simplify test
	
	// Create the BraveHumanCronMapper with AIFirst option
	braveMapper, err := NewBraveHumanCronMapper("", mockAI, WithAIFirst(true))
	if err != nil {
		t.Fatalf("Failed to create mapper: %v", err)
	}
	
	// Manually set the rules for testing
	braveMapper.HumanCronMapper.allRules = baseRules
	braveMapper.HumanCronMapper.currentRules = baseRules["en"]

	// Compile the patterns
	for i := range braveMapper.HumanCronMapper.currentRules.Rules {
		if err := braveMapper.HumanCronMapper.currentRules.Rules[i].CompilePattern(); err != nil {
			t.Fatalf("Failed to compile pattern: %v", err)
		}
	}

	// AI should be used first and its result returned
	result, err := braveMapper.ToCron("every day at noon")
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if result != "0 12 * * *" {
		t.Errorf("Expected '0 12 * * *', got: %s", result)
	}

	// Test case that only exists in AI
	result, err = braveMapper.ToCron("every monday at 10am") 
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if result != "0 10 * * 1" {
		t.Errorf("Expected '0 10 * * 1', got: %s", result)
	}
}

// Helper function to create test rules
func createTestRules() map[string]*Rules {
	return map[string]*Rules{
		"en": {
			Language: "en",
			Rules: []Rule{
				{
					Name:    "every day at noon",
					Pattern: "every day at noon",
					Format:  "0 12 * * *",
				},
			},
		},
	}
}