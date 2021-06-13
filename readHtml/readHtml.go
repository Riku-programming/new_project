package readHtml

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	TimeToString "pdf_new_app/timeToString"
	"strings"
	"time"
)

//Replacerで変換したいものをかえてる
//Matchで変換したいファイルを指定してる
// todo 将来的にこれを内包した関数にしたい

func Replace(path string, fi os.FileInfo, err error) error {
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
		read, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		CopyFile()
		replacer := strings.NewReplacer("Hello", "Goodbye", "world", "5000", "date", TimeToString.TimeToString("01/02", time.Now()))
		newContents := replacer.Replace(string(read))
		err = ioutil.WriteFile(path, []byte(newContents), 0)
		RenameTmpAndTemplate()
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func CopyFile() {
	//fixme 絶対パスを変更する
	b, err := ioutil.ReadFile("/Users/riku/go/src/pdf_new_app/template/template.html")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("/Users/riku/go/src/pdf_new_app/template/tmp.html", b, 0644)
	if err != nil {
		panic(err)
	}
}

func RenameTmpAndTemplate() {
	if err := os.Rename("/Users/riku/go/src/pdf_new_app/template/template.html", "/Users/riku/go/src/pdf_new_app/template/result.html"); err != nil {
		fmt.Println(err)
	}
	if err := os.Rename("/Users/riku/go/src/pdf_new_app/template/tmp.html", "/Users/riku/go/src/pdf_new_app/template/template.html"); err != nil {
		fmt.Println(err)
	}
}
