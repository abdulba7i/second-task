GREP = go run ./cmd/main.go
DATADIR = data

# простой поиск слова 'go'
case1:
	$(GREP) -F "go" $(DATADIR)/input.txt

# игнорирование регистров и вывод номеров строк
case2:
	$(GREP) -i -n "go" $(DATADIR)/input.txt

# инвертированный поиск без учёта регистра
case3:
	$(GREP) -i -v "go" $(DATADIR)/input.txt

# поиск контекста вокруг строки
case4:
	$(GREP) -C 1 "go" $(DATADIR)/input.txt

# счётчик совпадений
case5:
	$(GREP) -c "go" $(DATADIR)/input.txt

# полное совпадение строки
case6:
	$(GREP) -F "go is awesome." $(DATADIR)/input.txt

# контекст после строки
case7:
	$(GREP) -A 2 "go" $(DATADIR)/input.txt

test: case1 case2 case3 case4 case5 case6 case7
