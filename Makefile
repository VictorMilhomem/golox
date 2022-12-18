BINARY_NAME=glox.exe

build:
	go build -o glox/bin/${BINARY_NAME} glox/main.go

runexe:
	./glox/bin/${BINARY_NAME}

run:
	go run glox/main.go

