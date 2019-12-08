package mocking

import "fmt"

type ExternalDependency interface {
	DoSomething(input string) (string, error)
}

func NewExternalDependency() ExternalDependency {
	return &craptasticDep{
		prefix: "moo",
	}
}

type craptasticDep struct {
	prefix string
}

func (c *craptasticDep) DoSomething(input string) (string, error) {
	return fmt.Sprintf("%s: %s", c.prefix, input), nil
}
