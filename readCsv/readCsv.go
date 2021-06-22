package readCsvToObject

import (
	"fmt"
	"github.com/jszwec/csvutil"
	"io/ioutil"
)

type User struct {
	Name  string `csv:"name"`
	Price int    `csv:"price,omitempty"`
}

//todo こうしたい name = users[i].Name, age = users[i].Age

func Number(number int) User {
	var csvInput = useIoutilReadFile("/Users/riku/go/src/pdf_new_app/readCsv/addresses.csv")
	var users []User
	if err := csvutil.Unmarshal(csvInput, &users); err != nil {
		fmt.Println("error:", err)
	}
	return users[number]
}

func useIoutilReadFile(fileName string) []byte {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return bytes
}
