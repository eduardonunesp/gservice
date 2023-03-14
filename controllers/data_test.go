package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/eduardonunesp/gservice/models"
	"github.com/eduardonunesp/gservice/repos"
	"github.com/eduardonunesp/gservice/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
)

func TestGetData(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	db := models.GetTestDB()
	repo := repos.NewDataRepo(db)
	service := services.NewDataService(repo)

	name := "Title Test"
	stage := 1
	score := 100

	// Inserting register on db to test
	db.Create(models.Data{
		Name:          name,
		Stage:         stage,
		Score:         score,
		UUID4:         uuid.New().String(),
		UnixTimestamp: time.Now().UTC().Unix(),
	})

	controller := NewDataController(service)

	controller.GetData(c)

	assert.Equal(t, 200, w.Code)

	b, err := ioutil.ReadAll(w.Body)

	if err != nil {
		t.Error(err)
		return
	}

	var results []struct {
		UUID4     string `json:"UUID4"`
		Name      string `json:"name"`
		Stage     int    `json:"stage"`
		Score     int    `json:"score"`
		Timestamp string `json:"timestamp"`
	}

	err = json.Unmarshal(b, &results)

	if err != nil {
		t.Error(err)
		return
	}

	if len(results) != 1 {
		t.Error("Result should return one register")
	}

	if results[0].Name != name {
		t.Error("Wrong title returned from the http call")
	}
}

func TestData(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	db := models.GetTestDB()
	repo := repos.NewDataRepo(db)
	service := services.NewDataService(repo)
	controller := NewDataController(service)
	router.POST("/post-data", controller.PostData)

	buf := strings.NewReader(`{"name": "mydata", "stage": 1, "score": 100}`)
	req, err := http.NewRequest("POST", "/post-data", buf)

	if err != nil {
		t.Error(err)
	}

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 201)

	// Check if has inserted
	var result []models.Data
	err = db.Find(&result).Error

	if err != nil {
		t.Error(err)
		return
	}

	if len(result) == 0 {
		t.Error("Result should greater than 0")
	}
}
