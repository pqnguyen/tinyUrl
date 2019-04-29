package password

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pqnguyen/tinyUrl/models"
	"github.com/pqnguyen/tinyUrl/services/auth"
	"time"
)

type PasswordAuthority interface {
	VerifyUser(user *models.User, password string) bool
	GenerateJWT(user *models.User) string
	HashPassword(password string) string
}

type DefaultPasswordAuthority struct{}

func (authority *DefaultPasswordAuthority) VerifyUser(user *models.User, password string) bool {
	hashPassword := authority.HashPassword(password)
	return user.Password == hashPassword
}

func (authority *DefaultPasswordAuthority) GenerateJWT(user *models.User) string {
	claims := auth.Claims{}
	claims.Id = user.ID.Hex()
	claims.Name = user.Name
	claims.IssuedAt = time.Now().UTC().Unix()
	claims.ExpiresAt = time.Now().Add(30 * 24 * time.Hour).UTC().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(auth.SignedString))
	fmt.Println(signedToken)
	return signedToken
}

func (authority *DefaultPasswordAuthority) HashPassword(password string) string {
	return base64.StdEncoding.EncodeToString(md5.New().Sum([]byte(password)))
}
