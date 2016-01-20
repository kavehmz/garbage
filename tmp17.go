package main

import "fmt"

func mapper(operator func(interface{}) interface{}, m interface{}) (result interface{}) {
	switch m.(type) {
	case []int:
		result := m.([]int)
		for i, n := range result {
			result[i] = operator(n).(int)
		}
		return result
	}
	return result
}

func curriedMapper(operator func(interface{}) interface{}) func(interface{}) interface{} {
	return func(m interface{}) (result interface{}) {
		switch m.(type) {
		case []int:
			result := m.([]int)
			for i, n := range result {
				result[i] = operator(n).(int)
			}
			return result
		}
		return result
	}
}

func main() {
	square := func(value interface{}) interface{} {
		return value.(int) * value.(int)
	}

	squaring := curriedMapper(square)
	fmt.Println(squaring([]int{1, 2, 3}))

	fmt.Println(mapper(square, []int{1, 2, 3}))
}
