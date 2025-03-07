package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/flaticols/cronscribe"
)

func main() {
	// Check if OpenAI API key is set
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY environment variable must be set")
	}

	// Create a LangChain provider with default model (gpt-3.5-turbo)
	provider, err := cronscribe.NewLangChainProvider()
	if err != nil {
		log.Fatalf("Failed to create LangChain provider: %v", err)
	}

	// Create a brave mapper with LangChain provider
	mapper, err := cronscribe.NewBraveHumanCronMapper("../../rules", provider, cronscribe.WithAIFirst(true))
	if err != nil {
		log.Fatalf("Failed to create mapper: %v", err)
	}

	// Example expressions
	expressions := []string{
		"every day at 3 PM",
		"every weekday at 9 AM",
		"at 10:30 PM on the 1st and 15th of every month",
		"every hour from 9 AM to 5 PM on weekdays",
	}

	fmt.Println("Converting human-readable schedules to cron expressions using LangChain:")
	fmt.Println("-------------------------------------------------------------------")

	for _, expr := range expressions {
		cronExpr, err := mapper.ToCron(expr)
		if err != nil {
			fmt.Printf("❌ Failed to convert '%s': %v\n", expr, err)
			continue
		}
		fmt.Printf("✅ '%s' → '%s'\n", expr, cronExpr)
	}

	// Example using a different model
	fmt.Println("\nUsing a more capable model:")
	fmt.Println("-------------------------------------------------------------------")

	// Create a provider with a more capable model
	gpt4Provider, err := cronscribe.NewLangChainProvider(
		cronscribe.WithLangChainModel("gpt-4"),
	)
	if err != nil {
		log.Fatalf("Failed to create GPT-4 provider: %v", err)
	}

	// Create a mapper with the GPT-4 provider
	gpt4Mapper, err := cronscribe.NewBraveHumanCronMapper("../../rules", gpt4Provider, cronscribe.WithAIFirst(true))
	if err != nil {
		log.Fatalf("Failed to create mapper: %v", err)
	}

	// Test with a more complex expression
	complexExpr := "run at 3:45 AM on the second Tuesday of each month and on the last day of each quarter"
	cronExpr, err := gpt4Mapper.ToCron(complexExpr)
	if err != nil {
		fmt.Printf("❌ Failed to convert '%s': %v\n", complexExpr, err)
	} else {
		fmt.Printf("✅ '%s' → '%s'\n", complexExpr, cronExpr)
	}

	// Example of direct use
	fmt.Println("\nDirect use of the LangChain provider:")
	fmt.Println("-------------------------------------------------------------------")
	directResult, err := provider.GenerateCron(context.Background(), "every Monday at 7 AM")
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Printf("✅ Result: %s\n", directResult)
	}
}