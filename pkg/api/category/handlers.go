package category

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// Routes Exports all routes handled by this service
func Routes(router *gin.Engine, svc CategoryService) {
	subRouter := router.Group("/category")
	{
		subRouter.POST("", func(c *gin.Context) {
			CreateCategoryHandler(svc, c)
		})

		subRouter.GET("/:category_id", func(c *gin.Context) {
			GetCategoryHandler(svc, c)
		})

		subRouter.GET("", func(c *gin.Context) {
			GetCategorysHandler(svc, c)
		})

		subRouter.PUT("/:category_id", func(c *gin.Context) {
			UpdateCategoryHandler(svc, c)
		})

		subRouter.DELETE("/:category_id", func(c *gin.Context) {
			DeleteCategoryHandler(svc, c)
		})
	}
}

func DeleteCategoryHandler(svc CategoryService, c *gin.Context) {
	fmt.Println("Handling DELETE /category request")

	id := c.Param("category_id")
	log.Println(fmt.Sprintf("Requested ID = %v", id))

	categoryIDInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("The provided category id is not a valid number: %v", err)})
		return
	}
	_, err = svc.GetCategory(uint(categoryIDInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	err = svc.DeleteCategory(uint(categoryIDInt))
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "category deleted successfully"})
}

func UpdateCategoryHandler(svc CategoryService, c *gin.Context) {
	fmt.Println("Handling PUT /category request")

	category := Category{}
	if err := c.ShouldBind(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("The JSON body is invalid: %v", err)})
		return
	}

	id := c.Param("category_id")
	log.Println(fmt.Sprintf("Requested ID = %v", id))

	categoryIDInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("The provided category id is not a valid number: %v", err)})
		return
	}
	category.Id = uint(categoryIDInt)
	err = svc.UpdateCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Something went wrong updating the category: %v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "category updated successfully"})
}

func GetCategorysHandler(svc CategoryService, c *gin.Context) {
	fmt.Println("Handling GET /category request")

	categories, err := svc.GetCategories()
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Something went wrong getting the categories: %v", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": categories})
}

func GetCategoryHandler(svc CategoryService, c *gin.Context) {
	fmt.Println("Handling GET /category request")

	id := c.Param("category_id")
	log.Println(fmt.Sprintf("Requested ID = %v", id))

	categoryIDInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("The provided category id is not a valid number: %v", err)})
		return
	}
	category, err := svc.GetCategory(uint(categoryIDInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Something went wrong getting the category: %v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": category})
}

func CreateCategoryHandler(svc CategoryService, c *gin.Context) {
	fmt.Println("Handling POST /category func")

	category := Category{}
	if err := c.ShouldBind(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("The JSON body is invalid: %v", err)})
		return
	}

	_, err := svc.CreateCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Something went wrong creating the category: %v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "category created successfully"})
}
