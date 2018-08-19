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

package router

import (
	"net/http"

	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"

	"github.com/flowqio/proxy/handlers"

	_ "github.com/flowqio/proxy/utils"
)

func init() {

	log.Debug("Init Router succesful...")
	InitRouter()
}

//InitRouter provide single router config access
func InitRouter() {

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.Home)
	r.HandleFunc("/index", handlers.Home)
	r.PathPrefix("/").HandlerFunc(handlers.Home)

	//helper endporint
	r.HandleFunc("/favicon.ico", handlers.NotFound)
	r.HandleFunc("/api", handlers.API)
	r.HandleFunc("/api/v1", handlers.API)

	http.Handle("/", r)

}
