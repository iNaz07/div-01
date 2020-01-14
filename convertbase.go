package student

// import "fmt"

func ConvertBase(nbr, baseFrom, baseTo string) string {
	convStr := ""
	len := 0 
	for range baseTo {
		len++
	}
	num := ConvertTen(nbr, baseFrom) 
	for num > 0 {
		convStr = string(baseTo[num%len]) + convStr
				num /= len
	}
	return convStr	
}
func ConvertTen(nbr, base string) int {
	conv := 0
	len := 0 
	for range base {
		len++
	} 
 	for _, a := range nbr {
		for i, b := range base {
			if a == b {
				conv = len*conv + i			
			}
		}
	}
	return conv
}
// func main() {
// 	result := ConvertBase("10010001000011010000", "01", "0123456789ABCDEF")
// 	fmt.Println(result)
// }