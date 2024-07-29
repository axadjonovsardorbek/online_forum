// tag.go
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

// TagCreate handles the creation of a new tag.
// @Summary Create tag
// @Description Create a new tag
// @Tags tag
// @Accept json
// @Produce json
// @Param tag body pbf.TagCreateReq true "Tag data"
// @Success 200 {object} pbf.TagRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /tag [post]
func (h *Handler) TagCreate(c *gin.Context) {
	var req pbf.TagCreateReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	res, err := h.srvs.Tag.Create(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		_, filename, line, _ := runtime.Caller(1)
		log.Printf("[error] %s:%d %v", filename, line, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

// TagGet handles getting a tag by ID.
// @Summary Get tag
// @Description Get a tag by ID
// @Tags tag
// @Accept json
// @Produce json
// @Param id path string true "Tag ID"
// @Success 200 {object} pbf.TagRes
// @Failure 400 {object} string "Invalid tag ID"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /tag/{id} [get]
func (h *Handler) TagGetBydId(c *gin.Context) {
	id := &pbf.GetByIdReq{Id: c.Param("id")}
	res, err := h.srvs.Tag.GetById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get tag", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// TagGetAll handles getting all tags.
// @Summary Get all tags
// @Description Get all tags
// @Tags tag
// @Accept json
// @Produce json
// @Param limit query integer false "Limit"
// @Param offset query integer false "Offset"
// @Success 200 {object} pbf.TagGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /tags [get]
func (h *Handler) TagGetAll(c *gin.Context) {
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

	res, err := h.srvs.Tag.GetAll(context.Background(), &filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get tags", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// TagUpdate handles updating an existing tag.
// @Summary Update tag
// @Description Update an existing tag
// @Tags tag
// @Accept json
// @Produce json
// @Param id path string true "Tag ID"
// @Param tag body pbf.TagCreateReq true "Updated tag data"
// @Success 200 {object} pbf.TagRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Comment not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /tag/{id} [put]
func (h *Handler) TagUpdate(c *gin.Context) {
	id := c.Param("id")
	var req pbf.TagUpdate

	if err := c.BindJSON(&req.Tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	req.Id = id
	res, err := h.srvs.Tag.Update(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update tag", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// TagDelete handles deleting a tag by ID.
// @Summary Delete tag
// @Description Delete a tag by ID
// @Tags tag
// @Accept json
// @Produce json
// @Param id path string true "Tag ID"
// @Success 200 {object} string "Tag deleted"
// @Failure 400 {object} string "Invalid tag ID"
// @Failure 404 {object} string "Tag not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /tag/{id} [delete]
func (h *Handler) TagDelete(c *gin.Context) {
	id := &pbf.GetByIdReq{Id: c.Param("id")}
	_, err := h.srvs.Tag.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete tag", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tag deleted"})
}
