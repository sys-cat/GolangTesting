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
	f, e := os.Open("doc.png")
	if e != nil {
		return nil, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	g := gift.New()
	dst := image.NewRGBA(g.Bounds(img.Bounds()))
	g.Draw(dst, img)
	n, err := os.Create("test.png")
	if err != nil {
		return nil, err
	}
	err = png.Encode(n, dst)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	nfile, err := os.Stat("test.png")
	if err != nil {
		return nil, err
	}
	fmt.Println(nfile.Size())
	if nerr := os.Remove("test.png"); err != nil {
		return nil, nerr
	}

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
	if err != nil || ext == "" {
		return false
	}
	g := gift.New()
	dst := image.NewRGBA(g.Bounds(img_1.Bounds()))
	g.Draw(dst, img_1)
	b := new(bytes.Buffer)
	err = png.Encode(b, dst)
	if err != nil {
		return false
	}
	fmt.Println(len(b.Bytes()))
	return true
}
