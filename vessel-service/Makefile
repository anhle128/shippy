build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/anhle128/shippy/vessel-service \
    proto/vessel/vessel.proto
	docker build -t vessel-service .

run:
	docker run --name=vessel-service \ 
		-p 50052:50051 \ 
		-e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns \ 
		-e DB_HOST=mongodb://root:123456789@ds147659.mlab.com:47659/shippy \
		vessel-service
