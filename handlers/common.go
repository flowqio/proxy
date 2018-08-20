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
)

//IsWebSocket check connection is or not weboscket
func IsWebSocket(headers http.Header) bool {

	if _, ok := headers["Sec-Websocket-Version"]; ok {
		return true
	}
	conn := strings.Join(headers["Connection"], " ")

	if strings.Index(conn, "Upgrade") > -1 {
		return true
	}

	return false
}
