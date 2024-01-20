package main

import (
	"commentservice/comment"
	"commentservice/db"

	"github.com/gin-gonic/gin"
)

func main() {

	db := db.OpenDB(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
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
