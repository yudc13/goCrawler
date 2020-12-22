package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

// 获取网页的编码
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

// 获取城市信息
func getCityList(content []byte)  {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(content, -1)
	for _, match := range matches{
		fmt.Printf("city: %s  url: %d \n", match[2],  match[1])
	}
}

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error")
	}
	ecode := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, ecode.NewDecoder())
	contents, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	getCityList(contents)
}
