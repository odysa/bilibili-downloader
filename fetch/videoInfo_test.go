package fetch

import (
	"testing"
)

func TestGetVideoParts(t *testing.T) {
	data, err := fetchUrl("BV1iW411d7hd")
	if err != nil {
		panic(err)
	}
	res := getVideoParts(&data)
	if len(res) != 26 {
		t.Errorf("fetch video parts invalid")
	}
}

func TestFetch(t *testing.T) {
	res, err := Fetch("https://www.bilibili.com/video/BV1iW411d7hd?from=search&seid=4842193425456252344")
	if err != nil {
		panic(err)
	}
	if res.BV != "BV1iW411d7hd" || len(res.Parts) != 26 {
		t.Errorf("fetch video info error")
	}
}

func TestFetchUrl(t *testing.T) {
	data, err := fetchUrl("BV1iW411d7hd")
	if err != nil {
		panic(err)
	}
	if data.Get("code").Int() != 0 {
		t.Errorf("fetch video url error")
	}
}

func TestGetVideoInfo(t *testing.T) {
	testAid := int64(31289365)
	testTitle := "【精校中英字幕】2015 CMU 15-213 CSAPP 深入理解计算机系统 课程视频"
	data, err := fetchUrl("BV1iW411d7hd")
	if err != nil {
		panic(err)
	}
	aid, title, _, _ := getVideoInfo(&data)
	if aid != testAid {
		t.Errorf("failed to get detailed video information, aid shoule be %d, but got %d", testAid, aid)
	}
	if title != testTitle {
		t.Errorf("failed to get detailed video information, title shoule be %s, but got %s", testTitle, title)
	}
}
