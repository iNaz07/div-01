package student

func AtoiBase(s string, base string) int {
	lenBase := 0
	for range base {
		lenBase++
	}
	if lenBase < 2 {
		return 0
	}
	for i := 0; i < lenBase; i++ {
		for j := i + 1; j < lenBase; j++ {
			if base[i] == base[j] || base[i] == '-' || base[i] == '+' {
				return 0
			}
		}
	}
	atoi := 0
	for _, l := range s {
		for index, v := range base {
			if l == v {
				atoi = atoi * lenBase + index
			}
		}		
	}
	return atoi
}