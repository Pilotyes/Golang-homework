//** Напишите упрощенный аналог утилиты grep.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var file, matchString string

func init() {
	flag.StringVar(&file, "file", "", "input file")
	flag.StringVar(&matchString, "string", "", "match string")
}

func main() {
	flag.Parse()
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return
		}

		if strings.Contains(string(line), matchString) {
			fmt.Println(string(line))
		}
	}
}
