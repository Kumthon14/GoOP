package Controllers

import (
	"Go_OOP/Models"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthController struct{}

func (auth *AuthController) NewAuthController() {}

// @Summary Login User
// @Description Login User
// @ID LoginUser
// @Tags User
// @Param EnterDetails body Models.UserLogin true "Login"
// @Accept json
// @Success 200 {object} Models.UserLoginRes "Success"
// @Failure 400 {string} string "Error"
// @Router /auth-api/login [POST]
func (auth *AuthController) Login(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)
	err := Models.Login(&user)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
		})

		ss, err := token.SignedString([]byte("MySignature"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		userRes := Models.UserLoginRes{
			Name:     user.Name,
			Email:    user.Email,
			Phone:    user.Phone,
			Address:  user.Address,
			Username: user.Username,
			Token:    ss,
		}

		c.JSON(http.StatusOK, userRes)
	}

}

func (auth *AuthController) IsAccess(c *gin.Context) {
	s := c.Request.Header.Get("Authorization")

	token := strings.TrimPrefix(s, "Bearer")

	if err := auth.validateToken(token); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

func (auth *AuthController) validateToken(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte("MySignature"), nil
	})

	return err
}
