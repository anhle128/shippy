build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/anhle128/shippy/consignment-service \
		proto/consignment/consignment.proto
	CGO_ENABLED=0 GOOS=linux go build  -o consignment-service -a -installsuffix cgo main.go
	docker build -t consignment-service .

run:
	docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns -e DB_HOST=mongodb://root:123456789@ds147659.mlab.com:47659/shippy consignment-service
