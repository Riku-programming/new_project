package main

import (
	"path/filepath"
	"pdf_new_app/convertHtmlToPdf"
	"pdf_new_app/csvArray"
	"pdf_new_app/readHtml"
)

const dateColumn = 0
const nameColumn = 1
const priceColumn = 2

const templatePath = "/Users/riku/go/src/pdf_new_app/template"
const openCSVPath = "/Users/riku/go/src/pdf_new_app/readCsv/addresses.csv"

func main() {
	CsvArray.CsvArray(openCSVPath)
	err := filepath.Walk(templatePath, readHtml.Replace)
	if err != nil {
		panic(err)
	}
	ConvertHtmlToPdf.ConvertHtmlToPdf()
}
