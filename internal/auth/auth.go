package auth

import (
	"../config"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"strings"
	"time"
)

func GenerateToken() string {
	expires := time.Now().Unix() + int64(24*time.Hour)
	type JSTClaims struct {
		jwt.StandardClaims
	}
	claims := JSTClaims{
		jwt.StandardClaims{ExpiresAt: expires, Issuer: "goose"},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Config.Auth.Secret))
	if err != nil {
		log.Println(err)
	}
	return tokenString
}

func Hash(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func VerifyCredentials(user, pass string) bool {
	configPassHash := Hash(config.Config.Auth.Pass)
	passHash := Hash(pass)
	return configPassHash == passHash && config.Config.Auth.User == strings.ToLower(user)
}
func VerifyToken(tokenString string) bool {
	if _, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("jwt: unexpected signing method %v", token.Header["alg"])
		}
		return []byte(config.Config.Auth.Secret), nil
	}); err != nil {
		log.Println(err)
		return false
	}
	return true
}
