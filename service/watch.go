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

package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	log "github.com/sirupsen/logrus"

	"github.com/docker/go-connections/nat"

	client "docker.io/go-docker"
	types "docker.io/go-docker/api/types"
)

const EnvNotAvailable string = "enviroment  not available"

var endpoints = make(map[string]nat.PortMap)

//GetEndpoint , check cotnainer id , port , if not exits in endpoint return EnvNotAvailable error
func GetEndpoint(cid, port string) (nat.PortBinding, error) {

	if v, ok := endpoints[cid]; ok {
		for k, _ := range v {
			if port == k.Port() {
				log.Debugf("match %+v", v[k][0])
				return v[k][0], nil
			}
		}
	}
	return nat.PortBinding{}, fmt.Errorf(EnvNotAvailable)
}

//WatchContainerEvent watch docker container event
func WatchContainerEvent() error {

	log.Info("Starting docker event watch")
	c, err := client.NewEnvClient()

	if err != nil {
		log.Error(err)
		return err
	}

	containers, err := c.ContainerList(context.Background(), types.ContainerListOptions{})
	log.Info("Load containers information from docker")
	if err != nil {
		log.Error(err)
		return err
	}
	for _, v := range containers {
		sID := shortID(v.ID)
		endpoints[sID] = make(nat.PortMap)
		for _, v1 := range v.Ports {
			port, _ := nat.NewPort(v1.Type, fmt.Sprintf("%d", v1.PrivatePort))
			endpoints[sID][port] = []nat.PortBinding{}
			endpoints[sID][port] = append(endpoints[sID][port], nat.PortBinding{HostIP: v1.IP, HostPort: fmt.Sprintf("%d", v1.PublicPort)})
		}

	}

	data, _ := json.MarshalIndent(endpoints, "", " ")
	log.Debug(string(data))

	messages, errs := c.Events(context.Background(), types.EventsOptions{})

	for {
		select {
		case err := <-errs:
			if err != nil && err != io.EOF {
				log.Error(err)
			}

		case e := <-messages:

			switch e.Action {
			case "start":
				info, err := c.ContainerInspect(context.Background(), e.ID)
				if err != nil {
					log.Println(err)
				} else {
					log.Debugf("start event received %s ,  %+v %+v", e.ID, info.HostConfig, info.NetworkSettings)
					addPortBindings(shortID(e.ID), info.NetworkSettings.Ports)

				}
			case "stop":
				log.Debugf("stop event received  %s", e.ID)
				removePortBindnds(shortID(e.ID))
			default:
				log.Debugf("other event received %s & id %s ", e.Action, e.ID)
			}

		}
	}
}

func addPortBindings(cid string, ports nat.PortMap) {
	endpoints[cid] = ports
	data, _ := json.MarshalIndent(endpoints, "", " ")
	log.Debug(string(data))
}

func removePortBindnds(cid string) {
	delete(endpoints, cid)
	data, _ := json.MarshalIndent(endpoints, "", " ")
	log.Debug(string(data))
}

//shortID is helper func ,just return container id length 16
func shortID(id string) string {
	if len(id) > 16 {
		return id[:16]
	}
	return id
}
