package gen

import "time"

// PostgresServer is
type PostgresServer struct {
	// ID is
	ID    string
	Ready chan bool
}

// Postgres is
func Postgres(id string) (db PostgresServer) {
	db.ID = id
	db.Ready = make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * 1)
		db.Ready <- true
	}()
	return db
}
