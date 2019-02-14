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
	"gonla/nladbm"
	"gonla/nlalib"
	"gonla/nlamsg"
	"gonla/nlamsg/nlalink"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
)

const TunnelUpdateInterval = 1800 * time.Second

type NLAMasterService struct {
	Service nlactl.NLAService
	iptun   *NLAMasterIptun
	NId     uint8
}

func NewNLAMasterService(service nlactl.NLAService) *NLAMasterService {
	s := &NLAMasterService{
		Service: service,
		NId:     0,
	}
	s.iptun = NewNLAMasterIptun(s, TunnelUpdateInterval)
	return s
}

func (n *NLAMasterService) Start(nid uint8, chans *nlactl.NLAChannels) error {
	n.NId = nid

	if err := n.Service.Start(n.NId, chans); err != nil {
		return nil
	}

	go n.iptun.Serve()
	go SubscribeNetlinkResources(chans.NlMsg, 0)

	log.Infof("MasterService: START")
	return nil
}

func (n *NLAMasterService) Stop() {
	n.Service.Stop()
	log.Infof("MasterService: STOP")
}

func (n *NLAMasterService) NetlinkMessage(nlmsg *nlamsg.NetlinkMessage) {
	//log.Debugf("MasterService: NlMsg")
}

func (n *NLAMasterService) NetlinkLink(nlmsg *nlamsg.NetlinkMessage, link *nlamsg.Link) {
	log.Debugf("MasterService: LINK")

	switch nlmsg.Type() {
	case syscall.RTM_NEWLINK:
		if old := nladbm.Links().Insert(link); old != nil {
			nlmsg.Header.Type = syscall.RTM_SETLINK
		}

		nlamsg.DispatchLink(nlmsg, link, n.Service)

		if iptun := link.Iptun(); iptun != nil {
			// remote is treated as neigh.
			// if route to remote exists, generate NEWNEIGH.
			n.iptun.RemoteUp(link.NId, iptun.Remote)
		}

	case syscall.RTM_DELLINK:
		if iptun := link.Iptun(); iptun != nil {
			// remote is treated as neigh.
			// if route to remote exists, generate DELNEIGH.
			n.iptun.RemoteDown(link.NId, iptun.Remote)
		}

		if old := nladbm.Links().Delete(nladbm.LinkToKey(link)); old != nil {
			link.LnId = old.LnId
		}

		nlamsg.DispatchLink(nlmsg, link, n.Service)

	default:
		log.Errorf("MasterService: LINK Invalid message. %v", nlmsg)
	}
}

func (n *NLAMasterService) NetlinkAddr(nlmsg *nlamsg.NetlinkMessage, addr *nlamsg.Addr) {
	log.Debugf("MasterService: ADDR")

	switch nlmsg.Type() {
	case syscall.RTM_NEWADDR:
		if deladdr := nladbm.Addrs().Insert(addr); deladdr != nil {
			log.Warnf("MasterService: ADDR duplicate. %v", addr)
			return
		}

	case syscall.RTM_DELADDR:
		if old := nladbm.Addrs().Delete(nladbm.AddrToKey(addr)); old != nil {
			addr.AdId = old.AdId
		}

	default:
		log.Errorf("MasterService: ADDR Invalid message. %v", nlmsg)
		return
	}

	nlamsg.DispatchAddr(nlmsg, addr, n.Service)
}

func (n *NLAMasterService) NetlinkNeigh(nlmsg *nlamsg.NetlinkMessage, neigh *nlamsg.Neigh) {
	log.Debugf("MasterService: NEIG")

	if nlalib.IsInvalidHardwareAddr(neigh.HardwareAddr) {
		nlmsg.Header.Type = syscall.RTM_DELNEIGH
	}

	switch nlmsg.Type() {
	case syscall.RTM_NEWNEIGH:
		if delneigh := nladbm.Neighs().Insert(neigh); delneigh != nil {
			log.Warnf("MasterService: NEIGH duplicate. %v", neigh)
			return
		}

		nlamsg.DispatchNeigh(nlmsg, neigh, n.Service)

		// if neigh is tunnel remote, generate NEWROUTE
		iptunNlMsg := *nlmsg
		iptunNlMsg.Header.Type = syscall.RTM_NEWROUTE
		n.iptun.NewRoutes(neigh, func(iptunRoute *nlamsg.Route) {
			nlamsg.DispatchRoute(&iptunNlMsg, iptunRoute, n.Service)
		})

	case syscall.RTM_DELNEIGH:
		if old := nladbm.Neighs().Delete(nladbm.NeighToKey(neigh)); old != nil {
			neigh.NeId = old.NeId

			delRtNlMsg := *nlmsg
			delRtNlMsg.Header.Type = syscall.RTM_DELROUTE
			nladbm.Mplss().WalkByGwFree(neigh.NId, neigh.IP, func(route *nlamsg.Route) error {
				n.NetlinkRoute(&delRtNlMsg, route)
				return nil

			})
			nladbm.Routes().WalkByGwFree(neigh.NId, neigh.IP, func(route *nlamsg.Route) error {
				n.NetlinkRoute(&delRtNlMsg, route)
				return nil
			})
			n.iptun.NewRoutes(neigh, func(iptunRoute *nlamsg.Route) {
				nlamsg.DispatchRoute(&delRtNlMsg, iptunRoute, n.Service)
			})
		}

		nlamsg.DispatchNeigh(nlmsg, neigh, n.Service)

	default:
		log.Errorf("MasterService: NEIGH Invalid message. %v", nlmsg)
		return
	}
}

