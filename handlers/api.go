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
	"encoding/json"
	"net/http"

	"github.com/flowqio/proxy/version"
	log "github.com/sirupsen/logrus"
)

//API Information
func API(w http.ResponseWriter, r *http.Request) {
	log.Debugf("Connection : %s ", r.Header["Connection"])
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(version.VersionInfo())
	if err != nil {
		panic(err)
	}

}
