package comment

type CommentController struct {
	commentModel *CommentModel
}

func NewCommentController(commentModel *CommentModel) *CommentController {
	return &CommentController{commentModel}
}

// implement controller
