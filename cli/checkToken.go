package cli

import (
	"fmt"
	"net/http"
	"os"
)

func CheckToken(apiKey string) {
	endpoint := "https://api.openai.com/v1/images/generations"
	client := &http.Client{}
	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == 401 {
		fmt.Println("The token is invalid")
		os.Exit(0)
	} else {
		return
	}
}
