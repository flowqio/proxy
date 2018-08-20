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

package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/flowqio/proxy/model"

	log "github.com/sirupsen/logrus"
)

//handleHTTP internal func , process http request
func handleHTTP(env model.Env, w http.ResponseWriter, req *http.Request) {
	log.Debugf("ContainerID %s, Port: %s,Connection : %s ,URL: %s", env.ContainerID, env.Port, req.Header["Connection"], req.URL.Path)

	req.Host = env.PrivateIP + ":" + env.Port
	req.URL = &url.URL{Scheme: "http", Host: req.Host, Path: req.URL.Path}

	log.Debugf("REQ %+v", req)

	defaultRoundTripper := http.DefaultTransport
	defaultTransportPointer, ok := defaultRoundTripper.(*http.Transport)
	if !ok {
		panic(fmt.Sprintf("defaultRoundTripper not an *http.Transport"))
	}
	defaultTransport := *defaultTransportPointer // dereference it to get a copy of the struct that the pointer points to
	defaultTransportPointer.MaxIdleConns = 100
	defaultTransport.MaxIdleConnsPerHost = 100

	resp, err := defaultTransport.RoundTrip(req)

	//resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		log.Error(err)
		http.Error(w, "Backend Service Unavailable", http.StatusServiceUnavailable)
		return
	}

	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
