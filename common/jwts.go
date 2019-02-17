package common

import (
	"fmt"
	"time"
)
import "github.com/dgrijalva/jwt-go"

func Hello() {
	fmt.Println("hello")
}

func NewJwtToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	return token.SignedString("fdafsd")
}
