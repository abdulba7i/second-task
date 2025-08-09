package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Downloader struct {
	visited map[string]bool
}

func NewDownloader() *Downloader {
	return &Downloader{
		visited: make(map[string]bool),
	}
}

func (d *Downloader) Download(startURL string, depth int) error {
	if depth <= 0 {
		return nil
	}

	if d.visited[startURL] {
		return nil
	}
	d.visited[startURL] = true

	fmt.Printf("Загружаю: %s\n", startURL)

	// Загружаем страницу
	resp, err := http.Get(startURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Сохраняем файл
	parsedURL, _ := url.Parse(startURL)
	filename := parsedURL.Host + ".html"
	err = os.WriteFile(filename, content, 0644)
	if err != nil {
		return err
	}

	// Если это HTML, ищем ссылки и ресурсы
	if strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
		links := d.findLinks(string(content), startURL)
		resources := d.findResources(string(content), startURL)

		// Загружаем ресурсы
		for _, resource := range resources {
			d.Download(resource, depth-1)
		}

		// Рекурсивно загружаем ссылки
		for _, link := range links {
			d.Download(link, depth-1)
		}
	}

	return nil
}

func (d *Downloader) findLinks(html, baseURL string) []string {
	var links []string

	// Ищем href ссылки
	parts := strings.Split(html, "href=\"")
	for i := 1; i < len(parts); i++ {
		end := strings.Index(parts[i], "\"")
		if end != -1 {
			link := parts[i][:end]
			if strings.HasPrefix(link, "http") {
				links = append(links, link)
			}
		}
	}

	return links
}

func (d *Downloader) findResources(html, baseURL string) []string {
	var resources []string

	// Ищем CSS файлы
	parts := strings.Split(html, "href=\"")
	for i := 1; i < len(parts); i++ {
		end := strings.Index(parts[i], "\"")
		if end != -1 {
			link := parts[i][:end]
			if strings.HasSuffix(link, ".css") && strings.HasPrefix(link, "http") {
				resources = append(resources, link)
			}
		}
	}

	// Ищем JS файлы
	parts = strings.Split(html, "src=\"")
	for i := 1; i < len(parts); i++ {
		end := strings.Index(parts[i], "\"")
		if end != -1 {
			link := parts[i][:end]
			if strings.HasSuffix(link, ".js") && strings.HasPrefix(link, "http") {
				resources = append(resources, link)
			}
		}
	}

	// Ищем изображения
	parts = strings.Split(html, "src=\"")
	for i := 1; i < len(parts); i++ {
		end := strings.Index(parts[i], "\"")
		if end != -1 {
			link := parts[i][:end]
			if (strings.HasSuffix(link, ".jpg") || strings.HasSuffix(link, ".png") || strings.HasSuffix(link, ".gif")) && strings.HasPrefix(link, "http") {
				resources = append(resources, link)
			}
		}
	}

	return resources
}
