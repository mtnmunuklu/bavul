package handlers_test

import (
	"context"
	"encoding/json"
	"io"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mtnmunuklu/bavul/api/handlers"
	"github.com/mtnmunuklu/bavul/api/util"
	authUtil "github.com/mtnmunuklu/bavul/authentication/util"
	"github.com/mtnmunuklu/bavul/pb"
	"github.com/mtnmunuklu/bavul/security"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"gopkg.in/mgo.v2/bson"
)

// Custom mock client that implements the pb.AuthServiceClient interface
type MockAuthServiceClientWrapper struct {
	ChangeUserRoleFunc           func(ctx context.Context, req *pb.ChangeUserRoleRequest, opts ...grpc.CallOption) (*pb.User, error)
	GetUserRoleFunc              func(ctx context.Context, req *pb.GetUserRoleRequest, opts ...grpc.CallOption) (*pb.GetUserRoleResponse, error)
	ListUsersFunc                func(ctx context.Context, req *pb.ListUsersRequest, opts ...grpc.CallOption) (pb.AuthService_ListUsersClient, error)
	UpdateUserPasswordFunc       func(ctx context.Context, req *pb.UpdateUserPasswordRequest, opts ...grpc.CallOption) (*pb.User, error)
	UpdateUserEmailFunc          func(ctx context.Context, req *pb.UpdateUserEmailRequest, opts ...grpc.CallOption) (*pb.User, error)
	UpdateUserNameFunc           func(ctx context.Context, req *pb.UpdateUserNameRequest, opts ...grpc.CallOption) (*pb.User, error)
	SignUpFunc                   func(ctx context.Context, req *pb.SignUpRequest, opts ...grpc.CallOption) (*pb.User, error)
	SignInFunc                   func(ctx context.Context, req *pb.SignInRequest, opts ...grpc.CallOption) (*pb.SignInResponse, error)
	DeleteUserFunc               func(ctx context.Context, req *pb.DeleteUserRequest, opts ...grpc.CallOption) (*pb.DeleteUserResponse, error)
	GetUserFunc                  func(ctx context.Context, req *pb.GetUserRequest, opts ...grpc.CallOption) (*pb.User, error)
	ChangeUserRoleFuncCalled     bool
	GetUserRoleFuncCalled        bool
	ListUsersFuncCalled          bool
	UpdateUserPasswordFuncCalled bool
	UpdateUserEmailFuncCalled    bool
	UpdateUserNameFuncCalled     bool
	SignUpFuncCalled             bool
	SignInFuncCalled             bool
	DeleteUserFuncCalled         bool
	GetUserFuncCalled            bool
}

func (c *MockAuthServiceClientWrapper) ChangeUserRole(ctx context.Context, req *pb.ChangeUserRoleRequest, opts ...grpc.CallOption) (*pb.User, error) {
	c.ChangeUserRoleFuncCalled = true
	if c.ChangeUserRoleFunc != nil {
		return c.ChangeUserRoleFunc(ctx, req, opts...)
	}
	return nil, nil
}

func (c *MockAuthServiceClientWrapper) GetUserRole(ctx context.Context, req *pb.GetUserRoleRequest, opts ...grpc.CallOption) (*pb.GetUserRoleResponse, error) {
	c.GetUserRoleFuncCalled = true
	if c.GetUserRoleFunc != nil {
		return c.GetUserRoleFunc(ctx, req, opts...)
	}
	return nil, nil
}

func (c *MockAuthServiceClientWrapper) ListUsers(ctx context.Context, req *pb.ListUsersRequest, opts ...grpc.CallOption) (pb.AuthService_ListUsersClient, error) {
	c.ListUsersFuncCalled = true
	if c.ListUsersFunc != nil {
		return c.ListUsersFunc(ctx, req, opts...)
	}
	return nil, nil
}

