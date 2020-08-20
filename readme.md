## rmc-client

### 连接节点

```shell
client, err := ethclient.Dial("http://chain-node.galaxynetwork.vip")
```
### 发送交易（离线签名模式）

#### 使用私钥构造from地址
```go
//0xf5403E4F120901407eF221E2419583D1F3556953
    privateKey, err := crypto.HexToECDSA("0xb77de610fb69f929f9ce38e07bc003bb8dfffc9024c0af0da26ab2d0a052492e")
    if err != nil {
        log.Fatal(err)
    }
    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }
    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
```
#### 构造to地址
```go
   toAddress := common.HexToAddress("RMC6cBe9DF6DF54281D363e7a5e1790dc66212438C7")
```

#### 构造交易体
```go
    tx := types.NewTransaction(nonce, toAddress, value,gas, gasPrice, data)
```
#### 查询chainID
```go
   chainID, err := client.NetworkID(context.Background())
```
#### 交易签名
```go
    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(chainID.Int64())), privateKey)
```
#### 发送签名数据
```go
    err = client.SendTransaction(context.Background(), signedTx)
```
### 查询交易
```go
	transaction,isPending,err:=client.TransactionByHash(context.Background(),TransactionHash)
	if err != nil {
        log.Fatal(err)
    }
```
### 查询余额
```go
	Address := common.HexToAddress("RMC86056D210eA7Bc23337aCaBE96dE275E584a67ce")
	balance,err:=client.BalanceAt(context.Background(),Address,nil)
	if err != nil {
        log.Fatal(err)
    }
```
### 查询区块
```go
	blockHash := common.HexToHash("0x18c8c36ac3c285d7b276e59b1988d0632aec58ee7f70faa17cfe74de0c5484b5")
    block,err=client.BlockByHash(context.Background(),blockHash)
    //or ByNumber
	// block,err=client.BlockByNumber(context.Background(),big.NewInt(25591))
```
