package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	url := flag.String("url", "", "URL для загрузки")
	depth := flag.Int("depth", 1, "Глубина")
	flag.Parse()

	if *url == "" {
		fmt.Println("Использование: go run main.go -url=https://example.com")
		os.Exit(1)
	}

	downloader := NewDownloader()
	err := downloader.Download(*url, *depth)
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	fmt.Println("Готово!")
}
