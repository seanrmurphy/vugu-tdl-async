GOBUILD=go build
FUNCTION=ondisconnect

build:
	GOOS=linux GOARCH=amd64 ${GOBUILD} -o ${CURDIR}/${FUNCTION} ${FUNCTION}.go
	chmod +x ${CURDIR}/${FUNCTION}
	zip -j ${CURDIR}/${FUNCTION}.zip ${CURDIR}/${FUNCTION}

.PHONY: build
