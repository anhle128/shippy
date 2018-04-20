package main_test

import (
	"testing"

	cs "github.com/anhle128/shippy/consignment-service"
	pb "github.com/anhle128/shippy/consignment-service/proto/consignment"
)

func TestCreate(t *testing.T) {
	db, err := cs.CreateSession("mongodb://root:123456789@ds147659.mlab.com:47659/shippy")
	if err != nil {
		t.Error(err)
		return
	}

	repository := cs.InitConsignmentRepository(db)
	err = repository.Create(&pb.Consignment{
		Id:          "aslkdjflasjdlfk",
		Description: "mot con vit ahihihihi",
		Weight:      100,
		VesselId:    "sdfjaoij",
	})
	if err != nil {
		t.Error(err)
		return
	}
}
