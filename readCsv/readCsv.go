package main

import (
	"fmt"
	"github.com/jszwec/csvutil"
	"io/ioutil"
)

func main() {
	var csvInput = useIoutilReadFile("/Users/riku/go/src/pdf_new_app/readCsv/addresses.csv")

	type User struct {
		Name string `csv:"name"`
		Age  int    `csv:"age,omitempty"`
		//CreatedAt time.Time
	}

	var users []User
	if err := csvutil.Unmarshal(csvInput, &users); err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(users[0].Age)
}
func useIoutilReadFile(fileName string) []byte {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return bytes
}
