package jwty

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type Jwty struct {
	
}

func New()*Jwty{
	return new(Jwty)
}

func (j *Jwty) Run() string{
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("hmacSampleSecret"))
	if err != nil {

		fmt.Println("token.signedstring failed"); log.Fatal(err)
	}
	return tokenString
}

type AuthClaims struct {
	jwt.StandardClaims
	Email string
	Id    int
}



func (j *Jwty) FastJwt(id int, email string) (string, error) {
	issuer := "yale"
	sec := "12345"
	// expiresAt := time.Now().Add(10 * time.Second).Unix()
	claims := AuthClaims{
		Id:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			Issuer:    issuer,
			// ExpiresAt: expiresAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(sec))
	if err != nil {
		return "", err
	}
	return strings.Trim(tokenString,"nil") , nil
}
