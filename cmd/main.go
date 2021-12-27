package main

import (
	"millennium/pkg/anki"
	"millennium/pkg/content"
	"millennium/pkg/section"
)

func main() {
	section.DownloadSectionInfo()
	content.DownloadContents()
	anki.BuildAnki()
}
