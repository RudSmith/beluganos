# Install guide
This document shows how to install Beluganos in your systems. Automation scripts are prepared.

## Pre-requirements

### Resources
Generally, network OS is installed into the white-box switches. In Beluganos, you can also build Beluganos into the white-box switches. However, at first, preparing extra server (or VM) and installing here are recommended at first.

- Ubuntu server
	- **Ubuntu 18.04** (18.04-live-server-amd64) is strongly recommended.
	    - If you use Ubuntu 18.04.1 or later, additional settings are required before proceed. Please check **Appendix A** of this document.
	- **Two or more network interfaces** are required.
	- Some LXC instance will be created. More than **12GB HDD** is recommended.
- White-box switches
	- To use OF-DPA mode, **[OF-DPA 2.0](https://github.com/Broadcom-Switch/of-dpa/) switch** and OpenFlow agent are required. OF-DPA application in Edge-core switches is also available at [Edge-core's repository](https://github.com/edge-core/beluganos-forwarding-app).
		- If you don't have OF-DPA switches, any OpenFlow 1.3 switches are acceptable to try Beluganos. In this case, [Lagopus switch](http://www.lagopus.org/) is recommended.
	- To use OpenNSL mode, **[OpenNSL 3.5](https://github.com/Broadcom-Switch/OpenNSL)** switch is required. OpenNSL agent is included in this repository. OpenNSL application in Edge-core switches is also available at [Edge-core's blog](https://support.edge-core.com/hc/en-us/sections/360002115754-OpenNSL).

### LXC settings

If LXC have not configured yet, please set up LXC before starting to build Beluganos. Most of the settings may be default or may be changed if needed. However, please note that following points:

- The new bridge name should be default ( `lxdbr0` ).
- The size of new loop device is depend on number of VRF which you will configure at most. ( Num-of-VRF + 2 ) GB or more is required.

```
$ sudo lxd init
Would you like to use LXD clustering? (yes/no) [default=no]:
Do you want to configure a new storage pool? (yes/no) [default=yes]:
Name of the new storage pool [default=default]:
Name of the storage backend to use (btrfs, dir, lvm) [default=btrfs]:
Create a new BTRFS pool? (yes/no) [default=yes]:
Would you like to use an existing block device? (yes/no) [default=no]:
Size in GB of the new loop device (1GB minimum) [default=15GB]: 8
Would you like to connect to a MAAS server? (yes/no) [default=no]:
Would you like to create a new network bridge? (yes/no) [default=yes]:
What should the new bridge be called? [default=lxdbr0]:
What IPv4 address should be used? (CIDR subnet notation, “auto” or “none”) [default=auto]:
What IPv6 address should be used? (CIDR subnet notation, “auto” or “none”) [default=auto]:
Would you like LXD to be available over the network? (yes/no) [default=no]:
Would you like stale cached images to be updated automatically? (yes/no) [default=yes]
Would you like a YAML "lxd init" preseed to be printed? (yes/no) [default=no]:
```

## 1. Build
Using shell scripts (`create.sh`) is recommended for building Beluganos. Before starting scripts, setting file (`create.ini`) should be edited for your environments. This script will get the required resources including repository of [beluganos/netconf](https://github.com/beluganos/netconf) and [beluganos/go-opennsl](https://github.com/beluganos/go-opennsl) automatically.

```
$ cd ~
$ git clone https://github.com/beluganos/beluganos/ && cd beluganos/
$ vi create.ini

  #
  # Proxy
  #
  # PROXY=http://<ip>:<port>       # (Optional) Comment out if you need internet proxy server

  #
  # Host
  #
  BELUG_OFC_IFACE=ens4             # Set your secure channel interface name connected to switches
  BELUG_OFC_ADDR=172.16.0.55/24    # (Optional) You can change BELUG_OFC_IFACE's IP address and prefix-length if needed
  # ENABLE_VIRTUALENV=yes

$ ./create.sh
```

## 2. Register as a service

Generally, registering Beluganos's main module as a Linux service is recommended.

```
$ cd ~/beluganos
$ make install-service
```

If you will use NETCONF to configure beluganos, following steps are also required.

```
$ cd ~/netconf
$ sudo make install-service
```

## Next steps

You may choose two options.

### Quick start by example
If you want to try our example cases like [case 1 (IP/MPLS router)](example/case1/case1.md) or [case 2 (MPLS-VPN PE router)](example/case2/case2.md), please get back the example documentations.

### Step-by-step procedure
You should register your white-box switches (or OpenFlow switches) to Beluganos's main module. Please refer [setup-guide.md](setup-guide.md) for more details.


## Appendix
### Appendix A. Additional settings at Ubuntu18.04.1 or later

In Ubuntu18.04.01 or later, some settings of apt source are removed. In this case, additional apt source is required to install Beluganos.

```
$ sudo vi /etc/apt/sources.list.d/beluganos.list

deb http://archive.ubuntu.com/ubuntu/ bionic universe
deb http://archive.ubuntu.com/ubuntu/ bionic-updates universe
deb http://archive.ubuntu.com/ubuntu/ bionic multiverse
deb http://archive.ubuntu.com/ubuntu/ bionic-updates multiverse
deb http://security.ubuntu.com/ubuntu bionic-security universe
deb http://security.ubuntu.com/ubuntu bionic-security multiverse

$ sudo apt update
```

### Appendix B. Change the connection settings of white-box switches after installation

If you want to change the white-box switch's settings which specify `BELUG_OFC_IFACE` or `BELUG_OFC_ADDR` at `create.ini` after installation, you can use netplan.

```
$ sudoedit /etc/netplan/02-beluganos.yaml

# -*- coding: utf-8 -*-
network:
  version: 2
  renderer: networkd
  ethernets:
    ens4:  ## <= In case device name was changed
      addresses:
        - 172.16.0.55/24  ## <= In case IP address was changed

```

After editing, to reflect settings, please reboot OS or issue apply command.

```
$ sudo netplan apply
```
