package interfaces

import (
	"project/pkg/domain"
)

type CategoryUseCase interface {
	AddCategory(category domain.Category) (domain.Category, error)
	GetCategories() ([]domain.Category, error)
	UpdateCategory(current string, new string) (domain.Category, error)
	DeleteCategory(CategoryID string) error
}
