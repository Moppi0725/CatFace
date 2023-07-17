PACKAGE_LIST := $(shell go list ./...)
VERSION := 0.1.16
NAME := cwf
DIST := $(NAME)-$(VERSION)

cwf: coverage.out # cmd/main.go *.go
	go build -o $(NAME) cmd/main.go cmd/generate_completion.go
# go build -o $(NAME) cmd/main.go

coverage.out: #cmd/main_test.go
	go test -covermode=count \
		-coverprofile=coverage.out $(PACKAGE_LIST)

docker: cwf
#	docker build -t ghcr.io/Moppi0725/cwf:$(VERSION) -t ghcr.io/Moppi0725/cwf:latest .
	docker buildx build -t ghcr.io/Moppi0725/cwf:$(VERSION) \
		-t ghcr.io/Moppi0725/cwf:latest --platform=linux/arm64/v8,linux/amd64 --push .

# refer from https://pod.hatenablog.com/entry/2017/06/13/150342
define _createDist
	mkdir -p dist/$(1)_$(2)/$(DIST)
		GOOS=$1 GOARCH=$2 go build -o dist/$(1)_$(2)/$(DIST)/$(NAME)$(3) cmd/main.go cmd/generate_completion.go
	cp -r README.md LICENSE dist/$(1)_$(2)/$(DIST)
	cp -r cmd/completions dist/$(1)_$(2)/$(DIST)
	# cp -r docs/public dist/$(1)_$(2)/$(DIST)/docs
	tar cfz dist/$(DIST)_$(1)_$(2).tar.gz -C dist/$(1)_$(2) $(DIST)
endef

dist: cwf
	@$(call _createDist,darwin,amd64,)
	@$(call _createDist,darwin,arm64,)
	@$(call _createDist,windows,amd64,.exe)
	@$(call _createDist,windows,arm64,.exe)
	@$(call _createDist,linux,amd64,)
	@$(call _createDist,linux,arm64,)

distclean: clean
	rm -rf dist

clean:
	rm -f cwf coverage.out
	rm -rf completions cmd/completions
