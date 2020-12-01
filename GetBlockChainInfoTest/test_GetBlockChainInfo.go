package main

/*
#include "GetBlockChainInfo.h"

*/
import "C"
import (
	"encoding/json"
	"log"

	"github.com/btcsuite/btcd/rpcclient"
)

func Getblockchaininfo() {
	// create new client instance
	client, err := rpcclient.New(&rpcclient.ConnConfig{
		HTTPPostMode: true,
		DisableTLS:   true,
		Host:         "localhost:8332",
		User:         "bitcoind",
		Pass:         "B1tC01nD",
	}, nil)
	if err != nil {
		log.Fatalf("error creating new btc client: %v", err)
	}

	result, err := client.GetBlockChainInfo()
	if err != nil {
		log.Fatalf("error listing accounts: %v", err)
	}
	// iterate over accounts (map[string]btcutil.Amount) and write to stdout
	log.Printf("%s", result)

	b, rror := json.Marshal(result)
	_ = rror
	log.Printf("%s", b)
	//return C.struct_RawGetBlockChainInfoResult{result.Chain, result.Blocks, result.Headers, result.BestBlockHash, result.Difficulty, result.MedianTime, result.VerificationProgress, result.Pruned, result.PruneHeight, result.ChainWork}
}

func Getblockchaininfojson() {
	client, err := rpcclient.New(&rpcclient.ConnConfig{
		HTTPPostMode: true,
		DisableTLS:   true,
		Host:         "localhost:8332",
		User:         "bitcoind",
		Pass:         "B1tC01nD",
	}, nil)
	if err != nil {
		log.Fatalf("error creating new btc client: %v", err)
	}

	r := client.GetBlockChainInfoAsync()

	log.Printf("%s", r.Response.result)

	//res, err := receiveFuture(r.Response)

}

func main() {
	//Getblockchaininfo()

}
