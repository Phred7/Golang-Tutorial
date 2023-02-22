package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: File name should be passed in to go run as Arg[1]")
		os.Exit(1)
	}
	fn := os.Args[1]
	fmt.Println(fn)

	f, err := os.Open(fn)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
