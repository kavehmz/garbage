package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
)

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	client := gohbase.NewClient("localhost")

	// Values maps a ColumnFamily -> Qualifiers -> Values.
	values := map[string]map[string][]byte{"cf": map[string][]byte{"ab": []byte("kmztest")}}
	putRequest, err := hrpc.NewPutStr(context.Background(), "test", "key", values)
	checkErr(err)
	rsp, err := client.Put(putRequest)
	checkErr(err)
	fmt.Println(rsp)

	fmt.Println("=====================")

	getRequest, err := hrpc.NewGetStr(context.Background(), "test", "row")
	checkErr(err)
	getRsp, err := client.Get(getRequest)
	checkErr(err)
	fmt.Println(getRsp)

}
