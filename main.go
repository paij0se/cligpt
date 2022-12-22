package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/paij0se/cligpt/cli"
)

type Data struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int                    `json:"created"`
	Model   string                 `json:"model"`
	Choices []TextCompletionChoice `json:"choices"`
	Usage   TextCompletionUsage    `json:"usage"`
}

type TextCompletionChoice struct {
	Text         string  `json:"text"`
	Index        int     `json:"index"`
	LogProbs     *string `json:"logprobs"`
	FinishReason string  `json:"finish_reason"`
}

type TextCompletionUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

func main() {
	cli.CreateConfigDirectory()
	if len(cli.ReadFromYml("auth")) < 51 {
		log.Fatal("Ensure to insert a valid token in cligpt.yml file.")
	}
	client := &http.Client{}
	var data = strings.NewReader(`{
		  "model": "` + cli.ReadFromYml("model") + `",
		  "prompt": "` + os.Args[1] + `",
		  "temperature": 0.7,
		  "max_tokens": ` + cli.ReadFromYml("max_tokens") + `,
		  "top_p": 1,
		  "frequency_penalty": 0,
		  "presence_penalty": 0
		}`)
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", `Bearer `+cli.ReadFromYml("auth")+``)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		log.Println("The token is valid?")
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var response Data
	json.Unmarshal(bodyText, &response)
	if err != nil {
		log.Println(err)
	}
	choice := response.Choices[0]
	text := choice.Text
	fmt.Println(text)
	//fmt.Println(bodyText)
}