func (c *MockAuthServiceClientWrapper) UpdateUserPassword(ctx context.Context, req *pb.UpdateUserPasswordRequest, opts ...grpc.CallOption) (*pb.User, error) {
	c.UpdateUserPasswordFuncCalled = true
	if c.UpdateUserPasswordFunc != nil {
		return c.UpdateUserPasswordFunc(ctx, req, opts...)
	}
	return nil, nil
}

func (c *MockAuthServiceClientWrapper) UpdateUserEmail(ctx context.Context, req *pb.UpdateUserEmailRequest, opts ...grpc.CallOption) (*pb.User, error) {
	c.UpdateUserEmailFuncCalled = true
	if c.UpdateUserEmailFunc != nil {
		return c.UpdateUserEmailFunc(ctx, req, opts...)
	}
	return nil, nil
}

func (c *MockAuthServiceClientWrapper) UpdateUserName(ctx context.Context, req *pb.UpdateUserNameRequest, opts ...grpc.CallOption) (*pb.User, error) {
	c.UpdateUserNameFuncCalled = true
	if c.UpdateUserNameFunc != nil {
		return c.UpdateUserNameFunc(ctx, req, opts...)
	}
	return nil, nil
}

func (c *MockAuthServiceClientWrapper) SignUp(ctx context.Context, req *pb.SignUpRequest, opts ...grpc.CallOption) (*pb.User, error) {
	c.SignUpFuncCalled = true
	if c.SignUpFunc != nil {
		return c.SignUpFunc(ctx, req, opts...)
	}
	return nil, nil
}

func (c *MockAuthServiceClientWrapper) SignIn(ctx context.Context, req *pb.SignInRequest, opts ...grpc.CallOption) (*pb.SignInResponse, error) {
	c.SignInFuncCalled = true
	if c.SignInFunc != nil {
		return c.SignInFunc(ctx, req, opts...)
	}
	return nil, nil
}

func (c *MockAuthServiceClientWrapper) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest, opts ...grpc.CallOption) (*pb.DeleteUserResponse, error) {
	c.DeleteUserFuncCalled = true
	if c.DeleteUserFunc != nil {
		return c.DeleteUserFunc(ctx, req, opts...)
	}
	return nil, nil
}

func (c *MockAuthServiceClientWrapper) GetUser(ctx context.Context, req *pb.GetUserRequest, opts ...grpc.CallOption) (*pb.User, error) {
	c.GetUserFuncCalled = true
	if c.GetUserFunc != nil {
		return c.GetUserFunc(ctx, req, opts...)
	}
	return nil, nil
}

type MockAuthService_StreamClientWrapper struct {
	users []*pb.User
	idx   int
}

// Recv simulates receiving Users from the server.
func (m *MockAuthService_StreamClientWrapper) Recv() (*pb.User, error) {
	if m.idx < len(m.users) {
		user := m.users[m.idx]
		m.idx++
		return user, nil
	}
	return nil, io.EOF
}

// CloseSend simulates closing the send stream.
func (m *MockAuthService_StreamClientWrapper) CloseSend() error {
	return nil
}

// Context returns the client context.
func (m *MockAuthService_StreamClientWrapper) Context() context.Context {
	return context.Background()
}

// SendMsg simulates sending a message.
func (m *MockAuthService_StreamClientWrapper) SendMsg(m1 interface{}) error {
	return nil
}

// RecvMsg simulates receiving a message.
func (m *MockAuthService_StreamClientWrapper) RecvMsg(m1 interface{}) error {
	return nil
}

// Header returns the header metadata.
func (m *MockAuthService_StreamClientWrapper) Header() (metadata.MD, error) {
	return nil, nil
}

// Trailer returns the trailer metadata.
func (m *MockAuthService_StreamClientWrapper) Trailer() metadata.MD {
	return nil
}

// SendHeader simulates sending header metadata.
func (m *MockAuthService_StreamClientWrapper) SendHeader(m1 metadata.MD) error {
	return nil
}

// SetSendDeadline simulates setting the send deadline.
func (m *MockAuthService_StreamClientWrapper) SetSendDeadline(m1 time.Time) error {
	return nil
}

