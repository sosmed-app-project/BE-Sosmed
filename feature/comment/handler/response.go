package handler

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// // GetCommentsHandler mengambil data semua comment atau comment berdasarkan filter
// func GetCommentsHandler(c *gin.Context) {
// 	// Implementasi untuk mendapatkan data comments dari database
// 	// ...

// 	// Dummy data untuk contoh response
// 	comments := []Comment{
// 		{ID: 1, PostID: 1, UserID: 1, Content: "Great post!"},
// 		{ID: 2, PostID: 1, UserID: 2, Content: "I agree!"},
// 	}

// 	// Konversi model ke response
// 	var commentResponses []CommentResponse
// 	for _, comment := range comments {
// 		commentResponses = append(commentResponses, CommentModelToResponse(comment))
// 	}

// 	c.JSON(http.StatusOK, gin.H{"comments": commentResponses})
// }

// // CreateCommentHandler membuat comment baru
// func CreateCommentHandler(c *gin.Context) {
// 	var comment Comment
// 	if err := c.ShouldBindJSON(&comment); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
// 		return
// 	}

// 	// Implementasi untuk menyimpan comment ke database
// 	// ...

// 	// Konversi model ke response
// 	commentResponse := CommentModelToResponse(comment)

// 	c.JSON(http.StatusCreated, gin.H{"comment": commentResponse})
// }

// // Fungsi-fungsi ini harus diimplementasikan untuk berinteraksi dengan database
// // ...

// // CommentModelToResponse mengonversi model Comment ke dalam bentuk response yang diinginkan
// func CommentModelToResponse(model Comment) CommentResponse {
// 	return CommentResponse{
// 		ID:        model.ID,
// 		PostID:    model.PostID,
// 		UserID:    model.UserID,
// 		Content:   model.Content,
// 		CreatedAt: model.CreatedAt,
// 		UpdatedAt: model.UpdatedAt,
// 	}
// }
