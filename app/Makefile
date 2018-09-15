#!/bin/bash

export NOW=$(shell date --rfc-3339=ns)
export PKGS=$(shell go list ./... | grep -v vendor/ )

all: build run
quick: build run
docker-run: build wait run

test:
	@echo "${NOW} TESTING..."
	@go test -cover -race ${PKGS} | tee test.out
	@./test.sh test.out 15

build:
	@echo "${NOW} ===BUILDING API==="
	@go build -o shopeetax ./

run:
	@echo "${NOW} RUNNING..."
	@./shopeetax & wait

wait:
	@echo "Wait till mysql ready for 15 seconds..."
	@sleep 10

update:
	@echo "${NOW} UPDATING..."
	@dep ensure -v

clean:
	@echo "${NOW} CLEANING..."
	@rm -rf vendor/ 2>/dev/null || true

