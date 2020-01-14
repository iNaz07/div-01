package student 

func RecursivePower(nb int, power int) int {
	res := 0 
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else if power == 1 {
		return nb
	} else {
		res = nb * RecursivePower(nb, power-1)
	}
	return res	
}