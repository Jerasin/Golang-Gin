package middleware

import (
	"fmt"

	"github.com/Jerasin/app/constant"
	"github.com/Jerasin/app/pkg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/goforj/godump"
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

			godump.Dump(claims)

			userID, ok := claims["id"].(float64)
			if !ok {
				pkg.PanicException(constant.BadRequest)
			}
			godump.Dump(userID)

			c.Set("userID", uint(userID))
		} else {
			fmt.Println("testing")
			fmt.Println(err)
			pkg.PanicException(constant.Unauthorized)
		}

	}
}
