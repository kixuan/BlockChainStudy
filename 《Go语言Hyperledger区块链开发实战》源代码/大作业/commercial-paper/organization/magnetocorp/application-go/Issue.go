package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

func issue(issuer string, paperNumber string, issueDateTime string, maturityDateTime string, faceValue string) (int, error) {
	log.Println("============ 启动 Issue ============")
	var contractName string = "papercontract"
	var userName string = "User1@org2.example.com"
	name := os.Getenv("CONTRACT_NAME")
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
		"connection-org2.yaml",
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(connectionProfile))),
		gateway.WithIdentity(wallet, userName),
	)
	if err != nil {
		log.Fatalf("连接到网关失败: %v", err)
		return -2,err
	}
	log.Println(string("成功连接到网关"))
	defer gw.Close()

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		log.Fatalf("F获取通道的Network对象失败: %v", err)
		return -3,err
	}
	log.Println(string("成功获取通道的Network对象"))
	contract := network.GetContract(contractName)
	log.Println(string("成功获取智能合约对象"))
	log.Println("--> 提交交易: issue()函数, 发行商业票据")	
	log.Println("issuer:"+issuer)	
	log.Println("paperNumber:"+paperNumber)	
	log.Println("issueDateTime:"+issueDateTime)	
	log.Println("maturityDateTime:"+maturityDateTime)	
	log.Println("faceValue:"+faceValue)		
	result, err := contract.SubmitTransaction("issue", issuer, paperNumber, issueDateTime, maturityDateTime, faceValue)
	if err != nil {
		log.Fatalf("提交交易失败: %v", err)
		return -4, err
	}
	log.Println(string(result))
	return 0, nil
}
