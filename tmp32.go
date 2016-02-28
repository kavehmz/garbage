package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func body(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	for i := 0; i < 15; i++ {
		for j := 0; j < 15; j++ {
			fmt.Fprintf(w, "<img src='/image?%d' height='10' width='10' border=1> ", rand.Int63n(1000)+time.Now().UnixNano())
		}
		fmt.Fprintf(w, "<br /> ")
	}

}
func blueHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	var img image.Image = m
	writeImage(w, &img)
}

func writeImage(w http.ResponseWriter, img *image.Image) {

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}

func main() {

	http.HandleFunc("/image", blueHandler)
	http.HandleFunc("/body", body)
	keyBase := "/Users/kaveh/dev/home/share/secret/letsencrypt/archive/h2.kaveh.me/"
	log.Fatal(http.ListenAndServeTLS(":443", keyBase+"fullchain2.pem", keyBase+"privkey2.pem", nil))
	//log.Fatal(http.ListenAndServe(":8443", nil))
}
