package models

import (
	"testing"

	pb "github.com/danvalencia/golang/addressbook/models"
)

func TestShouldSerializeObject(t *testing.T) {
	person := &pb.Person{
		Id:    1234,
		Name:  "Daniel Valencia",
		Email: "danvalencia@gmail.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "331-347-3493", Type: pb.Person_HOME},
		},
	}

	if person == nil {
		t.Errorf("person should not be nil")
	}
}
