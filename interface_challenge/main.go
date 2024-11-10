package main

import (
	// "fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    // data := make([]byte, 100)
    // count, err := file.Read(data)
    // if err != nil {
    //     log.Fatal(err)
    // }
    io.Copy(os.Stdout, file)
    //fmt.Printf("read %d bytes: %q\n", count, data[:count])
}