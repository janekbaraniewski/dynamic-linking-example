package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/janekbaraniewski/dynamic-linking-example/loader"
)

var providers = []string{"Provider1", "Provider2", "Provider3"}

func getLibraryFileName(baseName string) string {
	switch runtime.GOOS {
	case "windows":
		return baseName + ".dll"
	case "darwin":
		return baseName + ".dylib"
	default: // Linux and others
		return baseName + ".so"
	}
}

func main() {
	arg := "Hello from main!"
	libName := getLibraryFileName("pro")
	fmt.Printf("Attempting to load library: %s\n", libName)
	if err := loader.LoadLibrary(libName); err != nil {
		fmt.Printf("Pro version not available: %v\n", err)
		libName = getLibraryFileName("oss")
		fmt.Printf("Attempting to load library: %s\n", libName)
		if err := loader.LoadLibrary(libName); err != nil {
			fmt.Printf("OSS version also not available: %v\n", err)
			fmt.Println("Exiting.")
			os.Exit(1)
		}
	}

	for _, provider := range providers {
		data, err := loader.RunProvider(provider, arg)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Received data: %+v\n", data)
		}
	}
}
