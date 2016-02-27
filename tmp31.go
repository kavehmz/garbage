package main

import (
	"encoding/json"
	"fmt"
)

// Δ is upper case so it needs some comments!
func Δ(x, y int) int {
	return x - y
}

func چاپ() {
	fmt.Println("آروین")
}

func 인쇄() {
	fmt.Println("실험")
}

func 打印() {
	fmt.Println("實驗")
}

func main() {
	var data = []byte(`{"status": 200, "name":"test", "t":{"a":"b"}}`)

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("error:", err)
		return
	}

	var status = uint64(result["status"].(float64)) //ok
	fmt.Println("status value:", status)

	fmt.Println("status value:", result["t"].(map[string]interface{})["a"])

	fmt.Println(Δ(8, 2))
	چاپ()
	인쇄()
	打印()

}
