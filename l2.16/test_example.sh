#!/bin/bash

echo "Тест wget"
echo "=========="

echo "Загружем страницу:"
go run main.go -url=https://httpbin.org/html -depth=1

echo "Готово!" 