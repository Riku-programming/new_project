package readHtml

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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

func Replace(path string, fi os.FileInfo, err error) error {
	name := readCsvToObject.Number(0).Name
	priceInt := readCsvToObject.Number(0).Price
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
		ReplaceToHtml(path, name, price)
	}
	return nil
}

// todo  goalはreplace("old_name, new_name, old_age, new_age)
//replacer は外で定義したい

func SelectReplaceWord(companyNameArea string, companyName string, priceArea string, price string) *strings.Replacer {
	return strings.NewReplacer(companyNameArea, companyName, priceArea, price)
}

func ReplaceToHtml(path string, name string, price string) {
	read, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	CopyFile()
	replacer := SelectReplaceWord("__NAME__", name, "__PRICE__", price)
	newContents := replacer.Replace(string(read))
	err = ioutil.WriteFile(path, []byte(newContents), 0)
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
