package controllers

import (
	"crud-mysql/dal"
	"crud-mysql/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

//SetupAnswerRouter for answer router
func SetupAnswerRouter(r *gin.Engine) {
	//Create
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

	//Read
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
