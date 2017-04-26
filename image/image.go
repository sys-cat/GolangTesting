package image

import (
	"fmt"
	"net/http"
	"os"
)

func Run() (os.FileInfo, error) {
	file, err := os.Stat("doc.png")
	if err != nil {
		fmt.Printf("error is %v\n", err)
		return nil, err
	}
	fmt.Printf("Stat is %v\n", file)
	return file, nil
}

func RunOverHttp() bool {
	// npm install -g http-server
	// running http server
	img, err := http.Get("http://localhost:8080/doc.png")
	if err != nil {
		return false
	}
	fmt.Println(img.ContentLength)
	return true
}
