package auth

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//Clients are able to authorization
var acceptedMachins = map[string]UserMeta{"test@test.com": UserMeta{"123", 0}}

var (
	hmacSecret  string //token secret
	expiresAuth int64  //expires seconds
)

//UserMeta is meta info about authorized clients
type UserMeta struct {
	Password        string
	ExpiresAuthTime int64 //0 - if user not be authenticated
}

func userCorrect(userName string, password string) bool {
	if len(userName) == 0 || len(password) == 0 {
		return false
	}

	userMeta, ok := acceptedMachins[userName]
	if ok && userMeta.Password == password {
		return true
	}

	return false
}

func generateToken(userName string, authTime int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"usr": userName,
		"iat": authTime,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(hmacSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

//InitConfig is configuration initialization
func InitConfig(secret string, expires int64) {
	hmacSecret = secret
	expiresAuth = expires
}

//CheckAuthentication is authentication process
func CheckAuthentication(tokenString string) (bool, error) {
	if len(tokenString) == 0 {
		return false, errors.New("token can not be empty")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(hmacSecret), nil
	})

	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userName := claims["usr"]

		if userMeta, okMeta := acceptedMachins[userName.(string)]; okMeta {
			nowTime := time.Now().Unix()

			if timeDiff := userMeta.ExpiresAuthTime - nowTime; timeDiff > 0 {
				return true, nil
			}
		}
	}

	return false, nil
}

//OnAuthentication process
func OnAuthentication(userName string, password string) (string, string, int64, error) {
	if userCorrect(userName, password) {
		authTime := time.Now().Unix()
		token, err := generateToken(userName, authTime)
		if err != nil {
			return "", "", 0, err
		}

		expiresAuthTime := authTime + expiresAuth

		var userMeta = acceptedMachins[userName]
		userMeta.ExpiresAuthTime = expiresAuthTime
		acceptedMachins[userName] = userMeta

		return token, "Bearer", expiresAuthTime, nil
	}

	return "", "", 0, errors.New("user not exist")
}

//LogOut process
func LogOut(tokenString string) (bool, error) {
	if len(tokenString) == 0 {
		return false, errors.New("token can not be empty")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(hmacSecret), nil
	})

	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userName := claims["usr"].(string)

		if userMeta, okMeta := acceptedMachins[userName]; okMeta {
			userMeta.ExpiresAuthTime = 0 // set auth time to 0 - reset auth
			acceptedMachins[userName] = userMeta
			return true, nil
		}
	}

	return false, nil
}
