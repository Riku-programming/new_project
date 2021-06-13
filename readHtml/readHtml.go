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
	matched, err := filepath.Match("test.html", fi.Name())
	if err != nil {
		panic(err)
		return err
	}
	if matched {
		read, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		fmt.Println(path)
		replacer := strings.NewReplacer("__NAME__", "Company", "__PRICE__", "5000", "__DATE__", TimeToString.TimeToString("01/02", time.Now()))
		fmt.Println(replacer)
		newContents := replacer.Replace(string(read))
		err = ioutil.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			panic(err)
		}
	}
	return nil
}
