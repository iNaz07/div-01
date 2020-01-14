package student

// import "fmt"

func ActiveBits(n int) uint {
	var bits uint
	conv := ""
	for n > 0 {
		conv = string(n%2+48) + conv
		n /= 2
	}	
	for _, i := range conv {
		if i == '1' {
			bits++
		}
	}
	return bits
}
// func main() {
// 	nbits := ActiveBits(16)
// 	fmt.Println(nbits)
// }
