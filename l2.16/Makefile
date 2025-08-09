.PHONY: build clean test demo-html demo-css demo-js demo-images demo-links demo-all

build:
	go build -o wget .

clean:
	rm -f wget *.html *.css *.js *.png *.jpg *.gif

demo-html:
	@echo Демонстрация загрузки HTML страниц
	@echo "Загружаем главную страницу Wildberries..."
	go run . -url=https://www.wildberries.ru -depth=1
	@echo "Загружаем главную страницу VK"
	go run . -url=https://vk.com -depth=1
	@echo "HTML страницы загружены!"

demo-css:
	@echo Демонстрация загрузки CSS файлов"
	@echo "Загружаем сайт с CSS ресурсами"
	go run . -url=https://gismeteo.ru -depth=2
	@echo "CSS файлы загружены!"

demo-js:
	@echo "Демонстрация загрузки JavaScript файлов"
	@echo "Загружаем сайт с JS ресурсами"
	go run . -url=https://mail.ru -depth=2
	@echo "JavaScript файлы загружены!"

demo-images:
	@echo "Демонстрация загрузки изображений"
	@echo "Загружаем сайт с изображениями"
	go run . -url=https://lenta.ru -depth=2
	@echo "Изображения загружены!"

demo-links:
	@echo "Демонстрация загрузки по ссылкам"
	@echo "Загружаем сайт с рекурсивными ссылками"
	go run . -url=https://rbc.ru -depth=3
	@echo "Ссылки обработаны!"