package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/quiz", quizHandler)

	return router
}

func quizHandler(c *gin.Context) {
	question := c.Query("question")
	answer := ""
	switch question {
	case "first-ques":
		answer = "The CBI stands for California Bureau of Investigation."
	case "second-ques":
		answer = "Patrick Jane has a keen sense of observation and deduction."
	default:
		answer = "Invalid question number."
	}
	c.JSON(http.StatusOK, gin.H{"answer": answer})
}
