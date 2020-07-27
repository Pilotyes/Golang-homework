//* Внести в телефонный справочник дополнительные данные. Реализовать сохранение json-файла на диске с помощью пакета ioutil при изменении данных.

package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var fileName string = "./output.json"

type info struct {
	FirstName, LastName string
	PhoneNumbers        []int
}

func main() {
	addressBook := make(map[string]info)

	addressBook["Alex"] = info{
		FirstName:    "Alex",
		LastName:     "Sidorov",
		PhoneNumbers: []int{89996543210},
	}
	addressBook["Bob"] = info{
		FirstName:    "Bob",
		LastName:     "Petrov",
		PhoneNumbers: []int{89167243812, 89155243627},
	}

	wb, err := json.Marshal(addressBook)
	if err != nil {
		panic(err)
	}

	rb, err := ioutil.ReadFile(fileName)
	if err != nil {
		if err.(*os.PathError).Err.Error() != "no such file or directory" {
			panic(err)
		} else {
			err := ioutil.WriteFile("./output.json", wb, 0644)
			if err != nil {
				panic(err)
			}
		}
	} else {
		if string(rb) != string(wb) {
			err = ioutil.WriteFile("./output.json", wb, 0644)
			if err != nil {
				panic(err)
			}
		}

	}

}

