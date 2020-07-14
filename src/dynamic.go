package main

import (
	"encoding/json"
	"fmt"
	"log"

	dynamicstruct "github.com/ompluscator/dynamic-struct"
)

func main() {
	instance := dynamicstruct.NewStruct().
		AddField("Integer", 0, `json:"int"`).
		AddField("Text", "", `json:"someText"`).
		AddField("Float", 0.0, `json:"double"`).
		AddField("Boolean", false, "").
		AddField("Slice", []int{}, "").
		AddField("Anonymous", "", `json:"-"`).
		Build().
		New()

	data := []byte(`
{
    "int": 123,
    "someText": "example",
    "double": 123.45,
    "Boolean": true,
    "Slice": [1, 2, 3],
    "Anonymous": "avoid to read"
}
`)
	// Pegou o valor de data e colocou na instancia
	err := json.Unmarshal(data, &instance)
	if err != nil {
		log.Fatal(err)
	}

	// Pegou o valor da instance e colocou em data
	dataM, err := json.Marshal(instance)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(dataM))
	// Out:
	// {"int":123,"someText":"example","double":123.45,"Boolean":true,"Slice":[1,2,3]}
}
