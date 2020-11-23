package main

import (
	"fmt"
	"os"

	"./brainf"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expecting File name")
	} else {
		brainf.RunFile(os.Args[1])
	}
}
