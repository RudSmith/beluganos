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

package ribctl

import (
	"fabricflow/fibc/api"
	"fmt"
	"gonla/nlamsg"
	"gonla/nlamsg/nlalink"
	"strconv"
	"strings"
	"syscall"
)

func NewVRFLabel(base uint32, nid uint8) uint32 {
	return base + uint32(nid)
}

func newLinkName(nid uint8, name string, useNId bool) string {
	if useNId {
		return fmt.Sprintf("%d/%s", nid, name)
	}

	return name
}

func NewLinkName(link *nlamsg.Link, useNId bool) string {
	return newLinkName(link.NId, link.Attrs().Name, useNId)
}

func ParseLinkName(name string) (uint8, string) {
	i := strings.Index(name, "/")
	if i < 0 {
		return 0, name
	}

	nid := func() uint8 {
		if n, err := strconv.Atoi(name[:i]); err == nil {
			return uint8(n)
		}

		return 0
	}()

	return nid, name[i+1:]
}

func NewAddrLinkName(addr *nlamsg.Addr, useNId bool) string {
	return newLinkName(addr.NId, addr.Label, useNId)
}

func NewPortId(link *nlamsg.Link) uint32 {
	if link != nil {
		return (uint32(link.NId) << 16) + uint32(link.LnId)
	} else {
		return 0
	}
}

func ParsePortId(linkId uint32) (uint8, uint16) {
	return uint8(linkId >> 16), uint16(linkId & 0xffff)
}

func NewPortStatus(link *nlamsg.Link) fibcapi.PortStatus_Status {
	switch link.Attrs().OperState.String() {
	case "up":
		return fibcapi.PortStatus_UP
	case "down":
		return fibcapi.PortStatus_DOWN
	default:
		return fibcapi.PortStatus_NOP
	}
}

func NewNeighId(neigh *nlamsg.Neigh) uint32 {
	if neigh != nil {
		return (uint32(neigh.NId) << 16) + uint32(neigh.NeId)
	} else {
		return 0
	}
}

func GetPortConfigCmd(t uint16) string {
	switch t {
	case syscall.RTM_NEWLINK:
		return "ADD"
	case syscall.RTM_DELLINK:
		return "DELETE"
	case syscall.RTM_SETLINK:
		return "MODIFY"
	default:
		return "NOP"
	}
}

func GetGroupCmd(t uint16) fibcapi.GroupMod_Cmd {
	switch t {
	case syscall.RTM_NEWLINK, syscall.RTM_NEWADDR, syscall.RTM_NEWNEIGH, syscall.RTM_NEWROUTE:
		return fibcapi.GroupMod_ADD

	case syscall.RTM_SETLINK, nlalink.RTM_SETADDR, nlalink.RTM_SETNEIGH, nlalink.RTM_SETROUTE:
		return fibcapi.GroupMod_MODIFY

	case syscall.RTM_DELLINK, syscall.RTM_DELADDR, syscall.RTM_DELNEIGH, syscall.RTM_DELROUTE:
		return fibcapi.GroupMod_DELETE

	default:
		return fibcapi.GroupMod_NOP
	}
}

func GetFlowCmd(t uint16) fibcapi.FlowMod_Cmd {
	switch t {
	case syscall.RTM_NEWLINK, syscall.RTM_NEWADDR, syscall.RTM_NEWNEIGH, syscall.RTM_NEWROUTE:
		return fibcapi.FlowMod_ADD

	case syscall.RTM_SETLINK, nlalink.RTM_SETADDR, nlalink.RTM_SETNEIGH, nlalink.RTM_SETROUTE:
		return fibcapi.FlowMod_MODIFY

	case syscall.RTM_DELLINK, syscall.RTM_DELADDR, syscall.RTM_DELNEIGH, syscall.RTM_DELROUTE:
		return fibcapi.FlowMod_DELETE

	default:
		return fibcapi.FlowMod_NOP
	}
}

func FlowCmdToGroupCmd(cmd fibcapi.FlowMod_Cmd) fibcapi.GroupMod_Cmd {
	switch cmd {
	case fibcapi.FlowMod_ADD:
		return fibcapi.GroupMod_ADD
	case fibcapi.FlowMod_MODIFY:
		return fibcapi.GroupMod_MODIFY
	case fibcapi.FlowMod_DELETE:
		return fibcapi.GroupMod_DELETE
	default:
		return fibcapi.GroupMod_NOP
	}
}
