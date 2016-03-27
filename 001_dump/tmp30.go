package main

import (
	"encoding/json"
	"fmt"
)

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
}
