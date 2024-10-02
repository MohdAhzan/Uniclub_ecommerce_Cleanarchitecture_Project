package interfaces

import (
	"project/pkg/utils/domain"
)

type CategoryUseCase interface {
	AddCategory(category string) (domain.Category, error)
	GetCategories() ([]domain.Category, error)
	UpdateCategory(current string, new string) (domain.Category, error)
	DeleteCategory(CategoryID string) error
}
