package magicnumber

import "fmt"

func NthMagicNumber(n int) int {
	pow := 1
	answer := 0
	for n != 0 {
		pow *= 5

		if n&1 == 1 {
			answer += pow
		}

		n >>= 1
	}
	return answer
}

func main() {
	fmt.Print(NthMagicNumber(5))
}
