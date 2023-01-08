package main

import (
	"net/http"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type PostTodoBody struct {
	Message string "json:message"
}

func main() {

	todos := []string{
		"buy carrots",
		"go for a walk",
		"buy some gin",
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/api/todos", func(context *gin.Context) {
		context.JSON(http.StatusOK, todos)
	})

	router.POST("/api/todos", func(context *gin.Context) {

		postTodoBody := PostTodoBody{}

		err := context.ShouldBind(&postTodoBody)

		if err != nil {
			panic("cannot bind!")
		}

		todos = append(todos, postTodoBody.Message)

		context.JSON(
			http.StatusOK,
			gin.H{
				"message": "OK",
			},
		)
	})

	router.DELETE("api/todos", func(context *gin.Context) {

		todos = []string{}

		context.JSON(
			http.StatusOK,
			gin.H{
				"message": "cleared",
			},
		)
	})

	router.Run()
}
