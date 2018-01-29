package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	pb "github.com/kavehmz/garbage/050_grpc/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var host *string
var max *int
var cun *int

func init() {
	host = flag.String("host", "localhost", "")
	max = flag.Int("max", 5000, "")
	cun = flag.Int("cun", 10, "")
}

func benchGRPC() {
	conn, err := grpc.Dial(*host+":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	buf := make(chan bool, *cun)
	for i := 0; i < *max; i++ {
		buf <- true
		go func() {
			_, err := c.SayHello(context.Background(), &pb.Request{Name: "name"})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			<-buf
		}()
	}

	for i := 0; i < *cun; i++ {
		buf <- true
	}
}

func benchHTTP() {

	buf := make(chan bool, *cun)
	for i := 0; i < *max; i++ {
		buf <- true
		go func() {
			resp, err := http.Get("http://" + *host + ":8080/test/test")
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil || string(body) == "" {
				log.Fatalf("could not greet: %v", err)
			}
			<-buf
		}()
	}

	for i := 0; i < *cun; i++ {
		buf <- true
	}

}

func main() {
	flag.Parse()
	s := time.Now()
	benchGRPC()
	fmt.Println("GRPC", time.Since(s))
	s = time.Now()
	benchHTTP()
	fmt.Println("HTTP", time.Since(s))
}
