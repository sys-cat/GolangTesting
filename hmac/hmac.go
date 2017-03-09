package hmac

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
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
var data, _ = json.Marshal(pay)

func Run() string {
	hash := hmac.New(sha512.New, []byte(key))
	hash.Write([]byte(data))
	signature := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return signature
}

func Auth(sig string) bool {
	hash := hmac.New(sha512.New, []byte(key))
	hash.Write(data)
	signature := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	if sig == signature {
		return true
	}
	return false
}
