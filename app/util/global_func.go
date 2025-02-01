package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"

	"github.com/Jerasin/app/constant"
	"github.com/Jerasin/app/pkg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type CustomError struct {
	message string
}

func (c CustomError) Error() string {
	return c.message
}

func FindElementByCondition[T any](slice []T, condition func(T) bool) (*T, error) {
	index := slices.IndexFunc(slice, condition)

	if index != -1 {
		return &slice[index], nil
	} else {
		pkg.PanicException(constant.BadRequest)
		return nil, CustomError{
			message: "FindElementByCondition Error",
		}
	}

}

func GetPayloadInToken(c *gin.Context, field string) (any, error) {
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
		value, ok := claims[field]

		if ok {
			return value, nil
		} else {
			return nil, errors.New("Error GetPayloadInToken")
		}

	} else {
		return nil, err
	}
}

func ReadFile(path string) any {
	var err error
	plan, _ := os.ReadFile(path)
	var data []map[string]interface{}
	err = json.Unmarshal(plan, &data)

	if err != nil {
		panic("ReadFile Error")
	}

	fmt.Printf("ReadFile = %T: %s\n", data, data)

	return data
}
