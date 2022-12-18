BINARY_NAME=glox
OUTPUTDIR=ast

compile:
	go build -o glox/bin/${BINARY_NAME}.exe glox/main.go

runexe:
	./glox/bin/${BINARY_NAME}

run:
	go run glox/main.go

build:
	GOOS=windows GOARCH=amd64 go build -o glox/bin/${BINARY_NAME}-windows.exe glox/main.go
	GOOS=linux GOARCH=amd64 go build -o glox/bin/${BINARY_NAME}-linux glox/main.go
	GOOS=darwin GOARCH=amd64 go build -o glox/bin/${BINARY_NAME}-darwin glox/main.go

gen:
	go run tools/main.go ${OUTPUTDIR}


clean:
	go clean
	rm glox/bin/${BINARY_NAME}-darwin
	rm glox/bin/${BINARY_NAME}-linux
	rm glox/bin/${BINARY_NAME}-windows.exe
