package repository

import (
	"github.com/mtnmunuklu/bavul/authentication/models"
	"github.com/mtnmunuklu/bavul/db"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const UserCollection = "users"

// UserRepository is the interface of the authentication backend.
type UserRepository interface {
	Save(user *models.User) error
	GetById(id string) (user *models.User, err error)
	GetByEmail(email string) (user *models.User, err error)
	GetAll() (user []*models.User, err error)
	Update(user *models.User) error
	DeleteById(id string) error
}

// userRepository provides a mongo collection for database job.
type userRepository struct {
	c *mgo.Collection
}

// NewUserRepository creates a new UserRepository instance.
func NewUserRepository(conn db.Connection) UserRepository {
	return &userRepository{c: conn.DB().C(UserCollection)}
}

// Save creates a user.
func (r *userRepository) Save(user *models.User) error {
	return r.c.Insert(user)
}

// GetById returns the user based on id.
func (r *userRepository) GetById(id string) (user *models.User, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// GetByEmail returns the user based on email.
func (r *userRepository) GetByEmail(email string) (user *models.User, err error) {
	err = r.c.Find(bson.M{"email": email}).One(&user)
	return user, err
}

// GetAll returns all users.
func (r *userRepository) GetAll() (user []*models.User, err error) {
	err = r.c.Find(bson.M{}).All(&user)
	return user, err
}

// Update updates the user.
func (r *userRepository) Update(user *models.User) error {
	return r.c.UpdateId(user.Id, user)
}

// Delete deletes the user based on id.
func (r *userRepository) DeleteById(id string) error {
	return r.c.RemoveId(bson.ObjectIdHex(id))
}

// DeleteAll drops users collection.
func (r *userRepository) DeleteAll() error {
	return r.c.DropCollection()
}
