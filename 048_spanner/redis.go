// gcloud Redis instances create user-db-test --config regional-XXXXX --description test --nodes 3
package main

import (
	"flag"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
)

var redisURL string
var redisPool *redis.Pool

func init() {
	flag.StringVar(&redisURL, "redis", "", "redisURL")
}

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     1000,
		IdleTimeout: 60 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.DialURL(redisURL) },
	}
}

func getRedisClient() *redis.Pool {
	if redisPool == nil {
		redisPool = newPool(redisURL)
	}
	return redisPool
}

func redisInsert(n int) {
	c := getRedisClient().Get()
	defer c.Close()
	(c).Do("SET", strconv.Itoa(n), strconv.Itoa(n))
}
