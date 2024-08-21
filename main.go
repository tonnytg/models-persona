package main

import (
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	log.Println("Start creating your Persona friend")

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GOOGLE_API_KEY")))
	model := client.GenerativeModel("gemini-1.5-flash")

  imgData, _ := os.ReadFile("photo.jpg")
	resp, err := model.GenerateContent(
		ctx,
		genai.Text("What's in this photo?"),
		genai.ImageData("jpeg", imgData))

  if err != nil {
    log.Println("Error:", err)
    os.Exit(1)
  }
  log.Println("Resp:", resp)
}
