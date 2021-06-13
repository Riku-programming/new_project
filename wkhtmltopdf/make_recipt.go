package main

import (
	"go_tour/readHtml"
	"path/filepath"
)

func main() {
	err := filepath.Walk(".", readHtml.Replace)
	if err != nil {
		panic(err)
	}
	//
	//
	//wkhtmltopdf := "/usr/local/bin/wkhtmltopdf"
	//root , err := filepath.Abs("./")
	//if err != nil {
	//	fmt.Printf(err.Error())
	//	return
	//}
	////if root == "" {
	////	root == "." if root == ""
	////	template_html := root + "/template.html"
	////	temp_file := root + "/__tmp.file"
	////	output_file := root + "/output.pdf"
	////}
	//
	//if root == ""{
	//	root := "."
	//	templateHtml := root + "/template.html"
	//	temp_file := root + "/__tmp.file"
	//	output_file := root + "/output.pdf"
	//	file,err :=  os.OpenFile(templateHtml, os.O_RDWR, syscall.O_RDWR)
	//	if err != nil {
	//		fmt.Printf(err.Error())
	//		return
	//	}
	//
	//
	//}
	//
	//
}