func (n *NLAMasterService) NetlinkIPRouteOnMIC(nlmsg *nlamsg.NetlinkMessage, route *nlamsg.Route) {
	log.Debugf("MasterService: ROUTE(IP/MIC) nid=%d", route.NId)

	if route.GetEncap() != nil {
		route.EnIds = []uint32{nladbm.Encaps().EncapId(route.Dst, 0)}
	}

	switch nlmsg.Type() {
	case syscall.RTM_NEWROUTE:
		if old := nladbm.Routes().Insert(route); old != nil {
			// if same route already exists, create DELROUTE message.
			delNlMsg := *nlmsg
			delNlMsg.Header.Type = syscall.RTM_DELROUTE

			// if nexthop is used by vpns, create the vpns DELROUTE messages.
			n.NewVpnRoutes(old, func(vpnRoute *nlamsg.Route) error {
				log.Debugf("MasterService: RTM_DELROUTE(VPN/OLD) %v", vpnRoute)
				nlamsg.DispatchRoute(&delNlMsg, vpnRoute, n.Service)
				return nil
			})

			log.Debugf("MasterService: RTM_DELROUTE(IP/MIC/OLD) %v", old)
			nlamsg.DispatchRoute(&delNlMsg, old, n.Service)

			// if dst is tunnel-remote, generate DELNEIGH
			n.iptun.RemoteRouteDown(old)

			// if route is <dst> dev <tunnel>, generate NEWROUTE(<dst> via <remote>)
			if iptunRoute := n.iptun.NewIptunRoute(old); iptunRoute != nil {
				log.Debugf("MasterService: RTM_DELROUTE(IPTUN/MIC/OLD) %v", iptunRoute)
				nlamsg.DispatchRoute(&delNlMsg, iptunRoute, n.Service)
			}
		}

		log.Debugf("MasterService: RTM_NEWROUTE(IP/MIC) %v", route)
		nlamsg.DispatchRoute(nlmsg, route, n.Service)

		// if nexthop is used by vpns, create the vpns NEWROUTE messages.
		n.NewVpnRoutes(route, func(vpnRoute *nlamsg.Route) error {
			log.Debugf("MasterService: RTM_NEWROUTE(VPN) %v", vpnRoute)
			nlamsg.DispatchRoute(nlmsg, vpnRoute, n.Service)
			return nil
		})

		// if dst is tunnel-remote, generate NEWNEIGH
		n.iptun.RemoteRouteUp(route)

		// if route is <dst> dev <tunnel>, generate NEWROUTE(<dst> via <remote>)
		if iptunRoute := n.iptun.NewIptunRoute(route); iptunRoute != nil {
			log.Debugf("MasterService: RTM_NEWROUTE(IPTUN) %v", iptunRoute)
			nlamsg.DispatchRoute(nlmsg, iptunRoute, n.Service)
		}

	case syscall.RTM_DELROUTE:
		// if nexthop is used by vpns, create the vpns DELROUTE messages.
		n.NewVpnRoutes(route, func(vpnRoute *nlamsg.Route) error {
			log.Debugf("MasterService: RTM_DELROUTE(VPN)%v", vpnRoute)
			nlamsg.DispatchRoute(nlmsg, vpnRoute, n.Service)
			return nil
		})

		log.Debugf("MasterService: RTM_DELROUTE(IP/MIC) %v", route)
		nlamsg.DispatchRoute(nlmsg, route, n.Service)

		nladbm.Routes().Delete(nladbm.RouteToKey(route))

		// if dst is tunnel-remote, generate DELNEIGH
		n.iptun.RemoteRouteDown(route)

		// if route is <dst> dev <tunnel>, generate NEWROUTE(<dst> via <remote>)
		if iptunRoute := n.iptun.NewIptunRoute(route); iptunRoute != nil {
			log.Debugf("MasterService: RTM_DELROUTE(IPTUN) %v", iptunRoute)
			nlamsg.DispatchRoute(nlmsg, iptunRoute, n.Service)
		}

	default:
		log.Errorf("MasterService: ROUTE(IP/MIC) Invalid message. %v %s", nlmsg, route)
	}
}

