package images

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"net/http"
	"os"

	"github.com/disintegration/gift"
)

func Run() (os.FileInfo, error) {
	file, err := os.Stat("doc.png")
	if err != nil {
		fmt.Printf("error is %v\n", err)
		return nil, err
	}
	fmt.Println(file.Size())
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
	body := img.Body
	img_1, ext, err := image.Decode(body)
	if err != nil || ext == nil {
		return false
	}
	g := gift.New()
	dst := image.NewRGBA(g.Bounds(img_1.Bounds()))
	g.Draw(g, img_1)
	b := new(bytes.Buffer)
	err = png.Encode(b, dst)
	if err != nil {
		return false
	}
	fmt.Println(len(b.Bytes()))
	return true
}
