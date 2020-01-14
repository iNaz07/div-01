package student

// import "fmt"

func SplitWhiteSpaces(str string) []string {
	strNew := ""
	for i, a := range str {
		if i == 0 && !isSpace(a) {
			strNew += string(a)
		} else if i == 0 && isSpace(a) {
			continue
		} else { 
			if isSpace(a) && !isSpace(rune(str[i-1])) {
				strNew += "*" 
			} else if !isSpace(a) {
				strNew += string(a)
			}
		}
	}
	lenStr := 1 
	for _, v := range strNew {
		if v == '*' {
			lenStr++
			
		}
	}
	arr := make([]string, lenStr)
	newS := ""
	index := 0
	lenarr := -1
	for _, v := range strNew {		
		if v == '*' {
			index++
			newS = ""
		} else {
			newS += string(v)
			arr[index] = newS		
		}		
	}
	for range arr {
		lenarr++
	}
	if arr[lenarr] == "" {
		arr = arr[:lenarr]
	}
	return arr
}
func isSpace(s rune) bool {
	return (s == ' ' || s == '\n' || s == '\t')
}
// func main() {
// 	str := " Ada Lovelace wrote,"
// 	fmt.Println(SplitWhiteSpaces(str))
// }