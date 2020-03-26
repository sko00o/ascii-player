package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	 "github.com/sko00o/ascii-player/player"
	 "github.com/sko00o/ascii-player/video"
)

var (
	addr   string
	file   string
	stdout bool
)

func init() {
	flag.StringVar(&addr, "addr", ":8080", "TCP address to listen on")
	flag.StringVar(&file, "file", "static/rick_roll.txt", "Text file containing the ASCII video")
	flag.BoolVar(&stdout, "stdout", false, "print to stdout")
	flag.Parse()
}

func main() {
	v := video.DefaultVideo()
	if err := v.Load(file); err != nil {
		log.Printf("load video error: %v", err)
		return
	}

	if stdout {
		p := player.DefaultPlayer(os.Stdout, v)
		p.Play()
	} else {
		fmt.Printf("listening on %s\n", addr)
		log.Fatal(http.ListenAndServe(addr, newServer(v)))
	}
}

type server struct {
	v *video.Video
}

func newServer(v *video.Video) *server {
	return &server{v}
}

func (h *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := player.DefaultPlayer(w, h.v)
	p.Play()
}
