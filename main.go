package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	r := bufio.NewReader(file)

	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			fmt.Printf("%q\n", line)
		}

		if err != nil {
			break
		}
	}
}
