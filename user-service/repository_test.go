package main_test

import (
	"testing"

	us "github.com/anhle128/shippy/user-service"
	pb "github.com/anhle128/shippy/user-service/proto/user"
)

func TestCreate(t *testing.T) {
	db, err := us.CreateConnection("localhost", "root", "123456789", "shoppy")
	if err != nil {
		t.Error(err)
		return
	}

	repo := us.InitUserRepository(db)

	err = repo.Create(&pb.User{
		Name:     "mot con vit",
		Email:    "anhle12892@gmail.com",
		Password: "ohlalalalala",
		Company:  "TTAB",
	})
	if err != nil {
		t.Error(err)
	}
}
