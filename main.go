package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"html/template"
	"image"
	"image/png"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/nfnt/resize"
)

var amg *amg8833.AMG88xx
var frame string
var colors []int
var pic []int
var grid []float64
var refresh *int
var minTemp, maxTemp *float64
var newSize *int
var dir string

func init() {
	refresh = flag.Int("f", 100, "refresh rate to capture and display the images")
	minTemp = flag.Float64("min", 26, "minimum temperature to measure from the sensor")
	maxTemp = flag.Float64("max", 32, "max temperature to measure from the sensor")
	newSize = flag.Int("s", 360, "new image size in pixel width")
	flag.Parse()
	var err error
	dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var err error
	amg, err = amg8833.NewAMG8833(&amg8833.Opts{
		Device: "/dev/i2c-1",
		Mode:   amg8833.AMG88xxNormalMode,
		Reset:  amg8833.AMG88xxInitialReset,
		FPS:    amg8833.AMG88xxFPS10,
	})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to AMG8833 module.")
	}
	go startThermalCam()

	mux := http.NewServeMux()
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(dir+"/public"))))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/frame", getFrame)
	server := &http.Server{
		Addr:    "0.0.0.0:12345",
		Handler: mux,
	}
	fmt.Println("Started AMG8833 Thermal Camera server at", server.Addr)
	server.ListenAndServe()

}

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(dir + "/public/index.html")
	go generateFrames()
	t.Execute(w, *refresh)
}

func generateFrames() {
	for {
		img := createImage(8, 8) 
		createFrame(img)      
		time.Sleep(time.Duration(*refresh) * time.Millisecond)
	}
}

func getFrame(w http.ResponseWriter, r *http.Request) {
	str := "data:image/png;base64," + frame
	w.Header().Set("Cache-Control", "no-cache")
	w.Write([]byte(str))
}

func getColorIndex(temp float64) int {
	if temp < *minTemp {
		return 0
	}
	if temp > *maxTemp {
		return len(colors) - 1
	}
	return int((temp - *minTemp) * float64(len(colors)-1) / (*maxTemp - *minTemp))
}

func createFrame(img image.Image) {
	var buf bytes.Buffer
	png.Encode(&buf, img)
	frame = base64.StdEncoding.EncodeToString(buf.Bytes())
}

func createImage(w, h int) image.Image {
	pixels := image.NewRGBA(image.Rect(0, 0, w, h))
	n := 0
	for _, i := range grid {
		color := colors[getColorIndex(i)]
		pixels.Pix[n] = getR(color)
		pixels.Pix[n+1] = getG(color)
		pixels.Pix[n+2] = getB(color)
		pixels.Pix[n+3] = 0xFF 
		n = n + 4
	}
	dest := resize.Resize(360, 0, pixels, resize.Lanczos3)
	return dest
}

func getR(i int) uint8 {
	return uint8((i >> 16) & 0x0000FF)
}

func getG(i int) uint8 {
	return uint8((i >> 8) & 0x0000FF)
}

func getB(i int) uint8 {
	return uint8(i & 0x0000FF)
}

func startThermalCam() {
	for {
		grid = amg.ReadPixels()
		time.Sleep(time.Duration(*refresh) * time.Millisecond)
	}
}