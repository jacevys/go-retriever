package main

import (
	"fmt"
	"go-retriever/internal/retriever"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("網頁內容擷取與摘要")

	urlEntry := widget.NewEntry()
	urlEntry.SetPlaceHolder("請輸入網址...")

	titleLabel := widget.NewLabel("標題將顯示於此")
	titleLabel.TextStyle.Bold = true

	contentLabel := widget.NewLabel("內容將顯示於此")
	contentLabel.Wrapping = fyne.TextWrapWord
	contentScroll := container.NewScroll(contentLabel)

	summaryLabel := widget.NewLabel("摘要將顯示於此")
	summaryLabel.Wrapping = fyne.TextWrapWord
	summaryScroll := container.NewScroll(summaryLabel)

	var fetchedArticleText string
	// 20251118ok

	fetchButton := widget.NewButton("網頁內容", func() {
		titleLabel.SetText("獲取中...")
		contentLabel.SetText("")
		summaryLabel.SetText("")
		fetchedArticleText = ""

		article, err := retriever.Fetch(urlEntry.Text)
		if err != nil {
			titleLabel.SetText("錯誤")
			contentLabel.SetText(fmt.Sprintf("獲取失敗: %v", err))
			return
		}

		titleLabel.SetText(article.Title)
		contentLabel.SetText(article.TextContent)
		fetchedArticleText = article.TextContent
	})

	summarizeButton := widget.NewButton("摘要內容", func() {
		if fetchedArticleText == "" {
			summaryLabel.SetText("請先擷取文章內容再進行摘要。")
			return
		}

		summaryLabel.SetText("摘要生成中...")
		go func() {
			apiKey := os.Getenv("GEMINI_API_KEY")
			if apiKey == "" {
				log.Println("錯誤: GEMINI_API_KEY 環境變數未設定。")
				summaryLabel.SetText("摘要失敗: GEMINI_API_KEY 未設定")
				return
			}

			summary, err := retriever.SummarizeWithGemini(apiKey, fetchedArticleText)
			if err != nil {
				log.Printf("生成摘要失敗： %v", err)
				summaryLabel.SetText(fmt.Sprintf("摘要失敗: %v", err))
				return
			}
			summaryLabel.SetText(summary)
		}()
	})

	topControls := container.NewVBox(
		widget.NewLabel("請輸入網址："),
		urlEntry,
		container.NewHBox(fetchButton, summarizeButton),
	)

	contentArea := container.NewHSplit(
		container.NewBorder(
			container.NewVBox(
				titleLabel,
				widget.NewSeparator(),
				widget.NewLabel("原始內容："),
			), nil, nil, nil,
			contentScroll,
		),

		container.NewBorder(
			widget.NewLabel("摘要內容："), nil, nil, nil,
			summaryScroll,
		),
	)

	w.SetContent(container.NewBorder(
		topControls,
		nil, nil, nil,
		contentArea,
	))

	w.Resize((fyne.NewSize(1200, 800)))
	w.ShowAndRun()
}

// 20251118ok
