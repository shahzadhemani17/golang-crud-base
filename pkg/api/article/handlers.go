package article

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"weight-tracker/pkg/api/author"
	"weight-tracker/pkg/api/category"
)

// Routes Exports all routes handled by this service
func Routes(router *gin.Engine, svc ArticleService, authorService author.AuthorService, categoryService category.CategoryService) {
	subRouter := router.Group("/articles")
	{
		subRouter.POST("", func(c *gin.Context) {
			CreateArticleHandler(svc, authorService, categoryService, c)
		})

		subRouter.GET("/:article_id", func(c *gin.Context) {
			GetArticleHandler(svc, authorService, categoryService, c)
		})

		subRouter.GET("", func(c *gin.Context) {
			GetArticlesHandler(svc, authorService, categoryService, c)
		})

		subRouter.PUT("/:article_id", func(c *gin.Context) {
			UpdateArticleHandler(svc, authorService, categoryService, c)
		})

		subRouter.DELETE("/:article_id", func(c *gin.Context) {
			DeleteArticleHandler(svc, c)
		})
	}
}

func DeleteArticleHandler(svc ArticleService, c *gin.Context) {
	fmt.Println("Handling DELETE /articles request")

	id := c.Param("article_id")
	log.Println(fmt.Sprintf("Requested ID = %v", id))

	articleIDInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("The provided article id is not a valid number: %v", err)})
		return
	}
	_, err = svc.GetArticle(uint(articleIDInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	err = svc.DeleteArticle(uint(articleIDInt))
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "article deleted successfully"})
}

func UpdateArticleHandler(svc ArticleService, authorService author.AuthorService, categoryService category.CategoryService, c *gin.Context) {
	fmt.Println("Handling PUT /articles request")

	article := Article{}
	if err := c.ShouldBind(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("The JSON body is invalid: %v", err)})
		return
	}

	id := c.Param("article_id")
	log.Println(fmt.Sprintf("Requested ID = %v", id))

	articleIDInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("The provided article id is not a valid number: %v", err)})
		return
	}
	//Check if author exists
	_, err = authorService.GetAuthor(article.AuthorId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("author not found: %v", err.Error())})
		return
	}
	//Check if category exists
	_, err = categoryService.GetCategory(article.CategoryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("category not found: %v", err.Error())})
		return
	}

	article.Id = uint(articleIDInt)
	err = svc.UpdateArticle(article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Something went wrong updating the article: %v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "article updated successfully"})
}

func GetArticlesHandler(svc ArticleService, authorService author.AuthorService, categoryService category.CategoryService, c *gin.Context) {
	fmt.Println("Handling GET /articles request")

	articlesResp := make([]ArticleResponse, 0)
	articles, err := svc.GetArticles()
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Something went wrong getting the articles: %v", err))
		return
	}
	for _, a := range articles {
		articleresp := ArticleResponse{}
		author, err := authorService.GetAuthor(a.AuthorId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("author not found for articel id: %d : %v", a.Id, err.Error())})
			return
		}
		category, err := categoryService.GetCategory(a.CategoryId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("category not found for articel id: %d : %v", a.Id, err.Error())})
			return
		}
		articleresp.Name = a.Name
		articleresp.Description = a.Description
		articleresp.Author = author.Name
		articleresp.Category = category.Name
		articlesResp = append(articlesResp, articleresp)
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": articlesResp})
}

func GetArticleHandler(svc ArticleService, authorService author.AuthorService, categoryService category.CategoryService, c *gin.Context) {
	fmt.Println("Handling GET /articles request")

	articleResp := ArticleResponse{}
	id := c.Param("article_id")
	log.Println(fmt.Sprintf("Requested UID = %v", id))

	articleIDInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("The provided article id is not a valid number: %v", err)})
		return
	}
	article, err := svc.GetArticle(uint(articleIDInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Something went wrong getting the articles: %v", err)})
		return
	}
	author, err := authorService.GetAuthor(article.AuthorId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("author not found: %v", err.Error())})
		return
	}
	category, err := categoryService.GetCategory(article.CategoryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("category not found: %v", err.Error())})
		return
	}
	articleResp.Name = article.Name
	articleResp.Description = article.Description
	articleResp.Author = author.Name
	articleResp.Category = category.Name
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": articleResp})
}

func CreateArticleHandler(svc ArticleService, authorService author.AuthorService, categoryService category.CategoryService, c *gin.Context) {
	fmt.Println("Handling POST /articles func")

	article := Article{}
	if err := c.ShouldBind(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("The JSON body is invalid: %v", err)})
		return
	}
	//Check if author exists
	_, err := authorService.GetAuthor(article.AuthorId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("author not found: %v", err.Error())})
		return
	}
	//Check if category exists
	_, err = categoryService.GetCategory(article.CategoryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("category not found: %v", err.Error())})
		return
	}
	_, err = svc.CreateArticle(article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Something went wrong creating the article: %v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "article created successfully"})
}
