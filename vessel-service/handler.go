package main

import (
	"context"

	pb "github.com/anhle128/shippy/vessel-service/proto/vessel"
)

// Our grpc service handler
type service struct {
	repo IRepository
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {

	// Find the next available vessel
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}

	// Set the vessel as part of the response message type
	res.Vessel = vessel
	return nil
}

func (s *service) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	_, err := s.repo.Create(req)
	return err
}
