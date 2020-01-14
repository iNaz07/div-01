package main

import (
	// "fmt"
	"github.com/01-edu/z01"
	"os"
)

func main() {
	args := os.Args[1:]
	v1 := args[0]
	v2 := args[2]
	op := args[1]
	result := 0
	lenArg := 0
	for range args {
		lenArg++
	}
	// fmt.Println(Atoi(v2))
	if lenArg == 3 {
		if !isValidOp(op) || Atoi(v1) == 0 || Atoi(v2) == 0 && v2 != "0" {
			z01.PrintRune('0')
			z01.PrintRune('\n')
			return
		}
		if op == "+" {
			result = Atoi(v1) + Atoi(v2)
			if Atoi(v1) > 0 && Atoi(v2) > 0 && result < 0 || Atoi(v1) < 0 && Atoi(v2) < 0 && result > 0 {
				z01.PrintRune('0')
				z01.PrintRune('\n')
				return
			}
			Print(itoa(result))
		}
		if op == "-" {
			result = Atoi(v1) - Atoi(v2)
			if isOverflow(Atoi(v1)) && Atoi(v2) < 0 {
				z01.PrintRune('0')
				z01.PrintRune('\n')
				return
			}
			Print(itoa(result))
		}
		if op == "*" {
			if Atoi(v1) == 0 || Atoi(v2) == 0 {
				z01.PrintRune('0')
				z01.PrintRune('\n')
				return
			}
			result = Atoi(v1) * Atoi(v2)
			if Atoi(v2) != result/Atoi(v1) || Atoi(v1) != result/Atoi(v2) || Atoi(v1) > 0 && Atoi(v2) > 0 && result < 0 {
				z01.PrintRune('0')
				z01.PrintRune('\n')
				return
			}
			Print(itoa(result))
		}
		if op == "/" {
			div := "No division by 0"
			if v2 == "0" {
				Print(div)
				return
			}
			result = Atoi(v1) / Atoi(v2)
			if Atoi(v1) < 0 && Atoi(v2) < 0 && result < 0 {
				z01.PrintRune('0')
				z01.PrintRune('\n')
				return
			}
			Print(itoa(result))
		}
		if op == "%" {
			mod := "No modulo by 0"
			if v2 == "0" {
				Print(mod)
				return
			}
			result = Atoi(v1) % Atoi(v2)
			Print(itoa(result))
		}
	}

}

func isValidOp(s string) bool {
	for _, a := range s {
		return a == '+' || a == '-' || a == '*' || a == '/' || a == '%'
	}
	return false
}
func Atoi(s string) int {
	lens := 0
	for range s {
		lens++
	}
	if lens == 0 || lens > 20 {
		return 0
	}
	if lens == 20 && s > "-9223372036854775808" {
		return 0
	}
	num := 0
	str := s

	if s[0] == '-' {
		s = s[1:]
	}
	for _, i := range s {

		if i >= 48 && i <= 57 {
			num = num*10 + int(i-48)
		} else {
			return 0
		}

	}
	if str[0] == '-' {
		num = num * -1
	}
	if str[0] != '-' && num < 0 {
		return 0
	}
	return num
}
func isOverflow(n int) bool {
	return n == 9223372036854775807 || n == -9223372036854775808
}
func Print(s string) {
	for _, a := range s {
		z01.PrintRune(a)
	}
	z01.PrintRune('\n')
}
func itoa(n int) string {
	res := ""
	nbr := n
	if n == 0 {
		res = "0"
	}
	for n != 0 {
		index := n % 10
		if index < 0 {
			index *= -1
		}
		res = string(index+48) + res
		n /= 10
	}
	if nbr < 0 {
		res = "-" + res
	}
	return res
}
