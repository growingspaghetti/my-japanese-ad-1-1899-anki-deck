package section

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

const (
	Dirname = "sections"
)

func buildUrl(y int) string {
	return strings.Join([]string{
		"https://ja.wikipedia.org/w/api.php",
		"?action=parse",
		"&format=json&page=",
		strconv.Itoa(y),
		url.QueryEscape("年"),
		"&prop=sections&disabletoc=1",
	}, "")
}

func download(y int, u string) {
	log.Printf("downloading %s\n", u)
	resp, err := http.Get(u)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	f := fmt.Sprintf("%s/%d.json", Dirname, y)
	if err := ioutil.WriteFile(f, bytes, 0644); err != nil {
		panic(err)
	}
}

func DownloadSectionInfo() {
	err := os.Mkdir(Dirname, 0744)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
	for y := 1; y < 2021; y++ {
		u := buildUrl(y)
		download(y, u)
	}
}
