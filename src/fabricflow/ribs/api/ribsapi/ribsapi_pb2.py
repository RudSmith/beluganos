# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ribsapi.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='ribsapi.proto',
  package='ribsapi',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\rribsapi.proto\x12\x07ribsapi\"%\n\tRibUpdate\x12\n\n\x02rt\x18\x01 \x01(\t\x12\x0c\n\x04path\x18\x02 \x01(\x0c\"\r\n\x0bModRibReply\"-\n\x11MonitorRibRequest\x12\n\n\x02rt\x18\x01 \x01(\t\x12\x0c\n\x04n_id\x18\x02 \x01(\r\"\x1c\n\x0eSyncRibRequest\x12\n\n\x02rt\x18\x01 \x01(\t\"\x0e\n\x0cSyncRibReply\"\x14\n\x12GetNexthopsRequest\"C\n\x07Nexthop\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\n\n\x02rt\x18\x02 \x01(\t\x12\x0c\n\x04\x61\x64\x64r\x18\x03 \x01(\t\x12\x11\n\tsource_id\x18\x04 \x01(\t\"\x10\n\x0eGetRicsRequest\"1\n\x08RicEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\x0c\n\x04n_id\x18\x02 \x01(\r\x12\n\n\x02rt\x18\x03 \x01(\t\"\x11\n\x0fGetIPMapRequest\"(\n\nIPMapEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t2\xc2\x01\n\x0bRIBSCoreApi\x12\x34\n\x06ModRib\x12\x12.ribsapi.RibUpdate\x1a\x14.ribsapi.ModRibReply\"\x00\x12@\n\nMonitorRib\x12\x1a.ribsapi.MonitorRibRequest\x1a\x12.ribsapi.RibUpdate\"\x00\x30\x01\x12;\n\x07SyncRib\x12\x17.ribsapi.SyncRibRequest\x1a\x15.ribsapi.SyncRibReply\"\x00\x32\xca\x01\n\x07RIBSApi\x12@\n\x0bGetNexthops\x12\x1b.ribsapi.GetNexthopsRequest\x1a\x10.ribsapi.Nexthop\"\x00\x30\x01\x12\x39\n\x07GetRics\x12\x17.ribsapi.GetRicsRequest\x1a\x11.ribsapi.RicEntry\"\x00\x30\x01\x12\x42\n\rGetNexthopMap\x12\x18.ribsapi.GetIPMapRequest\x1a\x13.ribsapi.IPMapEntry\"\x00\x30\x01\x62\x06proto3')
)




_RIBUPDATE = _descriptor.Descriptor(
  name='RibUpdate',
  full_name='ribsapi.RibUpdate',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='rt', full_name='ribsapi.RibUpdate.rt', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='path', full_name='ribsapi.RibUpdate.path', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=26,
  serialized_end=63,
)


_MODRIBREPLY = _descriptor.Descriptor(
  name='ModRibReply',
  full_name='ribsapi.ModRibReply',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=65,
  serialized_end=78,
)


_MONITORRIBREQUEST = _descriptor.Descriptor(
  name='MonitorRibRequest',
  full_name='ribsapi.MonitorRibRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='rt', full_name='ribsapi.MonitorRibRequest.rt', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='n_id', full_name='ribsapi.MonitorRibRequest.n_id', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=80,
  serialized_end=125,
)


_SYNCRIBREQUEST = _descriptor.Descriptor(
  name='SyncRibRequest',
  full_name='ribsapi.SyncRibRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='rt', full_name='ribsapi.SyncRibRequest.rt', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=127,
  serialized_end=155,
)


_SYNCRIBREPLY = _descriptor.Descriptor(
  name='SyncRibReply',
  full_name='ribsapi.SyncRibReply',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=157,
  serialized_end=171,
)


_GETNEXTHOPSREQUEST = _descriptor.Descriptor(
  name='GetNexthopsRequest',
  full_name='ribsapi.GetNexthopsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=173,
  serialized_end=193,
)


_NEXTHOP = _descriptor.Descriptor(
  name='Nexthop',
  full_name='ribsapi.Nexthop',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='ribsapi.Nexthop.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='rt', full_name='ribsapi.Nexthop.rt', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='addr', full_name='ribsapi.Nexthop.addr', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='source_id', full_name='ribsapi.Nexthop.source_id', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=195,
  serialized_end=262,
)


