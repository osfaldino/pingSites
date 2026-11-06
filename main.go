package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Response struct {
	Duration time.Duration
	URL      string
	Success  bool
}

func Ping(url string, ch chan Response) {
	start := time.Now()

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		ch <- Response{Duration: time.Since((start)), URL: url, Success: false}
		return
	}

	defer resp.Body.Close()

	ch <- Response{Duration: time.Since((start)), URL: url, Success: true}
}

func main() {
	exePath, err := os.Executable() // path to execute file
	if err != nil {
		fmt.Println("❌ Cannot found path to exec file:", err)
		return
	}

	configPath := filepath.Join(filepath.Dir(exePath), "sites.conf")

	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("❌ Error while reading config:", err)
		return
	}

	fileContent := string(data)
	sites := strings.Split(fileContent, "\n")

	ch := make(chan Response)
	counter := 0
	for _, val := range sites {
		val = strings.TrimSpace(val)
		if val == "" {
			continue
		}
		counter++
		go Ping(val, ch)
	}

	for i := 0; i < counter; i++ {
		res := <-ch

		if res.Success {
			fmt.Printf("✅ %-30s  |  Time: %v\n", res.URL, res.Duration)
		} else {
			fmt.Printf("❌ %s\n", res.URL)
		}
	}
}
