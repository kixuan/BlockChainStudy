package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

func redeem(issuer string, paperNumber string, redeemingOwner string, redeenDateTime string) (int, error) {
	log.Println("============ 启动 Redeem ============")
	const ENVKEY string = "CONTRACT_NAME"
	var contractName string = "papercontract"
	var userName string = "User1@org1.example.com-cert.pem"
	//获取环境变量CONTRACT_NAME的值，如果有，则作为智能合约的名字；否则以"papercontract"作为智能合约的名字
	name := os.Getenv(ENVKEY)
	if name != "" {
		contractName = name
	}
	walletPath := filepath.Join(
		".",
		"wallet",
	)
	wallet, err := gateway.NewFileSystemWallet(walletPath)
	if err != nil {
		log.Fatalf("创建钱包失败: %v", err)
		return -1, err
	}

	connectionProfile := filepath.Join(
		"..",
		"gateway",
		"connection-org1.yaml",
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(connectionProfile))),
		gateway.WithIdentity(wallet, userName),
	)
	if err != nil {
		log.Fatalf("连接到网关失败: %v", err)
		return -2, err
	}
	log.Println(string("成功连接到网关"))
	defer gw.Close()

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		log.Fatalf("F获取通道的Network对象失败: %v", err)
		return -3, err
	}
	log.Println(string("成功获取通道的Network对象"))
	contract := network.GetContract(contractName)
	log.Println(string("成功获取智能合约对象"))
	log.Println("--> 提交交易: redeem()函数, 兑换商业票据")
	result, err := contract.SubmitTransaction("redeem", issuer, paperNumber, redeemingOwner, redeenDateTime)
	if err != nil {
		log.Fatalf("提交交易失败: %v", err)
		return -4, err
	}
	log.Println(string(result))
	return 0, nil
}
