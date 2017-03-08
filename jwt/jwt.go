package jwt

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

type Payload struct {
	User    int
	Pass    string
	Payload string
	jwt.StandardClaims
}

var key = "this is secret key"
var pay = Payload{
	User:    1,
	Pass:    "password",
	Payload: "something payloads",
}

func Run() string {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS512"), pay)
	signature, err := token.SignedString([]byte(key))
	if err != nil {
		fmt.Println(err)
	}
	return signature
}
