package api

import (
	"api-gateway/grpc/client"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"

	_ "api-gateway/api/docs"
	"api-gateway/api/handlers"
	"api-gateway/api/middleware"
	"api-gateway/config/logger"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewRouter(connF *client.GrpcClients, logger logger.Logger) *gin.Engine {
	h := handlers.NewHandler(connF, logger)
	router := gin.Default()

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	protected := router.Group("/", middleware.JWTMiddleware())
	
	// Comment routes
	comment := protected.Group("/comment")
	comment.POST("/", h.CommentCreate)
	comment.GET("/:id", h.CommentGetBydId)
	comment.PUT("/:id", h.ComentUpdate)
	comment.DELETE("/:id", h.CommentDelete)
	protected.GET("/comments", h.CommentGetAll)
	
	// Category routes
	category := protected.Group("/category")
	category.POST("/", h.CategoryCreate)
	category.GET("/:id", h.CategoryGetBydId)
	category.PUT("/:id", h.CategoryUpdate)
	category.DELETE("/:id", h.CategoryDelete)
	protected.GET("/categories", h.CategoryGetAll)
	protected.GET("categories/:category_id/posts", h.GetPostByCategory)
	
	// Tag routes
	tag := protected.Group("/tag")
	tag.POST("/", h.TagCreate)
	tag.GET("/:id", h.TagGetBydId)
	tag.PUT("/:id", h.TagUpdate)
	tag.DELETE("/:id", h.TagDelete)
	protected.GET("/tags", h.TagGetAll)
	protected.GET("/tags/popular", h.GetPopularTag)
	protected.GET("/tags/:tag_id/posts", h.GetPostByTag)

	// Post routes
	post := protected.Group("/post")
	post.POST("/", h.PostCreate)
	post.GET("/:id", h.PostGetBydId)
	post.PUT("/:id", h.PostUpdate)
	post.DELETE("/:id", h.PostDelete)
	protected.GET("/posts", h.PostGetAll)
	protected.GET("/posts/search", h.SearchPost)
	protected.GET("/posts/:post_id/comments", h.GetCommentByPost)

	// PostTag routes
	post_tag := protected.Group("/post_tag")
	post_tag.POST("/", h.PostTagCreate)
	post_tag.GET("/:id", h.PostTagGetBydId)
	post_tag.PUT("/:id", h.PostTagUpdate)
	post_tag.DELETE("/:id", h.PostTagDelete)
	protected.GET("/post_tags", h.PostTagGetAll)

	//User routes
	protected.GET("users/:user_id/posts", h.GetPostByUser)

	return router
}
