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
	Succes   bool
}

func Ping(url string, ch chan Response) {
	start := time.Now()

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		ch <- Response{Duration: time.Since((start)), URL: url, Succes: false}
		return
	}

	defer resp.Body.Close()

	ch <- Response{Duration: time.Since((start)), URL: url, Succes: true}
}

func main() {
	exePath, err := os.Executable() // путь до бинарника
	if err != nil {
		fmt.Println("❌ Не удалось определить путь к программе:", err)
		return
	}

	configPath := filepath.Join(filepath.Dir(exePath), "sites.conf")

	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("❌ Ошибка при чтении конфига:", err)
		return
	}

	fileContent := string(data)
	sites := strings.Split(fileContent, "\n")

	ch := make(chan Response)

	for _, val := range sites {
		go Ping(val, ch)
	}

	for i := 0; i < len(sites); i++ {
		res := <-ch

		if res.Succes {
			fmt.Printf("✅ %-30s  |  Время: %v\n", res.URL, res.Duration)
		} else {
			fmt.Printf("❌ %s\n", res.URL)
		}
	}
}
