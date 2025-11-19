package retriever

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/go-shiori/go-readability"
)

// Article 包含了擷取到的文章內容。
type Article struct {
	Title       string
	TextContent string
}

// Fetch 根據給定的 URL 擷取網頁內容並解析成純文字。
func Fetch(inURL string) (*Article, error) {
	timeout := 15 * time.Second

	if inURL == "" {
		return nil, fmt.Errorf("請輸入網址")
	}

	u, err := url.Parse(inURL)
	if err != nil {
		return nil, fmt.Errorf("網址無效: %w", err)
	}

	client := &http.Client{Timeout: timeout}
	req, _ := http.NewRequest("GET", u.String(), nil)
	req.Header.Set("User-Agent", "Go-Retriever-Desktop/1.0")
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP 錯誤: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("非預期的 HTTP 狀態: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("讀取回應內容時發生錯誤: %w", err)
	}

	parsedArticle, err := readability.FromReader(bytes.NewReader(body), u)
	if err != nil {
		return nil, fmt.Errorf("解析文章時發生錯誤: %w", err)
	}

	// 返回我們自訂的 Article 結構
	return &Article{
		Title:       parsedArticle.Title,
		TextContent: parsedArticle.TextContent,
	}, nil
}
