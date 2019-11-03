package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type redditResult struct {
	Data []struct {
		ID          string        `json:"id"`
		Title       string        `json:"title"`
		Description interface{}   `json:"description"`
		Datetime    int           `json:"datetime"`
		Type        string        `json:"type"`
		Animated    bool          `json:"animated"`
		Width       int           `json:"width"`
		Height      int           `json:"height"`
		Size        int           `json:"size"`
		Views       int           `json:"views"`
		Bandwidth   int64         `json:"bandwidth"`
		Vote        interface{}   `json:"vote"`
		Favorite    bool          `json:"favorite"`
		Nsfw        bool          `json:"nsfw"`
		Section     string        `json:"section"`
		AccountURL  interface{}   `json:"account_url"`
		AccountID   interface{}   `json:"account_id"`
		IsAd        bool          `json:"is_ad"`
		InMostViral bool          `json:"in_most_viral"`
		HasSound    bool          `json:"has_sound"`
		Tags        []interface{} `json:"tags"`
		AdType      int           `json:"ad_type"`
		AdURL       string        `json:"ad_url"`
		Edited      int           `json:"edited"`
		InGallery   bool          `json:"in_gallery"`
		Link        string        `json:"link"`
		AdConfig    struct {
			SafeFlags       []string      `json:"safeFlags"`
			HighRiskFlags   []string      `json:"highRiskFlags"`
			UnsafeFlags     []interface{} `json:"unsafeFlags"`
			WallUnsafeFlags []interface{} `json:"wallUnsafeFlags"`
			ShowsAds        bool          `json:"showsAds"`
		} `json:"ad_config"`
		CommentCount  interface{} `json:"comment_count"`
		FavoriteCount interface{} `json:"favorite_count"`
		Ups           interface{} `json:"ups"`
		Downs         interface{} `json:"downs"`
		Points        interface{} `json:"points"`
		Score         int         `json:"score"`
		IsAlbum       bool        `json:"is_album"`
	} `json:"data"`
}

type redditPlugin struct {
	domain, endpoint string
}

func newRedditPlugin() *redditPlugin {
	return &redditPlugin{
		domain:   "https://api.imgur.com",
		endpoint: "/3/gallery/r/wallpapers/time",
	}
}

func (b *redditPlugin) URL() (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", b.domain, b.endpoint), nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", "Client-ID c4803d90cfb105f")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	reddit := new(redditResult)
	json.NewDecoder(resp.Body).Decode(&reddit)
	return reddit.Data[0].Link, nil
}
