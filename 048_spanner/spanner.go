// gcloud spanner instances create user-db-test --config regional-XXXXX --description test --nodes 3
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"

	"cloud.google.com/go/spanner"
)

var spannerClient *spanner.Client

var spannerProject string

func init() {
	flag.StringVar(&spannerProject, "project", "", "spanner project")
}

func getSpannerClient() *spanner.Client {
	if spannerClient == nil {
		databaseName := "projects/" + spannerProject + "/instances/user-db-test/databases/testdb"
		cfg := spanner.ClientConfig{
			NumChannels: *cun,
		}
		cfg.MinOpened = uint64(*cun)
		client, err := spanner.NewClientWithConfig(context.Background(), databaseName, cfg)
		if err != nil {
			log.Panicf("Failed to create client %v", err)
		}
		// defer client.Close()
		fmt.Println("Connected")
		spannerClient = client
	}
	return spannerClient
}

func spannerInsert(n int) {
	client := getSpannerClient()
	m := spanner.Insert("testtable", []string{"key", "value"}, []interface{}{"key" + strconv.Itoa(n), "value" + strconv.Itoa(n)})
	var ms []*spanner.Mutation
	ms = append(ms, m)
	_, err := client.Apply(context.Background(), ms)
	if err != nil {
		log.Panicf("Query failed with %v", err)
	}
}

func spannerSelect(n int) {
	client := getSpannerClient()
	stmt := spanner.Statement{SQL: "SELECT key, value FROM testtable WHERE key='key" + strconv.Itoa(n) + "'"}
	iter := client.Single().Query(context.Background(), stmt)
	defer iter.Stop()

	row, err := iter.Next()
	if err != nil {
		log.Panicf("Query failed with %v", err)
	}

	var key, value string
	if row.Columns(&key, &value) != nil {
		log.Panicf("Failed to parse row %v", err)
	}
}
