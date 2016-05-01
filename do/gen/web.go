package gen

import "time"

// WebServer is
type WebServer struct {
	// ID is
	ID    string
	Ready chan bool
}

//AddDB is
func (web *WebServer) AddDB(db PostgresServer) {
	state := false
	if <-db.Ready {
		state = true
	}
	go func() {
		time.Sleep(time.Second * 1)
		web.Ready <- state
	}()
}

//AddRedis is
func (web *WebServer) AddRedis(redis RedisServer) {
	state := false
	if <-redis.Ready {
		state = true
	}
	go func() {
		time.Sleep(time.Second * 1)
		web.Ready <- state
	}()
}

// Web is
func Web(id string) (web WebServer) {
	web.ID = id
	web.Ready = make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * 1)
		web.Ready <- true
	}()
	return web
}
