package main 

import (
	"os"
	// "github.com/01-edu/z01"
	"fmt"
)

func main() {
	size := 0
	arg := os.Args[1:]
	for range arg {
		size++
	}
	arrNum := []string{}
	arrName := []string{}
	for i := 0; i < size; i++ {
		argm := arg[i]
		argsNum := argm[2:]
		if arg[i] == "-c" {
			if isNumber(arg[i+1]) == true {
				arrNum = append(arrNum, arg[i+1])
				i++
			} else {
				return
			}
		} else if argm[0:2] == "-c" {
			if isNumber(argsNum) == true {
				arrNum = append(arrNum, argsNum)
			} else {
				return
			}
		} else {
			arrName = append(arrName, arg[i])
		}
	}
	numLen := -1 
	for range arrNum {
		numLen++
	}
	flag := false 
	for i := 0; i <=  numLen; i++ {
		for j, num := range arrNum[i] {
			if j == 0 && num == '+' {
				flag = true
			}
		}
	}
	if flag == true {
		arrNum[numLen] = "+" + arrNum[numLen]
	}
	number := arrNum[numLen]

	fileLen := 0
		for range arrName {
			fileLen++
		}
		plus := false 
		for i := 0; i < fileLen; i++ {
			file, err := os.Open(arrName[i])
			if err != nil {
				fmt.Printf("tail: cannot open '%v' for reading: No such file or directory\n", arrName[i])
			} else {
				defer file.Close()
				if fileLen > 1 {
					if plus == true && i != 0 {
						fmt.Printf("\n")
					}
					fmt.Printf("==> %v <==\n", arrName[i])
					plus = true
				}
				var size int 
				stat, _ := file.Stat()
				size = int(stat.Size())
				if size < AtoiTail(number) {
					number = string(size)
				}
				data := []byte{}
				for i := 0; i < size; i++ {
					data = append(data, 0)
				}
				a := bytes(number, size)
				file.Read(data)
				fmt.Printf(string(data[a:]))
			}
		}
}
func bytes(argByte string, size int) int {
	var a int 
	for index, l := range argByte {
		if index == 0 && l == '+' {
			a = AtoiTail(argByte) - 1 
		} else if index == 0 && l != '+' {
			a = size - AtoiTail(argByte)
		}
	}
	return a 
}
func AtoiTail(str string) int {
	var n int 
	for _, l := range str {
		if l >= '0' && l <= '9' {
			count := 0
			for l > '0' {
				count++
				l--
			}
			n = n*10 + int(count)
		}
	}
	return n
}
func isNumber(n string) bool {
	for index, l := range n {
		if l == '-' || l == '+' {
			if index == 0 {
				continue
			} else {
				fmt.Printf("tail: invalid number of bytes: '&q' \n", n)
				return false
			}
		} else if l < '0' || l > '9' {
			fmt.Printf("tail: invalid number of bytes: '%v' \n", n)
			return false
		}
	}
	return true
}
 