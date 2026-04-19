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

func Login(c *gin.Context) {

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(400, gin.H{
			"error": "invalid request",
		})

		return
	}

	var user models.User

	result := config.DB.Where(
		"username = ?",
		req.Username,
	).First(&user)

	if result.Error != nil {

		c.JSON(401, gin.H{
			"error": "username not found",
		})

		return
	}

	if user.Password != req.Password {

		c.JSON(401, gin.H{
			"error": "invalid password",
		})

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

		c.JSON(500, gin.H{
			"error": "token failed",
		})

		return
	}

	c.JSON(200, gin.H{
		"token": tokenString,
	})

}
