GOBUILD=go build
BINPATH=${PWD}

build:
	GOOS=linux GOARCH=amd64 ${GOBUILD} -o ${BINPATH}/ondisconnect ondisconnect.go
	chmod +x ${BINPATH}/ondisconnect
	zip -j ${BINPATH}/ondisconnect.zip ${BINPATH}/ondisconnect
