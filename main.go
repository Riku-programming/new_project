package main

import (
	"path/filepath"
	"pdf_new_app/convertHtmlToPdf"
	"pdf_new_app/readHtml"
)

const templatePath = "/Users/riku/go/src/pdf_new_app/template"

func main() {
	err := filepath.Walk(templatePath, readHtml.Replace)
	if err != nil {
		panic(err)
	}
	ConvertHtmlToPdf.ConvertHtmlToPdf()
}
