// banyak yang ngira kalau dah jago bahasa go bakal dianggap sepuh. halah memek
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Response struct {
	Message string `json:"message"`
}

func sendRequest(text string) (string, error) {
	data := url.Values{}
	data.Set("text", text)
	data.Set("lc", "id")

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.simsimi.vn/v1/simtalk", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: %v", resp.Status)
	}

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	if response.Message == "" {
		return "no msg", nil
	}

	return response.Message, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("you>: ")
		inputText, _ := reader.ReadString('\n')
		inputText = strings.TrimSpace(inputText)

		if inputText == "exit" || inputText == "quit" || inputText == "keluar" || inputText == "murtad" {
			fmt.Println("byebye!!!!")
			break
		}

		message, err := sendRequest(inputText)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("bot>:", message)
	}
}
