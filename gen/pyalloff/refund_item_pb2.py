# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: refund_item.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x11refund_item.proto\x12\nrefunditem\x1a\x1bgoogle/protobuf/empty.proto\"\x8e\x01\n\nRefundItem\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x12\n\nrefund_fee\x18\x02 \x01(\x05\x12\x15\n\rrefund_amount\x18\x03 \x01(\x05\x12\x12\n\ncreated_at\x18\x04 \x01(\t\x12\x12\n\nupdated_at\x18\x05 \x01(\t\x12\r\n\x05order\x18\x06 \x01(\x03\x12\x12\n\norder_item\x18\x07 \x01(\x03\"\x17\n\x15RefundItemListRequest\"\'\n\x19RefundItemRetrieveRequest\x12\n\n\x02id\x18\x01 \x01(\x03\x32\xdf\x02\n\x14RefundItemController\x12\x45\n\x04List\x12!.refunditem.RefundItemListRequest\x1a\x16.refunditem.RefundItem\"\x00\x30\x01\x12:\n\x06\x43reate\x12\x16.refunditem.RefundItem\x1a\x16.refunditem.RefundItem\"\x00\x12K\n\x08Retrieve\x12%.refunditem.RefundItemRetrieveRequest\x1a\x16.refunditem.RefundItem\"\x00\x12:\n\x06Update\x12\x16.refunditem.RefundItem\x1a\x16.refunditem.RefundItem\"\x00\x12;\n\x07\x44\x65stroy\x12\x16.refunditem.RefundItem\x1a\x16.google.protobuf.Empty\"\x00\x42\x37Z5github.com/lessbutter/alloff-grpc-protos/gen/goalloffb\x06proto3')



_REFUNDITEM = DESCRIPTOR.message_types_by_name['RefundItem']
_REFUNDITEMLISTREQUEST = DESCRIPTOR.message_types_by_name['RefundItemListRequest']
_REFUNDITEMRETRIEVEREQUEST = DESCRIPTOR.message_types_by_name['RefundItemRetrieveRequest']
RefundItem = _reflection.GeneratedProtocolMessageType('RefundItem', (_message.Message,), {
  'DESCRIPTOR' : _REFUNDITEM,
  '__module__' : 'refund_item_pb2'
  # @@protoc_insertion_point(class_scope:refunditem.RefundItem)
  })
_sym_db.RegisterMessage(RefundItem)

RefundItemListRequest = _reflection.GeneratedProtocolMessageType('RefundItemListRequest', (_message.Message,), {
  'DESCRIPTOR' : _REFUNDITEMLISTREQUEST,
  '__module__' : 'refund_item_pb2'
  # @@protoc_insertion_point(class_scope:refunditem.RefundItemListRequest)
  })
_sym_db.RegisterMessage(RefundItemListRequest)

RefundItemRetrieveRequest = _reflection.GeneratedProtocolMessageType('RefundItemRetrieveRequest', (_message.Message,), {
  'DESCRIPTOR' : _REFUNDITEMRETRIEVEREQUEST,
  '__module__' : 'refund_item_pb2'
  # @@protoc_insertion_point(class_scope:refunditem.RefundItemRetrieveRequest)
  })
_sym_db.RegisterMessage(RefundItemRetrieveRequest)

_REFUNDITEMCONTROLLER = DESCRIPTOR.services_by_name['RefundItemController']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z5github.com/lessbutter/alloff-grpc-protos/gen/goalloff'
  _REFUNDITEM._serialized_start=63
  _REFUNDITEM._serialized_end=205
  _REFUNDITEMLISTREQUEST._serialized_start=207
  _REFUNDITEMLISTREQUEST._serialized_end=230
  _REFUNDITEMRETRIEVEREQUEST._serialized_start=232
  _REFUNDITEMRETRIEVEREQUEST._serialized_end=271
  _REFUNDITEMCONTROLLER._serialized_start=274
  _REFUNDITEMCONTROLLER._serialized_end=625
# @@protoc_insertion_point(module_scope)
