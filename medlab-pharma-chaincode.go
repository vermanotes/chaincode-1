
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

	// Handle different functions
	if function == "init" {
		return t.init(stub, args)
	} else if function == "write" {
		return t.write(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)
	
	return nil, errors.New("Received unknown function invocation: " + function)
}

// Invoke isur entry point to invoke a chaincode function
func (t *MedLabPharmaChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *MedLabPharmaChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)
	
	// Handle different functions
	if function == "read" { //read a variable
		return t.read(stub, args)
	}
	fmt.Println("query did not find func: " + function)
	return nil, errors.New("Received unknown function query: " + function)
}


func (t *MedLabPharmaChaincode) init(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	err := stub.PutState("Container", []byte(args[0]))
	if err != nil {
		return nil, err
	}

	return nil, nil
}


// write - invoke function to write key/value pair
func (t *MedLabPharmaChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, value string
	var err error
	fmt.Println("running write()")

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}

	key = args[0] //rename for funsies
	value = args[1]
	err = stub.PutState(key, []byte(value)) //write the variable into the chaincode state
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// read - query function to read key/value pair
func (t *MedLabPharmaChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil
}
