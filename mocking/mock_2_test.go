package mocking

import (
	"github.com/stretchr/testify/mock"
)

type mockExternalDep2 struct {
	mock.Mock
}

func (d *mockExternalDep2) DoSomething(input string) (string, error) {
	retValues := d.Called(input)
	return retValues.String(0), retValues.Error(1)
}
