build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/anhle128/shippy/user-service \
		proto/user/user.proto
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo
	docker build -t user-service .

run:
	docker run -p 50053:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns -e DB_HOST=host.docker.internal  -e DB_USER=root -e DB_NAME=shoppy -e DB_PASSWORD=123456789 user-service
