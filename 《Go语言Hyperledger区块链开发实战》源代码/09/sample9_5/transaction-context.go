package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)
//定义接口CustomTransactionContextInterface，指定自定义交易上下文对象的结构体中必须实现的函数
type CustomTransactionContextInterface interface {
	contractapi.TransactionContextInterface
	GetData() []byte
	SetData([]byte)
}


// 自定义交易上下文对象
type CustomTransactionContext struct {
	contractapi.TransactionContext
	data []byte
}

// GetData return set data
func (ctc *CustomTransactionContext) GetData() []byte {
	return ctc.data
}

// SetData provide a value for data
func (ctc *CustomTransactionContext) SetData(data []byte) {
	ctc.data = data
}

