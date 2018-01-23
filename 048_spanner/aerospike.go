package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	aerospike "github.com/aerospike/aerospike-client-go"
)

var aerospikeClient *aerospike.Client

var aerospikeAddr string

func init() {
	flag.StringVar(&aerospikeAddr, "aerospike", "", "aerospike client ip")
}

func getAerospikeClient() *aerospike.Client {
	if aerospikeClient == nil {
		clientPolicy := aerospike.NewClientPolicy()
		clientPolicy.ConnectionQueueSize = 1500
		var err error
		aerospikeClient, err = aerospike.NewClientWithPolicy(clientPolicy, aerospikeAddr, 3000)
		if err != nil {
			log.Panic(err)
		}

	}
	return aerospikeClient
}

func aerospikeInsert(n int) {
	client := getAerospikeClient()

	key, err := aerospike.NewKey("test", "aerospike", "key"+strconv.Itoa(n))
	if err != nil {
		log.Panic(err)
	}

	bins := aerospike.BinMap{
		"bin1": n,
		"bin2": "An elephant is a mouse with an operating system",
		"bin3": []interface{}{"Go", 2009},
	}
	policy := new(aerospike.WritePolicy)
	policy.CommitLevel = aerospike.COMMIT_ALL

	err = client.Put(policy, key, bins)
	if err != nil {
		log.Panic(err)
	}
}

func aerospikeSelect(n int) {
	client := getAerospikeClient()

	key, err := aerospike.NewKey("test", "aerospike", "key"+strconv.Itoa(n))
	if err != nil {
		log.Panic(err)
	}
	policy := new(aerospike.BasePolicy)
	policy.ConsistencyLevel = aerospike.CONSISTENCY_ALL

	rec, err := client.Get(policy, key)
	if err != nil {
		log.Panic(err)
	}
	if rec == nil {
		fmt.Printf("record was nil", n)
	}
}
