VERSION=$(shell git describe --abbrev=0 --tags)
BUILD=$(shell git rev-parse HEAD)
DIRBASE=./bin
DIR=${DIRBASE}/${VERSION}

LDFLAGS=-ldflags "-s -w -buildid=${BUILD}"

$(shell mkdir -p ${DIR})

windows:
	env CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -trimpath -o ${DIR}/tsk-windows_amd64.exe
	env CGO_ENABLED=1 GOOS=windows GOARCH=386 go build ${LDFLAGS} -trimpath -o ${DIR}/tsk-windows-386.exe

linux:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath ${LDFLAGS} -o ${DIR}/tsk-linux_amd64
	env CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -trimpath ${LDFLAGS} -o ${DIR}/tsk-linux-386

darwin:
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -trimpath ${LDFLAGS} -o ${DIR}/tsk-darwin_amd64
	env CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -trimpath ${LDFLAGS} -o ${DIR}/tsk-darwin-arm64

clean:
	rm -rf ${DIRBASE}


.PHONY: windows linux darwin android clean
