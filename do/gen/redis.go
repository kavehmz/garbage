package gen

import "time"

// RedisServer is
type RedisServer struct {
	// ID is
	ID    string
	Ready chan bool
}

// Redis is
func Redis(id string) (redis RedisServer) {
	redis.ID = id
	redis.Ready = make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * 1)
		redis.Ready <- true
	}()
	return redis
}
