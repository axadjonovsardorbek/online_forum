// post_tag.go
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

// PostTagCreate handles the creation of a new post_tag.
// @Summary Create post_tag
// @Description Create a new post_tag
// @Tags post_tag
// @Accept json
// @Produce json
// @Param post_tag body pbf.PostTagsCreateReq true "Post data"
// @Success 200 {object} pbf.PostTags
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /post_tag [post]
func (h *Handler) PostTagCreate(c *gin.Context) {
	var req pbf.PostTagsCreateReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	res, err := h.srvs.PostTag.Create(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		_, filename, line, _ := runtime.Caller(1)
		log.Printf("[error] %s:%d %v", filename, line, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

// PostTagGet handles getting a post_tag by ID.
// @Summary Get post_tag
// @Description Get a post_tag by ID
// @Tags post_tag
// @Accept json
// @Produce json
// @Param id path string true "PostTag ID"
// @Success 200 {object} pbf.PostTagsRes
// @Failure 400 {object} string "Invalid post_tag ID"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /post_tag/{id} [get]
func (h *Handler) PostTagGetBydId(c *gin.Context) {
	id := &pbf.GetByIdReq{Id: c.Param("id")}
	res, err := h.srvs.PostTag.GetById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get post_tag", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// PostTagGetAll handles getting all post_tags.
// @Summary Get all post_tags
// @Description Get all post_tags
// @Tags post_tag
// @Accept json
// @Produce json
// @Param limit query integer false "Limit"
// @Param offset query integer false "Offset"
// @Success 200 {object} pbf.PostTagsGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /post_tags [get]
func (h *Handler) PostTagGetAll(c *gin.Context) {
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

	res, err := h.srvs.PostTag.GetAll(context.Background(), &filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get post_tags", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// PostTagUpdate handles updating an existing post_tag.
// @Summary Update post_tag
// @Description Update an existing post_tag
// @Tags post_tag
// @Accept json
// @Produce json
// @Param id path string true "PostTag ID"
// @Param post_tag body pbf.PostTagsCreateReq true "Updated post_tag data"
// @Success 200 {object} pbf.PostTags
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "PostTag not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /post_tag/{id} [put]
func (h *Handler) PostTagUpdate(c *gin.Context) {
	id := c.Param("id")
	var req pbf.PostTagsUpdate

	if err := c.BindJSON(&req.Posttag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	req.Id = id
	res, err := h.srvs.PostTag.Update(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update post_tag", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// PostTagDelete handles deleting a post_tag by ID.
// @Summary Delete post_tag
// @Description Delete a post_tag by ID
// @Tags post_tag
// @Accept json
// @Produce json
// @Param id path string true "PostTag ID"
// @Success 200 {object} string "PostTag deleted"
// @Failure 400 {object} string "Invalid post_tag ID"
// @Failure 404 {object} string "PostTag not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /post_tag/{id} [delete]
func (h *Handler) PostTagDelete(c *gin.Context) {
	id := &pbf.GetByIdReq{Id: c.Param("id")}
	_, err := h.srvs.PostTag.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete post_tag", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "PostTag deleted"})
}


// PostGetAll handles getting all posts by tag.
// @Summary Get all posts by tag
// @Description Get all posts by tag
// @Tags tag
// @Accept json
// @Produce json
// @Param tag_id query string false "TagId"
// @Param limit query integer false "Limit"
// @Param offset query integer false "Offset"
// @Success 200 {object} pbf.PostGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /tags/{tag_id}/posts [get]
func (h *Handler) GetPostByTag(c *gin.Context){
	tag_id := c.Query("tag_id")
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
		Id: tag_id,
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	res, err := h.srvs.PostTag.GetPostByTag(context.Background(), &filter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get posts", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetPopularTag handles getting tag.
// @Summary Get popular tag
// @Description Get popular tag
// @Tags tag
// @Accept json
// @Produce json
// @Success 200 {object} pbf.PopularTag
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /tags/popular [get]
func (h *Handler) GetPopularTag(c *gin.Context){
	v := pbf.Void{}
	res, err := h.srvs.PostTag.GetPopularTag(context.Background(), &v)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get post_tags", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}