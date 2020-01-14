package student 

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	Recursive(0, n+2, n, ", ")
	z01.PrintRune('\n')
}
func Recursive(start, end, n int, out string) {
	base := ", 012345678"
	if n == 0 {
		if out == base[:end] {
			for _, v := range out {
				if v >= '0' && v <= '9' {
					z01.PrintRune(v)					
				}
			}
			return
		}
		for _, v := range out {
			z01.PrintRune(v)			
		}		
	}
	for i := start; i <= 9; i++ {
		str := out + string(i+48) 
		Recursive(i + 1, end, n - 1, str)
	}
}
