//* Напишите утилиту для копирования файлов, используя пакет flag.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var src, dst string

func init() {
	flag.StringVar(&src, "src", "", "source file")
	flag.StringVar(&dst, "dst", "", "destination file")
}

func main() {
	flag.Parse()

	if src == "" {
		log.Fatal("No input file")
	}

	if dst == "" {
		log.Fatal("No output file")
	}

	if s, err := os.Open(src); err != nil {
		log.Fatal(err)
	} else {
		defer s.Close()

		d, err := os.Create(dst)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer d.Close()

		_, err = io.Copy(d, s)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
