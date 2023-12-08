package data

// import (
// 	"time"

// 	"gorm.io/gorm"
// )

// type CommentQuery struct {
// 	db *gorm.DB
// }

// func NewCommentQuery(db *gorm.DB) CommentDataInterface {
// 	return &CommentQuery{
// 		db: db,
// 	}
// }

// // CreateComment implements comment.CommentDataInterface.
// func (repo *CommentQuery) CreateComment(input CreateCommentInput) (*Comment, error) {
// 	commentModel := Comment{
// 		PostID:    input.PostID,
// 		UserID:    input.UserID,
// 		Content:   input.Content,
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}

// 	result := repo.db.Create(&commentModel)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return &commentModel, nil
// }

// // GetComment implements comment.CommentDataInterface.
// func (repo *CommentQuery) GetComment(commentID uint) (*Comment, error) {
// 	var commentModel Comment
// 	result := repo.db.First(&commentModel, commentID)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return &commentModel, nil
// }

// // UpdateComment implements comment.CommentDataInterface.
// func (repo *CommentQuery) UpdateComment(commentID uint, input UpdateCommentInput) (*Comment, error) {
// 	var commentModel Comment
// 	result := repo.db.First(&commentModel, commentID)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	commentModel.Content = input.Content
// 	commentModel.UpdatedAt = time.Now()

// 	updateResult := repo.db.Save(&commentModel)
// 	if updateResult.Error != nil {
// 		return nil, updateResult.Error
// 	}

// 	return &commentModel, nil
// }

// // DeleteComment implements comment.CommentDataInterface.
// func (repo *CommentQuery) DeleteComment(commentID uint) error {
// 	result := repo.db.Delete(&Comment{}, commentID)
// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	return nil
// }

// // GetCommentsByPost implements comment.CommentDataInterface.
// func (repo *CommentQuery) GetCommentsByPost(postID uint) ([]UserComment, error) {
// 	var comments []Comment
// 	result := repo.db.Where("post_id = ?", postID).Find(&comments)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	var userComments []UserComment
// 	for _, comment := range comments {
// 		userComment := UserComment{
// 			ID:        comment.ID,
// 			PostID:    comment.PostID,
// 			UserID:    comment.UserID,
// 			Username:  "PlaceholderUsername", // Gantilah dengan pengambilan username dari user service
// 			Content:   comment.Content,
// 			CreatedAt: comment.CreatedAt,
// 			UpdatedAt: comment.UpdatedAt,
// 		}
// 		userComments = append(userComments, userComment)
// 	}

// 	return userComments, nil
// }
