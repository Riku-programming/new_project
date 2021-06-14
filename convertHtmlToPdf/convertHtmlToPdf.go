package ConvertHtmlToPdf

import (
	"bytes"
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"html/template"
	"log"
	"strings"
)

type templateData struct {
	Datas []data
}

type data struct {
	Name string
	Cal  string
}

func ConvertHtmlToPdf() {
	datas := make([]data, 0)
	datas = append(datas, data{Name: "M16", Cal: "5.56"})
	datas = append(datas, data{Name: "AK47", Cal: "7.62"})
	datas = append(datas, data{Name: "MP5", Cal: "9.00"})
	tmplData := templateData{Datas: datas}

	t := template.Must(template.ParseFiles("/Users/riku/go/src/pdf_new_app/template/result.html"))

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, tmplData); err != nil {
		log.Fatal(err)
	}
	html := tpl.String()

	pdfg := wkhtmltopdf.NewPDFPreparer()
	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(html)))
	pdfg.Dpi.Set(600)

	jsonBytes, err := pdfg.ToJSON()
	if err != nil {
		log.Fatal(err)
	}
	pdfgFromJSON, err := wkhtmltopdf.NewPDFGeneratorFromJSON(bytes.NewReader(jsonBytes))
	if err != nil {
		log.Fatal(err)
	}

	err = pdfgFromJSON.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = pdfgFromJSON.WriteFile("./output-result.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Print Done\nPdf size %d bytes", pdfgFromJSON.Buffer().Len())
}
