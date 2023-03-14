package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eduardonunesp/gservice/services"

	"github.com/gin-gonic/gin"
)

type DataCreateRequest struct {
	Name  string `form:"name" json:"name" xml:"name" binding:"required"`
	Stage int    `form:"stage" json:"stage" xml:"stage" binding:"required"`
	Score int    `form:"score" json:"score" xml:"score" binding:"required"`
}

type DataCreateResponse struct {
	ID string `json:"id"`
}

type GetDataResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Stage int    `json:"stage"`
	Score int    `json:"score"`
}

type DataController interface {
	GetDataByName(*gin.Context)
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
	results, err := pdc.Service.GetAll()

	if err != nil {
		log.Printf("Internal error %+v\n", err)

		c.JSON(500, gin.H{
			"error": "failed to get data",
		})

		return
	}

	resultsData := make([]GetDataResponse, len(results))
	for i, result := range results {
		resultsData[i] = GetDataResponse{
			ID:    result.UUID4,
			Name:  result.Name,
			Stage: result.Stage,
			Score: result.Score,
		}
	}

	c.JSON(200, resultsData)
}

func (pdc dataController) GetDataByName(c *gin.Context) {
	name := c.Param("name")

	result, err := pdc.Service.GetByName(name)

	if err != nil {
		log.Printf("Internal error %+v\n", err)

		c.JSON(500, gin.H{
			"error": "failed to get data",
		})

		return
	}

	c.JSON(200, GetDataResponse{
		ID:    result.UUID4,
		Name:  result.Name,
		Stage: result.Stage,
		Score: result.Score,
	})
}

func (pdc dataController) PostData(c *gin.Context) {
	var json DataCreateRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := pdc.Service.Insert(json.Name, json.Stage, json.Score)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed to insert post data",
			"cause": fmt.Sprintf("Internal error %+v\n", err),
		})
		return
	}

	c.JSON(201, DataCreateResponse{
		ID: result.UUID4,
	})
}
