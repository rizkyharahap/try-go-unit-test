package repository

import (
	"github.com/rizkyharahap/try-go-unit-test/entity"
	"github.com/stretchr/testify/mock"
)

// Repository Mock Struct
type CategoryRepositoryMock struct {
	Mock mock.Mock
}

/**
* Implement Mock FindById returned entity.Category or nil
**/
func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	}

	category := arguments.Get(0).(entity.Category)
	return &category
}
