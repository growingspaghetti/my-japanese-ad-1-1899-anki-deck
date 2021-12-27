package content

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	m "millennium/pkg/models"
	"millennium/pkg/section"
)

type (
	sections m.Sections
	kind     string
)

const (
	Dirname      = "contents"
	events  kind = "e"
	births  kind = "b"
	deaths  kind = "d"
)

var (
	Kinds = [3]kind{events, births, deaths}
)

func buildUrl(pageId int, sectionId string) string {
	return strings.Join([]string{
		"https://ja.wikipedia.org/w/api.php",
		"?action=parse",
		"&pageid=",
		strconv.Itoa(pageId),
		"&format=json&prop=text&wrapoutputclass",
		"&section=",
		sectionId,
		"&disablelimitreport&disableeditsection",
	}, "")
}

func download(y, pageId int, sectionId string, k kind) {
	u := buildUrl(pageId, sectionId)
	log.Printf("downloading %s\n", u)
	resp, err := http.Get(u)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	f := fmt.Sprintf("%s/%d_%s.json", Dirname, y, k)
	if err := ioutil.WriteFile(f, bytes, 0644); err != nil {
		panic(err)
	}
}

func (s *sections) downloadContents(y int) {
	for _, section := range s.Parse.Sections {
		switch section.Line {
		case "できごと":
			fallthrough
		case "出来事":
			download(y, s.Parse.PageId, section.Index, events)
		case "誕生":
			download(y, s.Parse.PageId, section.Index, births)
		case "死去":
			download(y, s.Parse.PageId, section.Index, deaths)
		default:
		}
	}
}

func DownloadContents() {
	err := os.Mkdir(Dirname, 0744)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
	for y := 1; y < 2021; y++ {
		file := fmt.Sprintf("%s/%d.json", section.Dirname, y)
		bytes, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		sections := new(sections)
		if err := json.Unmarshal(bytes, sections); err != nil {
			panic(err)
		}
		log.Printf("%v\n", sections)
		sections.downloadContents(y)
	}
}
