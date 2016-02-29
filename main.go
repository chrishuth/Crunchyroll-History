package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func getHTML(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Cookie", os.Getenv("CRUNCHYCOOKIE"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func main() {
	crunchyHTML := getHTML("http://www.crunchyroll.com/home/history")

	titleRX := regexp.MustCompile(`<span itemprop="name" class="series-title block ellipsis">(.+)?<\/span>`)
	descRX := regexp.MustCompile(`<p class="short-desc">(\s.+)?<\/p>`)

	titleTags := titleRX.FindAllStringSubmatch(crunchyHTML, -1)
	descTags := descRX.FindAllStringSubmatch(crunchyHTML, -1)

	for r := range titleTags {
		fmt.Println(titleTags[r][1], "  ", strings.TrimSpace(descTags[r][1]))
	}
}
