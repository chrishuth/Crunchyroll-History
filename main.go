package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
	html := getHTML("http://www.crunchyroll.com/home/history")
	fmt.Println(html)
}
