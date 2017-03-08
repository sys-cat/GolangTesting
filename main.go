package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"net/url"
)

var key = "thish is secret key"
var data = map[string]string{
	"user":    "1",
	"pass":    "password1234",
	"payload": "anything payloads",
}

func Run() string {
	hash := hmac.New(sha512.New, []byte(key))
	hash.Write([]byte(fmt.Sprint(data)))
	signature := url.QueryEscape(base64.StdEncoding.EncodeToString(hash.Sum(nil)))
	return signature
}

func main() {
	fmt.Println(Run())
}
