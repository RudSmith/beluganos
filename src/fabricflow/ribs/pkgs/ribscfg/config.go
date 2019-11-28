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

package ribscfg

import (
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
)

const (
	// GoBGPdGrpcPort is default port of GoBGP API
	GoBGPdGrpcPort = 50051
	// VirtNexthops is default address of virtual nexthop.
	VirtNexthops = "127.0.0.1/32"
)

//
// NodeConfig is config of node.
//
type NodeConfig struct {
	NId       uint8  `toml:"nid"`
	Label     uint32 `toml:"label"`
	NIdIfname string `toml:"nid_from_ifaddr"`
}

//
// NLAConfig is config of nla.
//
type NLAConfig struct {
	API string `toml:"api"`
}

//
// BgpConfig is config gobgp api.
//
type BgpConfig struct {
	Addr string `toml:"addr"`
	Port uint16 `toml:"port"`

	RouteFamily string `toml:"route_family"`
}

//
// GetAddr returns gobgp api address.
//
func (c *BgpConfig) GetAddr() string {
	return fmt.Sprintf("%s:%d", c.Addr, c.Port)
}

//
// VrfConfig is config of VRF.
//
type VrfConfig struct {
	Iface string `toml:"iface"`
	Rt    string `toml:"rt"`
	Rd    string `toml:"rd"`
}

//
// NHConfig is config of nexthop translation.
//
type NHConfig struct {
	Mode string `toml:"mode"`
	Args string `toml:"args"`
}

//
// RibsConfig is config of RIBS.
//
type RibsConfig struct {
	Disable  bool      `toml:"disable"`
	Core     string    `toml:"core"`
	API      string    `toml:"api"`
	SyncTime int64     `toml:"resync"`
	Nexthops NHConfig  `toml:"nexthops"`
	Bgp      BgpConfig `toml:"bgpd"`
	Vrf      VrfConfig `toml:"vrf"`
}

//
// GetSyncTime returns interval time of sync rib.
//
func (c *RibsConfig) GetSyncTime() time.Duration {
	return time.Duration(c.SyncTime) * time.Millisecond
}

//
// Config is root config.
//
type Config struct {
	Node NodeConfig `toml:"node"`
	NLA  NLAConfig  `toml:"nla"`
	Ribs RibsConfig `toml:"ribs"`
}

//
// GetBgpdAddr returns gobgp api address.
//
func (c *Config) GetBgpdAddr() string {
	return c.Ribs.Bgp.GetAddr()
}

//
// VrfLabel returns vrf label.
//
func (c *Config) VrfLabel() uint32 {
	return uint32(c.Node.NId) + c.Node.Label
}

//
// ReadConfig reads config from file.
//
func ReadConfig(path string, cfg *Config) error {
	_, err := toml.DecodeFile(path, cfg)
	if err != nil {
		return err
	}

	if cfg.Ribs.Bgp.Port == 0 {
		cfg.Ribs.Bgp.Port = GoBGPdGrpcPort
	}

	if cfg.Ribs.Nexthops.Args == "" {
		cfg.Ribs.Nexthops.Args = VirtNexthops
	}
	return nil
}
