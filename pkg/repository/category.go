package repository

import (
	"errors"
	"fmt"
	"project/pkg/utils/domain"
	interfaces "project/pkg/repository/interface"
	"strconv"

	"gorm.io/gorm"
)

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(DB *gorm.DB) interfaces.CategoryRepository {
	return &categoryRepository{DB}
}

func (c *categoryRepository) GetCategories() ([]domain.Category, error) {
	var model []domain.Category
	err := c.DB.Raw("select * from categories").Scan(&model).Error
	if err != nil {
		return []domain.Category{}, err
	}
	return model, nil

}

func (p *categoryRepository) AddCategory(c domain.Category) (domain.Category, error) {
	var b string

	err := p.DB.Raw("INSERT INTO categories(category) VALUES (?) RETURNING category", c.Category).Scan(&b).Error
	if err != nil {
		return domain.Category{}, nil
	}

	var categoryResponse domain.Category
	err = p.DB.Raw(`SELECT p.id,p.category FROM categories p WHERE p.category = ?`, b).Scan(&categoryResponse).Error
	if err != nil {
		return domain.Category{}, err
	}

	return categoryResponse, nil
}

func (cat *categoryRepository) CheckCategory(current string) (bool, error) {
	var i int

	err := cat.DB.Raw("SELECT count(*) FROM categories where category = ? ", current).Scan(&i).Error

	if err != nil {
		return false, err
	}
	if i == 0 {
		return false, err
	}
	return true, nil

}

func (cat *categoryRepository) UpdateCategory(current, new string) (domain.Category, error) {

	if cat.DB == nil {
		return domain.Category{}, errors.New("database connection is nil")
	}

	// update the category

	if err := cat.DB.Exec("UPDATE categories SET category = $1  WHERE category  = $2", new, current).Error; err != nil {
		return domain.Category{}, err
	}

	var newcat domain.Category
	if err := cat.DB.First(&newcat, "category = ?", new).Error; err != nil {
		return domain.Category{}, err
	}

	return newcat, nil
}

func (cd *categoryRepository) DeleteCategory(CategoryID string) error {

	id, err := strconv.Atoi(CategoryID)
	if err != nil {
		return errors.New("error converting CategoryID to integer")
	}

	result := cd.DB.Exec("DELETE FROM categories WHERE id = ?", id)
	errDelete := fmt.Sprintf("category not exist with the id = %d", id)
	if result.RowsAffected < 1 {
		return errors.New(errDelete)
	}

	return nil
}
