package main

import (
	"fmt"
	"gservice/controllers"
	"gservice/models"
	"gservice/repos"
	"gservice/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	httpHost         string
	httpPort         string
	postgresUser     string
	postgresPassword string
	postgresDB       string
	postgresHost     string
	postgresPort     string
)

func init() {
	var found bool
	httpHost, found = os.LookupEnv("HTTP_HOST")

	if !found {
		httpHost = "localhost"
	}

	httpPort, found = os.LookupEnv("HTTP_PORT")

	if !found {
		httpPort = "3000"
	}

	postgresPort, found = os.LookupEnv("POSTGRES_PORT")

	if !found {
		postgresPort = "5432"
	}

	postgresHost, found = os.LookupEnv("POSTGRES_HOST")

	if !found {
		log.Fatal("env var POSTGRES_HOST not found")
	}

	postgresUser, found = os.LookupEnv("POSTGRES_USER")

	if !found {
		log.Fatal("env var POSTGRES_USER not found")
	}

	postgresPassword, found = os.LookupEnv("POSTGRES_PASSWORD")

	if !found {
		log.Fatal("env var POSTGRES_PASSWORD not found")
	}

	postgresDB, found = os.LookupEnv("POSTGRES_DB")

	if !found {
		log.Fatal("env var POSTGRES_DB not found")
	}
}

func main() {

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		postgresUser, postgresPassword, postgresDB,
		postgresHost, postgresPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error while loading the tables: %v", err)
		return
	}

	db.AutoMigrate(&models.PostData{})

	repo := repos.NewPostDataRepo(db)
	service := services.NewPostDataService(repo)
	postDataController := controllers.NewPostDataController(service)

	r := gin.Default()
	r.POST("/post-data", postDataController.SavePostData)
	r.GET("/post-data", postDataController.GetPostData)
	r.GET("/post-data/:title", postDataController.GetPostData)

	hostPort := fmt.Sprintf("%s:%s", httpHost, httpPort)
	log.Println("Server starting at ", hostPort)
	r.Run(hostPort)
}
