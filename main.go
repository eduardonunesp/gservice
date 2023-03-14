package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eduardonunesp/gservice/controllers"
	"github.com/eduardonunesp/gservice/models"
	"github.com/eduardonunesp/gservice/repos"
	"github.com/eduardonunesp/gservice/services"

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
	authUser         string
	authPass         string
)

// Testsss
//a Test 2aaa

func init() {
	var found bool
	httpHost, found = os.LookupEnv("HTTP_HOST")

	if !found {
		httpHost = "localhost"
	}

	httpPort, found = os.LookupEnv("PORT")

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

	authUser, found = os.LookupEnv("AUTH_USER")

	if !found {
		log.Fatal("env var AUTH_USER not found")
	}

	authPass, found = os.LookupEnv("AUTH_PASS")

	if !found {
		log.Fatal("env var AUTH_PASS not found")
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

	db.AutoMigrate(&models.Data{})

	repo := repos.NewDataRepo(db)
	service := services.NewDataService(repo)
	dataController := controllers.NewDataController(service)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the API",
		})
	})

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		authUser: authPass,
	}))

	authorized.POST("/data", dataController.PostData)
	authorized.GET("/data", dataController.GetData)
	authorized.GET("/data/:name", dataController.GetDataByName)

	hostPort := fmt.Sprintf("%s:%s", httpHost, httpPort)
	log.Println("Server starting at ", hostPort)
	r.Run(hostPort)
}
