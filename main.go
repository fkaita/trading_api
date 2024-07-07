package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// sales contract
type contract struct {
	ID       string  `json:"id"`
	Type     string  `json:"type"`
	Date     string  `json:"data"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

// sales contracts
var contracts = []contract{
	{ID: "1", Type: "sales", Date: "2024-07-03", Currency: "USD", Amount: 1000000},
	{ID: "2", Type: "purchase", Date: "2024-07-03", Currency: "USD", Amount: 800000},
}

func main() {
	router := gin.Default()
	router.GET("/contracts", getContracts)
	router.POST("/contracts", postContract)
	router.PUT("/contracts", updateContractByID)
	router.DELETE("/contracts", deleteContractByID)

	router.Run("0.0.0.0:8080")
}

// get contract responds with the list of all contract as JSON.
func getContracts(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		// If no ID is provided, return all contracts.
		c.IndentedJSON(http.StatusOK, contracts)
		return
	}

	// Loop through the list of contract, looking for
	// an album whose ID value matches the parameter.
	for _, crt := range contracts {
		if crt.ID == id {
			c.IndentedJSON(http.StatusOK, crt)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "contract not found"})
}

// post contract adds an album from JSON received in the request body.
func postContract(c *gin.Context) {
	var newContract contract

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newContract); err != nil {
		return
	}

	// Check if the ID already exists
	for _, crt := range contracts {
		if crt.ID == newContract.ID {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Contract with this ID already exists"})
			return
		}
	}

	// Add the new album to the slice.
	contracts = append(contracts, newContract)
	c.IndentedJSON(http.StatusCreated, newContract)
}

// updateContractByID updates a contract whose ID matches the ID in the request.
func updateContractByID(c *gin.Context) {
	var updatedContract contract

	// Call BindJSON to bind the received JSON to updatedContract.
	if err := c.BindJSON(&updatedContract); err != nil {
		return
	}

	// Loop through the list of contracts, looking for
	// a contract whose ID value matches the parameter.
	for i, crt := range contracts {
		if crt.ID == updatedContract.ID {
			// Update the contract details.
			contracts[i] = updatedContract
			c.IndentedJSON(http.StatusOK, updatedContract)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "contract not found"})
}

// deleteContractByID deletes a contract whose ID matches the ID in the request.
func deleteContractByID(c *gin.Context) {
	id := c.Query("id")

	// Loop through the list of contracts, looking for
	// a contract whose ID value matches the parameter.
	for i, crt := range contracts {
		if crt.ID == id {
			// Remove the contract from the slice.
			contracts = append(contracts[:i], contracts[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "contract deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "contract not found"})
}
