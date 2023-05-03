package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const BaseURL = "https://api.openweathermap.org"

func main() {
	key := os.Getenv("OPENWEATHERMAP_API_KEY")
	url := fmt.Sprintf("%s/data/2.5/weather?q=London,UK&appid=a474b5d959538fd0f942b27e2788bddd%s", BaseURL, key)
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Unexpected response status", resp.Status)

	}

	io.Copy(os.Stdout, resp.Body)
}
