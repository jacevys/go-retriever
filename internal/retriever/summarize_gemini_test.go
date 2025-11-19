package retriever

import (
	"os"
	"testing"
)

func TestSummarizeWithGemini(t *testing.T) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		t.Skip("Skipping test: GEMINI_API_KEY environment variable not set.")
	}

	t.Run("WithAPIKey", func(t *testing.T) {
		text := `人工智慧（英語：artificial intelligence，縮寫為 AI）亦稱智械、機器智慧，指由人製造出來的機器所表現出來的智慧。
		通常人工智慧是指透過普通電腦程式來呈現人類智慧的技術。該詞也指出研究這樣的智慧系統是否能夠實現，以及如何實現的科學領域。
		同時，人類的智慧涉及許多、對目前的AI來說還很複雜的能力，例如：自我意識、創造力、智慧和心靈。`

		summary, err := SummarizeWithGemini(apiKey, text)

		if err != nil {
			t.Fatalf("SummarizeWithGemini() with a valid API key failed with error: %v", err)
		}

		if summary == "" {
			t.Errorf("Expected a non-empty summary, but got an empty string.")
		}

		t.Logf("Received summary: %s", summary)
	})
}

// 20251119ok
