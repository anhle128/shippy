build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/anhle128/shippy/consignment-service \
		proto/consignment/consignment.proto
	docker build -t consignment-service .

run: 
    docker rundocke -p 50051:50051 consignment-service.
