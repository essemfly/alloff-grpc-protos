# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: courier.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\rcourier.proto\x12\x07\x63ourier\x1a\x1bgoogle/protobuf/empty.proto\"r\n\x07\x43ourier\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x17\n\x0fsweettracker_id\x18\x03 \x01(\t\x12\x1e\n\x11tracking_url_base\x18\x04 \x01(\tH\x00\x88\x01\x01\x42\x14\n\x12_tracking_url_base\"\x14\n\x12\x43ourierListRequest\"$\n\x16\x43ourierRetrieveRequest\x12\n\n\x02id\x18\x01 \x01(\x03\x32\xa6\x02\n\x11\x43ourierController\x12\x39\n\x04List\x12\x1b.courier.CourierListRequest\x1a\x10.courier.Courier\"\x00\x30\x01\x12.\n\x06\x43reate\x12\x10.courier.Courier\x1a\x10.courier.Courier\"\x00\x12?\n\x08Retrieve\x12\x1f.courier.CourierRetrieveRequest\x1a\x10.courier.Courier\"\x00\x12.\n\x06Update\x12\x10.courier.Courier\x1a\x10.courier.Courier\"\x00\x12\x35\n\x07\x44\x65stroy\x12\x10.courier.Courier\x1a\x16.google.protobuf.Empty\"\x00\x62\x06proto3')



_COURIER = DESCRIPTOR.message_types_by_name['Courier']
_COURIERLISTREQUEST = DESCRIPTOR.message_types_by_name['CourierListRequest']
_COURIERRETRIEVEREQUEST = DESCRIPTOR.message_types_by_name['CourierRetrieveRequest']
Courier = _reflection.GeneratedProtocolMessageType('Courier', (_message.Message,), {
  'DESCRIPTOR' : _COURIER,
  '__module__' : 'courier_pb2'
  # @@protoc_insertion_point(class_scope:courier.Courier)
  })
_sym_db.RegisterMessage(Courier)

CourierListRequest = _reflection.GeneratedProtocolMessageType('CourierListRequest', (_message.Message,), {
  'DESCRIPTOR' : _COURIERLISTREQUEST,
  '__module__' : 'courier_pb2'
  # @@protoc_insertion_point(class_scope:courier.CourierListRequest)
  })
_sym_db.RegisterMessage(CourierListRequest)

CourierRetrieveRequest = _reflection.GeneratedProtocolMessageType('CourierRetrieveRequest', (_message.Message,), {
  'DESCRIPTOR' : _COURIERRETRIEVEREQUEST,
  '__module__' : 'courier_pb2'
  # @@protoc_insertion_point(class_scope:courier.CourierRetrieveRequest)
  })
_sym_db.RegisterMessage(CourierRetrieveRequest)

_COURIERCONTROLLER = DESCRIPTOR.services_by_name['CourierController']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  _COURIER._serialized_start=55
  _COURIER._serialized_end=169
  _COURIERLISTREQUEST._serialized_start=171
  _COURIERLISTREQUEST._serialized_end=191
  _COURIERRETRIEVEREQUEST._serialized_start=193
  _COURIERRETRIEVEREQUEST._serialized_end=229
  _COURIERCONTROLLER._serialized_start=232
  _COURIERCONTROLLER._serialized_end=526
# @@protoc_insertion_point(module_scope)