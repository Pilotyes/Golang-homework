//Самостоятельно изучите пакет encoding/csv. Напишите пример, иллюстрирующий его применение.
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("./phonebook.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ';'

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(record)
	}
}
