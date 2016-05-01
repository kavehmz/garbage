package gen

import "time"

// ProxyServer is
type ProxyServer struct {
	// ID is
	ID    string
	Ready chan bool
}

//AddWeb is
func (proxy *ProxyServer) Add(web WebServer) {
	state := false
	if <-web.Ready {
		state = true
	}
	go func() {
		time.Sleep(time.Second * 1)
		web.Ready <- state
	}()
}

// Proxy is
func Proxy(id string) (Proxy ProxyServer) {
	Proxy.ID = id
	Proxy.Ready = make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * 1)
		Proxy.Ready <- true
	}()
	return Proxy
}
