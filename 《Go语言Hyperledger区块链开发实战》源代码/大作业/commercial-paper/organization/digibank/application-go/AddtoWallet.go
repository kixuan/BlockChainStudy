package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

var identityLabel string = "User1@org1.example.com-cert.pem"

func AddtoWallet() (int, error) {
	log.Println("============ AddtoWallet.go启动 ============")

	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	if err != nil {
		log.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environemnt variable: %v", err)
		return -1, err
	}
	//创建存储身份集合的钱包
	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		log.Fatalf("创建钱包: %v", err)
		return -2, err
	} else {
		log.Println("创建钱包成功！")
	}

	if !wallet.Exists(identityLabel) {
		err = populateWallet(wallet)
		if err != nil {
			log.Fatalf("构造钱包内容失败: %v", err)
			return -3, err
		} else {
			log.Println("构造钱包内容成功！")
		}
	}

	return 0, nil
}

func populateWallet(wallet *gateway.Wallet) error {
	log.Println("============ Populating wallet ============")
	credPath := filepath.Join(
		"..",
		"..",
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"users",
		"User1@org1.example.com",
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", identityLabel)
	log.Println("证书路径为: ", certPath)
	// 读取证书
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}
	// 存储私钥文件的目录
	keyDir := filepath.Join(credPath, "keystore")

	keyPath := filepath.Join(keyDir, "priv_sk")          //私钥文件的路径
	key, err := ioutil.ReadFile(filepath.Clean(keyPath)) //读取私钥的内容
	if err != nil {
		return err
	}
	// 创建身份标识
	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))

	err = wallet.Put(identityLabel, identity)
	return err
}
