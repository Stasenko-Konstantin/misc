// Example use for synchronous recognition API Yandex SpeechKit in Golang.
// Based on https://cloud.yandex.com/en/docs/speechkit/stt/api/request-examples
package misc

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

const (
	FOLDER_ID = "<folder ID>" // Folder ID
	IAM_TOKEN = "<IAM token>" // IAM token
)

type Result struct {
	Result string `json:"result"`
}

func Speech2Text(file []byte) (string, error) {
	data := bytes.NewReader(file)
	params := strings.Join([]string{
		"topic=general",
		"folderId=" + FOLDER_ID,
		"lang=ru-RU",
	}, "&")

	req, err := http.NewRequest("POST", "https://stt.api.cloud.yandex.net/speech/v1/stt:recognize?"+params, data)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+IAM_TOKEN)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result Result
	body, _ := io.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &result); err != nil {
		return "", err
	}
	return result.Result, nil
}