// SetRecvDeadline simulates setting the receive deadline.
func (m *MockAuthService_StreamClientWrapper) SetRecvDeadline(m1 time.Time) error {
	return nil
}

func TestSignUp(t *testing.T) {
	// Create a custom mock client wrapper for Auth Service
	mockAuthWrapper := &MockAuthServiceClientWrapper{}

	// Create handlers using the custom mock client wrapper
	handler := handlers.NewAuthHandlers(mockAuthWrapper)

	// Set Auth Service Client in the mockWrapper
	mockAuthWrapper.SignUpFunc = func(ctx context.Context, req *pb.SignUpRequest, opts ...grpc.CallOption) (*pb.User, error) {
		// Simulate the behavior of the gRPC service
		return &pb.User{Id: "1", Name: req.GetName(), Email: req.GetEmail(), Role: "user", Created: "2024-02-02T18:18:00", Updated: "2024-02-02T18:18:00"}, nil
	}

	// Create a Fiber context
	app := fiber.New()
	fiberContext := app.AcquireCtx(&fasthttp.RequestCtx{})

	// Set the request body in the Fiber context
	request := &pb.SignUpRequest{
		Name:     "Test User1",
		Email:    "testemail@test.com.tr",
		Password: "test123",
	}
	body, err := json.Marshal(request)
	assert.NoError(t, err)

	// Set the content-type to JSON
	fiberContext.Request().SetBody(body)
	fiberContext.Request().Header.Set("Content-Type", "application/json")

	// Test the SignUp handler
	err = handler.SignUp(fiberContext)
	assert.NoError(t, err)

	// Assert that the SignUp functions were called with the expected parameters
	assert.True(t, mockAuthWrapper.SignUpFuncCalled, "SignUp function of mockWrapper should be called")

	// Release the Fiber context
	app.ReleaseCtx(fiberContext)
}

func TestSignIn(t *testing.T) {
	// Create a custom mock client wrapper for Auth Service
	mockAuthWrapper := &MockAuthServiceClientWrapper{}

	// Create handlers using the custom mock client wrapper
	handler := handlers.NewAuthHandlers(mockAuthWrapper)

	// Set Auth Service Client in the mockWrapper
	mockAuthWrapper.SignInFunc = func(ctx context.Context, req *pb.SignInRequest, opts ...grpc.CallOption) (*pb.SignInResponse, error) {
		// Simulate the behavior of the gRPC service
		user := &pb.User{Id: "1", Name: "Test User1", Email: req.GetEmail(), Role: "user", Created: "2024-02-02T18:18:00", Updated: "2024-02-02T18:18:00"}
		token, err := security.NewToken(user.Id)
		if err != nil {
			return nil, authUtil.ErrFailedSignIn
		}
		return &pb.SignInResponse{User: user, Token: token}, nil
	}

	// Create a Fiber context
	app := fiber.New()
	fiberContext := app.AcquireCtx(&fasthttp.RequestCtx{})

	// Set the request body in the Fiber context
	request := &pb.SignInRequest{
		Email:    "testemail@test.com.tr",
		Password: "test123",
	}
	body, err := json.Marshal(request)
	assert.NoError(t, err)

	// Set the content-type to JSON
	fiberContext.Request().SetBody(body)
	fiberContext.Request().Header.Set("Content-Type", "application/json")

	// Test the SignIn handler
	err = handler.SignIn(fiberContext)
	assert.NoError(t, err)

	// Assert that the SignIn functions were called with the expected parameters
	assert.True(t, mockAuthWrapper.SignInFuncCalled, "SignIn function of mockWrapper should be called")

	// Release the Fiber context
	app.ReleaseCtx(fiberContext)
}

