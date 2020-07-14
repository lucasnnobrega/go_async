// hello.go

package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Application struct
type Application struct {
	App      string
	Company  string
	Category string
}

func main() {
	// read data from file
	jsonDataFromFile, err := ioutil.ReadFile("company.json")

	if err != nil {
		panic(err)
	}

	done := make(chan bool)

	go myFunc(jsonDataFromFile, done)

	fmt.Println("dentro da main, antes do channel")
	fmt.Println("dentro da main, depois do channel")

	// Create a go routine with a lambda function, and pass a parameter in the end

	for i := 0; i < 200; i++ {
		fmt.Printf("main-func: %d \n", i)
	}

	x := <-done // receive from done
	fmt.Println(x)
}

func myFunc(jsonDataFromFile []byte, done chan bool) {

	fmt.Println("Unmarshal JSON data")
	var jsonData []Application
	err := json.Unmarshal([]byte(jsonDataFromFile), &jsonData)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Create file")
	csvFile, err := os.Create("./data.csv")

	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	fmt.Println("Write on the file")
	writer := csv.NewWriter(csvFile)

	for index, usance := range jsonData {
		fmt.Println("#", index, "#")
		var row []string
		fmt.Println(row)
		row = append(row, usance.App)
		fmt.Println(row)
		row = append(row, usance.Company)
		fmt.Println(row)
		row = append(row, usance.Category)
		fmt.Println(row)
		writer.Write(row)
	}

	fmt.Println("remember to flush to the file!")
	writer.Flush()

	fmt.Println("dentro da funcao, antes do channel")
	done <- true // send sum to c
	fmt.Println("dentro da funcao, depois do channel")

}