_GETRICSREQUEST = _descriptor.Descriptor(
  name='GetRicsRequest',
  full_name='ribsapi.GetRicsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=264,
  serialized_end=280,
)


_RICENTRY = _descriptor.Descriptor(
  name='RicEntry',
  full_name='ribsapi.RicEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='ribsapi.RicEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='n_id', full_name='ribsapi.RicEntry.n_id', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='rt', full_name='ribsapi.RicEntry.rt', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=282,
  serialized_end=331,
)


_GETIPMAPREQUEST = _descriptor.Descriptor(
  name='GetIPMapRequest',
  full_name='ribsapi.GetIPMapRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=333,
  serialized_end=350,
)


_IPMAPENTRY = _descriptor.Descriptor(
  name='IPMapEntry',
  full_name='ribsapi.IPMapEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='ribsapi.IPMapEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='ribsapi.IPMapEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=352,
  serialized_end=392,
)

DESCRIPTOR.message_types_by_name['RibUpdate'] = _RIBUPDATE
DESCRIPTOR.message_types_by_name['ModRibReply'] = _MODRIBREPLY
DESCRIPTOR.message_types_by_name['MonitorRibRequest'] = _MONITORRIBREQUEST
DESCRIPTOR.message_types_by_name['SyncRibRequest'] = _SYNCRIBREQUEST
DESCRIPTOR.message_types_by_name['SyncRibReply'] = _SYNCRIBREPLY
DESCRIPTOR.message_types_by_name['GetNexthopsRequest'] = _GETNEXTHOPSREQUEST
DESCRIPTOR.message_types_by_name['Nexthop'] = _NEXTHOP
DESCRIPTOR.message_types_by_name['GetRicsRequest'] = _GETRICSREQUEST
DESCRIPTOR.message_types_by_name['RicEntry'] = _RICENTRY
DESCRIPTOR.message_types_by_name['GetIPMapRequest'] = _GETIPMAPREQUEST
DESCRIPTOR.message_types_by_name['IPMapEntry'] = _IPMAPENTRY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

RibUpdate = _reflection.GeneratedProtocolMessageType('RibUpdate', (_message.Message,), {
  'DESCRIPTOR' : _RIBUPDATE,
  '__module__' : 'ribsapi_pb2'
  # @@protoc_insertion_point(class_scope:ribsapi.RibUpdate)
  })
_sym_db.RegisterMessage(RibUpdate)

ModRibReply = _reflection.GeneratedProtocolMessageType('ModRibReply', (_message.Message,), {
  'DESCRIPTOR' : _MODRIBREPLY,
  '__module__' : 'ribsapi_pb2'
  # @@protoc_insertion_point(class_scope:ribsapi.ModRibReply)
  })
_sym_db.RegisterMessage(ModRibReply)

MonitorRibRequest = _reflection.GeneratedProtocolMessageType('MonitorRibRequest', (_message.Message,), {
  'DESCRIPTOR' : _MONITORRIBREQUEST,
  '__module__' : 'ribsapi_pb2'
  # @@protoc_insertion_point(class_scope:ribsapi.MonitorRibRequest)
  })
_sym_db.RegisterMessage(MonitorRibRequest)

SyncRibRequest = _reflection.GeneratedProtocolMessageType('SyncRibRequest', (_message.Message,), {
  'DESCRIPTOR' : _SYNCRIBREQUEST,
  '__module__' : 'ribsapi_pb2'
  # @@protoc_insertion_point(class_scope:ribsapi.SyncRibRequest)
  })
_sym_db.RegisterMessage(SyncRibRequest)

SyncRibReply = _reflection.GeneratedProtocolMessageType('SyncRibReply', (_message.Message,), {
  'DESCRIPTOR' : _SYNCRIBREPLY,
  '__module__' : 'ribsapi_pb2'
  # @@protoc_insertion_point(class_scope:ribsapi.SyncRibReply)
  })
_sym_db.RegisterMessage(SyncRibReply)

