# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: package_log.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x11package_log.proto\x12\x0bpackage_log\x1a\x1bgoogle/protobuf/empty.proto\"\x99\x01\n\nPackageLog\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x13\n\x0b\x64\x65scription\x18\x02 \x01(\t\x12\x12\n\ncreated_at\x18\x03 \x01(\t\x12\x12\n\nfield_name\x18\x04 \x01(\t\x12\x0e\n\x06\x62\x65\x66ore\x18\x05 \x01(\t\x12\r\n\x05\x61\x66ter\x18\x06 \x01(\t\x12\x12\n\ncreated_by\x18\x07 \x01(\x05\x12\x0f\n\x07package\x18\x08 \x01(\x03\"\x17\n\x15PackageLogListRequest\"\'\n\x19PackageLogRetrieveRequest\x12\n\n\x02id\x18\x01 \x01(\x03\x32\xe8\x02\n\x14PackageLogController\x12G\n\x04List\x12\".package_log.PackageLogListRequest\x1a\x17.package_log.PackageLog\"\x00\x30\x01\x12<\n\x06\x43reate\x12\x17.package_log.PackageLog\x1a\x17.package_log.PackageLog\"\x00\x12M\n\x08Retrieve\x12&.package_log.PackageLogRetrieveRequest\x1a\x17.package_log.PackageLog\"\x00\x12<\n\x06Update\x12\x17.package_log.PackageLog\x1a\x17.package_log.PackageLog\"\x00\x12<\n\x07\x44\x65stroy\x12\x17.package_log.PackageLog\x1a\x16.google.protobuf.Empty\"\x00\x42\x37Z5github.com/lessbutter/alloff-grpc-protos/gen/goalloffb\x06proto3')



_PACKAGELOG = DESCRIPTOR.message_types_by_name['PackageLog']
_PACKAGELOGLISTREQUEST = DESCRIPTOR.message_types_by_name['PackageLogListRequest']
_PACKAGELOGRETRIEVEREQUEST = DESCRIPTOR.message_types_by_name['PackageLogRetrieveRequest']
PackageLog = _reflection.GeneratedProtocolMessageType('PackageLog', (_message.Message,), {
  'DESCRIPTOR' : _PACKAGELOG,
  '__module__' : 'package_log_pb2'
  # @@protoc_insertion_point(class_scope:package_log.PackageLog)
  })
_sym_db.RegisterMessage(PackageLog)

PackageLogListRequest = _reflection.GeneratedProtocolMessageType('PackageLogListRequest', (_message.Message,), {
  'DESCRIPTOR' : _PACKAGELOGLISTREQUEST,
  '__module__' : 'package_log_pb2'
  # @@protoc_insertion_point(class_scope:package_log.PackageLogListRequest)
  })
_sym_db.RegisterMessage(PackageLogListRequest)

PackageLogRetrieveRequest = _reflection.GeneratedProtocolMessageType('PackageLogRetrieveRequest', (_message.Message,), {
  'DESCRIPTOR' : _PACKAGELOGRETRIEVEREQUEST,
  '__module__' : 'package_log_pb2'
  # @@protoc_insertion_point(class_scope:package_log.PackageLogRetrieveRequest)
  })
_sym_db.RegisterMessage(PackageLogRetrieveRequest)

_PACKAGELOGCONTROLLER = DESCRIPTOR.services_by_name['PackageLogController']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z5github.com/lessbutter/alloff-grpc-protos/gen/goalloff'
  _PACKAGELOG._serialized_start=64
  _PACKAGELOG._serialized_end=217
  _PACKAGELOGLISTREQUEST._serialized_start=219
  _PACKAGELOGLISTREQUEST._serialized_end=242
  _PACKAGELOGRETRIEVEREQUEST._serialized_start=244
  _PACKAGELOGRETRIEVEREQUEST._serialized_end=283
  _PACKAGELOGCONTROLLER._serialized_start=286
  _PACKAGELOGCONTROLLER._serialized_end=646
# @@protoc_insertion_point(module_scope)
