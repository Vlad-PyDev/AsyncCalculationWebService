package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

const (
	HmacSampleSecret string = "an7DkUH?L8iClxbVj5JZdbRVO2M$1Jc~D6CXsL@4"
)

func Generate(id int) (string, error) {
	currentTime := time.Now()
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"nbf": currentTime.Unix(),
		"exp": currentTime.Add(10 * time.Minute).Unix(),
		"iat": currentTime.Unix(),
	})

	signedToken, err := newToken.SignedString([]byte(HmacSampleSecret))
	if err != nil {
		log.Printf("token generation failed: %v", err)
		return "", err
	}

	return signedToken, nil
}

func Verify(tokenString string) (bool, int) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isHMAC := token.Method.(*jwt.SigningMethodHMAC); !isHMAC {
			return nil, fmt.Errorf("invalid signing method: %v", token.Header["alg"])
		}

		return []byte(HmacSampleSecret), nil
	})

	if err != nil {
		log.Printf("token parsing error: %v", err)
		return false, 0
	}

	if !parsedToken.Valid {
		log.Println("token is not valid")
		return false, 0
	}

	if claims, isValid := parsedToken.Claims.(jwt.MapClaims); isValid {
		userID, isIDValid := claims["id"].(float64)
		if !isIDValid {
			log.Println("invalid user ID format")
			return false, 0
		}
		log.Printf("verified user ID: %d", int(userID))
		return true, int(userID)
	}

	log.Println("invalid claims format")
	return false, 0
}
