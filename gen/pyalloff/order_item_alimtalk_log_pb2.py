# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: order_item_alimtalk_log.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1dorder_item_alimtalk_log.proto\x12\x14orderitemalimtalklog\x1a\x1bgoogle/protobuf/empty.proto\"\xb3\x01\n\x14OrderItemAlimtalkLog\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x11\n\tuser_uuid\x18\x02 \x01(\t\x12\x15\n\ruser_username\x18\x03 \x01(\t\x12\x15\n\ralimtalk_type\x18\x04 \x01(\t\x12\x12\n\nrequest_id\x18\x05 \x01(\t\x12\x12\n\ncreated_at\x18\x06 \x01(\t\x12\x12\n\norder_item\x18\x07 \x01(\x03\x12\x12\n\naction_log\x18\x08 \x01(\x03\"!\n\x1fOrderItemAlimtalkLogListRequest\"1\n#OrderItemAlimtalkLogRetrieveRequest\x12\n\n\x02id\x18\x01 \x01(\x03\x32\x9d\x04\n\x1eOrderItemAlimtalkLogController\x12m\n\x04List\x12\x35.orderitemalimtalklog.OrderItemAlimtalkLogListRequest\x1a*.orderitemalimtalklog.OrderItemAlimtalkLog\"\x00\x30\x01\x12\x62\n\x06\x43reate\x12*.orderitemalimtalklog.OrderItemAlimtalkLog\x1a*.orderitemalimtalklog.OrderItemAlimtalkLog\"\x00\x12s\n\x08Retrieve\x12\x39.orderitemalimtalklog.OrderItemAlimtalkLogRetrieveRequest\x1a*.orderitemalimtalklog.OrderItemAlimtalkLog\"\x00\x12\x62\n\x06Update\x12*.orderitemalimtalklog.OrderItemAlimtalkLog\x1a*.orderitemalimtalklog.OrderItemAlimtalkLog\"\x00\x12O\n\x07\x44\x65stroy\x12*.orderitemalimtalklog.OrderItemAlimtalkLog\x1a\x16.google.protobuf.Empty\"\x00\x42\x37Z5github.com/lessbutter/alloff-grpc-protos/gen/goalloffb\x06proto3')



_ORDERITEMALIMTALKLOG = DESCRIPTOR.message_types_by_name['OrderItemAlimtalkLog']
_ORDERITEMALIMTALKLOGLISTREQUEST = DESCRIPTOR.message_types_by_name['OrderItemAlimtalkLogListRequest']
_ORDERITEMALIMTALKLOGRETRIEVEREQUEST = DESCRIPTOR.message_types_by_name['OrderItemAlimtalkLogRetrieveRequest']
OrderItemAlimtalkLog = _reflection.GeneratedProtocolMessageType('OrderItemAlimtalkLog', (_message.Message,), {
  'DESCRIPTOR' : _ORDERITEMALIMTALKLOG,
  '__module__' : 'order_item_alimtalk_log_pb2'
  # @@protoc_insertion_point(class_scope:orderitemalimtalklog.OrderItemAlimtalkLog)
  })
_sym_db.RegisterMessage(OrderItemAlimtalkLog)

OrderItemAlimtalkLogListRequest = _reflection.GeneratedProtocolMessageType('OrderItemAlimtalkLogListRequest', (_message.Message,), {
  'DESCRIPTOR' : _ORDERITEMALIMTALKLOGLISTREQUEST,
  '__module__' : 'order_item_alimtalk_log_pb2'
  # @@protoc_insertion_point(class_scope:orderitemalimtalklog.OrderItemAlimtalkLogListRequest)
  })
_sym_db.RegisterMessage(OrderItemAlimtalkLogListRequest)

OrderItemAlimtalkLogRetrieveRequest = _reflection.GeneratedProtocolMessageType('OrderItemAlimtalkLogRetrieveRequest', (_message.Message,), {
  'DESCRIPTOR' : _ORDERITEMALIMTALKLOGRETRIEVEREQUEST,
  '__module__' : 'order_item_alimtalk_log_pb2'
  # @@protoc_insertion_point(class_scope:orderitemalimtalklog.OrderItemAlimtalkLogRetrieveRequest)
  })
_sym_db.RegisterMessage(OrderItemAlimtalkLogRetrieveRequest)

_ORDERITEMALIMTALKLOGCONTROLLER = DESCRIPTOR.services_by_name['OrderItemAlimtalkLogController']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z5github.com/lessbutter/alloff-grpc-protos/gen/goalloff'
  _ORDERITEMALIMTALKLOG._serialized_start=85
  _ORDERITEMALIMTALKLOG._serialized_end=264
  _ORDERITEMALIMTALKLOGLISTREQUEST._serialized_start=266
  _ORDERITEMALIMTALKLOGLISTREQUEST._serialized_end=299
  _ORDERITEMALIMTALKLOGRETRIEVEREQUEST._serialized_start=301
  _ORDERITEMALIMTALKLOGRETRIEVEREQUEST._serialized_end=350
  _ORDERITEMALIMTALKLOGCONTROLLER._serialized_start=353
  _ORDERITEMALIMTALKLOGCONTROLLER._serialized_end=894
# @@protoc_insertion_point(module_scope)
