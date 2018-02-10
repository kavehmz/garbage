package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kavehmz/jobber/handover"
	"google.golang.org/grpc"
)

type output struct {
	Out string `json:"out"`
}

type input struct {
	In string `json:"in"`
}

func handler(ctx context.Context, in input) (output, error) {

	worker()

	return output{Out: "Out1234567"}, nil
}

func main() {
	lambda.Start(handler)
}

func worker() {

	// creds, err := credentials.NewClientTLSFromFile("cert/server-cert.pem", "localhost")
	// if err != nil {
	// 	log.Fatalf("cert load error: %s", err)
	// }

	conn, err := grpc.Dial("184.73.151.80:50051", grpc.WithInsecure())
	if err != nil {
		log.Println("Failed to connect: %v", err)
		return
	}
	defer conn.Close()

	c := handover.NewHandoverClient(conn)
	stream, err := c.Join(context.Background())
	if err != nil {
		log.Println("Failed to connect: %v", err)
		return
	}

	log.Println("worker: Joined")
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Println("worker: received error", err)
				return
			}
			log.Println("worker: reveiced a task from server", in.String())

			// This is what worker does. The rest is the template how to write and strean in grpc
			time.Sleep(time.Millisecond * 50)

			if err = stream.Send(&handover.Result{Data: time.Now().String()}); err != nil {
				log.Println("startWorker: Failed to send, ", err)
				return
			}
		}
	}()
	<-time.After(time.Second * 5)
	stream.CloseSend()
	time.Sleep(time.Second)
}
