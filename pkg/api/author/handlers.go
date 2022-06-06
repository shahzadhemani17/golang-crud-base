package author

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// Routes Exports all routes handled by this service
func Routes(router *gin.Engine, svc AuthorService) {
	subRouter := router.Group("/authors")
	{
		subRouter.POST("", func(c *gin.Context) {
			CreateAuthorHandler(svc, c)
		})

		subRouter.GET("/:author_id", func(c *gin.Context) {
			GetAuthorHandler(svc, c)
		})

		subRouter.GET("", func(c *gin.Context) {
			GetAuthorsHandler(svc, c)
		})

		subRouter.PUT("/:author_id", func(c *gin.Context) {
			UpdateAuthorHandler(svc, c)
		})

		subRouter.DELETE("/:author_id", func(c *gin.Context) {
			DeleteAuthorHandler(svc, c)
		})
	}
}

func DeleteAuthorHandler(svc AuthorService, c *gin.Context) {
	fmt.Println("Handling DELETE /authors request")

	id := c.Param("author_id")
	log.Println(fmt.Sprintf("Requested ID = %v", id))

	authorIDInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("The provided author id is not a valid number: %v", err)})
		return
	}
	_, err = svc.GetAuthor(uint(authorIDInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	err = svc.DeleteAuthor(uint(authorIDInt))
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "author deleted successfully"})
}

func UpdateAuthorHandler(svc AuthorService, c *gin.Context) {
	fmt.Println("Handling PUT /authors request")

	author := Author{}
	if err := c.ShouldBind(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("The JSON body is invalid: %v", err)})
		return
	}

	id := c.Param("author_id")
	log.Println(fmt.Sprintf("Requested ID = %v", id))

	authorIDInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("The provided author id is not a valid number: %v", err)})
		return
	}
	author.Id = uint(authorIDInt)
	err = svc.UpdateAuthor(author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Something went wrong updating the author: %v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "author updated successfully"})
}

func GetAuthorsHandler(svc AuthorService, c *gin.Context) {
	fmt.Println("Handling GET /authors request")

	authors, err := svc.GetAuthors()
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Something went wrong getting the authors: %v", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": authors})
}

func GetAuthorHandler(svc AuthorService, c *gin.Context) {
	fmt.Println("Handling GET /authors request")

	id := c.Param("author_id")
	log.Println(fmt.Sprintf("Requested ID = %v", id))

	authorIDInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("The provided author id is not a valid number: %v", err)})
		return
	}
	author, err := svc.GetAuthor(uint(authorIDInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Something went wrong getting the author: %v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": author})
}

func CreateAuthorHandler(svc AuthorService, c *gin.Context) {
	fmt.Println("Handling POST /authors func")

	author := Author{}
	if err := c.ShouldBind(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("The JSON body is invalid: %v", err)})
		return
	}

	_, err := svc.CreateAuthor(author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Something went wrong creating the author: %v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "author created successfully"})
}
