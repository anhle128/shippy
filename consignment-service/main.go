// consignment-service/main.go
package main

import (

	// Import the generated protobuf code
	"fmt"
	"log"

	pb "github.com/anhle128/shippy/consignment-service/proto/consignment"
	vesselProto "github.com/anhle128/shippy/vessel-service/proto/vessel"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

const (
	defaultHost = "localhost:27017"
)

func main() {

	host := defaultHost

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name:   "mongo_db",
				EnvVar: "MONGO_DB",
				Usage:  "mongodb url",
			},
		),
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	// Init will parse the command line flags.
	srv.Init(
		micro.Action(func(c *cli.Context) {
			host = c.String("mongo_db")
		}),
	)
	session, err := CreateSession(host)
	if err != nil {

		// We're wrapping the error returned from our CreateSession
		// here to add some context to the error.
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}
	defer session.Clone()
	// Register handler
	pb.RegisterShippingServiceHandler(srv.Server(), &handler{vesselClient, session})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
