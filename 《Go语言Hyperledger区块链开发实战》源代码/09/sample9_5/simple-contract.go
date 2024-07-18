package main

import (
    "errors"
    "fmt"

    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SimpleContract contract for handling writing and reading from the world state
type SimpleContract struct {
    contractapi.Contract
}

// 在世界状态数据库中添加一个键值对
func (sc *SimpleContract) Create(ctx CustomTransactionContextInterface, key string, value string) error {
    existing := ctx.GetData()

    if existing != nil {
        return fmt.Errorf("Cannot create world state pair with key %s. Already exists", key)
    }

    err := ctx.GetStub().PutState(key, []byte(value))

    if err != nil {
        return errors.New("Unable to interact with world state")
    }

    return nil
}

// 在世界状态数据库中修改指定键的值
func (sc *SimpleContract) Update(ctx CustomTransactionContextInterface, key string, value string) error {
     existing := ctx.GetData()

    if existing == nil {
        return fmt.Errorf("Cannot update world state pair with key %s. Does not exist", key)
    }

    err := ctx.GetStub().PutState(key, []byte(value))

    if err != nil {
        return errors.New("Unable to interact with world state")
    }

    return nil
}

// 从世界状态数据库中根据键读取数据
func (sc *SimpleContract) Read(ctx CustomTransactionContextInterface, key string) (string, error) {
    existing := ctx.GetData()

    if existing == nil {
        return "", fmt.Errorf("Cannot read world state pair with key %s. Does not exist", key)
    }

    return string(existing), nil
}
