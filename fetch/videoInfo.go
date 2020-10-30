package fetch

import (
	"github.com/odysa/bilibili-downloader/model"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"regexp"
)

const cidUrl = "https://api.bilibili.com/x/web-interface/view?bvid="

func Fetch(url string) (result model.Video, err error) {

	BVRe := regexp.MustCompile(`.*/video/([^?]*).*`)
	result.BV = string(BVRe.FindSubmatch([]byte(url))[1])

	data, err := fetchUrl(result.BV)
	if err != nil {
		return model.Video{}, err
	}

	result.Aid, result.Title, result.Desc, result.Up = getVideoInfo(&data)
	result.Parts = getVideoParts(&data)
	return
}

func fetchUrl(bv string) (gjson.Result, error) {
	res, err := http.Get(cidUrl + bv)
	if err != nil {
		return gjson.Result{}, err
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	return gjson.GetBytes(data, "data"), err
}

func getVideoInfo(data *gjson.Result) (aid int64, title string, desc string, up model.Up) {
	aid = data.Get("aid").Int()
	title = data.Get("title").String()
	desc = data.Get("desc").String()

	upParse := data.Get("owner")
	up = model.Up{
		Name: upParse.Get("name").String(),
		Mid:  upParse.Get("mid").Int(),
		Face: upParse.Get("face").String(),
	}

	return
}

func getVideoParts(data *gjson.Result) (result []model.VideoPart) {
	pages := data.Get("pages").Array()
	for _, item := range pages {
		result = append(result, model.VideoPart{
			Cid:   item.Get("cid").Int(),
			Title: item.Get("part").String(),
		})
	}
	return
}
