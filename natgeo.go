package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type natGeoResult struct {
	GalleryTitle     string `json:"galleryTitle"`
	PreviousEndpoint string `json:"previousEndpoint"`
	Items            []struct {
		Title       string  `json:"title"`
		Caption     string  `json:"caption"`
		Credit      string  `json:"credit"`
		ProfileURL  string  `json:"profileUrl"`
		AltText     string  `json:"altText"`
		FullPathURL string  `json:"full-path-url"`
		URL         string  `json:"url"`
		OriginalURL string  `json:"originalUrl"`
		AspectRatio float64 `json:"aspectRatio"`
		Sizes       struct {
			Num240  string `json:"240"`
			Num320  string `json:"320"`
			Num500  string `json:"500"`
			Num640  string `json:"640"`
			Num800  string `json:"800"`
			Num1024 string `json:"1024"`
			Num1600 string `json:"1600"`
			Num2048 string `json:"2048"`
		} `json:"sizes"`
		Internal    bool   `json:"internal"`
		PageURL     string `json:"pageUrl"`
		PublishDate string `json:"publishDate"`
		YourShot    bool   `json:"yourShot"`
		Social      struct {
			OgTitle       string `json:"og:title"`
			OgDescription string `json:"og:description"`
			TwitterSite   string `json:"twitter:site"`
		} `json:"social"`
		Livefyre struct {
			PageGUID      string `json:"pageGuid"`
			Checksum      string `json:"checksum"`
			LfMetadata    string `json:"lfMetadata"`
			SiteSecret    string `json:"siteSecret"`
			LfSiteID      string `json:"lfSiteId"`
			LfNetworkName string `json:"lfNetworkName"`
			LfElement     string `json:"lfElement"`
		} `json:"livefyre"`
	} `json:"items"`
}

type natGeoPlugin struct {
	domain, endpoint string
}

func newNatGeoPlugin() *natGeoPlugin {
	return &natGeoPlugin{
		domain:   "http://www.nationalgeographic.com",
		endpoint: "/photography/photo-of-the-day/_jcr_content/.gallery.json",
	}
}

func (n *natGeoPlugin) URL() (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s%s", n.domain, n.endpoint))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	natgeo := new(natGeoResult)
	json.NewDecoder(resp.Body).Decode(&natgeo)
	if len(natgeo.Items) == 0 {
		return "", errors.New("no items")
	}

	item := natgeo.Items[0]
	return fmt.Sprintf("%s%s", item.URL, item.Sizes.Num2048), nil
}
