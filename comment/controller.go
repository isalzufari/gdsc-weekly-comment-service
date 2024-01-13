package comment

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentModel *CommentModel
}

func NewCommentController(commentModel *CommentModel) *CommentController {
	return &CommentController{commentModel}
}

// implement controller
func (c *CommentController) GetCommentsByProductId(ctx *gin.Context) {
	productID := ctx.Param("productId")

	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "gagal", "message": err.Error()})
		return
	}

	comments, err := c.commentModel.GetCommentsByProductId(productIDInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "gagal", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": comments})
}

func (c *CommentController) CreateComment(ctx *gin.Context) {
	productID := ctx.Param("productId")
	userID := ctx.Param("userId")

	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "gagal", "message": err.Error()})
		return
	}

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "gagal", "message": err.Error()})
		return
	}

	var comment Comment

	err = ctx.BindJSON(&comment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "gagal", "message": err.Error()})
		return
	}

	comment.ProductID = productIDInt
	comment.UserID = userIDInt

	id, err := c.commentModel.CreateComment(comment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "gagal", "message": err.Error()})
		return
	}

	comment.ID = int(id)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": comment})
}

func (c *CommentController) DeleteComment(ctx *gin.Context) {
	commentID := ctx.Param("commentId")

	commentIDInt, err := strconv.Atoi(commentID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "gagal", "message": err.Error()})
		return
	}

	_, err = c.commentModel.DeleteComment(commentIDInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "gagal", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
