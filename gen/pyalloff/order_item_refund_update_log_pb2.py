# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: order_item_refund_update_log.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\"order_item_refund_update_log.proto\x12\x18orderitemrefundupdatelog\x1a\x1bgoogle/protobuf/empty.proto\"\xb7\x01\n\x18OrderItemRefundUpdateLog\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x11\n\tuser_uuid\x18\x02 \x01(\t\x12\x15\n\ruser_username\x18\x03 \x01(\t\x12\x12\n\nrefund_fee\x18\x04 \x01(\x03\x12\x15\n\rrefund_amount\x18\x05 \x01(\x03\x12\x12\n\ncreated_at\x18\x06 \x01(\t\x12\x12\n\norder_item\x18\x07 \x01(\x03\x12\x12\n\naction_log\x18\x08 \x01(\x03\"%\n#OrderItemRefundUpdateLogListRequest\"5\n\'OrderItemRefundUpdateLogRetrieveRequest\x12\n\n\x02id\x18\x01 \x01(\x03\x32\xea\x04\n\"OrderItemRefundUpdateLogController\x12}\n\x04List\x12=.orderitemrefundupdatelog.OrderItemRefundUpdateLogListRequest\x1a\x32.orderitemrefundupdatelog.OrderItemRefundUpdateLog\"\x00\x30\x01\x12r\n\x06\x43reate\x12\x32.orderitemrefundupdatelog.OrderItemRefundUpdateLog\x1a\x32.orderitemrefundupdatelog.OrderItemRefundUpdateLog\"\x00\x12\x83\x01\n\x08Retrieve\x12\x41.orderitemrefundupdatelog.OrderItemRefundUpdateLogRetrieveRequest\x1a\x32.orderitemrefundupdatelog.OrderItemRefundUpdateLog\"\x00\x12r\n\x06Update\x12\x32.orderitemrefundupdatelog.OrderItemRefundUpdateLog\x1a\x32.orderitemrefundupdatelog.OrderItemRefundUpdateLog\"\x00\x12W\n\x07\x44\x65stroy\x12\x32.orderitemrefundupdatelog.OrderItemRefundUpdateLog\x1a\x16.google.protobuf.Empty\"\x00\x62\x06proto3')



_ORDERITEMREFUNDUPDATELOG = DESCRIPTOR.message_types_by_name['OrderItemRefundUpdateLog']
_ORDERITEMREFUNDUPDATELOGLISTREQUEST = DESCRIPTOR.message_types_by_name['OrderItemRefundUpdateLogListRequest']
_ORDERITEMREFUNDUPDATELOGRETRIEVEREQUEST = DESCRIPTOR.message_types_by_name['OrderItemRefundUpdateLogRetrieveRequest']
OrderItemRefundUpdateLog = _reflection.GeneratedProtocolMessageType('OrderItemRefundUpdateLog', (_message.Message,), {
  'DESCRIPTOR' : _ORDERITEMREFUNDUPDATELOG,
  '__module__' : 'order_item_refund_update_log_pb2'
  # @@protoc_insertion_point(class_scope:orderitemrefundupdatelog.OrderItemRefundUpdateLog)
  })
_sym_db.RegisterMessage(OrderItemRefundUpdateLog)

OrderItemRefundUpdateLogListRequest = _reflection.GeneratedProtocolMessageType('OrderItemRefundUpdateLogListRequest', (_message.Message,), {
  'DESCRIPTOR' : _ORDERITEMREFUNDUPDATELOGLISTREQUEST,
  '__module__' : 'order_item_refund_update_log_pb2'
  # @@protoc_insertion_point(class_scope:orderitemrefundupdatelog.OrderItemRefundUpdateLogListRequest)
  })
_sym_db.RegisterMessage(OrderItemRefundUpdateLogListRequest)

OrderItemRefundUpdateLogRetrieveRequest = _reflection.GeneratedProtocolMessageType('OrderItemRefundUpdateLogRetrieveRequest', (_message.Message,), {
  'DESCRIPTOR' : _ORDERITEMREFUNDUPDATELOGRETRIEVEREQUEST,
  '__module__' : 'order_item_refund_update_log_pb2'
  # @@protoc_insertion_point(class_scope:orderitemrefundupdatelog.OrderItemRefundUpdateLogRetrieveRequest)
  })
_sym_db.RegisterMessage(OrderItemRefundUpdateLogRetrieveRequest)

_ORDERITEMREFUNDUPDATELOGCONTROLLER = DESCRIPTOR.services_by_name['OrderItemRefundUpdateLogController']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  _ORDERITEMREFUNDUPDATELOG._serialized_start=94
  _ORDERITEMREFUNDUPDATELOG._serialized_end=277
  _ORDERITEMREFUNDUPDATELOGLISTREQUEST._serialized_start=279
  _ORDERITEMREFUNDUPDATELOGLISTREQUEST._serialized_end=316
  _ORDERITEMREFUNDUPDATELOGRETRIEVEREQUEST._serialized_start=318
  _ORDERITEMREFUNDUPDATELOGRETRIEVEREQUEST._serialized_end=371
  _ORDERITEMREFUNDUPDATELOGCONTROLLER._serialized_start=374
  _ORDERITEMREFUNDUPDATELOGCONTROLLER._serialized_end=992
# @@protoc_insertion_point(module_scope)
