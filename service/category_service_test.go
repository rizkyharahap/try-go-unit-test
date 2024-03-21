package service

import (
	"testing"

	"github.com/rizkyharahap/try-go-unit-test/entity"
	"github.com/rizkyharahap/try-go-unit-test/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Make mock CategoryRepositoryMock
var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}

// Inject CategoryRepositoryMock to CategoryService.Repository
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	// Mock repository returned nil
	categoryRepository.Mock.On("FindById", "0").Return(nil)

	category, err := categoryService.Get("0")
	assert.NotNil(t, err)
	assert.Nil(t, category)
}

func TestCategoryService_GetFound(t *testing.T) {
	category := entity.Category{
		Id:   "1",
		Name: "Laptop",
	}
	categoryRepository.Mock.On("FindById", "1").Return(category)

	result, err := categoryService.Get("1")
	assert.Nil(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
