package main

import (
	"fmt"
)

var Data *map[string]string

func main() {

	Data = hello()
	fmt.Println((*Data)["hi"])

}
func hello() *map[string]string {
	var cl = make(map[string]string)

	cl["hi"] = "hello"

	Clusters := &cl
	fmt.Println((*Clusters)["hi"])
	return &cl

}
