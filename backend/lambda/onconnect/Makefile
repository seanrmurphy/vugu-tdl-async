GOBUILD=go build
BINPATH=${PWD}

build:
	GOOS=linux GOARCH=amd64 ${GOBUILD} -o ${BINPATH}/onconnect onconnect.go
	chmod +x ${BINPATH}/onconnect
	zip -j ${BINPATH}/onconnect.zip ${BINPATH}/onconnect
