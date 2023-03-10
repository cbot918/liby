package jwty

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt"
)

type Jwty struct {}

func New()*Jwty{
	return new(Jwty)
}

type JwtClaims struct {
	jwt.StandardClaims
	Email string
	Id    int
}

const (
	issuer = "yale"
	sec = "12345"
)

func (j *Jwty) FastJwt(id int, email string) (string, error) {
	// expiresAt := time.Now().Add(10 * time.Second).Unix()
	claims := JwtClaims{
		Id:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			Issuer:    issuer,
			// ExpiresAt: expiresAt,
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	tokenString, err := token.SignedString([]byte(sec))
	if err != nil { return "", err }
	
	return strings.Trim(tokenString,"nil") , nil
}

func (j *Jwty) DecodeJwt(tokenString string) *JwtClaims{
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(sec), nil
	})
	
	claims, ok := token.Claims.(*JwtClaims)
	if ok && token.Valid {
		return claims
	} else {
		fmt.Println(err)
		return &JwtClaims{}
	}
}

func (j *Jwty) GetIdAndEmail(claims *JwtClaims) (int, string) {
	return claims.Id, claims.Email
}
