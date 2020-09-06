package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/eduardonunesp/gservice/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostData struct {
	Title string `form:"title" json:"title" xml:"title"  binding:"required"`
}

type PostDataController interface {
	GetPostData(*gin.Context)
	SavePostData(*gin.Context)
}

type postDataController struct {
	Service services.PostDataService
}

func NewPostDataController(service services.PostDataService) PostDataController {
	return &postDataController{service}
}

func (pdc postDataController) GetPostData(c *gin.Context) {
	title := c.Param("title")

	if len(title) == 0 {
		results, err := pdc.Service.GetAll()

		if err != nil {
			log.Printf("Internal error %+v\n", err)

			c.JSON(500, gin.H{
				"error": "failed to get data",
			})

			return
		}

		c.JSON(200, results)
		return
	}

	result, err := pdc.Service.GetByTitle(title)

	if err != nil {
		log.Printf("Internal error %+v\n", err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{
				"error": "title not found",
			})

			return
		}

		c.JSON(500, gin.H{
			"error": "failed to get post data",
		})

		return
	}

	c.JSON(200, result)
}

func (pdc postDataController) SavePostData(c *gin.Context) {
	var json PostData
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pdc.Service.Insert(json.Title); err != nil {
		c.JSON(500, gin.H{
			"error": "failed to insert post data",
			"cause": fmt.Sprintf("Internal error %+v\n", err),
		})
		return
	}

	c.JSON(201, gin.H{
		"msg": "post inserted with success",
	})
}
