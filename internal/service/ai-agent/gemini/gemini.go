package gemini

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Service interacts with the Gemini API to generate email drafts.
type GeminiService struct {
	apiKey string
}

// New creates a new Gemini service with the provided API key.
func NewGeminiService(apiKey string) GeminiService {
	return GeminiService{apiKey: apiKey}
}

// GenerateEmail uses the Gemini API to generate an email draft based on the input prompt.
func (s GeminiService) GenerateEmailReply(prompt string) (string, error) {
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent?key=%s", s.apiKey)

	requestBody := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]string{
					{"text": prompt},
				},
			},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to send request to Gemini API: %v", err)
	}
	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read Gemini response: %v", err)
	}
	var responseMap map[string]interface{}
	if err := json.Unmarshal(respData, &responseMap); err != nil {
		return "", fmt.Errorf("failed to decode Gemini response: %v", err)
	}

	candidates, ok := responseMap["candidates"].([]interface{})
	if !ok || len(candidates) == 0 {
		return "", fmt.Errorf("no candidates found in Gemini response")
	}

	content := candidates[0].(map[string]interface{})["content"].(map[string]interface{})
	parts := content["parts"].([]interface{})
	text := parts[0].(map[string]interface{})["text"].(string)

	return text, nil
}
