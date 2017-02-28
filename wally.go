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
	"time"
)

const (
	bingdomain = "https://www.bing.com"
	bingurl    = "http://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1"
)

var (
	url      = flag.String("url", "", "url of the wallpaper")
	path     = flag.String("path", "Pictures", "dir for wallpaper")
	force    = flag.Bool("force", false, "force download")
	filename = flag.String("filename", "wally.jpg", "filename for wallpaper")
)

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

func absfilepath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(usr.HomeDir, *path, *filename)
}

func init() {
	flag.Parse()
	fi, err := os.Stat(absfilepath())
	if err != nil {
		log.Fatal(err)
	}

	recent := fi.ModTime().After(time.Now().Add(-24 * time.Hour))
	if recent && !*force {
		os.Exit(0)
	}
	if *url != "" {
		return
	}

	resp, err := http.Get(bingurl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bing := new(BingResult)
	json.NewDecoder(resp.Body).Decode(&bing)
	*url = fmt.Sprintf("%s%s", bingdomain, bing.Images[0].URL)

	log.Printf("Using default url: %v", *url)
}

func main() {
	resp, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	out, err := os.Create(absfilepath())
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	io.Copy(out, resp.Body)
}
