package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type J map[string]interface{}

type Paste struct {
	gorm.Model
	Body string
}

type PasteRequest struct {
	Body string `json:"body"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&Paste{})

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:3000"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders: []string{"Content-Length"},
	}))
	r.GET("/", func(c *gin.Context) {
		var pastes []Paste
		db.Find(&pastes)
		c.JSON(200, J{"data": pastes})
	})

	r.GET("/paste/:id", func(c *gin.Context) {
		id := c.Param("id")
		parsed, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(422, J{"error": "invalid ID (ints only)"})
			return
		}

		var paste Paste
		result := db.First(&paste, parsed)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, J{"error": "not found"})
			return
		}

		c.JSON(200, J{"data": paste})
	})

	r.GET("/new", func(c *gin.Context) {
		body, ok := c.GetQuery("body")
		if !ok {
			c.JSON(422, J{"error": "missing body query (add '?body=<text>' )"})
			return
		}

		newPaste := Paste{Body: body}
		db.Create(&newPaste)

		c.JSON(200, J{"data": newPaste})
	})

	r.POST("/", func(c *gin.Context) {
		var newRequest PasteRequest
		if err := c.ShouldBindJSON(&newRequest); err != nil {
			c.JSON(422, J{"error": "malformed JSON (expected { \"body\": <xxx> }"})
			return
		}

		newPaste := Paste{Body: newRequest.Body}
		db.Create(&newPaste)

		c.JSON(200, J{"data": newPaste})
	})

	fmt.Println("\navailable routes")
	fmt.Println("=================")
	fmt.Println(`
GET / - Get all new pastes
GET /{id} - Get a specific paste
POST / - Create a new paste, expects a request: { "body": "sample text" }
GET /new - GET version of creating new pastes, body is passed in as a query string: /new?body=sample
`)

	r.Run("localhost:8000")
}
