# proxy
Simple HTTP/WebSocket proxy used access lab environment

This project use default golang http libaray, so have some performance issue, in the future should be changed.


flowqio/proxy use dep as default golang package manage tool.

proxy need access docker.sock (default docker access endpoint)

proxy is flowq compoent , deployed every env host through k8s as DaemonSet

# how to install

go get -u github.com/flowqio/proxy

dep ensure

go build

./proxy


# how is work

The endpoint follow spec rule :

http|https://{{user container id [:16]}}-{{ export port }}-{{env cluster}}.env.flowq.io

When proxy accept brower/client request , it will check local docker and proxy all data.


