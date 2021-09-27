.PHONY = server-no-handler server-with-handler deps

deps:
	go get -v ./...

server-no-handler: deps
	go run ./cmd/no-handler

server-with-handler: deps
	go run ./cmd/with-handler

request:
	curl localhost:8888
