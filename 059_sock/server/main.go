package main

import (
	"log"
	"net"
)

func echoServer(c net.Conn) {
	defer log.Println("done")
	for {
		log.Println("Got message")
		buf := make([]byte, 512)
		nr, err := c.Read(buf)
		if err != nil {
			log.Println("bye", err)
			return
		}

		data := buf[0:nr]
		println("Server got:", string(data))
		_, err = c.Write(data)
		if err != nil {
			log.Fatal("Write: ", err)
		}
	}
}

func main() {
	l, err := net.Listen("unix", "/tmp/echo.sock")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	for {
		fd, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		go echoServer(fd)
	}
}
