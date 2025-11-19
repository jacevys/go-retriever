# go-retriever：網頁內容擷取與 Gemini 摘要工具

## 專案概覽

go-retriever 是一個 Go 語言專案，旨在提供一個桌面應用程式，用於擷取網頁內容並使用 Google Gemini API 進行摘要。它也整合了 `go-llama.cpp` 子模組，用於本地語言模型操作（儘管目前摘要功能已切換至 Gemini API）。

## 功能

*   **網頁內容擷取**：從指定的 URL 擷取文章標題和主要文本內容。
*   **Gemini 摘要**：使用 Google Gemini 2.5 Flash 模型對擷取到的文章內容進行簡潔的繁體中文摘要。
*   **桌面應用程式**：提供一個基於 Fyne 的圖形使用者介面 (GUI)，方便使用者輸入 URL、擷取內容並查看摘要。

## 環境設定

在執行本專案之前，請確保您的系統已安裝以下工具：

1.  **Go 語言環境**：請從 [Go 官方網站](https://golang.org/dl/) 下載並安裝最新版本的 Go。
2.  **Git**：用於克隆專案和管理子模組。

### 步驟 1：克隆專案與初始化子模組

```bash
git clone https://github.com/your-username/go-retriever.git # 請替換為您的實際儲存庫位址
cd go-retriever
git submodule update --init --recursive
```

### 步驟 2：安裝 Go 模組依賴

```bash
go mod tidy
```

### 步驟 3：設定 Gemini API 金鑰

您需要一個 Google Gemini API 金鑰來使用摘要功能。請按照以下步驟操作：

1.  前往 [Google AI Studio](https://aistudio.google.com/app/apikey) 取得您的 API 金鑰。
2.  將您的 API 金鑰設定為環境變數 `GEMINI_API_KEY`。在 Linux/macOS 上，您可以在終端機中執行：
    ```bash
    export GEMINI_API_KEY="您的API金鑰"
    ```
    （請將 `"您的API金鑰"` 替換為您實際的金鑰）
    
    為了方便，您也可以將此行加入到您的 `~/.bashrc`, `~/.zshrc` 或其他 shell 設定檔中，使其永久生效。

## 如何執行桌面應用程式

設定完成後，您可以執行桌面應用程式：

```bash
go run ./cmd/desktop
```

應用程式啟動後，您可以在輸入框中輸入網址，點擊「擷取內容」按鈕，然後點擊「摘要內容」按鈕來查看摘要。

## 如何執行測試

您可以執行專案中的單元測試來驗證功能：

```bash
go test ./...
```

如果您只想測試 `retriever` 模組，可以執行：

```bash
go test ./internal/retriever
```

請注意，`TestSummarizeWithGemini` 測試需要設定 `GEMINI_API_KEY` 環境變數才能執行。

## 專案結構

```
go-retriever/
├── cmd/
│   ├── desktop/        # 桌面應用程式的進入點
│   └── server/         # 伺服器應用程式的進入點 (如果存在)
├── go-llama.cpp/       # LLaMA.cpp 的 Go 語言綁定子模組
├── internal/
│   └── retriever/      # 核心擷取和摘要邏輯
│       ├── fetch.go
│       ├── summarize.go
│       ├── summarize_gemini.go
│       └── summarize_gemini_test.go
├── models/             # 語言模型檔案 (例如 .gguf 檔案)
├── .gitignore          # Git 忽略檔案
├── go.mod
├── go.sum
└── README.md           # 本說明文件
```

## 貢獻

歡迎任何形式的貢獻！如果您有任何建議、錯誤報告或功能請求，請隨時提出 Issue 或提交 Pull Request。

## 授權

本專案採用 MIT 授權。詳情請參閱 `LICENSE` 檔案。
