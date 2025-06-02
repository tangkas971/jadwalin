package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("sdfsdfksdjflskdfjslkfj")

type JWTClaim struct {
	UserId uint		`json:"id"`
	Email  string 	`json:"email"`
	Role   string	`json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(userId uint, email string, role string)(string, error){
	// batas waktu
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &JWTClaim{
		UserId: userId,
		Email: email,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			// waktu token habis
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			// waktu token dibuat
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	//membuat token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// menandatangi token
	return token.SignedString(jwtKey)
}

func ValidateToken(signedToken string)(*JWTClaim, error){
	token, err := jwt.ParseWithClaims(
		signedToken, 
		&JWTClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}