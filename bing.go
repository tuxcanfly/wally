package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

type bingPlugin struct {
	domain, endpoint string
}

func newBingPlugin() *bingPlugin {
	return &bingPlugin{
		domain:   "https://www.bing.com",
		endpoint: "/HPImageArchive.aspx?format=js&idx=0&n=1",
	}
}

func (b *bingPlugin) URL() (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s%s", b.domain, b.endpoint))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bing := new(BingResult)
	json.NewDecoder(resp.Body).Decode(&bing)
	return fmt.Sprintf("%s%s", b.domain, bing.Images[0].URL), nil
}
