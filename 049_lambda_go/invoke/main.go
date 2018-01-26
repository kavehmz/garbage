package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type output struct {
	Out string `json:"out"`
}

type input struct {
	In string `json:"out"`
}

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	client := lambda.New(sess, &aws.Config{Region: aws.String("us-east-1")})

	request := input{}
	payload, err := json.Marshal(request)
	checkErr(err)

	result, err := client.Invoke(&lambda.InvokeInput{FunctionName: aws.String("MyTest"), Payload: payload})
	checkErr(err)

	var o output
	err = json.Unmarshal(result.Payload, &o)

	fmt.Println(err, o.Out)

}

func checkErr(e error) {
	if e != nil {
		log.Panic(e)
	}
}
