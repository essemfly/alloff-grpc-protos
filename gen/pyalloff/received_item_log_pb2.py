# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: received_item_log.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x17received_item_log.proto\x12\x0freceiveditemlog\x1a\x1bgoogle/protobuf/empty.proto\"\xba\x01\n\x0fReceivedItemLog\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x11\n\tuser_uuid\x18\x02 \x01(\t\x12\x15\n\ruser_username\x18\x03 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x04 \x01(\t\x12\x12\n\ncreated_at\x18\x05 \x01(\t\x12\x12\n\nfield_name\x18\x06 \x01(\t\x12\x0e\n\x06\x62\x65\x66ore\x18\x07 \x01(\t\x12\r\n\x05\x61\x66ter\x18\x08 \x01(\t\x12\x15\n\rreceived_item\x18\t \x01(\x03\"\x1c\n\x1aReceivedItemLogListRequest\",\n\x1eReceivedItemLogRetrieveRequest\x12\n\n\x02id\x18\x01 \x01(\x03\x32\xbe\x03\n\x19ReceivedItemLogController\x12Y\n\x04List\x12+.receiveditemlog.ReceivedItemLogListRequest\x1a .receiveditemlog.ReceivedItemLog\"\x00\x30\x01\x12N\n\x06\x43reate\x12 .receiveditemlog.ReceivedItemLog\x1a .receiveditemlog.ReceivedItemLog\"\x00\x12_\n\x08Retrieve\x12/.receiveditemlog.ReceivedItemLogRetrieveRequest\x1a .receiveditemlog.ReceivedItemLog\"\x00\x12N\n\x06Update\x12 .receiveditemlog.ReceivedItemLog\x1a .receiveditemlog.ReceivedItemLog\"\x00\x12\x45\n\x07\x44\x65stroy\x12 .receiveditemlog.ReceivedItemLog\x1a\x16.google.protobuf.Empty\"\x00\x42\x37Z5github.com/lessbutter/alloff-grpc-protos/gen/goalloffb\x06proto3')



_RECEIVEDITEMLOG = DESCRIPTOR.message_types_by_name['ReceivedItemLog']
_RECEIVEDITEMLOGLISTREQUEST = DESCRIPTOR.message_types_by_name['ReceivedItemLogListRequest']
_RECEIVEDITEMLOGRETRIEVEREQUEST = DESCRIPTOR.message_types_by_name['ReceivedItemLogRetrieveRequest']
ReceivedItemLog = _reflection.GeneratedProtocolMessageType('ReceivedItemLog', (_message.Message,), {
  'DESCRIPTOR' : _RECEIVEDITEMLOG,
  '__module__' : 'received_item_log_pb2'
  # @@protoc_insertion_point(class_scope:receiveditemlog.ReceivedItemLog)
  })
_sym_db.RegisterMessage(ReceivedItemLog)

ReceivedItemLogListRequest = _reflection.GeneratedProtocolMessageType('ReceivedItemLogListRequest', (_message.Message,), {
  'DESCRIPTOR' : _RECEIVEDITEMLOGLISTREQUEST,
  '__module__' : 'received_item_log_pb2'
  # @@protoc_insertion_point(class_scope:receiveditemlog.ReceivedItemLogListRequest)
  })
_sym_db.RegisterMessage(ReceivedItemLogListRequest)

ReceivedItemLogRetrieveRequest = _reflection.GeneratedProtocolMessageType('ReceivedItemLogRetrieveRequest', (_message.Message,), {
  'DESCRIPTOR' : _RECEIVEDITEMLOGRETRIEVEREQUEST,
  '__module__' : 'received_item_log_pb2'
  # @@protoc_insertion_point(class_scope:receiveditemlog.ReceivedItemLogRetrieveRequest)
  })
_sym_db.RegisterMessage(ReceivedItemLogRetrieveRequest)

_RECEIVEDITEMLOGCONTROLLER = DESCRIPTOR.services_by_name['ReceivedItemLogController']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z5github.com/lessbutter/alloff-grpc-protos/gen/goalloff'
  _RECEIVEDITEMLOG._serialized_start=74
  _RECEIVEDITEMLOG._serialized_end=260
  _RECEIVEDITEMLOGLISTREQUEST._serialized_start=262
  _RECEIVEDITEMLOGLISTREQUEST._serialized_end=290
  _RECEIVEDITEMLOGRETRIEVEREQUEST._serialized_start=292
  _RECEIVEDITEMLOGRETRIEVEREQUEST._serialized_end=336
  _RECEIVEDITEMLOGCONTROLLER._serialized_start=339
  _RECEIVEDITEMLOGCONTROLLER._serialized_end=785
# @@protoc_insertion_point(module_scope)
