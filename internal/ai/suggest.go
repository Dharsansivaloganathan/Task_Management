package ai

import (
    "bytes"
    "encoding/json"
    "net/http"
    "os"
)

type OpenAIRequest struct {
    Prompt string `json:"prompt"`
}

type OpenAIResponse struct {
    Choices []struct {
        Text string `json:"text"`
    } `json:"choices"`
}

func GetTaskSuggestion() (string, error) {
    apiKey := os.Getenv("OPENAI_API_KEY")
    url := "https://api.openai.com/v1/completions"

    reqBody, _ := json.Marshal(OpenAIRequest{Prompt: "Suggest a task for a project"})
    req, _ := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
    req.Header.Set("Authorization", "Bearer "+apiKey)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    var openAIResp OpenAIResponse
    json.NewDecoder(resp.Body).Decode(&openAIResp)
    return openAIResp.Choices[0].Text, nil
}
