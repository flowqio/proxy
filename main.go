// Copyright 2018 flowq Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/flowqio/proxy/service"
	"github.com/flowqio/proxy/version"

	_ "github.com/flowqio/proxy/router"
)

//build information  use -ldflags, etc. -ldflags "-X main.release=dev  -X main.githash=`git rev-parse HEAD`"
var release = "dev"
var builddate = time.Now().Format("200601021504")
var githash = ""

var port *int
var host *string

//main func
func main() {

	port = flag.Int("p", 8899, "listen port")
	host = flag.String("h", "localhost", "listen address")
	flag.Parse()

	//wtach container start and stop, in future will changed use API/etcd event, not need watch docker
	go service.WatchContainerEvent()

	serverStart()

}

func serverStart() {

	log.Infof("FlowQ Proxy Server ")
	log.Infof("version: %s.%s  build: %s  ", release, version.Version, builddate)

	if githash != "" {
		log.Infof("git hash: %s", githash)
	}

	log.Infof("Listen at %s:%d ", *host, *port)

	srv := &http.Server{
		Addr:         *host + ":" + strconv.Itoa(*port),
		WriteTimeout: 120 * time.Second,
		ReadTimeout:  120 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
