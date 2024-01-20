package main

import (
	"commentservice/comment"
	"commentservice/db"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	db := db.OpenDB(
		"root",
		"e5H53cc5AB6-2caAa3baE6Eh2hf5h6H4",
		"viaduct.proxy.rlwy.net",
		"5003",
		"railway",
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
