NAME=gopt
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
	gopt gen_gin_server -name=testGinServer -port=9100; \
	make; bin/testGinServer

test_gen_types: install
	@echo "test gen_types"; \
	rm -rf test; mkdir test; cd test; \
	gopt gen_types -types=../types.yaml;

test_gen_orm: install
	@echo "test gen_orm"; \
	rm -rf test; mkdir test; cd test; \
	gopt gen_orm -orm=../db.yaml;

test_gen_gin_controller: install
	@echo "test gen_gin_controller"; \
	rm -rf test; mkdir test; cd test; \
	gopt gen_gin_api -api=../api.yaml;

test_gen_js_api: install
	@echo "test gen_js_api"; \
	rm -rf test; mkdir test; cd test; \
	gopt gen_js_api -api=../api.yaml;



# add echo to disable make error
rebase:
	git version
	git stash
	git checkout develop
	git pull origin develop
	git checkout $(CURRENT)
	git rebase develop
	git stash pop | echo
