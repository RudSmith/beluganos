# -*- coding: utf-8 -*-

# Copyright (C) 2017 Nippon Telegraph and Telephone Corporation.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
# implied.
# See the License for the specific language governing permissions and
# limitations under the License.

"""
FIB Controller API message.
"""

import StringIO
import struct

_HEADER_FMT = "!HHI"
_HEADER_LEN = 8

def set_fib_header_type(hdr, typ):
    """
    Set type.
    """
    hdr[0] = typ

def get_fib_header_type(hdr):
    """
    Get type.
    """
    return hdr[0]

def set_fib_header_len(hdr, length):
    """
    Set length.
    """
    hdr[1] = length

def get_fib_header_len(hdr):
    """
    Get length.
    """
    return hdr[1]

def set_fib_header_xid(hdr, xid):
    """
    Set XID.
    """
    hdr[2] = xid

def get_fib_header_xid(hdr):
    """
    Get XID.
    """
    return hdr[2]

def unpack_fib_header(data):
    """
    Unpack FIBC API Header. binary->(type, length, xid)
    """
    return struct.unpack(_HEADER_FMT, data)


def pack_fib_header(hdr):
    """
    Pack IBC API Header. (type, length, xid) -> binary
    """
    return struct.pack(
        _HEADER_FMT,
        get_fib_header_type(hdr),
        get_fib_header_len(hdr),
        get_fib_header_xid(hdr)
    )


def read_data(soc, length):
    """
    Read data from socket.
    """
    datas = StringIO.StringIO()
    while length > 0:
        data = soc.recv(length)
        if not data:
            return ''  # connection closed by peer.

        length -= len(data)
        datas.write(data)

    result = datas.getvalue()
    datas.close()
    return result


def read_fib_header(soc):
    """
    Read FIBC API Header from socket.
    """
    data = read_data(soc, _HEADER_LEN)
    if not data:
        return None

    return unpack_fib_header(data)

def read_fib_msg(soc):
    """
    Read FIBC API header and body.
    """
    hdr = read_fib_header(soc)
    if hdr is not None:
        body_len = get_fib_header_len(hdr)
        body = read_data(soc, body_len)
        if len(body) == body_len:
            return hdr, body

    return None, None

def write_fib_msg(soc, mtype, xid, data):
    """
    Write FIBC API header and body.
    """
    soc.sendall(pack_fib_header((mtype, len(data), xid)))
    soc.sendall(data)
