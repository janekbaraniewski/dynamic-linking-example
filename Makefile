# Detect platform and set appropriate file extensions
OS := $(shell uname)
ifeq ($(OS), Linux)
	OSS_SHARED_LIB=oss.so
	PRO_SHARED_LIB=pro.so
	SHARED_BUILD_MODE=c-shared
else ifeq ($(OS), Darwin)
	OSS_SHARED_LIB=oss.dylib
	PRO_SHARED_LIB=pro.dylib
	SHARED_BUILD_MODE=c-shared
else ifeq ($(OS), Windows_NT)
	OSS_SHARED_LIB=oss.dll
	PRO_SHARED_LIB=pro.dll
	SHARED_BUILD_MODE=c-shared
else
	$(error Unsupported OS: $(OS))
endif

# Build the OSS shared library
build-oss:
	go build -o $(OSS_SHARED_LIB) -buildmode=$(SHARED_BUILD_MODE) ./oss/oss.go

# Build the Pro shared library
build-pro:
	go build -o $(PRO_SHARED_LIB) -buildmode=$(SHARED_BUILD_MODE) ./pro/pro.go

# Build the main application
build-main:
	go build -o main main.go

# Clean up build artifacts
clean:
	rm -f main $(OSS_SHARED_LIB) $(PRO_SHARED_LIB) oss.h pro.h

# Run the application with OSS version
run-oss: build-oss build-main
	./main

# Run the application with Pro version
run-pro: build-pro build-main
	./main

# By default, build and run with OSS version
all: run-oss
