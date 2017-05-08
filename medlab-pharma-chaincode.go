/**
@author: Sushil Verma
@version: 1.0.0
@date: 08-05-2017
**/

package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type MedLabPharmaChaincode struct {
}

func main() {
	fmt.Println("***** Inside MedLabPharmaChaincode main function")
	err := shim.Start(new(MedLabPharmaChaincode))
	if err != nil {
		fmt.Printf("Error starting MedLabPharma chaincode: %s", err)
	}
}

// Init resets all the things
func (t *MedLabPharmaChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("***** Initializing MedLabPharmaChaincode-->arshad ")

	// Handle different functions
	if function == "init" {
		return t.init(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Invoke isur entry point to invoke a chaincode function
func (t *MedLabPharmaChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("***** Invoke is running " + function)

	// Handle different functions
	if function == "TestInvokeFunction"{
		return t.TestInvokeFunction(stub, args[0])
	}else if function == "GetCertValues" {
		userType := t.GetCertValues(stub)
		fmt.Println("***** userType " + userType)
		return nil, nil
	}

	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *MedLabPharmaChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("***** Query is running " + function)
	fmt.Println("query did not find func: " + function)
	return nil, errors.New("Received unknown function query: " + function)
}

func (t *MedLabPharmaChaincode) init(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("***** Inside init() func...")
	return nil, nil
}

func (t *MedLabPharmaChaincode) TestInvokeFunction(stub shim.ChaincodeStubInterface, test_message string) ([]byte, error) {
	fmt.Println("***** Inside TestInvokeFunction() func...")
	fmt.Println("***** Hello " + test_message)
	fmt.Println("TestInvokeFunction success")
	return nil, nil
}

func (t *MedLabPharmaChaincode) GetCertValues(stub shim.ChaincodeStubInterface) (string) {
	username, err := stub.ReadCertAttribute("username")
	fmt.Println(username)
	fmt.Println(err)
	
	if err != nil {
		return "Couldn't get attribute 'username'. Error: ";
	}
	return string(username)
}

