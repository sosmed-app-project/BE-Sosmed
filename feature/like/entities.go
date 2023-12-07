package like

import "time"

// Like model
type Like struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	PostID    uint      `json:"post_id"`
	CountLike bool      `json:"count_like"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
