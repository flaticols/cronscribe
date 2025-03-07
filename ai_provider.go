package cronscribe

import (
	"context"
)

// AIProvider is an interface for services that can generate cron expressions from human text
type AIProvider interface {
	// GenerateCron generates a cron expression from a human-readable string
	GenerateCron(ctx context.Context, input string) (string, error)
}

// RecommendedSystemPrompt returns a recommended system prompt for AI models
func RecommendedSystemPrompt() string {
	return "You are a helpful assistant that converts human-readable schedule descriptions to cron expressions. Only respond with the valid cron expression, without any explanations."
}

// RecommendedUserPrompt formats input text into a recommended user prompt
func RecommendedUserPrompt(input string) string {
	return `Convert the following human-readable schedule description to a cron expression:
"` + input + `"

The response should be ONLY the valid cron expression in the standard 5-field format (minute hour day-of-month month day-of-week).
Do not include any explanations or additional text.`
}