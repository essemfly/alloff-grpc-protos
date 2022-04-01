# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: order_item_action_log.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from . import order_item_status_change_log_pb2 as order__item__status__change__log__pb2
from . import order_item_refund_update_log_pb2 as order__item__refund__update__log__pb2
from . import order_item_alimtalk_log_pb2 as order__item__alimtalk__log__pb2
from . import inventory_receipt_log_pb2 as inventory__receipt__log__pb2
from . import received_item_generation_log_pb2 as received__item__generation__log__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1border_item_action_log.proto\x12\x12orderitemactionlog\x1a\x1bgoogle/protobuf/empty.proto\x1a\"order_item_status_change_log.proto\x1a\"order_item_refund_update_log.proto\x1a\x1dorder_item_alimtalk_log.proto\x1a\x1binventory_receipt_log.proto\x1a\"received_item_generation_log.proto\"\xdf\x04\n\x12OrderItemActionLog\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x11\n\tuser_uuid\x18\x02 \x01(\t\x12\x15\n\ruser_username\x18\x03 \x01(\t\x12\x0e\n\x06\x64\x65tail\x18\x04 \x01(\t\x12\x13\n\x0b\x61\x63tion_type\x18\x05 \x01(\t\x12\x12\n\ncreated_at\x18\x06 \x01(\t\x12\x12\n\norder_item\x18\x07 \x01(\x03\x12N\n\rstatus_change\x18\x08 \x01(\x0b\x32\x32.orderitemstatuschangelog.OrderItemStatusChangeLogH\x00\x88\x01\x01\x12N\n\rrefund_update\x18\t \x01(\x0b\x32\x32.orderitemrefundupdatelog.OrderItemRefundUpdateLogH\x01\x88\x01\x01\x12\x41\n\x08\x61limtalk\x18\n \x01(\x0b\x32*.orderitemalimtalklog.OrderItemAlimtalkLogH\x02\x88\x01\x01\x12@\n\tinventory\x18\x0b \x01(\x0b\x32(.inventoryreceiptlog.InventoryReceiptLogH\x03\x88\x01\x01\x12P\n\rreceived_item\x18\x0c \x01(\x0b\x32\x34.receiveditemgenerationlog.ReceivedItemGenerationLogH\x04\x88\x01\x01\x42\x10\n\x0e_status_changeB\x10\n\x0e_refund_updateB\x0b\n\t_alimtalkB\x0c\n\n_inventoryB\x10\n\x0e_received_item\"\x1f\n\x1dOrderItemActionLogListRequest\"/\n!OrderItemActionLogRetrieveRequest\x12\n\n\x02id\x18\x01 \x01(\x03\x32\xf7\x03\n\x1cOrderItemActionLogController\x12\x65\n\x04List\x12\x31.orderitemactionlog.OrderItemActionLogListRequest\x1a&.orderitemactionlog.OrderItemActionLog\"\x00\x30\x01\x12Z\n\x06\x43reate\x12&.orderitemactionlog.OrderItemActionLog\x1a&.orderitemactionlog.OrderItemActionLog\"\x00\x12k\n\x08Retrieve\x12\x35.orderitemactionlog.OrderItemActionLogRetrieveRequest\x1a&.orderitemactionlog.OrderItemActionLog\"\x00\x12Z\n\x06Update\x12&.orderitemactionlog.OrderItemActionLog\x1a&.orderitemactionlog.OrderItemActionLog\"\x00\x12K\n\x07\x44\x65stroy\x12&.orderitemactionlog.OrderItemActionLog\x1a\x16.google.protobuf.Empty\"\x00\x62\x06proto3')



_ORDERITEMACTIONLOG = DESCRIPTOR.message_types_by_name['OrderItemActionLog']
_ORDERITEMACTIONLOGLISTREQUEST = DESCRIPTOR.message_types_by_name['OrderItemActionLogListRequest']
_ORDERITEMACTIONLOGRETRIEVEREQUEST = DESCRIPTOR.message_types_by_name['OrderItemActionLogRetrieveRequest']
OrderItemActionLog = _reflection.GeneratedProtocolMessageType('OrderItemActionLog', (_message.Message,), {
  'DESCRIPTOR' : _ORDERITEMACTIONLOG,
  '__module__' : 'order_item_action_log_pb2'
  # @@protoc_insertion_point(class_scope:orderitemactionlog.OrderItemActionLog)
  })
_sym_db.RegisterMessage(OrderItemActionLog)

OrderItemActionLogListRequest = _reflection.GeneratedProtocolMessageType('OrderItemActionLogListRequest', (_message.Message,), {
  'DESCRIPTOR' : _ORDERITEMACTIONLOGLISTREQUEST,
  '__module__' : 'order_item_action_log_pb2'
  # @@protoc_insertion_point(class_scope:orderitemactionlog.OrderItemActionLogListRequest)
  })
_sym_db.RegisterMessage(OrderItemActionLogListRequest)

OrderItemActionLogRetrieveRequest = _reflection.GeneratedProtocolMessageType('OrderItemActionLogRetrieveRequest', (_message.Message,), {
  'DESCRIPTOR' : _ORDERITEMACTIONLOGRETRIEVEREQUEST,
  '__module__' : 'order_item_action_log_pb2'
  # @@protoc_insertion_point(class_scope:orderitemactionlog.OrderItemActionLogRetrieveRequest)
  })
_sym_db.RegisterMessage(OrderItemActionLogRetrieveRequest)

_ORDERITEMACTIONLOGCONTROLLER = DESCRIPTOR.services_by_name['OrderItemActionLogController']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  _ORDERITEMACTIONLOG._serialized_start=249
  _ORDERITEMACTIONLOG._serialized_end=856
  _ORDERITEMACTIONLOGLISTREQUEST._serialized_start=858
  _ORDERITEMACTIONLOGLISTREQUEST._serialized_end=889
  _ORDERITEMACTIONLOGRETRIEVEREQUEST._serialized_start=891
  _ORDERITEMACTIONLOGRETRIEVEREQUEST._serialized_end=938
  _ORDERITEMACTIONLOGCONTROLLER._serialized_start=941
  _ORDERITEMACTIONLOGCONTROLLER._serialized_end=1444
# @@protoc_insertion_point(module_scope)
