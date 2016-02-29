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

	titleRegex := regexp.MustCompile(`<span itemprop="name" class="series-title block ellipsis">(.+)?<\/span>`)
	descRegex := regexp.MustCompile(`<p class="short-desc">(\s.+)?<\/p>`)

	titleTags := titleRegex.FindAllStringSubmatch(crunchyHTML, -1)
	descTags := descRegex.FindAllStringSubmatch(crunchyHTML, -1)

	for r := range titleTags {
		fmt.Printf("%-50s | %-50s\n", titleTags[r][1], strings.TrimSpace(descTags[r][1]))
	}
}
