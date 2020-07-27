GOBUILD=go build
BINPATH=${PWD}

build:
	GOOS=linux GOARCH=amd64 ${GOBUILD} -o ${BINPATH}/echo echo.go
	chmod +x ${BINPATH}/echo
	zip -j ${BINPATH}/echo.zip ${BINPATH}/echo
