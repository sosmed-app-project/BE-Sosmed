// comment/model.go

package data

import "time"

// Comment model
type Comment struct {
	ID        uint      `json:"id"`
	PostID    uint      `json:"post_id"`
	UserID    uint      `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserComment model
type UserComment struct {
	ID        uint      `json:"id"`
	PostID    uint      `json:"post_id"`
	UserID    uint      `json:"user_id"`
	Username  string    `json:"username"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateCommentInput model for creating a new comment
type CreateCommentInput struct {
	PostID  uint   `json:"post_id"`
	UserID  uint   `json:"user_id"`
	Content string `json:"content"`
}

// UpdateCommentInput model for updating an existing comment
type UpdateCommentInput struct {
	Content string `json:"content"`
}
