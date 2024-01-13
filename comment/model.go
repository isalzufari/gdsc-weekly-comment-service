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
func (m *CommentModel) GetCommentsByProductId(productID int) ([]Comment, error) {
	rows, err := m.db.Query(`
		SELECT id, userId, productId, comment
		FROM comments
		WHERE productId = ?
	`, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []Comment{}

	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.ID, &comment.UserID, &comment.ProductID, &comment.Comment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func (m *CommentModel) CreateComment(comment Comment) (int64, error) {
	result, err := m.db.Exec(`
		INSERT INTO comments (userId, productId, comment)
		VALUES (?, ?, ?)
	`, comment.UserID, comment.ProductID, comment.Comment)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (m *CommentModel) DeleteComment(commentID int) (int64, error) {
	result, err := m.db.Exec(`
		DELETE FROM comments
		WHERE id = ?
	`, commentID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
