// post.go
package handlers

import (
	"context"
	"log"
	"net/http"
	"runtime"
	"strconv"

	pbf "api-gateway/genproto/forum" // Import generated protobuf package

	"github.com/gin-gonic/gin"
)

// PostCreate handles the creation of a new post.
// @Summary Create post
// @Description Create a new post
// @Tags post
// @Accept json
// @Produce json
// @Param post body pbf.PostCreateReq true "Post data"
// @Success 200 {object} pbf.Post
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /post [post]
func (h *Handler) PostCreate(c *gin.Context) {
	var req pbf.PostCreateReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	res, err := h.srvs.Post.Create(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		_, filename, line, _ := runtime.Caller(1)
		log.Printf("[error] %s:%d %v", filename, line, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

// PostGet handles getting a post by ID.
// @Summary Get post
// @Description Get a post by ID
// @Tags post
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} pbf.PostRes
// @Failure 400 {object} string "Invalid post ID"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /post/{id} [get]
func (h *Handler) PostGetBydId(c *gin.Context) {
	id := &pbf.GetByIdReq{Id: c.Param("id")}
	res, err := h.srvs.Post.GetById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get post", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// PostGetAll handles getting all posts.
// @Summary Get all posts
// @Description Get all posts
// @Tags post
// @Accept json
// @Produce json
// @Param limit query integer false "Limit"
// @Param offset query integer false "Offset"
// @Success 200 {object} pbf.PostGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /posts [get]
func (h *Handler) PostGetAll(c *gin.Context) {
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	var limit, offset int
	var err error
	if limitStr == "" {
		limit = 10
	} else {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
			return
		}
	}
	if offsetStr == "" {
		offset = 0
	} else {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
			return
		}
	}

	filter := pbf.Filter{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	res, err := h.srvs.Post.GetAll(context.Background(), &filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get posts", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// PostUpdate handles updating an existing post.
// @Summary Update post
// @Description Update an existing post
// @Tags post
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Param post body pbf.PostCreateReq true "Updated post data"
// @Success 200 {object} pbf.Post
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Post not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /post/{id} [put]
func (h *Handler) PostUpdate(c *gin.Context) {
	id := c.Param("id")
	var req pbf.PostUpdate

	if err := c.BindJSON(&req.UpdatePost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	req.Id = id
	res, err := h.srvs.Post.Update(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update post", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// PostDelete handles deleting a post by ID.
// @Summary Delete post
// @Description Delete a post by ID
// @Tags post
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} string "Post deleted"
// @Failure 400 {object} string "Invalid post ID"
// @Failure 404 {object} string "Post not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /post/{id} [delete]
func (h *Handler) PostDelete(c *gin.Context) {
	id := &pbf.GetByIdReq{Id: c.Param("id")}
	_, err := h.srvs.Post.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete post", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}

// PostGetAll handles getting all posts by user.
// @Summary Get all posts by user
// @Description Get all posts user
// @Tags user
// @Accept json
// @Produce json
// @Param user_id query string false "UserId"
// @Param limit query integer false "Limit"
// @Param offset query integer false "Offset"
// @Success 200 {object} pbf.PostGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /users/{user_id}/posts [get]
func (h *Handler) GetPostByUser(c *gin.Context){

	user_id := c.Query("user_id")
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	var limit, offset int
	var err error
	if limitStr == "" {
		limit = 10
	} else {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
			return
		}
	}
	if offsetStr == "" {
		offset = 0
	} else {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
			return
		}
	}

	filter := pbf.GetFilter{
		Id: user_id,
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	res, err := h.srvs.Post.GetPostByUser(context.Background(), &filter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get posts", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}


// PostGetAll handles getting all posts by category.
// @Summary Get all posts by category
// @Description Get all posts by category
// @Tags category
// @Accept json
// @Produce json
// @Param category_id query string false "CategoryId"
// @Param limit query integer false "Limit"
// @Param offset query integer false "Offset"
// @Success 200 {object} pbf.PostGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /categories/{category_id}/posts [get]
func (h *Handler) GetPostByCategory(c *gin.Context){

	category_id := c.Query("category_id")
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	var limit, offset int
	var err error
	if limitStr == "" {
		limit = 10
	} else {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
			return
		}
	}
	if offsetStr == "" {
		offset = 0
	} else {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
			return
		}
	}

	filter := pbf.GetFilter{
		Id: category_id,
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	res, err := h.srvs.Post.GetPostByCategory(context.Background(), &filter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get posts", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// PostGetAll handles getting all posts.
// @Summary Get all posts
// @Description Get all posts
// @Tags post
// @Accept json
// @Produce json
// @Param title query string false "Title"
// @Param body query string false "Body"
// @Param limit query integer false "Limit"
// @Param offset query integer false "Offset"
// @Success 200 {object} pbf.PostGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /posts/search [get]
func (h *Handler) SearchPost(c *gin.Context){

	title := c.Query("title")
	body := c.Query("body")
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	var limit, offset int
	var err error
	if limitStr == "" {
		limit = 10
	} else {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
			return
		}
	}
	if offsetStr == "" {
		offset = 0
	} else {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
			return
		}
	}

	filter := pbf.PostSearch{
		Title: title,
		Body: body,
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	res, err := h.srvs.Post.SearchPost(context.Background(), &filter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get posts", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}