package repository_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/mtnmunuklu/bavul/authentication/models"
	"github.com/mtnmunuklu/bavul/authentication/repository"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
)

// MockUserRepository is a mock implementation of the repository.UserRepository interface.
type MockUserRepository struct {
	Users map[string]*models.User
}

var _ repository.UserRepository = &MockUserRepository{}

func (muc *MockUserRepository) Save(user *models.User) error {
	muc.Users[user.Id.Hex()] = user
	return nil
}

func (muc *MockUserRepository) GetById(id string) (*models.User, error) {
	user, ok := muc.Users[id]
	if !ok {
		return nil, errors.New("User not found")
	}
	return user, nil
}

func (muc *MockUserRepository) GetByEmail(email string) (*models.User, error) {
	for _, user := range muc.Users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("User not found")
}

func (muc *MockUserRepository) GetAll() ([]*models.User, error) {
	users := make([]*models.User, 0, len(muc.Users))
	for _, user := range muc.Users {
		users = append(users, user)
	}
	return users, nil
}

func (muc *MockUserRepository) Update(user *models.User) error {
	if _, ok := muc.Users[user.Id.Hex()]; !ok {
		return errors.New("User not found")
	}
	muc.Users[user.Id.Hex()] = user
	return nil
}

func (muc *MockUserRepository) DeleteById(id string) error {
	delete(muc.Users, id)
	return nil
}

// TestMockUsersRepositorySave tests the user create operation with a mock repository.
func TestMockUsersRepositorySave(t *testing.T) {
	mockRepo := &MockUserRepository{Users: make(map[string]*models.User)}
	id := bson.NewObjectId()

	user := &models.User{
		Id:       id,
		Name:     "TEST",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "123456789",
		Role:     "user",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	err := mockRepo.Save(user)
	assert.NoError(t, err)

	found, err := mockRepo.GetById(user.Id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, user, found)
}

// TestMockUsersRepositoryGetById tests the operation to return user based on id with a mock repository.
func TestMockUsersRepositoryGetById(t *testing.T) {
	mockRepo := &MockUserRepository{Users: make(map[string]*models.User)}
	id := bson.NewObjectId()

	user := &models.User{
		Id:       id,
		Name:     "TEST",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "123456789",
		Role:     "user",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	mockRepo.Save(user)

	found, err := mockRepo.GetById(user.Id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, user, found)

	notFound, err := mockRepo.GetById(bson.NewObjectId().Hex())
	assert.Error(t, err)
	assert.Nil(t, notFound)
}

// TestMockUsersRepositoryGetByEmail tests the operation to return user based on email with a mock repository.
func TestMockUsersRepositoryGetByEmail(t *testing.T) {
	mockRepo := &MockUserRepository{Users: make(map[string]*models.User)}
	id := bson.NewObjectId()

	user := &models.User{
		Id:       id,
		Name:     "TEST",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "123456789",
		Role:     "user",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	mockRepo.Save(user)

	found, err := mockRepo.GetByEmail(user.Email)
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, user, found)

	notFound, err := mockRepo.GetByEmail("nonexistent@email.test")
	assert.Error(t, err)
	assert.Nil(t, notFound)
}

// TestMockUsersRepositoryUpdate tests the user update operation with a mock repository.
func TestMockUsersRepositoryUpdate(t *testing.T) {
	mockRepo := &MockUserRepository{Users: make(map[string]*models.User)}
	id := bson.NewObjectId()

	user := &models.User{
		Id:       id,
		Name:     "TEST",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "123456789",
		Role:     "user",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	mockRepo.Save(user)

	updatedUser := &models.User{
		Id:       id,
		Name:     "UpdatedTest",
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
		Created:  user.Created,
		Updated:  time.Now(),
	}

	err := mockRepo.Update(updatedUser)
	assert.NoError(t, err)

	found, err := mockRepo.GetById(id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, updatedUser, found)
}

// TestMockUsersRepositoryDelete tests the user delete operation with a mock repository.
func TestMockUsersRepositoryDelete(t *testing.T) {
	mockRepo := &MockUserRepository{Users: make(map[string]*models.User)}
	id := bson.NewObjectId()

	user := &models.User{
		Id:       id,
		Name:     "TEST",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "123456789",
		Role:     "user",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	mockRepo.Save(user)

	err := mockRepo.DeleteById(id.Hex())
	assert.NoError(t, err)

	notFound, err := mockRepo.GetById(id.Hex())
	assert.Error(t, err)
	assert.Nil(t, notFound)
}

// TestMockUsersRepositoryGetAll tests the operation to return all users with a mock repository.
func TestMockUsersRepositoryGetAll(t *testing.T) {
	mockRepo := &MockUserRepository{Users: make(map[string]*models.User)}
	id := bson.NewObjectId()

	user := &models.User{
		Id:       id,
		Name:     "TEST",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "123456789",
		Role:     "user",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	mockRepo.Save(user)

	users, err := mockRepo.GetAll()
	assert.NoError(t, err)
	assert.NotEmpty(t, users)
	assert.Equal(t, []*models.User{user}, users)
}
