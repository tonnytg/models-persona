# Models Persona

## Overview
This project is a Go-based API that integrates with the Gemini model to create a persona friend. The API uses the Gemini model to generate content based on an image and a textual prompt. The goal of this project is to demonstrate how to integrate generative AI capabilities into Go applications using the Gemini backend.

## Requirements
- Go 1.18 or higher
- Google Cloud API key
- The `photo.jpg` image file for analysis

## Setup

1. Clone this repository:
    ```bash
    git clone https://github.com/yourusername/models-persona.git
    cd models-persona
    ```

2. Set your Google Cloud API Key as an environment variable:
    ```bash
    export GOOGLE_API_KEY="your-google-api-key"
    ```

3. Install necessary Go dependencies:
    ```bash
    go mod tidy
    ```

4. Ensure you have a photo named `photo.jpg` in the project directory. This image will be used as input for the Gemini model.

5. Run the application:
    ```bash
    go run main.go
    ```

6. The application will generate content based on the image and a textual prompt ("What's in this photo?"). The response from Gemini will be logged to the console.

## Description

This API connects to the Gemini generative model (version 1.5) via the `generative-ai-go` client. It sends an image to the model along with a question, "What's in this photo?", and the model generates a response. The API provides a way to interact with the Gemini model in a simple Go application, ideal for building AI-driven persona-based applications or interactive chatbots.

## Error Handling

- If the `GOOGLE_API_KEY`
