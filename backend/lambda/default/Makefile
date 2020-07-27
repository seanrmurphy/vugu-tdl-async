GOBUILD=go build
BINPATH=${PWD}

build:
	GOOS=linux GOARCH=amd64 ${GOBUILD} -o ${BINPATH}/default default.go
	chmod +x ${BINPATH}/default
	zip -j ${BINPATH}/default.zip ${BINPATH}/default
