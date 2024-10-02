package handler

import (
	"net/http"
	services "project/pkg/usecase/interface"
	response "project/pkg/utils/Response"
	"project/pkg/utils/domain"
	"project/pkg/utils/models"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	CategoryUseCase services.CategoryUseCase
}

func NewCategoryHandler(usecase services.CategoryUseCase) *CategoryHandler {

	return &CategoryHandler{
		CategoryUseCase: usecase,
	}

}

func (cat *CategoryHandler) AddCategory(c *gin.Context) {
	var category domain.Category
	if err := c.BindJSON(&category); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	CategoryResponse, err := cat.CategoryUseCase.AddCategory(category.Category)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add the Category", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added Category", CategoryResponse, nil)
	c.JSON(http.StatusOK, successRes)

}

func (Cat *CategoryHandler) GetCategory(c *gin.Context) {

	categories, err := Cat.CategoryUseCase.GetCategories()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "couldn't get categories", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully got all categories", categories, nil)
	c.JSON(http.StatusOK, successRes)

}
func (cat *CategoryHandler) UpdateCategory(c *gin.Context) {

	var update models.Rename

	if err := c.BindJSON(&update); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	categories, err := cat.CategoryUseCase.UpdateCategory(update.Current, update.New)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't update the  category", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfullly updated the categories", categories, nil)
	c.JSON(http.StatusOK, successRes)

}

func (cat *CategoryHandler) DeleteCategory(c *gin.Context) {
	CategoryID := c.Query("id")
	err := cat.CategoryUseCase.DeleteCategory(CategoryID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't delete the category", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	succesRes := response.ClientResponse(http.StatusOK, "successfully deleted category", nil, nil)
	c.JSON(http.StatusOK, succesRes)
}
