GO=go

DEP_LIST:=$(shell ls ./deployments)

.ONESHELL:
build-dep:
	@for var in $(DEP_LIST); do
		cd ./deployments/$$var
			make build
		cd ./../..
	done

.ONESHELL:
build-pkg:
	cd ./pkg/greetings
		go build .
	cd ./../..

build: build-dep build-pkg

.ONESHELL:
test-dep:
	@for var in $(DEP_LIST); do
		cd ./deployments/$$var
			make test
		cd ./../..
	done

.ONESHELL:
test-pkg:
	cd ./pkg/greetings
		go test -v ./...
	cd ./../..
test: test-dep test-pkg
