package models

import (
	"time"

	"github.com/mtnmunuklu/bavul/pb"

	"gopkg.in/mgo.v2/bson"
)

// User provides the user instance for authentication job.
type User struct {
	Id       bson.ObjectId `bson:"_id"`
	Name     string        `bson:"name"`
	Email    string        `bson:"email"`
	Password string        `bson:"password"`
	Role     string        `bson:"role"`
	Created  time.Time     `bson:"created"`
	Updated  time.Time     `bson:"updated"`
}

// ToProto converts the user structure into a protocol buffer user structure.
func (u *User) ToProto() *pb.User {
	return &pb.User{
		Id:      u.Id.Hex(),
		Name:    u.Name,
		Email:   u.Email,
		Role:    u.Role,
		Created: u.Created.Unix(),
		Updated: u.Updated.Unix(),
	}
}

// FromProto gets user from protocol buffer and converts to the user structure.
func (u *User) FromProto(user *pb.User) {
	u.Id = bson.ObjectIdHex(user.GetId())
	u.Name = user.GetName()
	u.Email = user.GetEmail()
	u.Role = user.GetRole()
	u.Created = time.Unix(user.GetCreated(), 0)
	u.Updated = time.Unix(user.GetUpdated(), 0)
}
