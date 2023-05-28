package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateOpeningRequest struct {
	Company  string `json:"company"`
	Link     string `json:"link"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Role     string `json:"role"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Salary >= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if r.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}
	if r.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}

	return nil

}

type UpdateOpeningRequest struct {
	Company  string `json:"company"`
	Link     string `json:"link"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Role     string `json:"role"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Salary >= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}

	// if any field is provided, validation is truthy
	if r.Company != "" || r.Link != "" || r.Location != "" || r.Remote != nil || r.Role != "" || r.Salary >= 0 {
		return nil
	}

	// if none of the fields are providade validation is falsy
	return fmt.Errorf("at least one valid field must be provided")
}
