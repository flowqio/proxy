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

package util

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	EnableDebug()
}

//EnableDebug is helper func , if os environment have DEBUG, open debug level
func EnableDebug() {
	if debug := os.Getenv("DEBUG"); debug != "" {
		log.SetLevel(log.DebugLevel)
	}
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, TimestampFormat: "2006-01-02 15:04:05"})
}
