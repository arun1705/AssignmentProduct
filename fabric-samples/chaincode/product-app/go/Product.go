package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"

	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
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

// AddProduct adds a new product to the world state with given details
func (s *SmartContract) AddProduct(ctx contractapi.TransactionContextInterface,id string, name string, description string, prize int, colour string) error {
	val, ok, err := cid.GetAttributeValue(ctx, "role")
	if err != nil {
		// There was an error trying to retrieve the attribute
		return fmt.Errorf("Error while retriving attributes")
	}
	if !ok {
		// The client identity does not possess the attribute
		return fmt.Errorf("Client identity doesnot posses the attribute")
	}
	
	if val != "approver" {
		fmt.Println("Attribute role: " + val)
		return fmt.Errorf("Only user with role as APPROVER have access this method!")
	} else {
	
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
}

// This function returns all the existing products 
func (s *SmartContract) QueryAllProduct(ctx contractapi.TransactionContextInterface) ([]*Product, error) {
	val, ok, err := cid.GetAttributeValue(ctx, "role")
	if err != nil {
		// There was an error trying to retrieve the attribute
		return fmt.Errorf("Error while retriving attributes")
	}
	if !ok {
		// The client identity does not possess the attribute
		return fmt.Errorf("Client identity doesnot posses the attribute")
	}
	
	if val != "approver" {
		fmt.Println("Attribute role: " + val)
		return fmt.Errorf("Only user with role as APPROVER have access this method!")
	} else {

	productIterator, err := ctx.GetStub().GetStateByRange("", "")
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
}


// This function helps to query the product by Id
func (s *SmartContract) QueryProductById(ctx contractapi.TransactionContextInterface, id string) (*Product, error) {
	val, ok, err := cid.GetAttributeValue(ctx, "role")
	if err != nil {
		// There was an error trying to retrieve the attribute
		return fmt.Errorf("Error while retriving attributes")
	}
	if !ok {
		// The client identity does not possess the attribute
		return fmt.Errorf("Client identity doesnot posses the attribute")
	}
	
	if val != "approver" {
		fmt.Println("Attribute role: " + val)
		return fmt.Errorf("Only user with role as APPROVER have access this method!")
	} else {

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
}


func (s *SimpleChaincode) DeleteProductById(ctx contractapi.TransactionContextInterface, id string) (*Product, error) {
	val, ok, err := cid.GetAttributeValue(ctx, "role")
	if err != nil {
		// There was an error trying to retrieve the attribute
		return fmt.Errorf("Error while retriving attributes")
	}
	if !ok {
		// The client identity does not possess the attribute
		return fmt.Errorf("Client identity doesnot posses the attribute")
	}
	
	if val != "approver" {
		fmt.Println("Attribute role: " + val)
		return fmt.Errorf("Only user with role as APPROVER have access this method!")
	} else {

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

	return ctx.GetStub().DeleteState(id)

    }
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
