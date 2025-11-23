package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nfsarch33/secure-auth-platform/backend/internal/api"
	"github.com/nfsarch33/secure-auth-platform/backend/internal/service"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var req api.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Error: "Invalid request body"})
		return
	}

	if req.Email == "" || len(req.Password) < 8 {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Error: "Invalid input"})
		return
	}

	user, err := h.service.SignUp(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			c.JSON(http.StatusConflict, api.ErrorResponse{Error: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Error: "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, api.AuthResponse{
		User: api.User{
			Id:        user.ID.String(),
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		},
		Token: nil,
	})
}

func (h *AuthHandler) SignIn(c *gin.Context) {
	var req api.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Error: "Invalid request body"})
		return
	}

	user, token, err := h.service.SignIn(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		if err == service.ErrInvalidCredentials {
			c.JSON(http.StatusUnauthorized, api.ErrorResponse{Error: "Invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Error: "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, api.AuthResponse{
		User: api.User{
			Id:        user.ID.String(),
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		},
		Token: &token,
	})
}
