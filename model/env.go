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

package model

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/flowqio/proxy/service"
)

//Env is environment model , include Container ID, Port ,Region/Zone , PrivateIP
type Env struct {
	ContainerID string
	Port        string
	Region      string
	PrivateIP   string
}

//NewEnv is builder func , build Env from string like  da8eff3898bb606a-80-env1.env.blockbox.io
func NewEnv(host string) (Env, error) {

	config := strings.Split(host, "-")

	if len(config) != 3 {
		log.Errorf("Host %s format is invalid", host)
		return Env{}, fmt.Errorf(service.EnvNotAvailable)
	}

	if strings.Index(config[2], "env") > -1 {
		//check container id and expose port
		port, err := service.GetEndpoint(config[0], config[1])
		if err != nil {
			log.Error(err)
			return Env{}, fmt.Errorf(service.EnvNotAvailable)
		}
		return Env{ContainerID: config[0], Port: port.HostPort, PrivateIP: port.HostIP, Region: config[2]}, nil
	} else {
		log.Errorf("Host %s format is invalid", host)
		return Env{}, fmt.Errorf(service.EnvNotAvailable)
	}

}
