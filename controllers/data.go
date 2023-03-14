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

type Data struct {
	Name  string `form:"name" json:"name" xml:"name"  binding:"required"`
	Stage int    `form:"stage" json:"stage" xml:"stage"  binding:"required"`
	Score int    `form:"score" json:"score" xml:"score"  binding:"required"`
}

type DataController interface {
	GetData(*gin.Context)
	PostData(*gin.Context)
}

type dataController struct {
	Service services.DataService
}

func NewDataController(service services.DataService) DataController {
	return &dataController{service}
}

func (pdc dataController) GetData(c *gin.Context) {
	name := c.Param("name")

	if len(name) == 0 {
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

	result, err := pdc.Service.GetByName(name)

	if err != nil {
		log.Printf("Internal error %+v\n", err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{
				"error": "name not found",
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

func (pdc dataController) PostData(c *gin.Context) {
	var json Data
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pdc.Service.Insert(json.Name, json.Stage, json.Score); err != nil {
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
