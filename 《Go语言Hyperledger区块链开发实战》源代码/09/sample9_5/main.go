package main

import (
    "errors"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)
//交易前处理函数
func GetWorldState(ctx CustomTransactionContextInterface) error {
	_, params := ctx.GetStub().GetFunctionAndParameters()

	if len(params) < 1 {
		return errors.New("Missing key for world state")
	}

	existing, err := ctx.GetStub().GetState(params[0])

	if err != nil {
		return errors.New("Unable to interact with world state")
	}

	ctx.SetData(existing)

	return nil
}

func main() {
   simpleContract := new(SimpleContract)
  simpleContract.BeforeTransaction = GetWorldState
   simpleContract.TransactionContextHandler = new(CustomTransactionContext)
    cc, err := contractapi.NewChaincode(simpleContract)

    if err != nil {
        panic(err.Error())
    }
   if err := cc.Start(); err != nil {
        panic(err.Error())
    }
}
