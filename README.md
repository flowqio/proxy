# proxy
Simple HTTP/WebSocket proxy used access lab environment

This project use default golang http libaray, so have some performance issue, in the future should be changed.


flowqio/proxy use dep as default golang package manage tool.

proxy need access docker.sock (default docker access endpoint)

# how to install

go get -u github.com/flowqio/proxy

dep ensure

go build

./proxy