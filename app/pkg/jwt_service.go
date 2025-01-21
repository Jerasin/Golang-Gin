package pkg

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Jerasin/app/config"
	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(username string) string
	ValidateToken(token string) (*jwt.Token, error)
	GenerateRefreshToken(username string) string
}

type jwtServices struct {
	secretKey string
	issure    string
}

type authCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func NewAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    "Bikash",
	}
}

func (service *jwtServices) GenerateRefreshToken(username string) string {
	config.EnvConfig()
	JWT_EXPIRE_MINUTE := config.GetEnv("JWT_EXPIRE_MINUTE", "15")

	expire_time, err := strconv.Atoi(JWT_EXPIRE_MINUTE)
	if err != nil {
		// ... handle error
		panic(err)
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["username"] = username
	rtClaims["exp"] = time.Now().Add(time.Minute * time.Duration(expire_time)).Unix()
	rt, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		panic(err)
	}

	return rt
}

func (service *jwtServices) GenerateToken(username string) string {
	config.EnvConfig()
	JWT_EXPIRE_MINUTE := config.GetEnv("JWT_EXPIRE_MINUTE", "15")

	expire_time, err := strconv.Atoi(JWT_EXPIRE_MINUTE)
	if err != nil {
		// ... handle error
		panic(err)
	}

	claims := &authCustomClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(expire_time)).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}
