package student

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	size := 0
	for i := range base {
		size = i + 1 
	}
	if size < 2 {
		Print("NV")
		return
	}
	for i := 0; i <= size; i++ {
		for j := i+1; j < size; j++ {
			if base[i] == base[j] || base[i] == '-' || base[i] == '+' {
				Print("NV")
				return
			}	
		}
	}
	num := nbr
	out := ""
	for nbr != 0 {
		index := nbr%size
		if index < 0 {
			index *= -1	
		}
		out = string(base[index]) + out
			nbr /= size
	}
	if num < 0 {
		out = "-" + out
	}
	Print(out)
}
func Print(str string) {
	for _, v := range str {
		z01.PrintRune(v)
	}
}