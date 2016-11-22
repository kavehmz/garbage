package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"strconv"

	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
)

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	client := gohbase.NewClient("localhost", gohbase.RpcQueueSize(1), gohbase.FlushInterval(time.Nanosecond))

	t := time.Now()
	for i := 0; i < 100; i++ {
		// Values maps a ColumnFamily -> Qualifiers -> Values.
		values := map[string]map[string][]byte{"cf": map[string][]byte{"ab": []byte("kmztest")}}
		putRequest, err := hrpc.NewPutStr(context.Background(), "test", "key"+strconv.Itoa(i), values)
		checkErr(err)
		_, err = client.Put(putRequest)
		checkErr(err)
	}
	fmt.Println(float64(time.Now().UnixNano()-t.UnixNano()) / 1000000)

	fmt.Println("=====================")

	getRequest, err := hrpc.NewGetStr(context.Background(), "test", "row")
	checkErr(err)
	getRsp, err := client.Get(getRequest)
	checkErr(err)
	fmt.Println(getRsp)

}
