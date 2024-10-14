package interfaces

import "github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/domain"

type CategoryRepository interface {
	GetCategories() ([]domain.Category, error)
	AddCategory(category string) (domain.Category, error)
	CheckCategory(current string) (bool, error)
	UpdateCategory(current string, new string) (domain.Category, error)
	DeleteCategory(CategoryID string) error
	CheckCategoryByID(categoryID int) (bool, error)
}
