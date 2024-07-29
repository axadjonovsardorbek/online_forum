// comment.go
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

// CommentCreate handles the creation of a new comment.
// @Summary Create comment
// @Description Create a new comment
// @Tags comment
// @Accept json
// @Produce json
// @Param comment body pbf.CommentCreateReq true "Comment data"
// @Success 200 {object} pbf.Comment
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /comment [post]
func (h *Handler) CommentCreate(c *gin.Context) {
	var req pbf.CommentCreateReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	res, err := h.srvs.Comment.Create(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		_, filename, line, _ := runtime.Caller(1)
		log.Printf("[error] %s:%d %v", filename, line, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

// CommentGet handles getting a comment by ID.
// @Summary Get comment
// @Description Get a comment by ID
// @Tags comment
// @Accept json
// @Produce json
// @Param id path string true "Comment ID"
// @Success 200 {object} pbf.CommentRes
// @Failure 400 {object} string "Invalid comment ID"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /comment/{id} [get]
func (h *Handler) CommentGetBydId(c *gin.Context) {
	id := &pbf.GetByIdReq{Id: c.Param("id")}
	res, err := h.srvs.Comment.GetById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get comment", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// CommentGetAll handles getting all comments.
// @Summary Get all comments
// @Description Get all comments
// @Tags comment
// @Accept json
// @Produce json
// @Param limit query integer false "Limit"
// @Param offset query integer false "Offset"
// @Success 200 {object} pbf.CommentGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /comments [get]
func (h *Handler) CommentGetAll(c *gin.Context) {
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

	res, err := h.srvs.Comment.GetAll(context.Background(), &filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get comments", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// CommentUpdate handles updating an existing comment.
// @Summary Update comment
// @Description Update an existing comment
// @Tags comment
// @Accept json
// @Produce json
// @Param id path string true "Comment ID"
// @Param comment body pbf.CommentCreateReq true "Updated comment data"
// @Success 200 {object} pbf.Comment
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Comment not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /comment/{id} [put]
func (h *Handler) ComentUpdate(c *gin.Context) {
	id := c.Param("id")
	var req pbf.CommentUpdate

	if err := c.BindJSON(&req.UpdateComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	req.Id = id
	res, err := h.srvs.Comment.Update(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update comment", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// CommentDelete handles deleting a comment by ID.
// @Summary Delete comment
// @Description Delete a comment by ID
// @Tags comment
// @Accept json
// @Produce json
// @Param id path string true "Comment ID"
// @Success 200 {object} string "Comment deleted"
// @Failure 400 {object} string "Invalid comment ID"
// @Failure 404 {object} string "Comment not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /comment/{id} [delete]
func (h *Handler) CommentDelete(c *gin.Context) {
	id := &pbf.GetByIdReq{Id: c.Param("id")}
	_, err := h.srvs.Comment.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete comment", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted"})
}


// CommentGetAll handles getting all comments by post.
// @Summary Get all comments by post
// @Description Get all comments post
// @Tags post
// @Accept json
// @Produce json
// @Param post_id query string false "PostId"
// @Param limit query integer false "Limit"
// @Param offset query integer false "Offset"
// @Success 200 {object} pbf.CommentGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /posts/{post_id}/comments [get]
func (h *Handler) GetCommentByPost(c *gin.Context){

	post_id := c.Query("post_id")
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
		Id: post_id,
		Limit:  int32(limit),
		Offset: int32(offset),
	}



	res, err := h.srvs.Comment.GetCommentByPost(context.Background(), &filter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get comments", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}