GetNexthopsRequest = _reflection.GeneratedProtocolMessageType('GetNexthopsRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETNEXTHOPSREQUEST,
  '__module__' : 'ribsapi_pb2'
  # @@protoc_insertion_point(class_scope:ribsapi.GetNexthopsRequest)
  })
_sym_db.RegisterMessage(GetNexthopsRequest)

Nexthop = _reflection.GeneratedProtocolMessageType('Nexthop', (_message.Message,), {
  'DESCRIPTOR' : _NEXTHOP,
  '__module__' : 'ribsapi_pb2'
  # @@protoc_insertion_point(class_scope:ribsapi.Nexthop)
  })
_sym_db.RegisterMessage(Nexthop)

GetRicsRequest = _reflection.GeneratedProtocolMessageType('GetRicsRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETRICSREQUEST,
  '__module__' : 'ribsapi_pb2'
  # @@protoc_insertion_point(class_scope:ribsapi.GetRicsRequest)
  })
_sym_db.RegisterMessage(GetRicsRequest)

RicEntry = _reflection.GeneratedProtocolMessageType('RicEntry', (_message.Message,), {
  'DESCRIPTOR' : _RICENTRY,
  '__module__' : 'ribsapi_pb2'
  # @@protoc_insertion_point(class_scope:ribsapi.RicEntry)
  })
_sym_db.RegisterMessage(RicEntry)

GetIPMapRequest = _reflection.GeneratedProtocolMessageType('GetIPMapRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETIPMAPREQUEST,
  '__module__' : 'ribsapi_pb2'
  # @@protoc_insertion_point(class_scope:ribsapi.GetIPMapRequest)
  })
_sym_db.RegisterMessage(GetIPMapRequest)

IPMapEntry = _reflection.GeneratedProtocolMessageType('IPMapEntry', (_message.Message,), {
  'DESCRIPTOR' : _IPMAPENTRY,
  '__module__' : 'ribsapi_pb2'
  # @@protoc_insertion_point(class_scope:ribsapi.IPMapEntry)
  })
_sym_db.RegisterMessage(IPMapEntry)



_RIBSCOREAPI = _descriptor.ServiceDescriptor(
  name='RIBSCoreApi',
  full_name='ribsapi.RIBSCoreApi',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=395,
  serialized_end=589,
  methods=[
  _descriptor.MethodDescriptor(
    name='ModRib',
    full_name='ribsapi.RIBSCoreApi.ModRib',
    index=0,
    containing_service=None,
    input_type=_RIBUPDATE,
    output_type=_MODRIBREPLY,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='MonitorRib',
    full_name='ribsapi.RIBSCoreApi.MonitorRib',
    index=1,
    containing_service=None,
    input_type=_MONITORRIBREQUEST,
    output_type=_RIBUPDATE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='SyncRib',
    full_name='ribsapi.RIBSCoreApi.SyncRib',
    index=2,
    containing_service=None,
    input_type=_SYNCRIBREQUEST,
    output_type=_SYNCRIBREPLY,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_RIBSCOREAPI)

DESCRIPTOR.services_by_name['RIBSCoreApi'] = _RIBSCOREAPI


_RIBSAPI = _descriptor.ServiceDescriptor(
  name='RIBSApi',
  full_name='ribsapi.RIBSApi',
  file=DESCRIPTOR,
  index=1,
  serialized_options=None,
  serialized_start=592,
  serialized_end=794,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetNexthops',
    full_name='ribsapi.RIBSApi.GetNexthops',
    index=0,
    containing_service=None,
    input_type=_GETNEXTHOPSREQUEST,
    output_type=_NEXTHOP,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetRics',
    full_name='ribsapi.RIBSApi.GetRics',
    index=1,
    containing_service=None,
    input_type=_GETRICSREQUEST,
    output_type=_RICENTRY,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetNexthopMap',
    full_name='ribsapi.RIBSApi.GetNexthopMap',
    index=2,
    containing_service=None,
    input_type=_GETIPMAPREQUEST,
    output_type=_IPMAPENTRY,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_RIBSAPI)

DESCRIPTOR.services_by_name['RIBSApi'] = _RIBSAPI

# @@protoc_insertion_point(module_scope)
