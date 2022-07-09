GO=go

PKG_LIST:=$(shell ls ./pkg)
APP_LIST:=$(shell ls ./app)

test:
	@for var in $(PKG_LIST);do cd ./pkg/$$var && go test && cd ./../.. ; done

build:
	@for var in $(APP_LIST);do cd ./app/$$var && go build . && cd ./../.. ; done