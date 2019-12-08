package mocking

import "github.com/stretchr/testify/mock"

type mockEncoder struct {
	mock.Mock
}

func (m mockEncoder) EncodeToString(bytes []byte) string {
	args := m.Called(bytes)
	return args.String(0)
}
