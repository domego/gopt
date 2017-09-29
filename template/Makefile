NAME=`cat NAME`
VERSION=`cat VERSION`
COMMIT=`git rev-parse HEAD`
ARCH=`uname | tr '[:upper:]' '[:lower:]'`
CURRENT=$(shell git rev-parse --abbrev-ref HEAD)
BRANCH_STR=`git rev-parse --abbrev-ref HEAD|sed "s/\//_/g"`
OS_NAME=$(shell echo $(shell uname) | tr '[A-Z]' '[a-z]')

.PHONY: build

build:
	@echo "run build"; \
	go build -o bin/$(NAME)

xbuild:
	@echo "run xbuild"; \
	GOOS=linux GOARCH=amd64 go build -o bin/$(NAME)

fresh:
	@fresh -c fresh.conf

# add echo to disable make error
rebase:
	git version
	git stash
	git checkout develop
	git pull origin develop
	git checkout $(CURRENT)
	git rebase develop
	git stash pop | echo
