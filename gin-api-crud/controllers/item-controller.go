package controllers

import (
	"crud/dal"
	"crud/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

//SetupItemRouter for item router
func SetupItemRouter(r *gin.Engine) {
	//Create
	r.POST("/item", func(c *gin.Context) {
		var item model.Item
		if err := c.ShouldBindJSON(&item); err == nil {
			rowsAffected, lastInsertedId, err := dal.InsertItem(item)
			if err != nil {
				c.JSON(500, gin.H{
					"messages": "Insert Item error",
				})
			} else {
				if rowsAffected > 0 {
					c.JSON(200, gin.H{
						"messages": "Insert Item complete",
						"itemId":   lastInsertedId,
					})
				}
			}
		}
	})

	//Read
	r.GET("/item/:id", func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		item, err := dal.GetItem(id)
		if err != nil {
			c.JSON(500, gin.H{
				"messages": "Item not found",
			})
		} else {
			c.JSON(200, item)
		}
	})

	//Update
	r.PUT("/item", func(c *gin.Context) {
		var item model.Item
		if err := c.ShouldBindJSON(&item); err == nil {
			rowsAffected, err := dal.UpdateItem(item)
			if err != nil {
				c.JSON(500, gin.H{
					"messages": "update Item error",
				})
			} else {
				if rowsAffected > 0 {
					c.JSON(200, gin.H{
						"messages": "update Item complete",
					})
				}
			}
		}
	})

	//Delete
	r.DELETE("/item/:id", func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		rowsDeletedAffected, err := dal.DeleteItem(id)
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
