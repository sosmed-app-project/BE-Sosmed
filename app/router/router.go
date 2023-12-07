// router.go
package main

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Endpoints for each entity
	userRoutes(r)
	postRoutes(r)
	commentRoutes(r)
	likeRoutes(r)

	return r
}

func userRoutes(r *gin.Engine) {
	userGroup := r.Group("/users")
	{
		userGroup.GET("/:user_id", GetUserHandler)
		// Tambahkan rute-rute lain untuk entitas user sesuai kebutuhan
	}
}

func postRoutes(r *gin.Engine) {
	postGroup := r.Group("/posting")
	{
		postGroup.GET("", GetPostsHandler)
		postGroup.POST("", CreatePostHandler)
		postGroup.GET("/:posting_id", GetPostHandler)
		postGroup.PUT("/:posting_id", UpdatePostHandler)
		postGroup.DELETE("/:posting_id", DeletePostHandler)
		// Tambahkan rute-rute lain untuk entitas posting sesuai kebutuhan
	}
}

func commentRoutes(r *gin.Engine) {
	commentGroup := r.Group("/comment")
	{
		commentGroup.GET("", GetCommentsHandler)
		commentGroup.POST("", CreateCommentHandler)
		commentGroup.GET("/:comment_id", GetCommentHandler)
		commentGroup.DELETE("/:comment_id", DeleteCommentHandler)
		// Tambahkan rute-rute lain untuk entitas comment sesuai kebutuhan
	}
}

func likeRoutes(r *gin.Engine) {
	likeGroup := r.Group("/like")
	{
		likeGroup.GET("", GetLikesHandler)
		likeGroup.POST("", CreateLikeHandler)
		likeGroup.DELETE("/:like_id", DeleteLikeHandler)
		// Tambahkan rute-rute lain untuk entitas like sesuai kebutuhan
	}
}

func main() {
	r := SetupRouter()

	// Run the server
	r.Run(":8080")
}
