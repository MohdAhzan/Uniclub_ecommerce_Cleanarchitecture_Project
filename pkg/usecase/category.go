package usecase

import (
	"errors"
	"fmt"
	"project/pkg/domain"
	interfaces "project/pkg/repository/interface"
	services "project/pkg/usecase/interface"
)

type categoryUseCase struct {
	repository interfaces.CategoryRepository
}

func NewCategoryUseCase(repo interfaces.CategoryRepository) services.CategoryUseCase {
	return &categoryUseCase{
		repository: repo,
	}
}

func (cat *categoryUseCase) AddCategory(category domain.Category) (domain.Category, error) {

	productResponse, err := cat.repository.AddCategory(category)
	if err != nil {
		return domain.Category{}, err
	}

	return productResponse, nil
}

func (cat *categoryUseCase) GetCategories() ([]domain.Category, error) {

	categories, err := cat.repository.GetCategories()
	if err != nil {
		return []domain.Category{}, err
	}
	return categories, nil

}

func (cat *categoryUseCase) UpdateCategory(current string, new string) (domain.Category, error) {

	Exists, err := cat.repository.CheckCategory(current)

	if err != nil {
		return domain.Category{}, err
	}
	catErr := fmt.Sprintf("There is no category named  %s", current)
	if !Exists {
		return domain.Category{}, errors.New(catErr)
	}
	updatedCategory, err := cat.repository.UpdateCategory(current, new)
	if err != nil {
		return domain.Category{}, err
	}
	return updatedCategory, nil
}

func (cd *categoryUseCase) DeleteCategory(CategoryID string) error {

	err := cd.repository.DeleteCategory(CategoryID)

	if err != nil {
		return err
	}

	return nil
}
