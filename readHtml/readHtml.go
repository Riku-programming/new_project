package readHtml

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	ConvertHtmlToPdf "pdf_new_app/convertHtmlToPdf"
	readCsvToObject "pdf_new_app/readCsv"
	"strconv"
	"strings"
)

//fixme 以下の関数を切り分ける
//Replacerで変換したいものをかえてる
//Matchで変換したいファイルを指定してる
const templateFilePath = "/Users/riku/go/src/pdf_new_app/template/template.html"
const tmpFilePath = "/Users/riku/go/src/pdf_new_app/template/readCsv.html"
const resultFilePath = "/Users/riku/go/src/pdf_new_app/template/result.html"
const openCSVPath = "/Users/riku/go/src/pdf_new_app/readCsv/addresses.csv"

// todo 将来的にReplacerを別関数に分けてユーザーが置換する文字列を選択できるようにしたい

//CsvをObjectにしたものを読み込んでそれをhtmlテンプレートに流し込み変換したものをmain関数に流す関数
//Replace関数の形式はいじれないためnumberで指定しているところをforで回す

func Replace(templatePath string, fi os.FileInfo, err error) error {
	name := readCsvToObject.Number(1).Name
	priceInt := readCsvToObject.Number(1).Price
	price := strconv.Itoa(priceInt)
	if err != nil {
		return err
	}
	if !!fi.IsDir() {
		return nil //
	}
	matched, err := filepath.Match("template.html", fi.Name())
	if err != nil {
		panic(err)
		return err
	}
	if matched {
		ReplaceToHtml(templatePath, name, price)
		//xxx for分で回して一枚ずつpdfを作るためにPdfにする処理をここに入れた
		ConvertHtmlToPdf.ConvertHtmlToPdf(resultFilePath)
	}
	return nil
}

// todo  goalはreplace("old_name, new_name, old_age, new_age)

func ReplaceToHtml(templatePath string, name string, price string) {
	read, err := ioutil.ReadFile(templatePath)
	if err != nil {
		panic(err)
	}
	CopyFile()
	replacer := SelectReplaceWord("__NAME__", name, "__PRICE__", price)
	newContents := replacer.Replace(string(read))
	err = ioutil.WriteFile(templatePath, []byte(newContents), 0)
	RenameTmpAndTemplate()
	if err != nil {
		panic(err)
	}
}

func CopyFile() {
	b, err := ioutil.ReadFile(templateFilePath)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = ioutil.WriteFile(tmpFilePath, b, 0644)
	if err != nil {
		panic(err)
	}
}

func RenameTmpAndTemplate() {
	if err := os.Rename(templateFilePath, resultFilePath); err != nil {
		fmt.Println(err)
	}
	if err := os.Rename(tmpFilePath, templateFilePath); err != nil {
		fmt.Println(err)
	}
}
func SelectReplaceWord(companyNameArea string, companyName string, priceArea string, price string) *strings.Replacer {
	return strings.NewReplacer(companyNameArea, companyName, priceArea, price)
}
