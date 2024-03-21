package repository

import "github.com/rizkyharahap/try-go-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
