package hmac

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
)

type Payload struct {
	User    int
	Pass    string
	Payload string
}

var key = "thish is secret key"
var pay = Payload{
	User:    1,
	Pass:    "password",
	Payload: "something payloads",
}
var data = fmt.Sprint(pay.User) + pay.Pass + pay.Payload

func Run() string {
	hash := hmac.New(sha512.New, []byte(key))
	hash.Write([]byte(data))
	signature := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return signature
}

func Auth(sig string) bool {
	hash := hmac.New(sha512.New, []byte(key))
	hash.Write([]byte(data))
	signature := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	if sig == signature {
		return true
	}
	return false
}
