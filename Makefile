# build an executable file named simpledependency

.PHONY: build run test install cleanup vendor vendor-dev

default: build

build: 
	go build -o ./bin/smpldep ./cmd/smpldep

run:
	go run ./cmd/smpldep/main.go

test: vendor


install:


cleanup:
	rm -dRf ${PWD}/vendor


vendor: cleanup
	glide install

vendor-dev:
