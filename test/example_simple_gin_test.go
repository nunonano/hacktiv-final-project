package test

import (
	"fmt"
	"testing"
)

type JSONer interface {
	JSON(code int)
}

type UserController struct{}

func (ctrl UserController) GetAll(c JSONer) {
	c.JSON(200)
}

type ContextMock struct {
	JSONCalled bool
}

func (c *ContextMock) JSON(code int) {
	c.JSONCalled = true
}

func TestSimpleGin(t *testing.T) {
	c := &ContextMock{false}
	ctrl := UserController{}

	ctrl.GetAll(c)

	if c.JSONCalled == false {
		fmt.Print("Failed")
	} else {
		fmt.Print("Tested successfully")
	}
}
