package interfaces

import "project/pkg/domain"

type CategoryRepository interface {
	GetCategories() ([]domain.Category, error)
	AddCategory(category domain.Category) (domain.Category, error)
	CheckCategory(current string) (bool, error)
	UpdateCategory(current string, new string) (domain.Category, error)
	DeleteCategory(CategoryID string) error
}
