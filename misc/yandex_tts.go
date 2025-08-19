package yandex

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Result struct {
	Result string `json:"result"`
}

func Text2Speech(folderId, iamToken, text, voice, output string) error {
	endpoint := "https://tts.api.cloud.yandex.net/speech/v1/tts:synthesize"

	form := url.Values{}
	form.Set("text", text)
	form.Set("lang", "ru-RU")
	form.Set("voice", voice)
	form.Set("folderId", folderId)
	form.Set("format", "lpcm")
	form.Set("sampleRateHertz", "48000")

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+iamToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("bad response: %d — %s", resp.StatusCode, string(bodyBytes))
	}

	outFile, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("cannot create output file: %w", err)
	}
	defer outFile.Close()

	if _, err := io.Copy(outFile, resp.Body); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
