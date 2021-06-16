package CsvArray

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func CsvArray(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	r.Comma = ';'

	//for {
	record, err := r.ReadAll()
	if err == io.EOF {
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(record[0][0])
}
