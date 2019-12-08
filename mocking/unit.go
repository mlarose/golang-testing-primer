package mocking

type Unit interface {
	SomeFunc(input string) (string, error)
}

type BytesToStringEncoder interface {
	EncodeToString([]byte) string
}

type unit struct {
	dep     ExternalDependency
	encoder BytesToStringEncoder
}

func (u *unit) SomeFunc(input string) (string, error) {
	out, err := u.dep.DoSomething(input)
	if err != nil {
		return out, err
	}

	return u.encoder.EncodeToString([]byte(out)), nil
}
