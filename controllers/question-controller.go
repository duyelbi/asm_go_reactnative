package controllers

import (
	"crud-mysql/dal"
	"crud-mysql/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

//SetupQuestionRouter for question router
func SetupQuestionRouter(r *gin.Engine) {
	//Create question
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

	//Read a question
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

	//Read all question
	r.GET("/questions", func(c *gin.Context) {
		question, err := dal.GetAllQuestion()
		if err != nil {
			c.JSON(500, gin.H{
				"messages": "Question not found ok",
			})
		} else {
			c.JSON(200, question)
		}
	})

	//Update a question
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

	//Delete a question
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

//SetupAnswerRouter for answer router
	//Create answer
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
						"messages":   "Insert Answer complete",
						"answerId": lastInsertedId,
					})
				}
			}
		}
	})

	//Read a answer
	r.GET("/answer/:id", func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		answer, err := dal.GetAnswer(id)
		if err != nil {
			c.JSON(500, gin.H{
				"messages": "Answer not found",
			})
		} else {
			c.JSON(200, answer)
		}
	})

	//Update a answer
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

	//Delete a answer
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

	//Read all answer
	r.GET("/answers", func(c *gin.Context) {
		answer, err := dal.GetAllAnswer()
		if err != nil {
			c.JSON(500, gin.H{
				"messages": "Answer not found ok",
			})
		} else {
			c.JSON(200, answer)
		}
	})


	//Create questionAnswer
	r.POST("/questionanswer", func(c *gin.Context) {
		var questionAnswer models.QuestionAnswer
		if err := c.ShouldBindJSON(&questionAnswer); err == nil {
			rowsAffected, lastInsertedId, err := dal.InsertQuestionAnswer(questionAnswer)
			if err != nil {
				c.JSON(500, gin.H{
					"messages": "Insert QuestionAnswer error",
				})
			} else {
				if rowsAffected > 0 {
					c.JSON(200, gin.H{
						"messages":   "Insert QuestionAnswer complete",
						"answerId": lastInsertedId,
					})
				}
			}
		}
	})

	//Read a questionAnswer
	r.GET("/questionanswer/:id", func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		questionAnswer, err := dal.GetQuestionAnswer(id)
		if err != nil {
			c.JSON(500, gin.H{
				"messages": "QuestionAnswer not found",
			})
		} else {
			c.JSON(200, questionAnswer)
		}
	})

	//Read all questionAnswer
	r.GET("/questionanswers", func(c *gin.Context) {
		questionAnswer, err := dal.GetAllQuestionAnswer()
		if err != nil {
			c.JSON(500, gin.H{
				"messages": "QuestionAnswer not found ok",
			})
		} else {
			c.JSON(200, questionAnswer)
		}
	})

	r.GET("/questionanswers/:id", func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		questionAnswer, err := dal.GetQuestionAnswerID(id)
		if err != nil {
			c.JSON(500, gin.H{
				"messages": "QuestionAnswer not found ok",
			})
		} else {
			c.JSON(200, questionAnswer)
		}
	})


	//Update a questionAnswer
	r.PUT("/questionanswer", func(c *gin.Context) {
		var questionAnswer models.QuestionAnswer
		if err := c.ShouldBindJSON(&questionAnswer); err == nil {
			rowsAffected, err := dal.UpdateQuestionAnswer(questionAnswer)
			if err != nil {
				c.JSON(500, gin.H{
					"messages": "update QuestionAnswer error",
				})
			} else {
				if rowsAffected > 0 {
					c.JSON(200, gin.H{
						"messages": "update QuestionAnswer complete",
					})
				}
			}
		}
	})

	//Delete a questionAnswer
	r.DELETE("/questionanswers/:id", func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		rowsDeletedAffected, err := dal.DeleteQuestionAnswer(id)
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