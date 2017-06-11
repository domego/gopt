NAME=gpt
COMMIT=`git rev-parse HEAD`
ARCH=`uname | tr '[:upper:]' '[:lower:]'`
CURRENT=$(shell git rev-parse --abbrev-ref HEAD)
BRANCH_STR=`git rev-parse --abbrev-ref HEAD|sed "s/\//_/g"`
OS_NAME=$(shell echo $(shell uname) | tr '[A-Z]' '[a-z]')

all: install

bindata:
	@cd template; go-bindata -o ../bindata.go ./...

install: bindata
	@go install

build: bindata
	@go build

test_gen_gin_server: install
	@echo "test gen_gin_server"; \
	rm -rf test; mkdir test; cd test; \
	ginpt gen_gin_server -name=testGinServer -port=9100; \
	cd ../

# add echo to disable make error
rebase:
	git version
	git stash
	git checkout develop
	git pull origin develop
	git checkout $(CURRENT)
	git rebase develop
	git stash pop | echo
