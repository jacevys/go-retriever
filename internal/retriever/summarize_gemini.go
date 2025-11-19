package retriever

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func SummarizeWithGemini(apikey string, question string) (string, error) {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(apikey))
	if err != nil {
		return "", fmt.Errorf("create client error: %w", err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.5-flash")

	resp, err := model.GenerateContent(ctx, genai.Text(question))
	if err != nil {
		return "", fmt.Errorf("API error: %w", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("Empty response")
	}

	return fmt.Sprint(resp.Candidates[0].Content.Parts[0]), nil
}

// 20251119ok
