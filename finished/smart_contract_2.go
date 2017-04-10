//a program to insert and update a JSON and validate with smart Contract.
/*
The following Code is used for insertion and parsing of json with blockchain.
The code inserts the smart contract schema in form of json in the blockchain and provide
functionality to pull the json edit the json and update the json in blockchain
this appending the block in the blockchain.
*/

package main

import (
	"errors"
	"fmt"
	//	"encoding/json"
	//	"os"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	//  "strconv"
)

//SimpleChaincode example simple Chaincode implementation

type SimpleChaincode struct {
}

//inputjson
/*func main(){
//data := `{"product_id":"IOT1124s","Contractid":"232241123","Parents":["Gomez","Morticia"]}`
//Contract
//data := `{"product_id":"IOT1124s","Contractid":"232241123","stake_holders":["Saurabh_id123","Vinit_Ajay123"],"sensor_value":"24","payment_percent":"20"}`
b := []byte(data)
}*/

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	err := stub.PutState("contract_json", []byte(args[0]))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Invoke isur entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "write" {
		return t.write(stub, args)
	} else if function == "putcontract" {
		return t.putcontract(stub, args)
	} else if function == "validate" {
		return t.validate(stub, args)
	}

	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" { //read a variable
		return t.read(stub, args)
	} else if function == "getcontract" {
		return t.getcontract(stub, args)
	}

	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}

// putcontract - Put the received bytearray smatcontract in the json
func (t *SimpleChaincode) putcontract(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
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

// write - invoke function to write key/value pair
func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
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

// getcontract - Get the smart Contract from the blockchain as bytearray
func (t *SimpleChaincode) getcontract(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
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

//  - query function to read key/value pair
func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
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

// read - query function to read key/value pair
func (t *SimpleChaincode) validate(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	//var key, jsonResp string
	//var err, err_state, err_contract error

	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	//parameters productid,contractkey,bytearray
	//data := `{"product_id":"IOT1124s","Contractid":"232241123","stake_holders":["Saurabh_id123","Vinit_Ajay123"],"sensor_value":"24","payment_percent":"20"}`

	/*StateJsonAsbytes := []byte(args[0])
		contractkey :=  args[1]
		//productid := args[2]


		ContractvalAsbytes, err := stub.GetState(contractkey)
		if err != nil {
			jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"

			return nil, errors.New(jsonResp)
		}

	    var f interface{}     //Interface for marshalling the data received from blockchain contract used for comparison.
	    var g interface{}      //Interface for receiving and marshalling the received data

			err_contract = json.Unmarshal(ContractvalAsbytes, &f)
			if (err_contract!=nil) {
				os.Exit(1)
			}

			err_state = json.Unmarshal(StateJsonAsbytes, &g)
			if (err_state!=nil) {
				os.Exit(1)
			}

			contract_json := f.(map[string]interface{})

		    state_json := g.(map[string]interface{})

		    // The Key value iteration can be done better for dynamicity as a seperate function. to loop over the two structs.

	        var sensor_value,sensor_contract string

			for k, v := range contract_json {
	    	   if k == "sensor_value" {

	                    fmt.Println(k, "is to be compared", v)
	                    sensor_value=v.(string)

	                 }

			}

	        for k, v := range state_json {
	    	   if k == "sensor_value" {

	                    fmt.Println(k, "is to be compared", v)
	                    sensor_contract=v.(string)
	                 }

			}


	        val1,_ := strconv.Atoi(sensor_value)
	        val2,_ := strconv.Atoi(sensor_contract)
	*/
	var key_temp string
	var exception string

	key_temp = "success"

	if 5 < 4 {
		exception = "{\"Error\":\"Failed to get state for " + key_temp + "\"}"
	} else {
		exception = "{\"Error\":\"Failed to get state for " + key_temp + "\"}"
	}

	exceptionAsBytes := []byte(exception)

	/*Section to validate the two jsons and put state only if data is validated*/

	//Smart Contract Rules :

	// case : blockchain.sensor_value==received.sensor_value

	// case : blockchain.expiry_max== received.expiry

	//if true : insert in to blockchain.

	return exceptionAsBytes, errors.New(exception)
}
