BINARY = rpic-server
DATE = $(shell date +%FT%T%Z)
BUILD_DIR = build/bin
LOG_DIR= build/log
PWD = $(shell pwd)

LDFLAGS =

.PHONY: test clean arm amd64 run install uninstall

# Build the project
all: clean test amd64 arm

run:
	$(shell build/bin/rpic-server-linux-amd64)

amd64:
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${BUILD_DIR}/${BINARY}-linux-amd64 -v
arm:
	GOOS=linux GOARCH=arm go build ${LDFLAGS} -o ${BUILD_DIR}/${BINARY}-linux-arm -v
install:
	#$(shell cp ./build/bin/rpic-server-linux-arm /usr/bin/)
	#$(shell cp ./systemd/rpicserver.service /etc/systemd/system)

uninstall:
	$(shell rm /etc/systemd/system/rpicserver.service)
test:
	@echo "Write testlog..."
	$(shell go test -v ./test... > ${LOG_DIR}/test_${DATE}.log)
viz:
	go-callvis -minlen 3 -focus sockets -group pkg,type ./ | dot -Tpng -o sockets.png
clean:
	-rm -f ${BUILD_DIR}/${BINARY}-*

