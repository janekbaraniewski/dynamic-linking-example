package main

import "C"
import (
	"encoding/json"
	"fmt"

	"github.com/janekbaraniewski/dynamic-linking-example/shared"
)

//export Provider1
func Provider1(arg *C.char) *C.char {
	fmt.Printf("Pro Provider 1 activated with arg: %s\n", C.GoString(arg))
	data := shared.Data{
		Message: "Hello from Pro Provider 1",
		Value:   200,
	}
	jsonData, _ := json.Marshal(data)
	return C.CString(string(jsonData))
}

//export Provider2
func Provider2(arg *C.char) *C.char {
	fmt.Printf("Pro Provider 2 activated with arg: %s\n", C.GoString(arg))
	data := shared.Data{
		Message: "Hello from Pro Provider 2",
		Value:   300,
	}
	jsonData, _ := json.Marshal(data)
	return C.CString(string(jsonData))
}

//export Provider3
func Provider3(arg *C.char) *C.char {
	fmt.Printf("Pro Provider 3 activated with arg: %s\n", C.GoString(arg))
	data := shared.Data{
		Message: "Hello from Pro Provider 3",
		Value:   400,
	}
	jsonData, _ := json.Marshal(data)
	return C.CString(string(jsonData))
}

func main() {}
