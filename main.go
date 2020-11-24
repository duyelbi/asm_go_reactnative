package main

import (
	"crud-mysql/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	controllers.SetupQuestionRouter(r)
	// controllers.SetupAnswerRouter(r)
	r.Run(":8080")
}
