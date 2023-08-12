package vocabulary

import "github.com/stretchr/testify/mock"

var _ FileHandler = &Filesystem{}

var _ FileHandler = &MockFileHandler{}

type MockFileHandler struct {
	mock.Mock
}

func (m *MockFileHandler) Read(dir, filename string) (string, error) {
	args := m.Called(dir, filename)
	return args.String(0), args.Error(1)
}

func (m *MockFileHandler) Write(dir, filename, contents string) error {
	args := m.Called(dir, filename, contents)
	return args.Error(0)
}
