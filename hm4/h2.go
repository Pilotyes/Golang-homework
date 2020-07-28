// Создать тип, описывающий контакт в телефонной книге. Создать псевдоним типа телефонной книги (массив контактов) и реализовать для него интерфейс Sort{}.
package main

import (
	"fmt"
	"sort"
)

type info struct {
	firstName, lastName string
	phoneNumber         int
}

type phoneBook []info

func (pb *phoneBook) Sort() {
	sort.Sort(pb)
}

func (pb *phoneBook) Len() int {
	return len(*pb)
}

func (pb *phoneBook) Swap(i, j int) {
	(*pb)[i], (*pb)[j] = (*pb)[j], (*pb)[i]
}

func (pb *phoneBook) Less(i, j int) bool {
	//return (*pb)[i].phoneNumber < (*pb)[j].phoneNumber
	return (*pb)[i].firstName < (*pb)[j].firstName
	//return (*pb)[i].lastName < (*pb)[j].lastName
}

func main() {
	pb := phoneBook{
		{
			firstName:   "Dominik",
			lastName:    "Smith",
			phoneNumber: 89231467231,
		},
		{
			firstName:   "Angelina",
			lastName:    "Jolie",
			phoneNumber: 89211141413,
		},
		{
			firstName:   "Kventin",
			lastName:    "Tarantino",
			phoneNumber: 89111337654,
		},
	}
	fmt.Println(pb)
	pb.Sort()
	fmt.Println(pb)
}
