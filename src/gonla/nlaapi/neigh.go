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

package nlaapi

import (
	"gonla/nladbm"
	"gonla/nlamsg"
	"net"

	"github.com/vishvananda/netlink"
)

func (n *Neigh) GetIP() net.IP {
	return net.IP(n.Ip)
}

func (n *Neigh) GetLLIPAddr() net.IP {
	return net.IP(n.LlIpAddr)
}

func (n *Neigh) NetHardwareAddr() net.HardwareAddr {
	return net.HardwareAddr(n.HardwareAddr)
}

func (n *Neigh) SetNotun(notun *NeighNotun) {
	n.Tunnel = &Neigh_Notun{
		Notun: notun,
	}
}

func (n *Neigh) SetIptun(iptun *NeighIptun) {
	n.Tunnel = &Neigh_Iptun{
		Iptun: iptun,
	}
}

func (n *Neigh) ToNetlink() *netlink.Neigh {
	return &netlink.Neigh{
		LinkIndex:    int(n.LinkIndex),
		Family:       int(n.Family),
		State:        int(n.State),
		Type:         int(n.Type),
		Flags:        int(n.Flags),
		IP:           n.GetIP(),
		HardwareAddr: n.NetHardwareAddr(),
		LLIPAddr:     n.GetLLIPAddr(),
		Vlan:         int(n.VlanId),
		VNI:          int(n.Vni),
	}
}

func (n *Neigh) ToNative() *nlamsg.Neigh {
	return &nlamsg.Neigh{
		Neigh:   n.ToNetlink(),
		PhyLink: int(n.PhyLink),
		Tunnel:  Neigh_TunnelToNative(n.Tunnel),
		NeId:    uint16(n.NeId),
		NId:     uint8(n.NId),
	}
}

func NewNeighFromNative(n *nlamsg.Neigh) *Neigh {
	return &Neigh{
		LinkIndex:    int32(n.LinkIndex),
		Family:       int32(n.Family),
		State:        int32(n.State),
		Type:         int32(n.Type),
		Flags:        int32(n.Flags),
		Ip:           n.IP,
		HardwareAddr: n.HardwareAddr,
		LlIpAddr:     n.LLIPAddr,
		VlanId:       int32(n.Vlan),
		Vni:          int32(n.VNI),
		PhyLink:      int32(n.PhyLink),
		Tunnel:       NewNeigh_TunnelFromNative(n.Tunnel),
		NId:          uint32(n.NId),
		NeId:         uint32(n.NeId),
	}
}

func Neigh_TunnelToNative(src isNeigh_Tunnel) nlamsg.NeighTunnel {
	if src == nil {
		return nil
	}

	switch n := src.(type) {
	case *Neigh_Notun:
		return n.Notun.ToNative()

	case *Neigh_Iptun:
		return n.Iptun.ToNative()

	default:
		return nil
	}
}

func NewNeigh_TunnelFromNative(src nlamsg.NeighTunnel) isNeigh_Tunnel {
	if src == nil {
		return &Neigh_Notun{
			Notun: &NeighNotun{},
		}
	}

	switch n := src.(type) {
	case *nlamsg.NeighIptun:
		return &Neigh_Iptun{
			Iptun: NewNeighIptunFromNative(n),
		}

	default:
		return &Neigh_Notun{
			Notun: &NeighNotun{},
		}
	}
}

//
// Neigh (Key)
//

func (k *NeighKey) ToNative() *nladbm.NeighKey {
	return &nladbm.NeighKey{
		NId:     uint8(k.NId),
		Addr:    k.Addr,
		Ifindex: int(k.Ifindex),
		Vid:     int(k.VlanId),
	}
}

func NewNeighKeyFromNative(n *nladbm.NeighKey) *NeighKey {
	return &NeighKey{
		NId:     uint32(n.NId),
		Addr:    n.Addr,
		Ifindex: int32(n.Ifindex),
		VlanId:  int32(n.Vid),
	}
}

//
// Neighs
//
func NewGetNeighsRequest(nid uint8) *GetNeighsRequest {
	return &GetNeighsRequest{
		NId: uint32(nid),
	}
}