func TestGetUser(t *testing.T) {
	// Create a custom mock client wrapper for Auth Service
	mockAuthWrapper := &MockAuthServiceClientWrapper{}

	// Create handlers using the custom mock client wrapper
	handler := handlers.NewAuthHandlers(mockAuthWrapper)

	// Set Auth Service Client in the mockWrapper
	mockAuthWrapper.GetUserRoleFunc = func(ctx context.Context, req *pb.GetUserRoleRequest, opts ...grpc.CallOption) (*pb.GetUserRoleResponse, error) {
		// Simulate the behavior of the gRPC service
		return &pb.GetUserRoleResponse{Role: "admin"}, nil
	}

	mockAuthWrapper.GetUserFunc = func(ctx context.Context, req *pb.GetUserRequest, opts ...grpc.CallOption) (*pb.User, error) {
		// Simulate the behavior of the gRPC service
		return &pb.User{Id: "1", Name: "Test User1", Email: req.GetEmail(), Role: "user", Created: "2024-02-02T18:18:00", Updated: "2024-02-02T18:18:00"}, nil
	}

	// Create a Fiber context
	app := fiber.New()
	fiberContext := app.AcquireCtx(&fasthttp.RequestCtx{})

	// Set the request headers in the Fiber context
	userId := bson.NewObjectId()
	token, err := security.NewToken(userId.Hex())
	assert.NoError(t, err)
	fiberContext.Request().Header.Set("Authorization", "Bearer "+token+"")
	fiberContext.Request().Header.Set("Email", "testemail1@test.com.tr")

	// Test the GetUser handler for the first time
	err = handler.GetUser(fiberContext)
	assert.NoError(t, err)

	// Assert that the GetUserRole and GetUser functions were called with the expected parameters
	assert.True(t, mockAuthWrapper.GetUserRoleFuncCalled, "GetUserRole function of mockWrapper should be called")
	assert.True(t, mockAuthWrapper.GetUserFuncCalled, "GetUser function of mockWrapper should be called")

	// Get the cached result for the first time
	cachedDataFirstTime, foundFirstTime := util.GetFromCache("GetUser:testemail1@test.com.tr")
	assert.True(t, foundFirstTime, "Result should be in cache for the first time")

	// Test the GetUser handler for the second time
	err = handler.GetUser(fiberContext)
	assert.NoError(t, err)

	// Assert that the GetUserRole and GetUser functions were called again (second time) with the expected parameters
	assert.True(t, mockAuthWrapper.GetUserRoleFuncCalled, "GetUserRole function of mockWrapper should be called again (second time)")
	assert.True(t, mockAuthWrapper.GetUserFuncCalled, "GetUser function of mockWrapper should be called again (second time)")

	// Get the cached result for the second time
	cachedDataSecondTime, foundSecondTime := util.GetFromCache("GetUser:testemail1@test.com.tr")
	assert.True(t, foundSecondTime, "Result should be in cache for the second time")

	// Assert that the cached results for the first and second times are the same
	assert.Equal(t, cachedDataFirstTime, cachedDataSecondTime, "Cached results for the first and second times should be the same")

	// Release the Fiber context
	app.ReleaseCtx(fiberContext)
}

func TestDeleteUser(t *testing.T) {
	// Create a custom mock client wrapper for Auth Service
	mockAuthWrapper := &MockAuthServiceClientWrapper{}

	// Create handlers using the custom mock client wrapper
	handler := handlers.NewAuthHandlers(mockAuthWrapper)

	// Set Auth Service Client in the mockWrapper
	mockAuthWrapper.GetUserRoleFunc = func(ctx context.Context, req *pb.GetUserRoleRequest, opts ...grpc.CallOption) (*pb.GetUserRoleResponse, error) {
		// Simulate the behavior of the gRPC service
		return &pb.GetUserRoleResponse{Role: "admin"}, nil
	}

	mockAuthWrapper.DeleteUserFunc = func(ctx context.Context, req *pb.DeleteUserRequest, opts ...grpc.CallOption) (*pb.DeleteUserResponse, error) {
		// Simulate the behavior of the gRPC service
		return &pb.DeleteUserResponse{Email: req.GetEmail()}, nil
	}

	// Create a Fiber context
	app := fiber.New()
	fiberContext := app.AcquireCtx(&fasthttp.RequestCtx{})

	// Set the request headers in the Fiber context
	userId := bson.NewObjectId()
	token, err := security.NewToken(userId.Hex())
	assert.NoError(t, err)
	fiberContext.Request().Header.Set("Authorization", "Bearer "+token+"")
	fiberContext.Request().Header.Set("Email", "testemail1@test.com.tr")

	// Test the DeleteUser handler
	err = handler.DeleteUser(fiberContext)
	assert.NoError(t, err)

	// Assert that the GetUserRole and DeleteUser functions were called with the expected parameters
	assert.True(t, mockAuthWrapper.GetUserRoleFuncCalled, "GetUserRole function of mockWrapper should be called")
	assert.True(t, mockAuthWrapper.DeleteUserFuncCalled, "DeleteUser function of mockWrapper should be called")

	// Release the Fiber context
	app.ReleaseCtx(fiberContext)
}

