package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
)

var url = flag.String("url", "", "url of the wallpaper")
var path = flag.String("path", "Pictures", "path to write the wallpaper")
var filename = flag.String("filename", "wally.jpg", "filename to write the wallpaper")

type BingResult struct {
	Images []struct {
		Startdate     string        `json:"startdate"`
		Fullstartdate string        `json:"fullstartdate"`
		Enddate       string        `json:"enddate"`
		URL           string        `json:"url"`
		Urlbase       string        `json:"urlbase"`
		Copyright     string        `json:"copyright"`
		Copyrightlink string        `json:"copyrightlink"`
		Quiz          string        `json:"quiz"`
		Wp            bool          `json:"wp"`
		Hsh           string        `json:"hsh"`
		Drk           int           `json:"drk"`
		Top           int           `json:"top"`
		Bot           int           `json:"bot"`
		Hs            []interface{} `json:"hs"`
	} `json:"images"`
	Tooltips struct {
		Loading  string `json:"loading"`
		Previous string `json:"previous"`
		Next     string `json:"next"`
		Walle    string `json:"walle"`
		Walls    string `json:"walls"`
	} `json:"tooltips"`
}

func init() {
	flag.Parse()
	if *url == "" {
		resp, err := http.Get("http://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		bing := new(BingResult)
		json.NewDecoder(resp.Body).Decode(&bing)
		*url = fmt.Sprintf("https://www.bing.com%s", bing.Images[0].URL)
		log.Printf("Using default url: %v", *url)
	}
}

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
