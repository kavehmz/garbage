package main

import "fmt"
import "strconv"

var sum int64

func divisibile(digits []int) {
	if (digits[1]*100+digits[2]*10+digits[3])%2 == 0 && (digits[2]*100+digits[3]*10+digits[4])%3 == 0 && (digits[3]*100+digits[4]*10+digits[5])%5 == 0 && (digits[4]*100+digits[5]*10+digits[6])%7 == 0 && (digits[5]*100+digits[6]*10+digits[7])%11 == 0 && (digits[6]*100+digits[7]*10+digits[8])%13 == 0 && (digits[7]*100+digits[8]*10+digits[9])%17 == 0 {
		var c string
		for i := 0; i <= 9; i++ {
			c += strconv.Itoa(digits[i])
		}
		n, _ := strconv.ParseInt(c, 10, 64)
		sum += n
		fmt.Println(sum)
	}
}

func permute(items []int, located []int) {
	if len(items) == 0 {
		divisibile(located)
	}
	for i := 0; i < len(items); i++ {
		slice1 := make([]int, i)
		slice2 := make([]int, len(items)-1-i)
		copy(slice1, items[:i])
		copy(slice2, items[i+1:])
		permute(append(slice1, slice2...), append(located, items[i]))
	}
}

func main() {
	permute([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, []int{})
}
