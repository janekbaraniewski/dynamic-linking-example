//go:build darwin || linux
// +build darwin linux

package loader

/*
#include <dlfcn.h>
#include <stdlib.h>

void* handle = NULL;

int loadLibrary(const char* libName) {
	handle = dlopen(libName, RTLD_LAZY);
	if (!handle) {
		return -1;
	}
	return 0;
}

const char* runProvider(const char* providerName, const char* arg) {
	const char* (*providerFunc)(const char*);
	providerFunc = (const char* (*)(const char*))dlsym(handle, providerName);
	if (!providerFunc) {
		return NULL;
	}
	return providerFunc(arg);
}
*/
import "C"
import (
	"encoding/json"
	"errors"
	"fmt"
	"unsafe"

	"github.com/janekbaraniewski/dynamic-linking-example/shared"
)

func LoadLibrary(libName string) error {
	cLibName := C.CString(libName)
	defer C.free(unsafe.Pointer(cLibName))
	if C.loadLibrary(cLibName) != 0 {
		return errors.New("failed to load library")
	}
	return nil
}

func RunProvider(providerName string, arg string) (*shared.Data, error) {
	cProviderName := C.CString(providerName)
	defer C.free(unsafe.Pointer(cProviderName))
	cArg := C.CString(arg)
	defer C.free(unsafe.Pointer(cArg))

	result := C.runProvider(cProviderName, cArg)
	if result == nil {
		return nil, fmt.Errorf("provider %s not implemented", providerName)
	}

	var data shared.Data
	if err := json.Unmarshal([]byte(C.GoString(result)), &data); err != nil {
		return nil, err
	}

	return &data, nil
}
