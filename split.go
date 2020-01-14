package student

// import "fmt"

func Split(str, charset string) []string {
	lenCharset := 0
	lenStr := 0
	str = str
	for range charset {
		lenCharset++
	}
	for range str {
		lenStr++
	}
	// fmt.Println(lenStr)
	newStr := ""
	n := 0
 	for i := 0; i < lenStr; i++ {
		 if i == lenStr-lenCharset && str[i:lenStr] != charset {
			newStr += str[n:lenStr]
			break
		 }
		 if str[i:i+lenCharset] == charset {			 
			 newStr += str[n:i]
			 newStr += " "
			 n = i + lenCharset
		 }
	}
	space := 1 
	for _, a := range newStr {
		if a == ' ' {
			space++
		}		
	}
	index := 0
	newS := ""
	arr := make([]string, space)
	for _, a := range newStr {
		if a == ' ' {
			index++
			newS = ""
		} else {
			newS += string(a)
			arr[index] = newS
		}		
	}
	return arr
}

/*
func main() {
	str := "!==!the!==!term!==!computer!==!came!==!to!==!refer!==!to!==!the!==!machines!==!rather!==!than!==!their!==!human!==!predecessors.[12]!==!As!==!it!==!became!==!clear!==!that!==!computers!==!could!==!be!==!used!==!for!==!more!==!than!==!just!==!mathematical!==!calculations,!==!"
	fmt.Println(Split(str, "!==!"))
}
*/
