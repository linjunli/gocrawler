// 获取网页信息
package download

import (
	"errors"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

var (
	ErrorNil = errors.New("response is nil")
	ErrorWrongCode = errors.New("http response code is wrong")
)

func Download(url string) (*goquery.Document, error) {
	var (
		resp *http.Response
		err error
	)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil); if err != nil {
		return nil, ErrorNil
	}
	req.Header.Set("User-Agent","Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Mobile Safari/537.36")

	if resp, err = client.Do(req); err != nil {
		return nil, ErrorNil
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, ErrorWrongCode
	}

	return goquery.NewDocumentFromReader(resp.Body)
}