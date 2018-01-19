package main

import (
	"flag"
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

	err = client.Put(nil, key, bins)
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

	_, err = client.Get(nil, key)
	if err != nil {
		log.Panic(err)
	}
}
