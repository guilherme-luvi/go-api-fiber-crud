package handlers

import "fmt"

// helper function
func errParamMissing(param, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", param, typ)
}

// Struct to hold the request body for creating a user
type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *CreateUserRequest) validate() error {
	if req.Name == "" {
		return errParamMissing("name", "string")
	}

	if req.Email == "" {
		return errParamMissing("email", "string")
	}

	if req.Password == "" {
		return errParamMissing("password", "string")
	}

	return nil
}

// Struct to hold the request body for logging in
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *LoginRequest) validate() error {
	if req.Email == "" {
		return errParamMissing("email", "string")
	}

	if req.Password == "" {
		return errParamMissing("password", "string")
	}

	return nil
}

// Struct to hold the request body for updating a user
type UpdateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *UpdateUserRequest) validate() error {
	if req.Name == "" {
		return errParamMissing("name", "string")
	}

	if req.Email == "" {
		return errParamMissing("email", "string")
	}

	if req.Password == "" {
		return errParamMissing("password", "string")
	}

	return nil
}
