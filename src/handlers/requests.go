package handlers

import "fmt"

func errParamMissing(param, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", param, typ)
}

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