func TestChangeUserRole(t *testing.T) {
	// Create a custom mock client wrapper for Auth Service
	mockAuthWrapper := &MockAuthServiceClientWrapper{}

	// Create handlers using the custom mock client wrapper
	handler := handlers.NewAuthHandlers(mockAuthWrapper)

	// Set Auth Service Client in the mockWrapper
	mockAuthWrapper.GetUserRoleFunc = func(ctx context.Context, req *pb.GetUserRoleRequest, opts ...grpc.CallOption) (*pb.GetUserRoleResponse, error) {
		// Simulate the behavior of the gRPC service
		return &pb.GetUserRoleResponse{Role: "admin"}, nil
	}

	mockAuthWrapper.ChangeUserRoleFunc = func(ctx context.Context, req *pb.ChangeUserRoleRequest, opts ...grpc.CallOption) (*pb.User, error) {
		// Simulate the behavior of the gRPC service
		return &pb.User{Id: "1", Name: "Test User1", Email: req.GetEmail(), Role: req.GetRole(), Created: "2024-02-02T18:18:00", Updated: "2024-02-02T18:18:00"}, nil
	}

	// Create a Fiber context
	app := fiber.New()
	fiberContext := app.AcquireCtx(&fasthttp.RequestCtx{})

	// Set the request body in the Fiber context
	request := &pb.ChangeUserRoleRequest{
		Email: "testemail@test.com.tr",
		Role:  "admin",
	}
	body, err := json.Marshal(request)
	assert.NoError(t, err)

	// Set the content-type to JSON
	fiberContext.Request().SetBody(body)
	fiberContext.Request().Header.Set("Content-Type", "application/json")

	// Set the request headers in the Fiber context
	userId := bson.NewObjectId()
	token, err := security.NewToken(userId.Hex())
	assert.NoError(t, err)
	fiberContext.Request().Header.Set("Authorization", "Bearer "+token+"")

	// Test the ChangeUserRole handler
	err = handler.ChangeUserRole(fiberContext)
	assert.NoError(t, err)

	// Assert that the GetUserRole and ChangeUserRole functions were called with the expected parameters
	assert.True(t, mockAuthWrapper.GetUserRoleFuncCalled, "GetUserRole function of mockWrapper should be called")
	assert.True(t, mockAuthWrapper.ChangeUserRoleFuncCalled, "ChangeUserRole function of mockWrapper should be called")

	// Release the Fiber context
	app.ReleaseCtx(fiberContext)
}

