package handlers

import (
	"context"
	"encoding/json"
	"io"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
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

// Custom mock client that implements the pb.VulnServiceClient interface
type MockVulnServiceClientWrapper struct {
	AddCVEFunc              func(ctx context.Context, req *pb.AddCVERequest, opts ...grpc.CallOption) (*pb.CVE, error)
	SearchCVEFunc           func(ctx context.Context, req *pb.SearchCVERequest, opts ...grpc.CallOption) (pb.VulnService_SearchCVEClient, error)
	GetAllCVEsFunc          func(ctx context.Context, req *pb.GetAllCVEsRequest, opts ...grpc.CallOption) (pb.VulnService_GetAllCVEsClient, error)
	DeleteCVEFunc           func(ctx context.Context, req *pb.DeleteCVERequest, opts ...grpc.CallOption) (*pb.DeleteCVEResponse, error)
	UpdateCVEFunc           func(ctx context.Context, req *pb.UpdateCVERequest, opts ...grpc.CallOption) (*pb.CVE, error)
	FetchNVDFeedsFunc       func(ctx context.Context, req *pb.FetchNVDFeedsRequest, opts ...grpc.CallOption) (pb.VulnService_FetchNVDFeedsClient, error)
	AddCVEFuncCalled        bool
	SearchCVEFuncCalled     bool
	GetAllCVEsFuncCalled    bool
	DeleteCVEFuncCalled     bool
	UpdateCVEFuncCalled     bool
	FetchNVDFeedsFuncCalled bool
}

func (c *MockVulnServiceClientWrapper) AddCVE(ctx context.Context, req *pb.AddCVERequest, opts ...grpc.CallOption) (*pb.CVE, error) {
	c.AddCVEFuncCalled = true
	if c.AddCVEFunc != nil {
		return c.AddCVEFunc(ctx, req, opts...)
	}
	return nil, nil
}

func (c *MockVulnServiceClientWrapper) SearchCVE(ctx context.Context, req *pb.SearchCVERequest, opts ...grpc.CallOption) (pb.VulnService_SearchCVEClient, error) {
	c.SearchCVEFuncCalled = true
	if c.SearchCVEFunc != nil {
		return c.SearchCVEFunc(ctx, req, opts...)
	}
	return nil, nil
}

func (c *MockVulnServiceClientWrapper) GetAllCVEs(ctx context.Context, req *pb.GetAllCVEsRequest, opts ...grpc.CallOption) (pb.VulnService_GetAllCVEsClient, error) {
	c.GetAllCVEsFuncCalled = true
	if c.GetAllCVEsFunc != nil {
		return c.GetAllCVEsFunc(ctx, req, opts...)
	}
	return nil, nil
}

func (c *MockVulnServiceClientWrapper) DeleteCVE(ctx context.Context, req *pb.DeleteCVERequest, opts ...grpc.CallOption) (*pb.DeleteCVEResponse, error) {
	c.DeleteCVEFuncCalled = true
	if c.DeleteCVEFunc != nil {
		return c.DeleteCVEFunc(ctx, req, opts...)
	}
	return nil, nil
}

func (c *MockVulnServiceClientWrapper) UpdateCVE(ctx context.Context, req *pb.UpdateCVERequest, opts ...grpc.CallOption) (*pb.CVE, error) {
	c.UpdateCVEFuncCalled = true
	if c.UpdateCVEFunc != nil {
		return c.UpdateCVEFunc(ctx, req, opts...)
	}
	return nil, nil
}

func (c *MockVulnServiceClientWrapper) FetchNVDFeeds(ctx context.Context, req *pb.FetchNVDFeedsRequest, opts ...grpc.CallOption) (pb.VulnService_FetchNVDFeedsClient, error) {
	c.FetchNVDFeedsFuncCalled = true
	if c.FetchNVDFeedsFunc != nil {
		return c.FetchNVDFeedsFunc(ctx, req, opts...)
	}
	return nil, nil
}

type MockVulnService_GetAllCVEsClientWrapper struct {
	cves []*pb.CVE
	idx  int
}

// Recv simulates receiving CVEs from the server.
func (m *MockVulnService_GetAllCVEsClientWrapper) Recv() (*pb.CVE, error) {
	if m.idx < len(m.cves) {
		cve := m.cves[m.idx]
		m.idx++
		return cve, nil
	}
	return nil, io.EOF
}

// CloseSend simulates closing the send stream.
func (m *MockVulnService_GetAllCVEsClientWrapper) CloseSend() error {
	return nil
}

// Context returns the client context.
func (m *MockVulnService_GetAllCVEsClientWrapper) Context() context.Context {
	return context.Background()
}

// SendMsg simulates sending a message.
func (m *MockVulnService_GetAllCVEsClientWrapper) SendMsg(m1 interface{}) error {
	return nil
}

// RecvMsg simulates receiving a message.
func (m *MockVulnService_GetAllCVEsClientWrapper) RecvMsg(m1 interface{}) error {
	return nil
}

// Header returns the header metadata.
func (m *MockVulnService_GetAllCVEsClientWrapper) Header() (metadata.MD, error) {
	return nil, nil
}

// Trailer returns the trailer metadata.
func (m *MockVulnService_GetAllCVEsClientWrapper) Trailer() metadata.MD {
	return nil
}

// SendHeader simulates sending header metadata.
func (m *MockVulnService_GetAllCVEsClientWrapper) SendHeader(m1 metadata.MD) error {
	return nil
}

// SetSendDeadline simulates setting the send deadline.
func (m *MockVulnService_GetAllCVEsClientWrapper) SetSendDeadline(m1 time.Time) error {
	return nil
}

// SetRecvDeadline simulates setting the receive deadline.
func (m *MockVulnService_GetAllCVEsClientWrapper) SetRecvDeadline(m1 time.Time) error {
	return nil
}

// MockVulnService_GetAllCVEsClient basit bir mock implementasyonu
func TestAddCVE(t *testing.T) {

	// Create a custom mock client wrapper for Auth Service
	mockAuthWrapper := &MockAuthServiceClientWrapper{}

	// Create a custom mock client wrapper for Vuln Service
	mockVulnWrapper := &MockVulnServiceClientWrapper{}

	// Create handlers using the custom mock client wrapper
	handler := NewVulnHandlers(mockAuthWrapper, mockVulnWrapper)

	// Set Auth Service Client in the mockWrapper
	mockAuthWrapper.GetUserRoleFunc = func(ctx context.Context, req *pb.GetUserRoleRequest, opts ...grpc.CallOption) (*pb.GetUserRoleResponse, error) {
		// Simulate the behavior of the gRPC service
		return &pb.GetUserRoleResponse{Role: "admin"}, nil
	}

	// Set Vuln Service Client in the mockWrapper
	mockVulnWrapper.AddCVEFunc = func(ctx context.Context, req *pb.AddCVERequest, opts ...grpc.CallOption) (*pb.CVE, error) {
		// Simulate the behavior of the gRPC service
		return &pb.CVE{Id: "123", CveId: req.CveId, Description: "Test CVE", Severity: "High", Product: "Test Product", Vendor: "Test Vendor", Published: "2024-01-27T10:10:10", Modified: "2024-01-27T10:10:10"}, nil
	}

	// Create a Fiber context
	app := fiber.New()
	fiberContext := app.AcquireCtx(&fasthttp.RequestCtx{})

	// Set the request body in the Fiber context
	request := &pb.AddCVERequest{
		CveId:       "test123",
		Description: "Test CVE",
		Severity:    "High",
		Product:     "Test Product",
		Vendor:      "Test Vendor",
		Links:       []string{"https://example.com"},
		Published:   "2024-01-27T10:10:10",
		Modified:    "2024-01-27T10:10:10",
	}
	body, err := json.Marshal(request)
	assert.NoError(t, err)

	// Set the content-type to JSON
	fiberContext.Request().SetBody(body)
	fiberContext.Request().Header.Set("Content-Type", "application/json")

	userId := bson.NewObjectId()
	token, err := security.NewToken(userId.Hex())
	assert.NoError(t, err)
	fiberContext.Request().Header.Set("Authorization", "Bearer "+token+"")

	// Your test code here
	err = handler.AddCVE(fiberContext)
	assert.NoError(t, err)

	// Assert that the GetUserRole and AddCVE functions were called with the expected parameters
	assert.True(t, mockAuthWrapper.GetUserRoleFuncCalled, "GetUserRole function of mockWrapper should be called")
	assert.True(t, mockVulnWrapper.AddCVEFuncCalled, "AddCVE function of mockWrapper should be called")

	// Release the Fiber context
	app.ReleaseCtx(fiberContext)
}

func TestGetAllCVEs(t *testing.T) {
	// Create a custom mock client wrapper for Auth Service
	mockAuthWrapper := &MockAuthServiceClientWrapper{}

	// Create a custom mock client wrapper for Vuln Service
	mockVulnWrapper := &MockVulnServiceClientWrapper{}

	// Create handlers using the custom mock client wrapper
	handler := NewVulnHandlers(mockAuthWrapper, mockVulnWrapper)

	// Set Auth Service Client in the mockWrapper
	mockAuthWrapper.GetUserRoleFunc = func(ctx context.Context, req *pb.GetUserRoleRequest, opts ...grpc.CallOption) (*pb.GetUserRoleResponse, error) {
		// Simulate the behavior of the gRPC service
		return &pb.GetUserRoleResponse{Role: "admin"}, nil
	}

	// Set Vuln Service Client in the mockWrapper
	mockVulnWrapper.GetAllCVEsFunc = func(ctx context.Context, req *pb.GetAllCVEsRequest, opts ...grpc.CallOption) (pb.VulnService_GetAllCVEsClient, error) {
		// Simulate the behavior of the gRPC service
		cves := []*pb.CVE{
			{Id: "1", CveId: "CVE-2022-1234", Description: "Test CVE 1", Severity: "High", Product: "Test Product 1", Vendor: "Test Vendor 1", Published: "2024-01-27T10:10:10", Modified: "2024-01-27T10:10:10"},
			{Id: "2", CveId: "CVE-2022-5678", Description: "Test CVE 2", Severity: "Medium", Product: "Test Product 2", Vendor: "Test Vendor 2", Published: "2024-01-27T10:20:20", Modified: "2024-01-27T10:20:20"},
		}

		mockStream := &MockVulnService_GetAllCVEsClientWrapper{
			cves: cves,
		}

		return mockStream, nil
	}

	// Create a dummy Fiber context
	app := fiber.New()
	fiberContext := app.AcquireCtx(&fasthttp.RequestCtx{})

	// Set the request headers in the Fiber context
	userId := bson.NewObjectId()
	token, err := security.NewToken(userId.Hex())
	assert.NoError(t, err)
	fiberContext.Request().Header.Set("Authorization", "Bearer "+token+"")

	// Your test code here
	err = handler.GetAllCVEs(fiberContext)
	assert.NoError(t, err)

	// Assert that the GetUserRole and GetAllCVEs functions were called with the expected parameters
	assert.True(t, mockAuthWrapper.GetUserRoleFuncCalled, "GetUserRole function of mockAuthWrapper should be called")
	assert.True(t, mockVulnWrapper.GetAllCVEsFuncCalled, "GetAllCVEs function of mockVulnWrapper should be called")

	// Additional assertions based on the expected response can be added if needed

	// Release the Fiber context
	app.ReleaseCtx(fiberContext)
}