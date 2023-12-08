package like

// import (
// 	"time"

// 	"gorm.io/gorm"
// )

// // Like model
// type Like struct {
// 	ID        uint      `json:"id"`
// 	UserID    uint      `json:"user_id"`
// 	PostID    uint      `json:"post_id"`
// 	CountLike bool      `json:"count_like"`
// 	CreatedAt time.Time `json:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at"`
// }

// type LikeDataInterface interface {
// 	GetLikes() ([]Like, error)
// 	CreateLike(input Like) error
// 	// Fungsi-fungsi lainnya sesuai kebutuhan
// }

// type LikeData struct {
// 	db *gorm.DB
// }

// func NewLikeData(db *gorm.DB) LikeDataInterface {
// 	return &LikeData{
// 		db: db,
// 	}
// }

// // Implementasikan fungsi-fungsi dari LikeDataInterface
