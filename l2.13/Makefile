GOCMD = go run ./cmd/main.go
DATADIR = data

case1:
	$(GOCMD) -f "1" -d "," $(DATADIR)/input.txt

case2:
	$(GOCMD) -f "1,3-4" -d "," $(DATADIR)/input.txt

case3:
	$(GOCMD) -f "2" -s $(DATADIR)/input.txt

case4:
	$(GOCMD) -f "1-3" $(DATADIR)/input.txt

test:
	go test ./... -v
