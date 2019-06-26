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

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

type Port struct {
	Id     string
	Port   uint32
	HwAddr string
}

func (p *Port) String() string {
	return fmt.Sprintf("{'%s', 0x%x, '%s'}", p.Id, p.Port, p.HwAddr)
}

type Link struct {
	ReId string
	Name string
}

func (l *Link) String() string {
	if l == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{'%s', '%s'}", l.ReId, l.Name)
}

type PortMap struct {
	Name    Link
	Vm      Port
	Vs      Port
	Dp      Port
	Link    *Link
	DpEnter bool
	NoVS    bool
}

func (p *PortMap) String() string {
	var cnt int

	vm := func() string {
		if p.Vm.Port != 0 {
			cnt++
			return fmt.Sprintf("+%s", &p.Vm)
		}
		return fmt.Sprintf("-%s", &p.Vm)
	}()

	dp := func() string {
		if p.DpEnter {
			cnt++
			return fmt.Sprintf("+%s", &p.Dp)
		}
		if p.Link != nil {
			cnt++
			return fmt.Sprintf("|%s", &p.Dp)
		}
		return fmt.Sprintf("-%s", &p.Dp)
	}()

	vs := func() string {
		if p.NoVS {
			cnt++
			return fmt.Sprintf("*%s", &p.Vs)
		}
		if p.Link != nil {
			cnt++
			return fmt.Sprintf("|%s", &p.Vs)
		}
		if p.Vs.Port != 0 {
			cnt++
			return fmt.Sprintf("+%s", &p.Vs)
		}
		return fmt.Sprintf("-%s", &p.Vs)
	}()

	name := func() string {
		if cnt == 3 {
			return fmt.Sprintf("%s +", &p.Name)
		}
		return fmt.Sprintf("%s -", &p.Name)
	}()

	return fmt.Sprintf("IF=%s VM%s DP%s VS%s Link=%s", name, vm, dp, vs, p.Link)
}

func DecodePortMap(datas map[string]interface{}) *PortMap {
	pm := &PortMap{}
	for key, val := range datas {
		switch key {
		case "name":
			field := val.([]interface{})
			pm.Name = Link{
				ReId: field[0].(string),
				Name: field[1].(string),
			}

		case "vm":
			field := val.([]interface{})
			pm.Vm = Port{
				Id:     field[0].(string),
				Port:   uint32(field[1].(float64)),
				HwAddr: pm.Vm.HwAddr,
			}

		case "dp":
			field := val.([]interface{})
			pm.Dp = Port{
				Id:     fmt.Sprintf("%d", uint32(field[0].(float64))),
				Port:   uint32(field[1].(float64)),
				HwAddr: pm.Dp.HwAddr,
			}

		case "vs":
			field := val.([]interface{})
			pm.Vs = Port{
				Id:     fmt.Sprintf("%d", uint32(field[0].(float64))),
				Port:   uint32(field[1].(float64)),
				HwAddr: pm.Vs.HwAddr,
			}

		case "link":
			if val != nil {
				field := val.([]interface{})
				pm.Link = &Link{
					ReId: field[0].(string),
					Name: field[1].(string),
				}
			}

		case "dpenter":
			field := val.(bool)
			pm.DpEnter = field

		case "no_vs":
			field := val.(bool)
			pm.NoVS = field

		case "vs_hw_addr":
			pm.Vs.HwAddr = val.(string)

		default:
			fmt.Printf("unknown field. %s %v\n", key, val)
		}
	}
	return pm
}

type Dpmap struct {
	dp_id uint32
	name  string
	mode  string
}

func (d *Dpmap) String() string {
	return fmt.Sprintf("DP={'%s', %d, '%s'}", d.name, d.dp_id, d.mode)
}

func DecodeDpmap(datas map[string]interface{}) *Dpmap {
	d := &Dpmap{}
	for key, val := range datas {
		switch key {
		case "dp_id":
			d.dp_id = uint32(val.(float64))
		case "name":
			d.name = val.(string)
		case "mode":
			d.mode = val.(string)
		default:
			fmt.Printf("unknown field. %s %v\n", key, val)
		}
	}
	return d
}

func HttpGet(url string) (io.Reader, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, err
	}

	return res.Body, nil
}

func showPortmap(addr string) {
	r, err := HttpGet(fmt.Sprintf("http://%s/fib%s", addr, "/portmap"))
	if err != nil {
		fmt.Printf("HttpGet error. %s", err)
		return
	}

	datas := make([]map[string]interface{}, 0)
	json.NewDecoder(r).Decode(&datas)

	for _, data := range datas {
		pm := DecodePortMap(data)
		fmt.Println(pm)
	}
}

type Idmap struct {
	DpId     uint64
	ReId     string
	DpStatus bool
	VmStatus bool
}

func (i *Idmap) String() string {
	return fmt.Sprintf("ID={'%s', %d, VM=%t, DP=%t}", i.ReId, i.DpId, i.VmStatus, i.DpStatus)
}

func DecodeIdMap(datas interface{}) *Idmap {
	im := &Idmap{}
	for key, val := range datas.(map[string]interface{}) {
		switch key {
		case "dp_status":
			im.DpStatus = val.(bool)
		case "vm_status":
			im.VmStatus = val.(bool)
		case "dp_id":
			im.DpId = uint64(val.(float64))
		case "re_id":
			im.ReId = val.(string)
		}
	}

	return im
}

func showIdmap(addr string) {
	r, err := HttpGet(fmt.Sprintf("http://%s/fib%s", addr, "/idmap"))
	if err != nil {
		fmt.Printf("HttpGet error. %s", err)
		return
	}

	datas := make([]interface{}, 0)
	json.NewDecoder(r).Decode(&datas)

	for _, data := range datas {
		im := DecodeIdMap(data)
		fmt.Println(im)
	}
}

func showDpmap(addr string) {
	r, err := HttpGet(fmt.Sprintf("http://%s/fib%s", addr, "/dpmap"))
	if err != nil {
		fmt.Printf("HttpGet error. %s", err)
		return
	}

	datas := make(map[string]interface{}, 0)
	json.NewDecoder(r).Decode(&datas)
	for _, data := range datas {
		d := DecodeDpmap(data.(map[string]interface{}))
		fmt.Printf("%v\n", d)
	}
}

func main() {
	var addr string
	flag.StringVar(&addr, "addr", "127.0.0.1:8080", "ryu addr")
	flag.Parse()

	showPortmap(addr)
	showIdmap(addr)
	showDpmap(addr)
}
