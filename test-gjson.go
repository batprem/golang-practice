package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

const json = `{
	"name":{"first":"Janet","last":"Prichard"},
	"age":47,
	"1":["test"]
}`

func main() {
	defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
	value := gjson.Get(json, "1.0")
	if value.String() == ""{
		panic("path not found")
	}
	println("test: ", value.String())
}
