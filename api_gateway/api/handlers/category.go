// category.go
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

// CategoryCreate handles the creation of a new category.
// @Summary Create category
// @Description Create a new category
// @Tags category
// @Accept json
// @Produce json
// @Param category body pbf.CategoryCreateReq true "Category data"
// @Success 200 {object} pbf.CategoryRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /category [post]
func (h *Handler) CategoryCreate(c *gin.Context) {
	var req pbf.CategoryCreateReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	res, err := h.srvs.Category.Create(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		_, filename, line, _ := runtime.Caller(1)
		log.Printf("[error] %s:%d %v", filename, line, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

// CategoryGet handles getting a category by ID.
// @Summary Get category
// @Description Get a category by ID
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} pbf.CategoryRes
// @Failure 400 {object} string "Invalid category ID"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /category/{id} [get]
func (h *Handler) CategoryGetBydId(c *gin.Context) {
	id := &pbf.GetByIdReq{Id: c.Param("id")}
	res, err := h.srvs.Category.GetById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get category", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// CategoryGetAll handles getting all categories.
// @Summary Get all categories
// @Description Get all categories
// @Tags category
// @Accept json
// @Produce json
// @Param limit query integer false "Limit"
// @Param offset query integer false "Offset"
// @Success 200 {object} pbf.CategoryGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /categories [get]
func (h *Handler) CategoryGetAll(c *gin.Context) {
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

	res, err := h.srvs.Category.GetAll(context.Background(), &filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get categories", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// CategoryUpdate handles updating an existing category.
// @Summary Update category
// @Description Update an existing category
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Param category body pbf.CategoryCreateReq true "Updated category data"
// @Success 200 {object} pbf.CategoryRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Comment not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /category/{id} [put]
func (h *Handler) CategoryUpdate(c *gin.Context) {
	id := c.Param("id")
	var req pbf.CategoryUpdate

	if err := c.BindJSON(&req.Category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	req.Id = id
	res, err := h.srvs.Category.Update(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update category", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// CategoryDelete handles deleting a category by ID.
// @Summary Delete category
// @Description Delete a category by ID
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} string "Category deleted"
// @Failure 400 {object} string "Invalid category ID"
// @Failure 404 {object} string "Category not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /category/{id} [delete]
func (h *Handler) CategoryDelete(c *gin.Context) {
	id := &pbf.GetByIdReq{Id: c.Param("id")}
	_, err := h.srvs.Category.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete category", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}
