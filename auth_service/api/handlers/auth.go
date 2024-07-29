package handlers

import (
	"auth-service/api/token"
	"auth-service/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	_ "github.com/swaggo/swag"
	"golang.org/x/crypto/bcrypt"
)

// Register godoc
// @Summary Register a new user
// @Description Register a new user with email, username, and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.RegisterReq true "User registration request"
// @Success 201 {object} token.Tokens "JWT tokens"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Router /register [post]
func (h *HTTPHandler) Register(c *gin.Context) {
	var req models.RegisterReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid request payload": err.Error()})
		return
	}

	exists, err := h.US.EmailExists(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	req.Password = string(hashedPassword)

	err = h.US.Register(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	tokens := token.GenerateJWTToken(req.ID, req.Email, req.Username)

	c.JSON(http.StatusCreated, tokens)
}

// Login godoc
// @Summary Login a user
// @Description Authenticate user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body models.LoginReq true "User login credentials"
// @Success 200 {object} token.Tokens "JWT tokens"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Invalid email or password"
// @Router /login [post]
func (h *HTTPHandler) Login(c *gin.Context) {
	req := models.LoginReq{}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid request payload": err.Error()})
		return
	}

	user, err := h.US.GetProfile(&models.GetProfileReq{Email: req.Email})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Invalid email or password": err.Error()})
		return
	}
	fmt.Println(user.Password, req.Password)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Invalid email or password...": err.Error()})
		return
	}

	tokens := token.GenerateJWTToken(user.ID, user.Email, user.Username)

	c.JSON(http.StatusOK, tokens)
}

// GetProfile godoc
// @Summary Get user profile
// @Description Get the profile of the authenticated user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} models.GetProfileResp
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "User not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /profile [get]
func (h *HTTPHandler) Profile(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	email := claims.(jwt.MapClaims)["email"].(string)
	user, err := h.US.GetProfile(&models.GetProfileReq{Email: email})
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"Server error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *HTTPHandler) GetByID(c *gin.Context) {
	id := &models.GetProfileByIdReq{ID: c.Param("id")}
	user, err := h.US.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Server error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
