# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: shipping_notice_log.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x19shipping_notice_log.proto\x12\x11shippingnoticelog\x1a\x1bgoogle/protobuf/empty.proto\"\xb5\x01\n\x11ShippingNoticeLog\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x11\n\tuser_uuid\x18\x02 \x01(\t\x12\x15\n\ruser_username\x18\x03 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x04 \x01(\t\x12\x12\n\ncreated_at\x18\x05 \x01(\t\x12\x12\n\nfield_name\x18\x06 \x01(\t\x12\x0e\n\x06\x62\x65\x66ore\x18\x07 \x01(\t\x12\r\n\x05\x61\x66ter\x18\x08 \x01(\t\x12\x0e\n\x06notice\x18\t \x01(\x03\"\x1e\n\x1cShippingNoticeLogListRequest\".\n ShippingNoticeLogRetrieveRequest\x12\n\n\x02id\x18\x01 \x01(\x03\x32\xe4\x03\n\x1bShippingNoticeLogController\x12\x61\n\x04List\x12/.shippingnoticelog.ShippingNoticeLogListRequest\x1a$.shippingnoticelog.ShippingNoticeLog\"\x00\x30\x01\x12V\n\x06\x43reate\x12$.shippingnoticelog.ShippingNoticeLog\x1a$.shippingnoticelog.ShippingNoticeLog\"\x00\x12g\n\x08Retrieve\x12\x33.shippingnoticelog.ShippingNoticeLogRetrieveRequest\x1a$.shippingnoticelog.ShippingNoticeLog\"\x00\x12V\n\x06Update\x12$.shippingnoticelog.ShippingNoticeLog\x1a$.shippingnoticelog.ShippingNoticeLog\"\x00\x12I\n\x07\x44\x65stroy\x12$.shippingnoticelog.ShippingNoticeLog\x1a\x16.google.protobuf.Empty\"\x00\x62\x06proto3')



_SHIPPINGNOTICELOG = DESCRIPTOR.message_types_by_name['ShippingNoticeLog']
_SHIPPINGNOTICELOGLISTREQUEST = DESCRIPTOR.message_types_by_name['ShippingNoticeLogListRequest']
_SHIPPINGNOTICELOGRETRIEVEREQUEST = DESCRIPTOR.message_types_by_name['ShippingNoticeLogRetrieveRequest']
ShippingNoticeLog = _reflection.GeneratedProtocolMessageType('ShippingNoticeLog', (_message.Message,), {
  'DESCRIPTOR' : _SHIPPINGNOTICELOG,
  '__module__' : 'shipping_notice_log_pb2'
  # @@protoc_insertion_point(class_scope:shippingnoticelog.ShippingNoticeLog)
  })
_sym_db.RegisterMessage(ShippingNoticeLog)

ShippingNoticeLogListRequest = _reflection.GeneratedProtocolMessageType('ShippingNoticeLogListRequest', (_message.Message,), {
  'DESCRIPTOR' : _SHIPPINGNOTICELOGLISTREQUEST,
  '__module__' : 'shipping_notice_log_pb2'
  # @@protoc_insertion_point(class_scope:shippingnoticelog.ShippingNoticeLogListRequest)
  })
_sym_db.RegisterMessage(ShippingNoticeLogListRequest)

ShippingNoticeLogRetrieveRequest = _reflection.GeneratedProtocolMessageType('ShippingNoticeLogRetrieveRequest', (_message.Message,), {
  'DESCRIPTOR' : _SHIPPINGNOTICELOGRETRIEVEREQUEST,
  '__module__' : 'shipping_notice_log_pb2'
  # @@protoc_insertion_point(class_scope:shippingnoticelog.ShippingNoticeLogRetrieveRequest)
  })
_sym_db.RegisterMessage(ShippingNoticeLogRetrieveRequest)

_SHIPPINGNOTICELOGCONTROLLER = DESCRIPTOR.services_by_name['ShippingNoticeLogController']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  _SHIPPINGNOTICELOG._serialized_start=78
  _SHIPPINGNOTICELOG._serialized_end=259
  _SHIPPINGNOTICELOGLISTREQUEST._serialized_start=261
  _SHIPPINGNOTICELOGLISTREQUEST._serialized_end=291
  _SHIPPINGNOTICELOGRETRIEVEREQUEST._serialized_start=293
  _SHIPPINGNOTICELOGRETRIEVEREQUEST._serialized_end=339
  _SHIPPINGNOTICELOGCONTROLLER._serialized_start=342
  _SHIPPINGNOTICELOGCONTROLLER._serialized_end=826
# @@protoc_insertion_point(module_scope)
