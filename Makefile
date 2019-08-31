GO=go
NAME := qrg
VERSION := 1.0.0

all: test build

deps:

setup: deps update_version

update_version:
	@for i in README.md ; do\
	    sed -e 's!Version-[0-9.]*-yellowgreen!Version-${VERSION}-yellowgreen!g' -e 's!tag/v[0-9.]*!tag/v${VERSION}!g' $$i > a ; mv a $$i; \
	done
	@sed 's/const VERSION = .*/const VERSION = "${VERSION}"/g' main.go > a
	@mv a main.go
	@echo "Replace version to \"${VERSION}\""

test: setup
	$(GO) test -covermode=count -coverprofile=coverage.out $$(go list ./...)

build: setup
	$(GO) build -o $(NAME) -v

clean:
	$(GO) clean
	rm -rf $(NAME) qrcode.png
