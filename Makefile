build:
	go build -ldflags "-X main.release=dev  -X main.githash=`git rev-parse HEAD`"

build_linux:
	GOOS=linux go build -o images/proxy -ldflags "-X main.release=dev  -X main.githash=`git rev-parse HEAD`"
