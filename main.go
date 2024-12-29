package main

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	log.Println("Starting the process to create your Persona friend...")

	// Verifying if the API key is set in the environment variables
	apiKey := os.Getenv("GOOGLE_API_KEY")
	if apiKey == "" {
		log.Fatal("GOOGLE_API_KEY is not set in the environment variables.")
	}

	// Initialize the context and the Gemini client
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Failed to create Gemini client: %v", err)
	}

	// Use the "gemini-1.5-flash" model for generating content
	model := client.GenerativeModel("gemini-1.5-flash")

	// Reading image data from the file system
	imgData, err := os.ReadFile("photo.jpg")
	if err != nil {
		log.Fatalf("Failed to read the image file: %v", err)
	}

	// Generating content by passing the image data and a text prompt
	resp, err := model.GenerateContent(
		ctx,
		genai.Text("What's in this photo?"),
		genai.ImageData("jpeg", imgData),
	)
	if err != nil {
		log.Fatalf("Error generating content: %v", err)
	}

	// Printing the response from Gemini model
	log.Printf("Response: %v", resp)
}
