package download

import (
	"fmt"
	"github.com/odysa/bilibili-downloader/model"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func DownloadVideo(video model.Video) {
	for _, item := range video.Parts {
		go func() {
			_, err := getDownloadUrl(video.BV, item.Cid)
			if err != nil {
				return
			}

		}()
	}
}

func Download(url string) error {
	client := http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36")
	request.Header.Set("Accept", "*/*")
	request.Header.Set("Accept-Language", "en-US,en;q=0.5")
	request.Header.Set("Accept-Encoding", "gzip, deflate, br")
	request.Header.Set("Range", "bytes=0-")
	request.Header.Set("Referer", "https://api.bilibili.com/x/web-interface/view?bvid=BV1xV41127rf")
	request.Header.Set("Origin", "https://www.bilibili.com")
	request.Header.Set("Connection", "keep-alive")

	log.Println("request send")
	res, err := client.Do(request)
	defer res.Body.Close()

	log.Println("video is downloading")
	file, err := os.Create("./1.flv")

	counter := &WriteCounter{}
	_, err = io.Copy(file, io.TeeReader(res.Body, counter))
	if err != nil {
		return err
	}
	log.Println("finished")

	return nil
}

const downloadAPI = `https://api.bilibili.com/x/player/playurl?bvid=%s&cid=%d`

func getDownloadUrl(BV string, cid int64) (string, error) {
	res, err := http.Get(fmt.Sprintf(downloadAPI, BV, cid))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	data := gjson.Get(string(content), "data")
	return data.Get("durl").Get("url").String(), nil
}

func genCheckRedirectFunc(referer string) func(req *http.Request, via []*http.Request) error {
	return func(req *http.Request, via []*http.Request) error {
		req.Header.Set("Referer", referer)
		return nil
	}
}
