
/**
@author: Sushil Verma
@version: 1.0.0
@date: 07/04/2017
@Description: MedLab-Pharma chaincode
**/

package main

import(
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
	"errors"
)

type MedLabPharmaChaincode struct {
}

func main(){
	fmt.Println("Inside MedLabPharmaChaincode main function");
	err := shim.Start(new(MedLabPharmaChaincode))
	if err != nil {
		fmt.Printf("Error starting MedLabPharma chaincode: %s", err)
	}
}

// Init resets all the things
func (t *MedLabPharmaChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error){
	fmt.Printf("Initializing MedLabPharmaChaincode")

	return nil, nil
}

// Invoke isur entry point to invoke a chaincode function
func (t *MedLabPharmaChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *MedLabPharmaChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	return nil, errors.New("Received unknown function query: " + function)
}
