package main

import (
    "context"
    "crypto/ecdsa"
    "fmt"
    "log"
    "math/big"
    "github.com/rmc-code/rmc-client/ethereumRMC/common"
    "github.com/rmc-code/rmc-client/ethereumRMC/core/types"
    "github.com/rmc-code/rmc-client/ethereumRMC/crypto"
    "github.com/rmc-code/rmc-client/ethereumRMC/ethclient"
)
func main() {
    //连接节点
    client, err := ethclient.Dial("http://chain-node.galaxynetwork.vip")
    if err != nil {
        log.Fatal(err)
    }

//构造from地址
//RMC30095Bb2A16CC8f4b897F511D2B62Fb8a0c2F0ec
    privateKey, err := crypto.HexToECDSA("b77de610fb69f929f9ce38e07bc003bb8dfffc9024c0af0da26ab2d0a052492e")
    if err != nil {
        log.Fatal(err)
    }
    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }
    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    if err != nil {
        log.Fatal(err)
    }
//设置to地址
toAddress := common.HexToAddress("RMC6cBe9DF6DF54281D363e7a5e1790dc66212438C7")

        
//设置value
    value,_:= new(big.Int).SetString("1",10)
//设置gasPrice
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }
//设置nonce
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
//设置data   
    data:=[]byte("")
//构造交易体
    tx := types.NewTransaction(nonce, toAddress, value,3000000, big.NewInt(gasPrice.Int64()), data)
//查询chainID
    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        log.Fatal(err)
    }
//交易签名
    var signedTx *types.Transaction
    signedTx, err = types.SignTx(tx, types.NewEIP155Signer(big.NewInt(chainID.Int64())), privateKey)
    if err != nil {
        log.Fatal(err)    
    }
//发送签名数据  
    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("tx Hash: %v\n", signedTx.Hash().Hex())
    log.Println("Waiting for the transaction, about 4 minutes...")
//等待交易
	for {
    tx, isPending, err := client.TransactionByHash(context.Background(), signedTx.Hash())
    if err != nil {
        log.Fatal(err)
    }
    if isPending==false{
         fmt.Println("transaction is successful!!")
		 receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
        }
        if receipt.Status==0{
            log.Fatal( "Error: Transaction has been reverted by the EVM")
        }
		fmt.Printf("receipt.Status:%v\n",receipt.Status)
		return 
    }
   }
}