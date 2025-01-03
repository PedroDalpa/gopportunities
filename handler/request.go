package handler

import "fmt"

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Salary   int64  `json:"salary"`
	Link     string `json:"link"`
}

func errParamIsRequired(n, t string) error {
	return fmt.Errorf("param: %s (type: %s) is required", n, t)
}

func (r *CreateOpeningRequest) Validate() error {

	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}

	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}

	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}

	if r.Salary <= 0 {
		return fmt.Errorf("salary must be a positive integer")
	}

	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}

	if r.Remote == nil {
		return errParamIsRequired("remote", "boolean")
	}

	return nil

}
