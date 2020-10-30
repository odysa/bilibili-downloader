package fetch

import (
	"testing"
)

func TestGetVideoParts(t *testing.T) {
	data, err := fetchUrl("BV1iW411d7hd")
	if err != nil {
		panic(err)
	}
	res := getVideoParts(data)
	if len(res) != 26 {
		t.Errorf("fetch video parts invalid")
	}
}

func TestFetchVideo(t *testing.T) {
	res, err := Fetch("https://www.bilibili.com/video/BV1iW411d7hd?from=search&seid=4842193425456252344")
	if err != nil {
		panic(err)
	}
	if res.BV != "BV1iW411d7hd" || len(res.Parts) != 26 {
		t.Errorf("fetch video info error")
	}
}
