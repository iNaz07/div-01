package main

import (
	"io/ioutil"
	"os"
	"io"
	"github.com/01-edu/z01"
)

func main() {
	size := 0 
	for i := range os.Args[1:] {
		size = i + 1 
	}
	if size == 0 {
		io.Copy(os.Stdout, os.Stdin)	
	} else {
		for _, v := range os.Args[1:] {
			check(v)
		}
	}
}

func check(s string) {
	error := "cat "+s+": no such file or directory"
	content, err := ioutil.ReadFile(s)
	if err != nil {
		for _, v := range error {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
	for _, v := range string(content) {
		z01.PrintRune(v)
	}
	// z01.PrintRune('\n')
}
