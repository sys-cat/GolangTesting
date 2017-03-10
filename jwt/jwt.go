package jwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Payload struct {
	User    string `json:"User"`
	Pass    string `json:"Pass"`
	Payload string `json:"Payload"`
	jwt.StandardClaims
}

var key = "this is secret key"
var pay = Payload{
	"1",
	"password",
	"something payloads",
	jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + 15000,
		Issuer:    "AnotherCompany",
	},
}

func Run() map[string]string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, pay)
	signature, err := token.SignedString([]byte(key))
	if err != nil {
		fmt.Println(err)
	}
	parsed := strings.Split(signature, ".")
	return map[string]string{
		"body":   parsed[1],
		"verify": parsed[2],
	}
}

func DeRun(data []string) (*jwt.Token, error, Payload) {
	// header : "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9"
	// body : eyJVc2VyIjoiMSIsIlBhc3MiOiJwYXNzd29yZCIsIlBheWxvYWQiOiJzb21ldGhpbmcgcGF5bG9hZHMiLCJleHAiOjE0ODkwNDQ5NjcsImlzcyI6IkFub3RoZXJDb21wYW55In0
	// verify : 6vXmzSIVtWy4XU4GIExR6IRLVO54Do8gNobCRLzZiuFknpKdRCDFconsTMYwMNIPPTF_OESYGp45jp-PEAmKJw
	hjson, err := json.Marshal(map[string]interface{}{
		"typ": "JWT",
		"alg": jwt.SigningMethodHS512.Alg(),
	})
	if err != nil {
		fmt.Println(err)
	}
	head := strings.TrimRight(base64.URLEncoding.EncodeToString(hjson), "=")
	input := strings.Join([]string{head, data[0], data[1]}, ".")
	json := Payload{}
	decode, err := jwt.ParseWithClaims(input, &json, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	return decode, err, json
}

func DecodeOrigin(data string) (string, string) {
	parsed := strings.Split(data, ".")
	de, _ := base64.StdEncoding.DecodeString(parsed[0])
	pa, _ := base64.StdEncoding.DecodeString(parsed[1])
	return string(de), string(pa)
}
