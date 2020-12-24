package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// 根据url抓取网页数据
func Fetch(u string) (contents []byte, err error) {
	client := &http.Client{}
	newUrl := strings.Replace(u, "http://", "https://", 1)
	req, err := http.NewRequest("GET", newUrl, nil)
	if err != nil {
		return nil, err
	}
	cookie := "sid=bL6TAtqBcgPE7TRWCwKQ; ec=fBQ6u8jh-1608570524649-611118dabd0cf1613101641; _pc_user_info_=%7B%22gender%22%3A1%2C%22marriage%22%3A1%2C%22workCity%22%3A%2210102013%22%2C%22birthday%22%3A327254400000%7D; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1608570952; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1608820037; _exid=9Oa9ViK3V65jKSxCw9Whva0oETYFGaF5hSoXW7fWI8EmvdacjIc9G6I%2BihOCjWspEzVXfK6qRQatkB%2FoKjv5SQ%3D%3D; _efmdata=wR8TB7c%2FAC6qjZLX8fVrzRHwQXVgYhaNxV5%2B28K70oK%2Bc%2F6YVyYjrkeGeK6rBfdClmO9pL%2BWAp8j5G8eRd%2BjE7vCOPjsgNI%2F%2FDoRBBDVc74%3D"
	unescape, err := url.QueryUnescape(cookie)
	if err != nil {
		fmt.Println("url.QueryUnescape error ", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")
	req.Header.Add("cookie", unescape)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code %d", resp.StatusCode)
	}
	ecode := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, ecode.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

// 获取网页的编码
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetcher error %s", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
