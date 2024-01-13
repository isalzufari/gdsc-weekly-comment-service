package main

import (
	"commentservice/comment"
	"commentservice/db"

	"github.com/gin-gonic/gin"
)

func main() {

	db := db.OpenDB(
		"root",
		"root",
		"127.0.0.1",
		"3306",
		"micro_web_gdsc",
	)
	defer db.Close()

	commentModel := comment.NewCommentModel(db)
	commentController := comment.NewCommentController(commentModel)

	r := gin.Default()

	r.GET("/:productId", commentController.GetCommentsByProductId)
	r.POST("/:productId/user/:userId", commentController.CreateComment)
	r.DELETE("/:commentId/user/:userId", commentController.DeleteComment)

	port := "5003"
	r.Run(":" + port)
}
