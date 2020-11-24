package controllers

import (
	"crud-mysql/dal"
	"crud-mysql/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

//SetupQuestionRouter for question router
func SetupQuestionRouter(r *gin.Engine) {
	//Create
	r.POST("/question", func(c *gin.Context) {
		var question models.Question
		if err := c.ShouldBindJSON(&question); err == nil {
			rowsAffected, lastInsertedId, err := dal.InsertQuestion(question)
			if err != nil {
				c.JSON(500, gin.H{
					"messages": "Insert Question error",
				})
			} else {
				if rowsAffected > 0 {
					c.JSON(200, gin.H{
						"messages":   "Insert Question complete",
						"questionId": lastInsertedId,
					})
				}
			}
		}
	})

	//Read
	r.GET("/question/:id", func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		question, err := dal.GetQuestion(id)
		if err != nil {
			c.JSON(500, gin.H{
				"messages": "Question not found",
			})
		} else {
			c.JSON(200, question)
		}
	})

	//Read all
	r.GET("/question", func(c *gin.Context) {
		question, err := dal.GetAllQuestion()
		if err != nil {
			c.JSON(500, gin.H{
				"messages": "Question not found ok",
			})
		} else {
			c.JSON(200, question)
		}
	})

	//Update
	r.PUT("/question", func(c *gin.Context) {
		var question models.Question
		if err := c.ShouldBindJSON(&question); err == nil {
			rowsAffected, err := dal.UpdateQuestion(question)
			if err != nil {
				c.JSON(500, gin.H{
					"messages": "update Question error",
				})
			} else {
				if rowsAffected > 0 {
					c.JSON(200, gin.H{
						"messages": "update Question complete",
					})
				}
			}
		}
	})

	//Delete
	r.DELETE("/question/:id", func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		rowsDeletedAffected, err := dal.DeleteQuestion(id)
		if err != nil {
			c.JSON(500, gin.H{
				"messages": "delete error.",
			})
		} else {
			if rowsDeletedAffected > 0 {
				c.JSON(200, gin.H{
					"messages": "delete completed.",
				})
			}
		}
	})

	r.POST("/answer", func(c *gin.Context) {
		var answer models.Answer
		if err := c.ShouldBindJSON(&answer); err == nil {
			rowsAffected, lastInsertedId, err := dal.InsertAnswer(answer)
			if err != nil {
				c.JSON(500, gin.H{
					"messages": "Insert Answer error",
				})
			} else {
				if rowsAffected > 0 {
					c.JSON(200, gin.H{
						"messages": "Insert Answer complete",
						"answerId": lastInsertedId,
					})
				}
			}
		}
	})

	//Read
	r.GET("/answer/:id", func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		answer, err := dal.GetAnswer(id)
		if err != nil {
			c.JSON(500, gin.H{
				"messages": "Answer not found ok",
			})
		} else {
			c.JSON(200, answer)
		}
	})


	//Update
	r.PUT("/answer", func(c *gin.Context) {
		var answer models.Answer
		if err := c.ShouldBindJSON(&answer); err == nil {
			rowsAffected, err := dal.UpdateAnswer(answer)
			if err != nil {
				c.JSON(500, gin.H{
					"messages": "update Answer error",
				})
			} else {
				if rowsAffected > 0 {
					c.JSON(200, gin.H{
						"messages": "update Answer complete",
					})
				}
			}
		}
	})

	//Delete
	r.DELETE("/answer/:id", func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		rowsDeletedAffected, err := dal.DeleteAnswer(id)
		if err != nil {
			c.JSON(500, gin.H{
				"messages": "delete error.",
			})
		} else {
			if rowsDeletedAffected > 0 {
				c.JSON(200, gin.H{
					"messages": "delete completed.",
				})
			}
		}
	})
}
