package main

import "C"
import (
	"encoding/json"
	"fmt"

	"github.com/janekbaraniewski/dynamic-linking-example/shared"
)

//export Provider1
func Provider1(arg *C.char) *C.char {
	fmt.Printf("OSS Provider 1 activated with arg: %s\n", C.GoString(arg))
	data := shared.Data{
		Message: "Hello from OSS Provider 1",
		Value:   100,
	}
	jsonData, _ := json.Marshal(data)
	return C.CString(string(jsonData))
}

func main() {}
