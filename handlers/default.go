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
	"net/http"
	"strings"

	"github.com/flowqio/proxy/model"
	log "github.com/sirupsen/logrus"
)

//Home default handle func
func Home(w http.ResponseWriter, r *http.Request) {

	host := strings.Split(r.Host, ":")

	log.Debugf("Host: %v, Path : %s", host[0], r.URL.Path)
	log.Debugf("Header %+v", r.Header)

	prefix := strings.Split(host[0], ".")[0]

	env, err := model.NewEnv(prefix)

	if err != nil {

		log.Error(err)

		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(err.Error()))

		return
	}

	if IsWebSocket(r.Header) {
		log.Debug(env, r.Host, r.URL.Path)
		handleWS(env, w, r)
	} else {
		handleHTTP(env, w, r)
	}

}
