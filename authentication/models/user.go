package models

import (
	"time"

	"github.com/mtnmunuklu/bavul/pb"
	"github.com/mtnmunuklu/bavul/vulnerability/util"

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
		Created: util.FormatTime(u.Created),
		Updated: util.FormatTime(u.Updated),
	}
}

// FromProto gets user from protocol buffer and converts to the user structure.
func (u *User) FromProto(user *pb.User) {
	u.Id = bson.ObjectIdHex(user.GetId())
	u.Name = user.GetName()
	u.Email = user.GetEmail()
	u.Role = user.GetRole()
	u.Created = util.ParseTime(user.GetCreated())
	u.Updated = util.ParseTime(user.GetUpdated())
}
