package main

import (
        "tokenmarket/chaincode"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
        err := shim.Start(new(chaincode.NekoTokenCC))
	if err != nil {
	        fmt.Printf("Error starting NekoTokenCC: %s", err)
	}
}