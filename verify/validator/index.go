package validator

import (
	"github.com/go-playground/validator"
	"github.com/xm-chentl/go-mvc/verify"
)

type validateInst struct {
	inst *validator.Validate
}

func (v *validateInst) Args(cfg string) verify.IVerify {
	return v
}

func (v *validateInst) Execute(structData interface{}) bool {
	if err := v.inst.Struct(structData); err != nil {
		return false
	}

	return true
}

func New() verify.IVerify {
	return &validateInst{
		inst: validator.New(),
	}
}
