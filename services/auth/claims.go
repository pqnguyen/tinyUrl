package auth

import "github.com/dgrijalva/jwt-go"

var SignedString = "123#123$123%"

type Claims struct {
	Name string
	jwt.StandardClaims
}
