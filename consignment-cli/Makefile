build-docker:
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-cli .

run-docer:
	docker run -e MICRO_REGISTRY=mdns consignment-cli

run:
	go run cli.go --registry=mdns
