<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# authentication

```go
import "github.com/mtnmunuklu/bavul/authentication"
```

## Index



# models

```go
import "github.com/mtnmunuklu/bavul/authentication/models"
```

## Index

- [type User](<#User>)
  - [func \(u \*User\) FromProto\(user \*pb.User\)](<#User.FromProto>)
  - [func \(u \*User\) ToProto\(\) \*pb.User](<#User.ToProto>)


<a name="User"></a>
## type [User](<https://github.com/mtnmunuklu/bavul/blob/main/authentication/models/user.go#L13-L21>)

User provides the user instance for authentication job.

```go
type User struct {
    Id       bson.ObjectId `bson:"_id"`
    Name     string        `bson:"name"`
    Email    string        `bson:"email"`
    Password string        `bson:"password"`
    Role     string        `bson:"role"`
    Created  time.Time     `bson:"created"`
    Updated  time.Time     `bson:"updated"`
}
```

<a name="User.FromProto"></a>
### func \(\*User\) [FromProto](<https://github.com/mtnmunuklu/bavul/blob/main/authentication/models/user.go#L36>)

```go
func (u *User) FromProto(user *pb.User)
```

FromProto gets user from protocol buffer and converts to the user structure.

<a name="User.ToProto"></a>
### func \(\*User\) [ToProto](<https://github.com/mtnmunuklu/bavul/blob/main/authentication/models/user.go#L24>)

```go
func (u *User) ToProto() *pb.User
```

ToProto converts the user structure into a protocol buffer user structure.

# repository

```go
import "github.com/mtnmunuklu/bavul/authentication/repository"
```

## Index

- [Constants](<#constants>)
- [type UserRepository](<#UserRepository>)
  - [func NewUserRepository\(conn db.Connection\) UserRepository](<#NewUserRepository>)


## Constants

<a name="UserCollection"></a>

```go
const UserCollection = "users"
```

<a name="UserRepository"></a>
## type [UserRepository](<https://github.com/mtnmunuklu/bavul/blob/main/authentication/repository/user.go#L14-L21>)

UserRepository is the interface of the authentication backend.

```go
type UserRepository interface {
    Save(user *models.User) error
    GetById(id string) (user *models.User, err error)
    GetByEmail(email string) (user *models.User, err error)
    GetAll() (user []*models.User, err error)
    Update(user *models.User) error
    DeleteById(id string) error
}
```

<a name="NewUserRepository"></a>
### func [NewUserRepository](<https://github.com/mtnmunuklu/bavul/blob/main/authentication/repository/user.go#L29>)

```go
func NewUserRepository(conn db.Connection) UserRepository
```

NewUserRepository creates a new UserRepository instance.

# service

```go
import "github.com/mtnmunuklu/bavul/authentication/service"
```

## Index

- [func NewAuthService\(userRepository repository.UserRepository\) pb.AuthServiceServer](<#NewAuthService>)
- [type AuthService](<#AuthService>)
  - [func \(s \*AuthService\) ChangeUserRole\(ctx context.Context, req \*pb.ChangeUserRoleRequest\) \(\*pb.User, error\)](<#AuthService.ChangeUserRole>)
  - [func \(s \*AuthService\) DeleteUser\(ctx context.Context, req \*pb.DeleteUserRequest\) \(\*pb.DeleteUserResponse, error\)](<#AuthService.DeleteUser>)
  - [func \(s \*AuthService\) GetUser\(ctx context.Context, req \*pb.GetUserRequest\) \(\*pb.User, error\)](<#AuthService.GetUser>)
  - [func \(s \*AuthService\) GetUserRole\(ctx context.Context, req \*pb.GetUserRoleRequest\) \(\*pb.GetUserRoleResponse, error\)](<#AuthService.GetUserRole>)
  - [func \(s \*AuthService\) ListUsers\(req \*pb.ListUsersRequest, stream pb.AuthService\_ListUsersServer\) error](<#AuthService.ListUsers>)
  - [func \(s \*AuthService\) SignIn\(ctx context.Context, req \*pb.SignInRequest\) \(\*pb.SignInResponse, error\)](<#AuthService.SignIn>)
  - [func \(s \*AuthService\) SignUp\(ctx context.Context, req \*pb.SignUpRequest\) \(\*pb.User, error\)](<#AuthService.SignUp>)
  - [func \(s \*AuthService\) UpdateUserEmail\(ctx context.Context, req \*pb.UpdateUserEmailRequest\) \(\*pb.User, error\)](<#AuthService.UpdateUserEmail>)
  - [func \(s \*AuthService\) UpdateUserName\(ctx context.Context, req \*pb.UpdateUserNameRequest\) \(\*pb.User, error\)](<#AuthService.UpdateUserName>)
  - [func \(s \*AuthService\) UpdateUserPassword\(ctx context.Context, req \*pb.UpdateUserPasswordRequest\) \(\*pb.User, error\)](<#AuthService.UpdateUserPassword>)


<a name="NewAuthService"></a>
## func [NewAuthService](<https://github.com/mtnmunuklu/bavul/blob/main/authentication/service/service.go#L24>)

```go
func NewAuthService(userRepository repository.UserRepository) pb.AuthServiceServer
```

NewAuthService creates a new AuthService instance.

<a name="AuthService"></a>
## type [AuthService](<https://github.com/mtnmunuklu/bavul/blob/main/authentication/service/service.go#L19-L21>)

AuthService provides usersRepository for authentication service.

```go
type AuthService struct {
    // contains filtered or unexported fields
}
```

<a name="AuthService.ChangeUserRole"></a>
### func \(\*AuthService\) [ChangeUserRole](<https://github.com/mtnmunuklu/bavul/blob/main/authentication/service/service.go#L132>)

```go
func (s *AuthService) ChangeUserRole(ctx context.Context, req *pb.ChangeUserRoleRequest) (*pb.User, error)
```

ChangeUserRole performs change the user role.

<a name="AuthService.DeleteUser"></a>
### func \(\*AuthService\) [DeleteUser](<https://github.com/mtnmunuklu/bavul/blob/main/authentication/service/service.go#L112>)

```go
func (s *AuthService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error)
```

DeleteUser performs delete the user.

<a name="AuthService.GetUser"></a>
### func \(\*AuthService\) [GetUser](<https://github.com/mtnmunuklu/bavul/blob/main/authentication/service/service.go#L97>)

```go
func (s *AuthService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error)
```

GetUser performs return the user by id.

<a name="AuthService.GetUserRole"></a>
### func \(\*AuthService\) [GetUserRole](<https://github.com/mtnmunuklu/bavul/blob/main/authentication/service/service.go#L161>)

```go
func (s *AuthService) GetUserRole(ctx context.Context, req *pb.GetUserRoleRequest) (*pb.GetUserRoleResponse, error)
```

GetUserRole performs return the user role by id.

<a name="AuthService.ListUsers"></a>
### func \(\*AuthService\) [ListUsers](<https://github.com/mtnmunuklu/bavul/blob/main/authentication/service/service.go#L298>)

```go
func (s *AuthService) ListUsers(req *pb.ListUsersRequest, stream pb.AuthService_ListUsersServer) error
```

ListUser list all users.

<a name="AuthService.SignIn"></a>
### func \(\*AuthService\) [SignIn](<https://github.com/mtnmunuklu/bavul/blob/main/authentication/service/service.go#L72>)

```go
func (s *AuthService) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error)
```

SignIn performs the user login process.

<a name="AuthService.SignUp"></a>
### func \(\*AuthService\) [SignUp](<https://github.com/mtnmunuklu/bavul/blob/main/authentication/service/service.go#L29>)

```go
func (s *AuthService) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.User, error)
```

SignUp performs the user registration process.

<a name="AuthService.UpdateUserEmail"></a>
### func \(\*AuthService\) [UpdateUserEmail](<https://github.com/mtnmunuklu/bavul/blob/main/authentication/service/service.go#L221>)

```go
func (s *AuthService) UpdateUserEmail(ctx context.Context, req *pb.UpdateUserEmailRequest) (*pb.User, error)
```

UpdateUser performs update the password.

<a name="AuthService.UpdateUserName"></a>
### func \(\*AuthService\) [UpdateUserName](<https://github.com/mtnmunuklu/bavul/blob/main/authentication/service/service.go#L260>)

```go
func (s *AuthService) UpdateUserName(ctx context.Context, req *pb.UpdateUserNameRequest) (*pb.User, error)
```

UpdateUser performs update the username.

<a name="AuthService.UpdateUserPassword"></a>
### func \(\*AuthService\) [UpdateUserPassword](<https://github.com/mtnmunuklu/bavul/blob/main/authentication/service/service.go#L175>)

```go
func (s *AuthService) UpdateUserPassword(ctx context.Context, req *pb.UpdateUserPasswordRequest) (*pb.User, error)
```

UpdateUser performs update the password.

# util

```go
import "github.com/mtnmunuklu/bavul/authentication/util"
```

## Index

- [Variables](<#variables>)
- [func FormatTime\(t time.Time\) string](<#FormatTime>)
- [func NormalizeEmail\(email string\) string](<#NormalizeEmail>)
- [func ParseTime\(dateStr string\) time.Time](<#ParseTime>)
- [func ValidateSignUp\(user \*pb.SignUpRequest\) error](<#ValidateSignUp>)


## Variables

<a name="ErrInvalidUserId"></a>Contains error codes for authentication service.

```go
var (
    ErrInvalidUserId         = errors.New("invalid user id")
    ErrEmptyName             = errors.New("name can't be empty")
    ErrEmptyEmail            = errors.New("email can't be empty")
    ErrEmptyNewEmail         = errors.New("new email can't be empty")
    ErrEmptyPassword         = errors.New("password can't be empty")
    ErrEmptyNewPassword      = errors.New("new password can't be empty")
    ErrEmptyUserRole         = errors.New("user role can't be empty")
    ErrExistEmail            = errors.New("email already exist")
    ErrNotFoundEmail         = errors.New("email did not found")
    ErrNotFoundUserId        = errors.New("user id could not be found")
    ErrFailedSignIn          = errors.New("signin failed")
    ErrMismatchedPassword    = errors.New("password did not match")
    ErrCreateUser            = errors.New("user could not be created")
    ErrDeleteUser            = errors.New("user could not be deleted")
    ErrUpdateUser            = errors.New("user could not be updated")
    ErrEncryptPassword       = errors.New("password could not be encrypted")
    ErrNotPerformedOperation = errors.New("operation could not be performed")
)
```

<a name="FormatTime"></a>
## func [FormatTime](<https://github.com/mtnmunuklu/bavul/blob/main/authentication/util/util.go#L59>)

```go
func FormatTime(t time.Time) string
```



<a name="NormalizeEmail"></a>
## func [NormalizeEmail](<https://github.com/mtnmunuklu/bavul/blob/main/authentication/util/util.go#L46>)

```go
func NormalizeEmail(email string) string
```

NormalizeEmail normalizes the user email address.

<a name="ParseTime"></a>
## func [ParseTime](<https://github.com/mtnmunuklu/bavul/blob/main/authentication/util/util.go#L50>)

```go
func ParseTime(dateStr string) time.Time
```



<a name="ValidateSignUp"></a>
## func [ValidateSignUp](<https://github.com/mtnmunuklu/bavul/blob/main/authentication/util/util.go#L33>)

```go
func ValidateSignUp(user *pb.SignUpRequest) error
```

ValidateSingnUp validates the user information for user registration process.

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
