//go:build windows
// +build windows

package loader

import (
	"encoding/json"
	"fmt"
	"syscall"
	"unsafe"

	"github.com/janekbaraniewski/dynamic-linking-example/shared"
)

var dll *syscall.DLL

func LoadLibrary(libName string) error {
	var err error
	dll, err = syscall.LoadDLL(libName)
	if err != nil {
		return err
	}
	return nil
}

func RunProvider(providerName string, arg string) (*shared.Data, error) {
	proc, err := dll.FindProc(providerName)
	if err != nil {
		return nil, fmt.Errorf("provider %s not implemented", providerName)
	}

	// Convert the string argument to a byte pointer that Windows can understand
	argPtr := uintptr(unsafe.Pointer(syscall.StringBytePtr(arg)))
	// Call the provider function, expecting it to return a pointer to a JSON string
	ret, _, err := proc.Call(argPtr)
	if err != syscall.Errno(0) {
		return nil, fmt.Errorf("error calling provider %s: %v", providerName, err)
	}

	// Convert the returned uintptr (which is a pointer to a JSON string) to a Go string
	jsonData := uintptrToString(ret)

	// Unmarshal the JSON data into a Go structure
	var data shared.Data
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		return nil, err
	}

	return &data, nil
}

// uintptrToString converts a pointer returned from a syscall to a Go string
func uintptrToString(ptr uintptr) string {
	return syscall.UTF16ToString((*[1 << 16]uint16)(unsafe.Pointer(ptr))[:])
}
