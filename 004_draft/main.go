package main

import (
	"fmt"
	"log"

	"github.com/kavehmz/garbage/do/gen"
)

func main() {
	web := gen.Web("jp-ui")
	redis := gen.Redis("jp-cache")
	db := gen.Postgres("jp-db")
	proxy := gen.Proxy("jp-proxy")

	if <-web.Ready {
		web.AddRedis(redis)
		web.AddDB(db)
	} else {
		log.Println("Something went wrong")
	}

	if <-web.Ready && <-proxy.Ready {
		proxy.Add(web)
	}
	fmt.Println(web, redis, db)

}
