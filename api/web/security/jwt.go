package security

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"strings"
	"time"
)

type claims struct {
	UserID uint64 `json:"user_id"`
	jwt.StandardClaims
}

func ValidateToken(value string) error {
	token, err := extractToken(value)
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("the user must provid a valid token")
	}
	return nil
}

func extractToken(value string) (*jwt.Token, error) {
	if !strings.Contains(value, AuthorizationPrefix) {
		return nil, fmt.Errorf("authorization token must have Bearer prefix")
	}
	token := strings.Split(value, " ")[1]
	tkn, err := jwt.ParseWithClaims(token, &claims{}, func(token *jwt.Token) (interface{}, error) {
		jwtKey, _ := os.LookupEnv("SECRET_KEY")
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("an error occurred extracting token")
	}
	return tkn, err
}

func GenerateToken(userID string) (string, error) {
	key, _ := os.LookupEnv("SECRET_KEY")
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userID
	atClaims["expiration_time"] = time.Now().Add(time.Hour * 8).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(key))
	if err != nil {
		return "", fmt.Errorf("an error occurred generating authorization token")
	}
	return fmt.Sprintf("Bearer %s", token), nil
}