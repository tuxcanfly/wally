package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
)

var url = flag.String("url", "https://www.bing.com/az/hprichbg/rb/ShengshanIsland_EN-US13597723185_1920x1080.jpg", "url of the wallpaper")
var path = flag.String("path", "Pictures", "path to write the wallpaper")
var filename = flag.String("filename", "wally.jpg", "filename to write the wallpaper")

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	out, err := os.Create(filepath.Join(usr.HomeDir, *path, *filename))
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	io.Copy(out, resp.Body)
}
