package requests

import (
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/validate"
	c "github.com/qbhy/goal/validate/checkers"
)

type LoginRequest struct {
	Username string
	Password string
}

func (this LoginRequest) Assure() {
	this.Validate().Assure()
}

func (this LoginRequest) Validate() contracts.ValidatedResult {
	return validate.Make(this, contracts.Checkers{
		"username": {c.StringLength{Min: 4, Max: 16}},
		"password": {c.StringLength{Min: 6, Max: 10}},
	}).Validate()
}
