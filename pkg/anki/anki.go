package anki

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"millennium/pkg/content"
	m "millennium/pkg/models"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type (
	text m.Content
)

var (
	href    = regexp.MustCompile(`<a.*?>`)
	toc     = regexp.MustCompile(`</div><ul>.*?</ul></div>`)
	cleaner = strings.NewReplacer("\t", "", "\n", "", "//upload.wikimedia.org", "https://upload.wikimedia.org", "</a>", "")
)

func readContents(f string) string {
	if _, err := os.Stat(f); errors.Is(err, os.ErrNotExist) {
		return ""
	}
	bytes, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}
	text := new(text)
	if err := json.Unmarshal(bytes, text); err != nil {
		panic(err)
	}
	return text.Parse.Text.Html
}

func clean(s string) string {
	s = href.ReplaceAllString(s, "")
	s = cleaner.Replace(s)
	s = toc.ReplaceAllString(s, "</div>")
	return s
}

func BuildAnki() {
	anki, err := os.OpenFile("anki.html", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer anki.Close()
	for y := 1; y < 2021; y++ {
		var sb strings.Builder
		sb.WriteString(strconv.Itoa(y))
		sb.WriteString("年\t\t")
		for _, k := range content.Kinds {
			f := fmt.Sprintf("%s/%d_%s.json", content.Dirname, y, k)
			s := readContents(f)
			sb.WriteString(clean(s))
			sb.WriteString("\t")
		}
		sb.WriteString("\n")
		anki.WriteString(sb.String())
	}
}
