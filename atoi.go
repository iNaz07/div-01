package student 

func Atoi(s string) int {
	num := 0
	str := s
	lens := 0
	for range s {
		lens ++
	}
	if lens <= 0 {
		return 0
	}
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
	}
	if s == "" {
		return 0
	}
 	for _, i := range []byte(s) {
		nb := int(i - 48)
		if nb >= 0 && nb <= 9 {
			num = num*10 + nb
		} else {
			return 0
		}
	}
	if str[0] == '-' {
		num = num*-1
	}
	return num
}