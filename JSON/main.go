package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	FullName string `json:"Name"`
	Age int
}

func main() {
	var jsonString = `{"Name": "Muhammad Irvan", "Age": 17}`
	var jsonData = []byte(jsonString)
	
	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	var data2 interface{}
	json.Unmarshal(jsonData, &data2)
	var decodeData = data2.(map[string]interface{}) 

	var jsonString1 = `[
		{"Name" : "Muhammad irvan", "Age": 17},
		{"Name" : "Irene Lutfia Puspita", "Age": 16}
	]`

	var data user
	var newData []user

	var newErr= json.Unmarshal([]byte(jsonString1), &newData)
	if newErr != nil {
		fmt.Println(newErr.Error())
		return
	}

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var object = []user{{"Muhammad Irvan", 17}, {"Irene Lutfia Puspita", 16}}
	var jsonData1, errr = json.Marshal(object)
	if errr != nil {
		fmt.Println(errr.Error())
		return
	}

	var jsonString2 = string(jsonData1)
	fmt.Println(jsonString2)

	fmt.Println("user :", data.FullName)
	fmt.Println("age  :", data.Age)
	
	fmt.Println("user :", data1["Name"])
	fmt.Println("age  :", data1["Age"])

	fmt.Println("user :", decodeData["Name"])
	fmt.Println("age  :", decodeData["Age"])

	fmt.Println("User 1:", newData[0].FullName)
	fmt.Println("User 2:", newData[1].FullName)
}