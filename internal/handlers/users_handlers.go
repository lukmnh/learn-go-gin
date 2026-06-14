package handlers

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"todo_rest_api/internal/apperrors"
	"todo_rest_api/internal/dto"
	"todo_rest_api/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UsersHandler struct {
	service service.UsersService
}

func NewUsersHandler(service service.UsersService) *UsersHandler {
	return &UsersHandler{service: service}
}

func (h *UsersHandler) Create(c *gin.Context) {
	var req dto.UserCreateRequest
	if err := c.ShouldBind(&req); !(err == nil) {
		if ve, ok := errors.AsType[validator.ValidationErrors](err); ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": formatValidationError(ve[0])})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users, err := h.service.Create(c.Request.Context(), req)
	if err != nil {
		log.Println("could not create user", err)
		switch {
		case errors.Is(err, apperrors.ErrDuplicateEmail):
			c.JSON(http.StatusConflict, gin.H{"error": "email sudah digunakan"})
		case errors.Is(err, apperrors.ErrHashPassword):
			c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal proses password"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "terjadi kesalahan pada server"})
		}
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *UsersHandler) Login(c *gin.Context) {
	var req dto.UserLoginRequest
	if err := c.ShouldBind(&req); !(err == nil) {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			c.JSON(http.StatusBadRequest, gin.H{"error": formatValidationError(ve[0])})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.Login(c.Request.Context(), req)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrInvalidCredentials):
			c.JSON(http.StatusUnauthorized, gin.H{"error": "email atau password salah"})
		default:
			log.Println("login error:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "terjadi kesalahan pada server"})
		}
		return
	}
	c.JSON(http.StatusOK, res)
}

func formatValidationError(fe validator.FieldError) string {
	field := strings.ToLower(fe.Field())

	switch fe.Tag() {
	case "required":
		return field + " wajib diisi"
	case "email":
		return field + " harus berupa email yang valid"
	case "min":
		return field + " minimal " + fe.Param() + " karakter"
	case "password":
		return "password harus mengandung huruf besar, huruf kecil, dan angka"
	default:
		return field + " tidak valid"
	}
}
