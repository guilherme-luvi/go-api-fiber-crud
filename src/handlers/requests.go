package handlers

import "fmt"

func errParamMissing(param, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", param, typ)
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (req *CreateUserRequest) validate() error {
	if req.Name == "" {
		return errParamMissing("name", "string")
	}

	if req.Email == "" {
		return errParamMissing("email", "string")
	}

	return nil
}
