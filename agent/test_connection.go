package main

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/genai"
)

func main() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		fmt.Println("Client error:", err)
		os.Exit(1)
	}
	result, err := client.Models.GenerateContent(ctx,
		"gemini-3.5-flash", genai.Text("Say hello in one word"), nil)
	if err != nil {
		fmt.Println("API error:", err)
		os.Exit(1)
	}
	fmt.Println("Connected! Gemini says:", result.Candidates[0].Content.Parts[0])
}
