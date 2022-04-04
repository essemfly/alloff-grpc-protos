# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: order.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from . import payment_pb2 as payment__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0border.proto\x12\x05order\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\rpayment.proto\"\xdd\x03\n\x05Order\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x17\n\x0f\x61lloff_order_id\x18\x02 \x01(\t\x12\x14\n\x0corder_status\x18\x03 \x01(\t\x12\x0f\n\x07user_id\x18\x04 \x01(\t\x12%\n\x04user\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tuser_memo\x18\x06 \x01(\t\x12\x15\n\rproduct_price\x18\x07 \x01(\x05\x12\x16\n\x0e\x64\x65livery_price\x18\x08 \x01(\x05\x12\x13\n\x0btotal_price\x18\t \x01(\x05\x12\x14\n\x0crefund_price\x18\n \x01(\x05\x12\x12\n\ncreated_at\x18\x0b \x01(\t\x12\x12\n\nupdated_at\x18\x0c \x01(\t\x12\x12\n\nordered_at\x18\r \x01(\t\x12!\n\x07payment\x18\x0e \x01(\x0b\x32\x10.payment.Payment\x12\x16\n\x0erecipient_name\x18\x0f \x01(\t\x12\x18\n\x10recipient_mobile\x18\x10 \x01(\t\x12\x1a\n\x12recipient_postcode\x18\x11 \x01(\t\x12\x19\n\x11recipient_address\x18\x12 \x01(\t\x12\x14\n\x0corderer_name\x18\x13 \x01(\t\x12\x16\n\x0eorderer_mobile\x18\x14 \x01(\t\"\x12\n\x10OrderListRequest\"\"\n\x14OrderRetrieveRequest\x12\n\n\x02id\x18\x01 \x01(\x03\x32\x80\x02\n\x0fOrderController\x12\x31\n\x04List\x12\x17.order.OrderListRequest\x1a\x0c.order.Order\"\x00\x30\x01\x12&\n\x06\x43reate\x12\x0c.order.Order\x1a\x0c.order.Order\"\x00\x12\x37\n\x08Retrieve\x12\x1b.order.OrderRetrieveRequest\x1a\x0c.order.Order\"\x00\x12&\n\x06Update\x12\x0c.order.Order\x1a\x0c.order.Order\"\x00\x12\x31\n\x07\x44\x65stroy\x12\x0c.order.Order\x1a\x16.google.protobuf.Empty\"\x00\x42\x37Z5github.com/lessbutter/alloff-grpc-protos/gen/goalloffb\x06proto3')



_ORDER = DESCRIPTOR.message_types_by_name['Order']
_ORDERLISTREQUEST = DESCRIPTOR.message_types_by_name['OrderListRequest']
_ORDERRETRIEVEREQUEST = DESCRIPTOR.message_types_by_name['OrderRetrieveRequest']
Order = _reflection.GeneratedProtocolMessageType('Order', (_message.Message,), {
  'DESCRIPTOR' : _ORDER,
  '__module__' : 'order_pb2'
  # @@protoc_insertion_point(class_scope:order.Order)
  })
_sym_db.RegisterMessage(Order)

OrderListRequest = _reflection.GeneratedProtocolMessageType('OrderListRequest', (_message.Message,), {
  'DESCRIPTOR' : _ORDERLISTREQUEST,
  '__module__' : 'order_pb2'
  # @@protoc_insertion_point(class_scope:order.OrderListRequest)
  })
_sym_db.RegisterMessage(OrderListRequest)

OrderRetrieveRequest = _reflection.GeneratedProtocolMessageType('OrderRetrieveRequest', (_message.Message,), {
  'DESCRIPTOR' : _ORDERRETRIEVEREQUEST,
  '__module__' : 'order_pb2'
  # @@protoc_insertion_point(class_scope:order.OrderRetrieveRequest)
  })
_sym_db.RegisterMessage(OrderRetrieveRequest)

_ORDERCONTROLLER = DESCRIPTOR.services_by_name['OrderController']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z5github.com/lessbutter/alloff-grpc-protos/gen/goalloff'
  _ORDER._serialized_start=97
  _ORDER._serialized_end=574
  _ORDERLISTREQUEST._serialized_start=576
  _ORDERLISTREQUEST._serialized_end=594
  _ORDERRETRIEVEREQUEST._serialized_start=596
  _ORDERRETRIEVEREQUEST._serialized_end=630
  _ORDERCONTROLLER._serialized_start=633
  _ORDERCONTROLLER._serialized_end=889
# @@protoc_insertion_point(module_scope)
