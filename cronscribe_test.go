package cronscribe

import (
	"slices"
	"testing"

	"github.com/flaticols/cronscribe/pkg/langs"
	"github.com/flaticols/cronscribe/pkg/rules"
	"github.com/stretchr/testify/require"
)

// Helper function to mock a CronScribe instance with specific rules for testing
func createMockCronScribe() *CronScribe {
	// Create a basic English rule set for testing
	englishRules := &rules.RuleSet{
		Language: "en",
		Rules: []rules.Rule{
			{
				Name:    "daily_at_time",
				Pattern: `(?i)(?:each|every)\s+day\s+at\s+(\d+)(?::(\d+))?\s*(am|pm)?`,
				Variables: map[string]int{
					"hour":   1,
					"minute": 2,
					"ampm":   3,
				},
				Format: "%minute %hour * * *",
				DefaultValues: map[string]string{
					"minute": "0",
				},
				Transformations: map[string][]rules.Transformation{
					"hour": {
						{
							Condition: "ampm == \"pm\" && hour < 12",
							Operation: "hour + 12",
						},
						{
							Condition: "ampm == \"am\" && hour == 12",
							Operation: "0",
						},
					},
				},
			},
			{
				Name:    "weekly_day",
				Pattern: `(?i)(?:each|every)\s+(monday|tuesday|wednesday|thursday|friday|saturday|sunday)(?:\s+at\s+(\d+)(?::(\d+))?\s*(am|pm)?)?`,
				Variables: map[string]int{
					"weekday": 1,
					"hour":    2,
					"minute":  3,
					"ampm":    4,
				},
				Format: "%minute %hour * * %weekday",
				DefaultValues: map[string]string{
					"minute": "0",
					"hour":   "0",
				},
				Dictionaries: map[string]string{
					"weekday": "weekdays",
					"ampm":    "time_ampm",
				},
				Transformations: map[string][]rules.Transformation{
					"hour": {
						{
							Condition: "ampm == \"pm\" && hour < 12",
							Operation: "hour + 12",
						},
						{
							Condition: "ampm == \"am\" && hour == 12",
							Operation: "0",
						},
					},
				},
			},
			{
				Name:    "last_day_of_month",
				Pattern: `(?i)(?:each|every|the)\s+last\s+day\s+of\s+(?:the\s+)?month`,
				Format:  "0 0 L * *",
			},
		},
		Dictionaries: map[string]rules.Dictionary{
			"weekdays": {
				"monday":    "1",
				"tuesday":   "2",
				"wednesday": "3",
				"thursday":  "4",
				"friday":    "5",
				"saturday":  "6",
				"sunday":    "0",
			},
			"time_ampm": {
				"am": "am",
				"pm": "pm",
			},
		},
	}

	// Compile patterns
	for i := range englishRules.Rules {
		_ = englishRules.Rules[i].CompilePattern()
	}

	// Create a mapper with mock rules
	mockMapper := &Mapper{
		langs: map[string]*rules.RuleSet{
			"en": englishRules,
		},
		currentRules: englishRules,
	}

	// Create a CronScribe with the mock mapper
	return &CronScribe{
		mapper: mockMapper,
	}
}

func TestNew(t *testing.T) {
	cs, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	if cs == nil {
		t.Fatalf("New() returned nil CronScribe instance")
	}
	if cs.mapper == nil {
		t.Fatalf("New() returned CronScribe with nil mapper")
	}
}

func TestCronScribe_Convert(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		want       string
		wantErr    bool
		setupFunc  func(*CronScribe) // Optional setup function
	}{
		{
			name:       "Every day at 9am",
			expression: "every day at 9am",
			want:       "0 9 * * *",
			wantErr:    false,
		},
		{
			name:       "Every last day of month",
			expression: "every last day of month",
			want:       "0 0 L * *",
			wantErr:    false,
		},
		{
			name:       "Invalid expression",
			expression: "some invalid expression",
			want:       "",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Use the mock CronScribe for testing
			cs := createMockCronScribe()

			// Apply any test-specific setup
			if tt.setupFunc != nil {
				tt.setupFunc(cs)
			}

			got, err := cs.Convert(tt.expression)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCronScribe_GetSupportedLanguages(t *testing.T) {
	// Use the mock CronScribe for testing
	cs := createMockCronScribe()

	langs := cs.GetSupportedLanguages()
	if len(langs) == 0 {
		t.Errorf("GetSupportedLanguages() returned empty array")
	}

	// Check that it contains at least English
	found := slices.Contains(langs, "en")
	if !found {
		t.Errorf("GetSupportedLanguages() does not contain 'en'")
	}
}

func TestCronScribe_RU(t *testing.T) {
	cs, err := New()
	require.NoError(t, err)
	cs.SetLanguage(langs.LangRU)

	tests := []struct {
		name       string
		expression string
		want       string
		wantErr    bool
	}{
		{
			name:       "Every day at 19",
			expression: "каждый день в 19",
			want:       "0 19 * * *",
			wantErr:    false,
		},
		{
			name:       "Every day at 19",
			expression: "каждый день в 13 часов",
			want:       "0 13 * * *",
			wantErr:    false,
		},
		{
			name:       "Invalid expression",
			expression: "some invalid expression",
			want:       "",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, err := cs.Convert(tt.expression)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && e != tt.want {
				t.Errorf("Convert() = %v, want %v", e, tt.want)
			}
		})
	}
}
