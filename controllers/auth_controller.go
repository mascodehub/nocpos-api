package controllers

import (
	"os"
	"time"

	"nocpos/config"
	"nocpos/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseRequest struct {
	Status  int `json:"status"`
	Message string `json:"message"`
	Data   string    `json:"data,omitempty"` // omitempty untuk mengabaikan jika kosong
}

func Login(c *gin.Context) {

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		response := ResponseRequest{
			Status: 400,
			Message: "NOT-OK",
			Data: "invalid request",
		}

		return
	}

	var user models.User

	result := config.DB.Where(
		"username = ?",
		req.Username,
	).First(&user)

	if result.Error != nil {

		response := ResponseRequest{
			Status: 401,
			Message: "NOT-OK",
			Data: "username not found",
		}

		return
	}

	if user.Password != req.Password {

		response := ResponseRequest{
			Status: 401,
			Message: "NOT-OK",
			Data: "invalid password",
		}

		return
	}

	claims := jwt.MapClaims{
		"iduser":   user.IDUser,
		"username": user.Username,
		"exp": time.Now().
			Add(time.Hour * 24).
			Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	tokenString, err := token.SignedString(
		[]byte(
			os.Getenv("JWT_SECRET"),
		),
	)

	if err != nil {

		response := ResponseRequest{
			Status: 500,
			Message: "NOT-OK",
			Data: "token failed",
		}

		return
	}

	response := ResponseRequest{
		Status: 200,
		Message: "OK",
		Data: tokenString,
	}

	c.JSON(200, response)

}
