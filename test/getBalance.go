package main

import (
    "context"
    "log"
    "github.com/rmc-code/rmc-client/common"
    "github.com/rmc-code/rmc-client/ethclient"
)
func main() {
    client, err := ethclient.Dial("http://192.168.1.127:8545")
    if err != nil {
        log.Fatal(err)
    }

	Address := common.HexToAddress("RMC86056D210eA7Bc23337aCaBE96dE275E584a67ce")

	balance,err:=client.BalanceAt(context.Background(),Address,nil)
	if err != nil {
        log.Fatal(err)
	}
    log.Println("balance===>",balance)

}