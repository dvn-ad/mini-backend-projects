package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error){
	bytes,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	return string(bytes),err
}

func CheckPasswordHash(password, hash string) bool {
	err:=bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err==nil
}

func GenerateJWT (username string) (string,error) {
	jwtSecret:=[]byte(os.Getenv("JWT_SECRET"))
	if len(jwtSecret)==0{
		return "", errors.New("JWT_SECRET environment variable is not set")
	}

	claims:=jwt.MapClaims{
		"username"	: username,
		"exp"		: time.Now().Add(24 * time.Hour).Unix(),
	}

	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenStr string) (string, error) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {                              
			return nil, errors.New("unexpected signing method")                               
		}                                                                                     
		return jwtSecret, nil                                                                 
	})

	if err != nil {
		return "", err                                                                        
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if username, ok := claims["username"].(string); ok {                                  
			return username, nil                                                              
		}                                                                                     
	}

	return "", errors.New("invalid token claims")
}