func TestUpdateUserPassword(t *testing.T) {
	// Create a custom mock client wrapper for Auth Service
	mockAuthWrapper := &MockAuthServiceClientWrapper{}

	// Create handlers using the custom mock client wrapper
	handler := handlers.NewAuthHandlers(mockAuthWrapper)

	mockAuthWrapper.UpdateUserPasswordFunc = func(ctx context.Context, req *pb.UpdateUserPasswordRequest, opts ...grpc.CallOption) (*pb.User, error) {
		// Simulate the behavior of the gRPC service
		return &pb.User{Id: "1", Name: "Test User1", Email: req.GetEmail(), Role: "user", Created: "2024-02-02T18:18:00", Updated: "2024-02-02T18:18:00"}, nil
	}

	// Create a Fiber context
	app := fiber.New()
	fiberContext := app.AcquireCtx(&fasthttp.RequestCtx{})

	// Set the request body in the Fiber context
	request := &pb.UpdateUserPasswordRequest{
		Email:       "testemail@test.com.tr",
		Password:    "old-password",
		NewPassword: "new-paswword",
	}
	body, err := json.Marshal(request)
	assert.NoError(t, err)

	// Set the content-type to JSON
	fiberContext.Request().SetBody(body)
	fiberContext.Request().Header.Set("Content-Type", "application/json")

	// Test the UpdateUserPassword handler
	err = handler.UpdateUserPassword(fiberContext)
	assert.NoError(t, err)

	// Assert that the UpdateUserPassword functions were called with the expected parameters
	assert.True(t, mockAuthWrapper.UpdateUserPasswordFuncCalled, "UpdateUserPassword function of mockWrapper should be called")

	// Release the Fiber context
	app.ReleaseCtx(fiberContext)

}

func TestUpdateUserEmail(t *testing.T) {
	// Create a custom mock client wrapper for Auth Service
	mockAuthWrapper := &MockAuthServiceClientWrapper{}

	// Create handlers using the custom mock client wrapper
	handler := handlers.NewAuthHandlers(mockAuthWrapper)

	mockAuthWrapper.UpdateUserEmailFunc = func(ctx context.Context, req *pb.UpdateUserEmailRequest, opts ...grpc.CallOption) (*pb.User, error) {
		// Simulate the behavior of the gRPC service
		return &pb.User{Id: "1", Name: "Test User1", Email: req.GetEmail(), Role: "user", Created: "2024-02-02T18:18:00", Updated: "2024-02-02T18:18:00"}, nil
	}

	// Create a Fiber context
	app := fiber.New()
	fiberContext := app.AcquireCtx(&fasthttp.RequestCtx{})

	// Set the request body in the Fiber context
	request := &pb.UpdateUserEmailRequest{
		Email:    "testemail@test.com.tr",
		NewEmail: "new-testemail@test.com.tr",
		Password: "paswword",
	}
	body, err := json.Marshal(request)
	assert.NoError(t, err)

	// Set the content-type to JSON
	fiberContext.Request().SetBody(body)
	fiberContext.Request().Header.Set("Content-Type", "application/json")

	// Test the UpdateUserEmail handler
	err = handler.UpdateUserEmail(fiberContext)
	assert.NoError(t, err)

	// Assert that the UpdateUserEmail functions were called with the expected parameters
	assert.True(t, mockAuthWrapper.UpdateUserEmailFuncCalled, "UpdateUserEmail function of mockWrapper should be called")

	// Release the Fiber context
	app.ReleaseCtx(fiberContext)

}

func TestUpdateUserName(t *testing.T) {
	// Create a custom mock client wrapper for Auth Service
	mockAuthWrapper := &MockAuthServiceClientWrapper{}

	// Create handlers using the custom mock client wrapper
	handler := handlers.NewAuthHandlers(mockAuthWrapper)

	mockAuthWrapper.UpdateUserNameFunc = func(ctx context.Context, req *pb.UpdateUserNameRequest, opts ...grpc.CallOption) (*pb.User, error) {
		// Simulate the behavior of the gRPC service
		return &pb.User{Id: "1", Name: req.GetName(), Email: req.GetEmail(), Role: "user", Created: "2024-02-02T18:18:00", Updated: "2024-02-02T18:18:00"}, nil
	}

	// Create a Fiber context
	app := fiber.New()
	fiberContext := app.AcquireCtx(&fasthttp.RequestCtx{})

	// Set the request body in the Fiber context
	request := &pb.UpdateUserNameRequest{
		Email:    "testemail@test.com.tr",
		Name:     "New Test User 1",
		Password: "paswword",
	}
	body, err := json.Marshal(request)
	assert.NoError(t, err)

	// Set the content-type to JSON
	fiberContext.Request().SetBody(body)
	fiberContext.Request().Header.Set("Content-Type", "application/json")

	// Test the UpdateUserName handler
	err = handler.UpdateUserName(fiberContext)
	assert.NoError(t, err)

	// Assert that the UpdateUserName functions were called with the expected parameters
	assert.True(t, mockAuthWrapper.UpdateUserNameFuncCalled, "UpdateUserName function of mockWrapper should be called")

	// Release the Fiber context
	app.ReleaseCtx(fiberContext)

}

