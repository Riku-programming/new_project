package main

import (
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
)

//func main() {
//	// Create new PDF generator
//	pdfg, err := wkhtmltopdf.NewPDFGenerator()
//	if err != nil {
//		log.Fatal(err)
//	}
//	url := "https://google.com/"
//
//	pdfg.AddPage(wkhtmltopdf.NewPage(url))
//
//	// PDF作成
//	err = pdfg.Create()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// 出力
//	err = pdfg.WriteFile("./google.pdf")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println("tada!")
//}

func main() {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}
 	p := "./test.html"
 	page := wkhtmltopdf.NewPage(p)
 	page.UserStyleSheet.Set("./style.css")
 	pdfg.AddPage(page)

	// PDF作成
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// 出力
	err = pdfg.WriteFile("./test.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("tada!")
}

//func main() {
//	// Create new PDF generator
//	pdfg, err := wkhtmltopdf.NewPDFGenerator()
//	if err != nil {
//		log.Fatal(err)
//	}
//	tmp1 := "./test.html"
//	tmp2 := "./test2.html"
//	tmp3 := "./test3.html"
//
//	pdfg.AddPage(wkhtmltopdf.NewPage(tmp1))
//	pdfg.AddPage(wkhtmltopdf.NewPage(tmp2))
//	pdfg.AddPage(wkhtmltopdf.NewPage(tmp3))
//
//	// PDF作成
//	err = pdfg.Create()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// 出力
//	err = pdfg.WriteFile("./test.pdf")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println("tada!")
//}
//
