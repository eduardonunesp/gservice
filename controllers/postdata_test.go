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
	repo := repos.NewPostDataRepo(db)
	service := services.NewPostDataService(repo)

	title := "Title Test"

	// Inserting register on db to test
	db.Create(models.PostData{
		Title:         title,
		UUID4:         uuid.New().String(),
		UnixTimestamp: time.Now().UTC().Unix(),
	})

	controller := NewPostDataController(service)

	controller.GetPostData(c)

	assert.Equal(t, 200, w.Code)

	b, err := ioutil.ReadAll(w.Body)

	if err != nil {
		t.Error(err)
		return
	}

	var results []struct {
		UUID4     string `json:"UUID4"`
		Title     string `json:"Title"`
		Timestamp string `json:"Timestamp"`
	}

	err = json.Unmarshal(b, &results)

	if err != nil {
		t.Error(err)
		return
	}

	if len(results) != 1 {
		t.Error("Result should return one register")
	}

	if results[0].Title != title {
		t.Error("Wrong title returned from the http call")
	}
}

func TestPostData(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	db := models.GetTestDB()
	repo := repos.NewPostDataRepo(db)
	service := services.NewPostDataService(repo)
	controller := NewPostDataController(service)
	router.POST("/post-data", controller.SavePostData)

	buf := strings.NewReader(`{"title": "mydata"}`)
	req, err := http.NewRequest("POST", "/post-data", buf)

	if err != nil {
		t.Error(err)
	}

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 201)

	// Check if has inserted
	var result []models.PostData
	err = db.Find(&result).Error

	if err != nil {
		t.Error(err)
		return
	}

	if len(result) == 0 {
		t.Error("Result should greater than 0")
	}
}
