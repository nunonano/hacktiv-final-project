package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nunonano/hacktiv-final-project/helper"
	"github.com/nunonano/hacktiv-final-project/model/entity"
	"github.com/stretchr/testify/assert"
)

// type Biodata struct {
// 	Name string `name:json`
// 	Age  int    `age:json`
// }

func UnitTestingExampleGET(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  "successfuly",
		"message": "get request successfuly",
	})
}

func UnitTestingExamplePOST(c *gin.Context) {

	var requestBody entity.Order
	c.ShouldBindJSON(&requestBody)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  "successfuly",
		"message": "post request successfuly",
		"data":    requestBody,
	})
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", UnitTestingExampleGET)
	router.POST("/", UnitTestingExamplePOST)
	return router
}

var router = SetupRouter()

func TestUnitTestingExampleGET(t *testing.T) {

	rr, req := helper.HttpTestRequest("GET", "/", nil)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(rr, req)

	response := helper.Parse(rr.Body.Bytes())

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "successfuly", response.Status)
	assert.Equal(t, "get request successfuly", response.Message)
	assert.Equal(t, nil, response.Data)
}

func TestUnitTestingExamplePOST(t *testing.T) {
	// defining a struct instance
	var order entity.Order
	// raw data in JSON format which
	// is to be decoded
	rawData := `{
        "name": "wisnu"
    }`
	// decoding bidata struct
	// from json format
	err := json.Unmarshal([]byte(rawData), &order)
	if err != nil {
		// if error is not nil
		// print error
		fmt.Println(err)
	}

	// printing details of
	// decoded data
	// Example struct
	// Biodata{
	// 	Name: "wisnu",
	// 	Age:  17,
	// }
	t.Log("Struct is:", order)

	input, _ := json.Marshal(order)

	payload := gin.H{
		"name": order.Name,
	}

	rr, req := helper.HttpTestRequest("POST", "/", helper.Strigify(payload))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(rr, req)

	response := helper.Parse(rr.Body.Bytes())
	b, _ := json.Marshal(response.Data)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "successfuly", response.Status)
	assert.Equal(t, "post request successfuly", response.Message)
	assert.JSONEq(t, string(input), string(b))
}
