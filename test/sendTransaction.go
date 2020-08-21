package main

import (
    "context"
    // "crypto/ecdsa"
    // "fmt"
    "log"
    "math/big"
    "github.com/rmc-code/rmc-client/common"
    "github.com/rmc-code/rmc-client/common/hexutil"
    "github.com/rmc-code/rmc-client/core/types"
    // "github.com/rmc-code/rmc-client/crypto"
    "github.com/rmc-code/rmc-client/ethclient"
)
func main() {
//Connect node
    client, err := ethclient.Dial("http://localhost:1234")
    if err != nil {
        log.Fatal(err)
	}
	
//Construct fromAddress
fromAddress := common.HexToAddress("RMC30095Bb2A16CC8f4b897F511D2B62Fb8a0c2F0ec")
//Construct toAddress
toAddress := common.HexToAddress("RMC6cBe9DF6DF54281D363e7a5e1790dc66212438C7")

//nonce
nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
//Construct transaction	
tx:=types.SendTxArgs{
	From:       fromAddress,
	To :        toAddress,
	Gas:        hexutil.Uint64(21000),
	GasPrice:   hexutil.Big(*big.NewInt(6000000000)),
	Value:      hexutil.Big(*big.NewInt(1)),
	Nonce:      hexutil.Uint64(nonce),
   
	Data:      hexutil.Bytes([]byte{}),
//Type must is 0
	Type:      0,
}
//Send transaction
	hash,err:=client.SendTransaction(context.Background(),tx)
	if err != nil {
        log.Fatal(err)
	}
	log.Println(hash)
}