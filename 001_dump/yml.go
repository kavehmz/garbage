package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"time"

	"io/ioutil"

	"github.com/garyburd/redigo/redis"

	"github.com/lib/pq"
	"gopkg.in/yaml.v2"
)

type ClientDB struct {
	Password string
	Company  map[string]struct {
		Write struct {
			Name string
			IP   string
		}
	} `yaml:",inline"`
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func publish(conn redis.Conn, payload []string) {
	m := make(map[string]string)

	m["id"] = payload[0]
	m["account_id"] = payload[1]
	m["action_type"] = payload[2]
	m["referrer_type"] = payload[3]
	m["financial_market_bet_id"] = payload[4]
	m["payment_id"] = payload[5]
	m["amount"] = payload[6]
	m["balance_after"] = payload[7]

	jsonVal, _ := json.Marshal(m)
	msg := string(jsonVal)
	conn.Do("PUBLISH", "balance_"+m["account_id"], msg)
	conn.Do("PUBLISH", "buy_"+m["account_id"], msg)
	conn.Do("PUBLISH", "sell_"+m["account_id"], msg)
	conn.Do("PUBLISH", "transaction_"+m["account_id"], msg)
	conn.Do("PUBLISH", "payment_"+m["account_id"], msg)
}

func waitForNotification(clientdb ClientDB, company string) {
	conninfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=require", "postgres", clientdb.Password, clientdb.Company[company].Write.IP, "test")
	listener := pq.NewListener(conninfo, 10*time.Second, time.Minute, nil)
	err := listener.Listen("getwork")
	checkErr(err)

	redisdb, err := redis.DialURL(os.Getenv("REDIS_URL"))
	checkErr(err)
	var notification *pq.Notification
	for {
		select {
		case notification = <-listener.Notify:
			if notification != nil {
				publish(redisdb, regexp.MustCompile(",").Split(notification.Extra, -1))
			}

		case <-time.After(60 * time.Second):
			fmt.Println("no notifications for 60 seconds...")
		}
	}
}

func main() {
	var clientdb ClientDB
	source, _ := ioutil.ReadFile("clientdb.yml")
	yaml.Unmarshal(source, &clientdb)

	for company, _ := range clientdb.Company {
		go waitForNotification(clientdb, company)
	}

	// Here we just wait for kill signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	s := <-c
	fmt.Println("Got a kill signal:", s)
}