func (n *NLAMasterService) NetlinkIPRouteOnRIC(nlmsg *nlamsg.NetlinkMessage, route *nlamsg.Route) {
	log.Debugf("MasterService: ROUTE(IP/RIC) nid=%d", route.NId)

	switch nlmsg.Type() {
	case syscall.RTM_NEWROUTE:
		if old := nladbm.Routes().Insert(route); old != nil {
			// if same route already exists, create DELROUTE message.
			delNlMsg := *nlmsg
			delNlMsg.Header.Type = syscall.RTM_DELROUTE

			if vpnRoute := n.NewVpnRoute(old); vpnRoute != nil {
				old = vpnRoute
			}

			log.Debugf("MasterService: RTM_DELROUTE(IP/RIC/OLD) %v", route)
			nlamsg.DispatchRoute(&delNlMsg, old, n.Service)
		}

	case syscall.RTM_DELROUTE:
		nladbm.Routes().Delete(nladbm.RouteToKey(route))

	default:
		log.Errorf("MasterService: ROUTE(IP/RIC) Invalid message. %v", nlmsg)
		return
	}

	// if dst is vpn route, modify nexthop and labels.
	if vpnRoute := n.NewVpnRoute(route); vpnRoute != nil {
		route = vpnRoute
	}

	log.Debugf("MasterService: %s (IP/RIC) %v", nlamsg.NlMsgTypeStr(nlmsg.Type()), route)
	nlamsg.DispatchRoute(nlmsg, route, n.Service)
}

func (n *NLAMasterService) NetlinkMplsRoute(nlmsg *nlamsg.NetlinkMessage, route *nlamsg.Route) {
	log.Debugf("MasterService: ROUTE(MPLS)")

	switch nlmsg.Type() {
	case syscall.RTM_NEWROUTE:
		if nladbm.Mplss().Insert(route) != nil {
			log.Warnf("MasterService: ROUTE(MPLS) duplicate. %v", route)
			return
		}

	case syscall.RTM_DELROUTE:
		nladbm.Mplss().Delete(nladbm.MplsToKey(route))

	default:
		log.Errorf("MasterService: ROUTE(MPLS) Invalid message. %v", nlmsg)
		return
	}

	log.Debugf("MasterService: %s (MPLS) %v", nlamsg.NlMsgTypeStr(nlmsg.Type()), route)
	nlamsg.DispatchRoute(nlmsg, route, n.Service)
}

func (n *NLAMasterService) NetlinkRoute(nlmsg *nlamsg.NetlinkMessage, route *nlamsg.Route) {
	if route.Table != 254 {
		log.Debugf("MasterService: ROUTE(bad table) %v", route)
		return
	}

	switch {
	case route.Dst != nil:
		if route.NId == n.NId {
			n.NetlinkIPRouteOnMIC(nlmsg, route)
		} else {
			n.NetlinkIPRouteOnRIC(nlmsg, route)
		}

	case route.MPLSDst != nil:
		n.NetlinkMplsRoute(nlmsg, route)

	default:
		log.Errorf("MasterService: ROUTE Invalid Dst %v", route)
	}
}

func (n *NLAMasterService) NetlinkNode(nlmsg *nlamsg.NetlinkMessage, node *nlamsg.Node) {
	if nlmsg.Type() == nlalink.RTM_DELNODE {

		nid := nlmsg.NId

		nlmsg.Header.Type = syscall.RTM_DELROUTE
		nladbm.Routes().WalkFree(func(route *nlamsg.Route) error {
			if route.NId == nid {
				n.NetlinkRoute(nlmsg, route)
			}
			return nil
		})

		nlmsg.Header.Type = syscall.RTM_DELNEIGH
		nladbm.Neighs().WalkFree(func(neigh *nlamsg.Neigh) error {
			if neigh.NId == nid {
				n.NetlinkNeigh(nlmsg, neigh)
			}
			return nil
		})

		nlmsg.Header.Type = syscall.RTM_DELADDR
		nladbm.Addrs().WalkFree(func(addr *nlamsg.Addr) error {
			if addr.NId == nid {
				n.NetlinkAddr(nlmsg, addr)
			}
			return nil
		})

		nlmsg.Header.Type = syscall.RTM_DELLINK
		nladbm.Links().WalkFree(func(link *nlamsg.Link) error {
			if link.NId == nid {
				n.NetlinkLink(nlmsg, link)
			}
			return nil
		})

		nlmsg.Header.Type = nlalink.RTM_DELNODE
	}

	nlamsg.DispatchNode(nlmsg, node, n.Service)
}

func (n *NLAMasterService) NetlinkVpn(nlmsg *nlamsg.NetlinkMessage, vpn *nlamsg.Vpn) {
	log.Debugf("MasterService: VPN")

	switch nlmsg.Type() {
	case nlalink.RTM_NEWVPN:
		if delvpn := nladbm.Vpns().Insert(vpn); delvpn != nil {
			log.Warnf("MasterService VPN Updated. %v", vpn)
		}

	case nlalink.RTM_DELVPN:
		nladbm.Vpns().Delete(nladbm.VpnToKey(vpn))

	default:
		log.Errorf("MasterService: VPN Invalid message. %v", nlmsg)
		return
	}

	nlamsg.DispatchVpn(nlmsg, vpn, n.Service)
}
