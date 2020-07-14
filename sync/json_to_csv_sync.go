// hello.go

package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

// Application struct
type Application struct {
	App      string
	Company  string
	Category string
}

// Pior caso com tudo com ponteiros
// valores nao usado apontava pro nul
// verificar se o campo eh nulo ou n

//struct anonimas

func main() {

	// read data from file
	jsonDataFromFile, err := ioutil.ReadFile("company.json")

	if err != nil {
		panic(err)
	}

	// Create a WaitGroup
	var wg sync.WaitGroup

	// Increment the wait group counter
	wg.Add(1)

	// Create a go routine with a lambda function, and pass a parameter in the end
	go func(jsonDataFromFile []byte) {
		// Decrement the counter when the go routine completes
		defer wg.Done()

		// Call the function that generate the file
		myFunc(jsonDataFromFile)

	}(jsonDataFromFile)

	for i := 0; i < 200; i++ {
		fmt.Printf("%d \n", i)
	}

	wg.Wait()
}

func myFunc(jsonDataFromFile []byte) {

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
}
