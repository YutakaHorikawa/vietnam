package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fp, err := os.Open("./ac.log")
	if err != nil {
		panic("can't open ac.log")
	}

	defer fp.Close()
	reader := bufio.NewReaderSize(fp, 4096)
	t, err := reader.ReadBytes(2)
	if err != nil {
		panic("die")
	}

	fmt.Println(t)
}
