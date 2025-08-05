SORT = go run ./cmd/main.go
DATADIR = internal/data

case1:
	$(SORT) $(DATADIR)/data.txt -k 1

case2:
	$(SORT) $(DATADIR)/data.txt -k 2 -n

case3:
	$(SORT) $(DATADIR)/data.txt -k 3 -M

case4:
	$(SORT) $(DATADIR)/data.txt -k 4 -h

case5:
	$(SORT) $(DATADIR)/data.txt -k 1 -r

case6:
	$(SORT) $(DATADIR)/data.txt -k 1 -u

case7:
	$(SORT) $(DATADIR)/data.txt -k 2 -b

test: case1 case2 case3 case4 case5 case6 case7 