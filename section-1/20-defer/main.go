package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := createFile("file.tmp")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer closeFile(f)
	writeFile(f, "Hello")
}

func createFile(name string) (*os.File, error) {
	fmt.Println("create file")
	return os.Create(name)
}

func closeFile(f *os.File) {
	fmt.Println("closing file")
	f.Close()
}

func writeFile(f *os.File, s string) {
	fmt.Println("write file")
	f.WriteString(s)
}
