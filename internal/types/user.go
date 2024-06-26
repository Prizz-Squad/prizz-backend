package types

import "fmt"

const (
	minUsernameLen = 3
	minPasswordLen = 7
)
const (
	MemberRole  = 0
	ClientRole  = 1
	ManagerRole = 2
	AdminRole   = 3
)

type User struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Department int    `json:"department"`
	Role       int    `json:"role"`
}
type UserCreateRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Department int    `json:"department"`
	Role       int    `json:"role"`
}
type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserResponse struct {
	ID string `json:"id"`
}

func (params *UserCreateRequest) Validate() map[string]string {
	errors := map[string]string{}
	if len(params.Username) < minUsernameLen {
		errors["username"] = fmt.Sprintf("username length should be at least %d characters", minUsernameLen)
	}
	if len(params.Password) < minPasswordLen {
		errors["password"] = fmt.Sprintf("password length should be at least %d characters", minPasswordLen)
	}
	return errors
}
