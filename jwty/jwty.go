package jwty

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt"
)

type Jwty struct {

}

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
	if err != nil {
		return "", err
	}
	return strings.Trim(tokenString,"nil") , nil
}

func (j *Jwty) DecodeJwt(tokenString string){
	// tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJpc3MiOiJ0ZXN0IiwiYXVkIjoic2luZ2xlIn0.QAWg1vGvnqRuCFTMcPkjZljXHh8U3L_qUjszOtQbeaA"
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(sec), nil
	})
	
	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.Email, claims.StandardClaims.Issuer)
	} else {
		fmt.Println("h")
		fmt.Println(err)
	}
}

// func (j *Jwty) GetData() (int, string) {
	
// }
