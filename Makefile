# Set the build dir, where built cross-compiled binaries will be output
BUILDDIR := bin

build:
	go build -o $(BUILDDIR)/main cmd/back-front-auth/main.go

run:
	go run main.go
