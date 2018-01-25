package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type output struct {
	Out string `json:"out"`
}

type input struct {
	In string `json:"out"`
}

func handler(ctx context.Context, in input) (output, error) {
	return output{Out: "Out1234"}, nil
}

func main() {
	lambda.Start(handler)
}
