package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	byt := []byte(`{"num":6.13,"strs":["a","b"], "contact":{"name":1, "family":"zamani"}}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	contact := (dat["contact"]).(map[string]interface{})
	fmt.Println((contact["name"]).(float64))
}
