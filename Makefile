PACKAGE_LIST := $(shell go list ./...)
VERSION := 0.1.16
NAME := cwf
DIST := $(NAME)-$(VERSION)

cwf: coverage.out
	go build -o cwf $(PACAKAGE_LIST)

test:
	go test -covermode=count -coverprofile=coverage.out $(PACKAGE_LIST)
	
clean:
	rm -f cwf coverage.out
