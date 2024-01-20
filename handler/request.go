package handler

import "fmt"

func errParameterIsRequred(name string, parameterType string) error {
	return fmt.Errorf("param %s (type: %s) is required", name, parameterType)
}

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {

	if r.Role == "" {
		return errParameterIsRequred("role", "string")
	}
	if r.Company == "" {
		return errParameterIsRequred("company", "string")
	}
	if r.Location == "" {
		return errParameterIsRequred("location", "string")
	}
	if r.Link == "" {
		return errParameterIsRequred("link", "string")
	}
	if r.Remote == nil {
		return errParameterIsRequred("remote", "bool")
	}

	if r.Salary <= 0 {
		return errParameterIsRequred("salary", "int64")
	}

	return nil
}
