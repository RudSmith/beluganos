// -*- coding: utf-8 -*-

// Copyright (C) 2017 Nippon Telegraph and Telephone Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nlasvc

import (
	"gonla/nlactl"

	log "github.com/sirupsen/logrus"
)

type NLACoreApiService struct {
	Addr string
	log  *log.Entry
}

func NewNLACoreApiService(addr string) *NLACoreApiService {
	return &NLACoreApiService{
		Addr: addr,
		log:  NewLogger("NLACoreApiService"),
	}
}

func (n *NLACoreApiService) Start(nid uint8, chans *nlactl.NLAChannels) error {
	s := NewNLACoreApiServer(n.Addr)
	if err := s.Start(chans.NlMsg); err != nil {
		return err
	}

	n.log.Infof("Start:")
	return nil
}

func (n *NLACoreApiService) Stop() {
	n.log.Infof("Stop:")
}
