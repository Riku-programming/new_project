package main

import (
	"path/filepath"
	"pdf_new_app/readHtml"
)

const templatePath = "/Users/riku/go/src/pdf_new_app/template"
const openCSVPath = "/Users/riku/go/src/pdf_new_app/readCsv/addresses.csv"
const inputFilePath = "/Users/riku/go/src/pdf_new_app/template/result.html"

func main() {
	err := filepath.Walk(templatePath, readHtml.Replace)
	if err != nil {
		panic(err)
	}
}
