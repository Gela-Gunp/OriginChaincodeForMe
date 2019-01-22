package tokenmarket

import (
        "github.com/hyperledger/fabric/core/chaincode/shim"
)

type Status struct {
        Id      string `json:"id"`
	Name    string `json:"name"`
	TokenQuan int  `json:"TokenQuan"`
	MoneyQuan int  `json:"MoneyQuan"`
}

type NekoTokenCC interface {
        check(stub shim.ChaincodeStubInterface, id string) error
	buy(stub shim.ChaincodeStubInterface, id1 string, id2 string, token string) error
	sell(stub shim.ChaincodeStubInterface, id1 string, id2 string, token string) error
	charge(stub shim.ChaincodeStubInterface, id string, money string) error
}