func TestListUsers(t *testing.T) {
	// Create a custom mock client wrapper for Auth Service
	mockAuthWrapper := &MockAuthServiceClientWrapper{}

	// Create handlers using the custom mock client wrapper
	handler := handlers.NewAuthHandlers(mockAuthWrapper)

	// Set Auth Service Client in the mockWrapper
	mockAuthWrapper.GetUserRoleFunc = func(ctx context.Context, req *pb.GetUserRoleRequest, opts ...grpc.CallOption) (*pb.GetUserRoleResponse, error) {
		// Simulate the behavior of the gRPC service
		return &pb.GetUserRoleResponse{Role: "admin"}, nil
	}

	mockAuthWrapper.ListUsersFunc = func(ctx context.Context, req *pb.ListUsersRequest, opts ...grpc.CallOption) (pb.AuthService_ListUsersClient, error) {
		// Simulate the behavior of the gRPC service
		users := []*pb.User{
			{Id: "1", Name: "Test User1", Email: "testuser1@test.com.tr", Role: "user", Created: "2024-02-02T18:18:00", Updated: "2024-02-02T18:18:00"},
			{Id: "2", Name: "Test User2", Email: "testuser@test.com.tr", Role: "user", Created: "2024-02-02T18:18:00", Updated: "2024-02-02T18:18:00"},
		}

		mockStream := &MockAuthService_StreamClientWrapper{
			users: users,
		}

		return mockStream, nil
	}

	// Create a Fiber context
	app := fiber.New()
	fiberContext := app.AcquireCtx(&fasthttp.RequestCtx{})

	// Set the request headers in the Fiber context
	userId := bson.NewObjectId()
	token, err := security.NewToken(userId.Hex())
	assert.NoError(t, err)
	fiberContext.Request().Header.Set("Authorization", "Bearer "+token+"")

	// Test the ListUsers handler for the first time
	err = handler.ListUsers(fiberContext)
	assert.NoError(t, err)

	// Assert that the GetUserRole and ListUsers functions were called with the expected parameters
	assert.True(t, mockAuthWrapper.GetUserRoleFuncCalled, "GetUserRole function of mockWrapper should be called")
	assert.True(t, mockAuthWrapper.ListUsersFuncCalled, "ListUsers function of mockWrapper should be called")

	// Get the cached result for the first time
	cachedDataFirstTime, foundFirstTime := util.GetFromCache("ListUsers")
	assert.True(t, foundFirstTime, "Result should be in cache for the first time")

	// Test the ListUsers handler for the second time
	err = handler.ListUsers(fiberContext)
	assert.NoError(t, err)

	// Assert that the GetUserRole and ListUsers functions were called again (second time) with the expected parameters
	assert.True(t, mockAuthWrapper.GetUserRoleFuncCalled, "GetUserRole function of mockWrapper should be called again (second time)")
	assert.True(t, mockAuthWrapper.ListUsersFuncCalled, "ListUsers function of mockWrapper should be called again (second time)")

	// Get the cached result for the second time
	cachedDataSecondTime, foundSecondTime := util.GetFromCache("ListUsers")
	assert.True(t, foundSecondTime, "Result should be in cache for the second time")

	// Assert that the cached results for the first and second times are the same
	assert.Equal(t, cachedDataFirstTime, cachedDataSecondTime, "Cached results for the first and second times should be the same")

	// Release the Fiber context
	app.ReleaseCtx(fiberContext)
}
