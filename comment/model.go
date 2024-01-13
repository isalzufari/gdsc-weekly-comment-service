package comment

import (
	"database/sql"
)

// Comment is a struct that represents a comment
type Comment struct {
	ID        int
	UserID    int
	ProductID int
	Comment   string
}

type CommentModel struct {
	db *sql.DB
}

func NewCommentModel(db *sql.DB) *CommentModel {
	return &CommentModel{db}
}

// implement model
