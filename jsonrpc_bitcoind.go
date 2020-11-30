package main

/*

struct RawGetBlockChainInfoResult {
	string
}

*/
import (
	"C"
	//~ "github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
	//~ "github.com/btcsuite/btcutil"
	"log"
)
import (
	"fmt"
	"time"
)

func GetBlockChainInfo_T() {
	// create new client instance
	start := time.Now()

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

	fmt.Println(time.Since(start))
	start = time.Now()

	result, err := client.GetBlockChainInfo()

	fmt.Println(time.Since(start))
	start = time.Now()

	result2, err := client.GetBlockChainInfo()
	_ = result2
	fmt.Println(time.Since(start))
	start = time.Now()

	result3, err := client.GetBlockChainInfo()
	_ = result3
	fmt.Println(time.Since(start))
	start = time.Now()
	// iterate over accounts (map[string]btcutil.Amount) and write to stdout
	log.Printf("%s", result)
}

func main() {
	GetBlockChainInfo_T()
	log.Printf("1")
	GetBlockChainInfo_T()
	log.Printf("2")
	GetBlockChainInfo_T()
	log.Printf("3")

	// list accounts
	//~ accounts, err := client.ListAccounts()
	//~ if err != nil {
	//~ log.Fatalf("error listing accounts: %v", err)
	//~ }
	//~ // iterate over accounts (map[string]btcutil.Amount) and write to stdout
	//~ for label, amount := range accounts {
	//~ log.Printf("%s: %s", label, amount)
	//~ }

	//~ // prepare a sendMany transaction
	//~ receiver1, err := btcutil.DecodeAddress("1someAddressThatIsActuallyReal", &chaincfg.MainNetParams)
	//~ if err != nil {
	//~ log.Fatalf("address receiver1 seems to be invalid: %v", err)
	//~ }
	//~ receiver2, err := btcutil.DecodeAddress("1anotherAddressThatsPrettyReal", &chaincfg.MainNetParams)
	//~ if err != nil {
	//~ log.Fatalf("address receiver2 seems to be invalid: %v", err)
	//~ }
	//~ receivers := map[btcutil.Address]btcutil.Amount{
	//~ receiver1: 42,  // 42 satoshi
	//~ receiver2: 100, // 100 satoshi
	//~ }

	//~ // create and send the sendMany tx
	//~ txSha, err := client.SendMany("some-account-label-from-which-to-send", receivers)
	//~ if err != nil {
	//~ log.Fatalf("error sendMany: %v", err)
	//~ }
	//~ log.Printf("sendMany completed! tx sha is: %s", txSha.String())
}
