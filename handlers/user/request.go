package user

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// Create User

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (r CreateUserRequest) Validate() error {
	if r.Username == "" && r.Password == "" && r.Email == "" {
		return fmt.Errorf("empty request body or malformed")
	}
	if r.Username == "" {
		return errParamIsRequired("username", "string")
	}
	if r.Password == "" {
		return errParamIsRequired("password", "string")
	}
	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}
	return nil
}

// Update User

type UpdateUserRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

func (r UpdateUserRequest) Validate() error {
	if r.Username != "" || r.Email != "" {
		return nil
	}
	return fmt.Errorf("empty request body at least one field is required")
}
