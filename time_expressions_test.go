package cronscribe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeExpressions(t *testing.T) {
	cronScribe, err := New()
	assert.NoError(t, err)

	tests := []struct {
		name        string
		language    string
		expression  string
		expected    string
		shouldError bool
	}{
		// English 24h time tests
		{
			name:       "en: daily at 24h time",
			language:   "en",
			expression: "every day at 14:30",
			expected:   "30 14 * * *",
		},
		{
			name:       "en: daily at 24h time with 'hours'",
			language:   "en",
			expression: "every day at 14:30 hours",
			expected:   "30 14 * * *",
		},
		{
			name:       "en: weekly with 24h time",
			language:   "en",
			expression: "every monday at 14:30",
			expected:   "30 14 * * 1",
		},
		// English specific time points (noon, midnight)
		{
			name:       "en: daily at noon",
			language:   "en",
			expression: "every day at noon",
			expected:   "0 12 * * *",
		},
		{
			name:       "en: daily at midnight",
			language:   "en",
			expression: "every day at midnight",
			expected:   "0 0 * * *",
		},
		{
			name:       "en: weekly at specific time",
			language:   "en",
			expression: "every monday at noon",
			expected:   "0 12 * * 1",
		},
		// English time periods (morning, afternoon, evening, night)
		{
			name:       "en: daily in morning",
			language:   "en",
			expression: "every day in the morning",
			expected:   "0 5-11 * * *",
		},
		{
			name:       "en: daily in afternoon",
			language:   "en",
			expression: "every day in the afternoon",
			expected:   "0 12-17 * * *",
		},
		{
			name:       "en: daily in evening",
			language:   "en",
			expression: "every day in the evening",
			expected:   "0 18-21 * * *",
		},
		{
			name:       "en: daily at night",
			language:   "en",
			expression: "every day in the night",
			expected:   "0 22-4 * * *",
		},
		{
			name:       "en: specific weekday in time period",
			language:   "en",
			expression: "every friday in the evening",
			expected:   "0 18-21 * * 5",
		},
		{
			name:       "en: specific day of month at 24h time",
			language:   "en",
			expression: "every 15th of the month at 14:30",
			expected:   "30 14 15 * *",
		},
		{
			name:       "en: specific month day at 24h time",
			language:   "en",
			expression: "every january 10 at 08:30",
			expected:   "30 8 10 1 *",
		},
		{
			name:       "en: last day of month at noon",
			language:   "en",
			expression: "the last day of the month at noon",
			expected:   "0 12 L * *",
		},
		{
			name:       "en: weekday nearest day at 24h time",
			language:   "en",
			expression: "the monday nearest to 15 at 16:45",
			expected:   "45 16 15W * 1",
		},

		// Russian 24h time tests
		{
			name:       "ru: daily at 24h time",
			language:   "ru",
			expression: "каждый день в 14:30",
			expected:   "30 14 * * *",
		},
		{
			name:       "ru: daily at 24h time with hours",
			language:   "ru",
			expression: "каждый день в 14:30 часов",
			expected:   "30 14 * * *",
		},
		{
			name:       "ru: weekly with 24h time",
			language:   "ru",
			expression: "каждый понедельник в 14:30",
			expected:   "30 14 * * 1",
		},
		// Russian specific time points (полдень, полночь)
		{
			name:       "ru: daily at noon",
			language:   "ru",
			expression: "каждый день в полдень",
			expected:   "0 12 * * *",
		},
		{
			name:       "ru: daily at midnight",
			language:   "ru",
			expression: "каждый день в полночь",
			expected:   "0 0 * * *",
		},
		{
			name:       "ru: weekly at specific time",
			language:   "ru",
			expression: "каждый понедельник в полдень",
			expected:   "0 12 * * 1",
		},
		// Russian time periods (утро, день, вечер, ночь)
		{
			name:       "ru: daily in morning",
			language:   "ru",
			expression: "каждый день утром",
			expected:   "0 5-11 * * *",
		},
		{
			name:       "ru: daily in afternoon",
			language:   "ru",
			expression: "каждый день днем",
			expected:   "0 12-17 * * *",
		},
		{
			name:       "ru: daily in evening",
			language:   "ru",
			expression: "каждый день вечером",
			expected:   "0 18-21 * * *",
		},
		{
			name:       "ru: daily at night",
			language:   "ru",
			expression: "каждый день ночью",
			expected:   "0 22-4 * * *",
		},

		// Dutch 24h time tests
		{
			name:       "nl: daily at 24h time",
			language:   "nl",
			expression: "elke dag om 14:30",
			expected:   "30 14 * * *",
		},
		{
			name:       "nl: daily at 24h time with uur",
			language:   "nl",
			expression: "elke dag om 14:30 uur",
			expected:   "30 14 * * *",
		},
		{
			name:       "nl: weekly with 24h time",
			language:   "nl",
			expression: "elke maandag om 14:30",
			expected:   "30 14 * * 1",
		},
		// Dutch specific time points (middernacht, middag)
		{
			name:       "nl: daily at noon",
			language:   "nl",
			expression: "elke dag om middag",
			expected:   "0 12 * * *",
		},
		{
			name:       "nl: daily at midnight",
			language:   "nl",
			expression: "elke dag om middernacht",
			expected:   "0 0 * * *",
		},
		{
			name:       "nl: weekly at specific time",
			language:   "nl",
			expression: "elke maandag om middag",
			expected:   "0 12 * * 1",
		},
		// Dutch time periods (ochtend, namiddag, avond, nacht)
		{
			name:       "nl: daily in morning",
			language:   "nl",
			expression: "elke dag in de ochtend",
			expected:   "0 5-11 * * *",
		},
		{
			name:       "nl: daily in afternoon",
			language:   "nl",
			expression: "elke dag in de namiddag",
			expected:   "0 12-17 * * *",
		},
		{
			name:       "nl: daily in evening",
			language:   "nl",
			expression: "elke dag in de avond",
			expected:   "0 18-21 * * *",
		},
		{
			name:       "nl: daily at night",
			language:   "nl",
			expression: "elke dag 's nacht",
			expected:   "0 22-4 * * *",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := cronScribe.SetLanguage(tt.language)
			assert.NoError(t, err)

			result, err := cronScribe.Convert(tt.expression)
			if tt.shouldError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTimeExpressionsAutoDetect(t *testing.T) {
	cronScribe, err := New()
	assert.NoError(t, err)

	tests := []struct {
		name        string
		expression  string
		expected    string
		shouldError bool
	}{
		// Auto-detect various languages
		{
			name:       "auto: english 24h time",
			expression: "every day at 14:30",
			expected:   "30 14 * * *",
		},
		{
			name:       "auto: english time period",
			expression: "every monday in the afternoon",
			expected:   "0 12-17 * * 1",
		},
		{
			name:       "auto: russian 24h time",
			expression: "каждый день в 14:30",
			expected:   "30 14 * * *",
		},
		{
			name:       "auto: russian time period",
			expression: "каждый понедельник утром",
			expected:   "0 5-11 * * 1",
		},
		{
			name:       "auto: dutch 24h time",
			expression: "elke dag om 14:30",
			expected:   "30 14 * * *",
		},
		{
			name:       "auto: dutch time period",
			expression: "elke maandag in de ochtend",
			expected:   "0 5-11 * * 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := cronScribe.AutoDetect(tt.expression)
			if tt.shouldError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}
