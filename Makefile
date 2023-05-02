GO=go
SHELL=/bin/bash
NAME := cwf
VERSION := 1.2.0
DIST := $(NAME)-$(VERSION)
PACAKAGE_LIST := $(shell go list ./...)
cwf:
	go build -o cwf $(PACAKAGE_LIST)

test:
	go test $(PACAKAGE_LIST) -covermode=count -coverprofile=coverage.out $(PACKAGE_LIST)
	
clean:
	rm -f cwf
