package handler

import "fmt"

func errParamIsReuired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// CreateOpening

type CreateopeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateopeningRequest) validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Salary <= 0{
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Role == "" {
		return errParamIsReuired("role", "string")
	}
	if r.Company == "" {
		return errParamIsReuired("company", "string")
	}
	if r.Location == "" {
		return errParamIsReuired("location", "string")
	}
	if r.Remote ==  nil {
		return errParamIsReuired("remote", "bool")
	}
	if r.Link == "" {
		return errParamIsReuired("link", "string")
	}
	if r.Salary <= 100 {
		return errParamIsReuired("salary", "int64")
	}

	return nil
}
