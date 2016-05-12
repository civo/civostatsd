# This is how we want to name the binary output
BINARY=civostatsd

# These are the values we want to pass for Version and BuildTime
VERSION=1.0.0
BUILD_TIME=`date +%FT%T%z`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME}"

all:
	go build ${LDFLAGS} -o ${BINARY} main.go
test:
	make all
	./civostatsd --test --config civostatsd-dummy.conf
