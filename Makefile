VERSION := $(shell git describe --always --tags --dirty)
GIT_COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null)
GIT_DATE := $(shell git log -1 --format=%cI)

ALL_FILES := $(shell find . -type f -name '*.go')
BINARY := md-slides

# make the main binary
.PHONY: dev
dev: $(BINARY)

# tasks for creating dist directory
dist/:
	mkdir dist

# tasks for creating binaries
DISTRIBUTABLES = $(BINARY) dist/$(BINARY).linux.amd64 dist/$(BINARY).darwin.amd64 dist/$(BINARY).windows.amd64.exe
GOOS = $(shell echo $@ | grep -oP "$(BINARY).\K\w+(?=\.)")
GOARCH = $(shell echo $@ | grep -oP "$(BINARY).$(GOOS).\K\w+(?=\.|$$)")

# official distributable version
$(DISTRIBUTABLES): dist/ $(ALL_FILES)
	@echo Building $@..
	@CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build \
		-o $@ \
		-ldflags "-X main.gitHash=$(GIT_COMMIT) -X main.gitDate=$(GIT_DATE) -X main.version=$(VERSION)"

# shasums file for dist
dist/SHA256SUMS: $(DISTRIBUTABLES)
	cd dist && sha256sum md-slides.* > SHA256SUMS

pages/:
	mkdir pages

pages/index.html: pages/ $(BINARY) README.md
	@./$(BINARY) serve --export-to=$@ README.md
	@cp -v windmill.jpeg pages/

# release build
.PHONY: release
release: dist/SHA256SUMS $(DISTRIBUTABLES)

# clean will just remove the main bits and pieces
.PHONY: clean
clean:
	@rm -rfv $(DISTRIBUTABLES)
	@rm -rfv dist
	@rm -rfv pages/index.html