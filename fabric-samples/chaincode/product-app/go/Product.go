package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing a car
type SmartContract struct {
	contractapi.Contract
}

// Mobile describes basic details of what makes up a mobile
type Product struct { 
	Id     string `json:"id"`
	Name   string `json:"name"`
	Description  string `json:"description"`
	Prize int `json:"prize"`
	Colour string `json:"colour"`
}

// QueryResult structure used for handling result of query
type QueryResult struct {
	Key    string `json:"Key"`
	Record *Product
}

// InitLedger adds a base set of cars to the ledger
// func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
// 	cars := []Car{
// 		Car{Make: "Toyota", Model: "Prius", Colour: "blue", Owner: "Tomoko"},
// 		Car{Make: "Ford", Model: "Mustang", Colour: "red", Owner: "Brad"},
// 		Car{Make: "Hyundai", Model: "Tucson", Colour: "green", Owner: "Jin Soo"},
// 		Car{Make: "Volkswagen", Model: "Passat", Colour: "yellow", Owner: "Max"},
// 		Car{Make: "Tesla", Model: "S", Colour: "black", Owner: "Adriana"},
// 		Car{Make: "Peugeot", Model: "205", Colour: "purple", Owner: "Michel"},
// 		Car{Make: "Chery", Model: "S22L", Colour: "white", Owner: "Aarav"},
// 		Car{Make: "Fiat", Model: "Punto", Colour: "violet", Owner: "Pari"},
// 		Car{Make: "Tata", Model: "Nano", Colour: "indigo", Owner: "Valeria"},
// 		Car{Make: "Holden", Model: "Barina", Colour: "brown", Owner: "Shotaro"},
// 	}

// 	for i, car := range cars {
// 		carAsBytes, _ := json.Marshal(car)
// 		err := ctx.GetStub().PutState("CAR"+strconv.Itoa(i), carAsBytes)

// 		if err != nil {
// 			return fmt.Errorf("Failed to put to world state. %s", err.Error())
// 		}
// 	}

// 	return nil
// }

// AddProduct adds a new product to the world state with given details
func (s *SmartContract) AddProduct(ctx contractapi.TransactionContextInterface,id string, name string, description string, prize int, colour string) error {
	productJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
        return fmt.Errorf("Failed to read the data from world state", err)
    }

	if productJSON != nil {
		return fmt.Errorf("the product %s already exists", id)
    }

	product := Product{
		Id:  id,
		Name:   name,
		Description:  description,
		Prize: prize,
		Colour: colour,
	}

	productAsBytes, err := json.Marshal(product)	
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, productAsBytes)
}

// This function returns all the existing products 
func (s *SmartContract) QueryAllProduct(ctx contractapi.TransactionContextInterface,startKey string,endKey string) ([]*Product, error) {
	productIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, err
	}
	
	defer productIterator.Close()

	var products []*Product
	for productIterator.HasNext() {
		productResponse, err := productIterator.Next()
		if err != nil {
			return nil, err
		}

		var product *Product
		err = json.Unmarshal(productResponse.Value, &products)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}


// This function helps to query the product by Id
func (s *SmartContract) QueryProductById(ctx contractapi.TransactionContextInterface, id string) (*Product, error) {
    productJSON, err := ctx.GetStub().GetState(id)
    if err != nil {
        return nil, fmt.Errorf("Failed to read the data from world state", err)
    }
	
    if productJSON == nil {
		return nil, fmt.Errorf("the product %s does not exist", id)
    }
	
	var product *Product
	err = json.Unmarshal(productJSON, &product)
	
	if err != nil {
		return nil, err
	}
	return product, nil
}


func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create fabcar chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting fabcar chaincode: %s", err.Error())
	}
}
