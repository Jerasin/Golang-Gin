package middleware

import (
	"fmt"

	"github.com/Jerasin/app/constant"
	"github.com/Jerasin/app/pkg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer pkg.PanicHandler(c)

		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			pkg.PanicException(constant.Unauthorized)
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := pkg.NewAuthService().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println("claims", claims)
		} else {
			fmt.Println("testing")
			fmt.Println(err)
			pkg.PanicException(constant.Unauthorized)
		}

	}
}
