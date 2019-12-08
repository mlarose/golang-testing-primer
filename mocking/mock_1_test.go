package mocking

import "fmt"

type mockExternalDep1 struct{}

func (d *mockExternalDep1) DoSomething(input string) (string, error) {
	switch input {
	case "foo":
		return "bar", nil
	case "invalid":
		return "invalid", fmt.Errorf("some error!")
	}

	return "", nil
}
