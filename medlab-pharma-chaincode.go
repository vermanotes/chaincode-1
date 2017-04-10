
/**
@author: Arshad Sarfarz
@version: 1.0.0
@date: 10/04/2017
@Description: MedLab-Pharma chaincode v1
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
	}
	fmt.Println("invoke did not find func: " + function)
	
	return nil, errors.New("Received unknown function invocation: " + function)
}

// Invoke isur entry point to invoke a chaincode function
func (t *MedLabPharmaChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "shipContainerUsingLogistics" {
		return t.ShipContainerUsingLogistics(stub, args[0], args[1])
	}
	fmt.Println("invoke did not find func: " + function)
	
	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *MedLabPharmaChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)
	
	// Handle different functions
	if function == "getcontainerdetails" { //read a variable
		return t.GetContainerDetails(stub, args)
	}
	fmt.Println("query did not find func: " + function)
	return nil, errors.New("Received unknown function query: " + function)
}


func (t *MedLabPharmaChaincode) init(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	err := stub.PutState("Settings", []byte(args[0]))
	if err != nil {
		return nil, err
	}

	return nil, nil
}


// write - invoke function to write key/value pair
func (t *MedLabPharmaChaincode) ShipContainerUsingLogistics(stub shim.ChaincodeStubInterface, container_id string, elements_json string) ([]byte, error) {
	var key, value string
	var err error
	fmt.Println("running write()")

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}

	key = container_id
	value = elements_json
	err = stub.PutState(key, []byte(value)) //write the variable into the chaincode state
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// read - query function to read key/value pair
func (t *MedLabPharmaChaincode) GetContainerDetails(stub shim.ChaincodeStubInterface, container_id []string) ([]byte, error) {
	var key, jsonResp string
	var err error

	if len(container_id) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = container_id[0]
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil
}
