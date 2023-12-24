package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	dat, err := os.ReadFile("text")
	check(err)
	fmt.Println(string(dat))

	f, err := os.Open("text")
	check(err)

	defer f.Close()

	r4 := bufio.NewReader(f)
	for {
		if char, _, err := r4.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				check(err)
			}
		} else {
			fmt.Print(string(char))
		}
	}
}
