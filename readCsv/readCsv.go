package ReadCsv

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCsv(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var line []string

	for {
		line, err = reader.Read()
		if err != nil {
			break
		}
		fmt.Println(line)
	}
	// fixme  配列にlineを格納してindex[0][0]みたいな感じにした方が良さそう
}
