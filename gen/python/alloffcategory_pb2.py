# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: alloffcategory.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x14\x61lloffcategory.proto\x12\x06protos\"~\n\x15\x41lloffCategoryMessage\x12\x13\n\x0b\x63\x61tegory_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x0f\n\x07keyname\x18\x03 \x01(\t\x12\r\n\x05level\x18\x04 \x01(\x05\x12\x11\n\tparent_id\x18\x05 \x01(\t\x12\x0f\n\x07img_url\x18\x06 \x01(\t\".\n\x19ListAlloffCategoryRequest\x12\x11\n\tparent_id\x18\x01 \x01(\t\"O\n\x1aListAlloffCategoryResponse\x12\x31\n\ncategories\x18\x01 \x03(\x0b\x32\x1d.protos.AlloffCategoryMessage2m\n\x0e\x41lloffCategory\x12[\n\x12ListAlloffCategory\x12!.protos.ListAlloffCategoryRequest\x1a\".protos.ListAlloffCategoryResponseB8Z6github.com/lessbutter/alloff-grpc-protos/gen/go/protosb\x06proto3')



_ALLOFFCATEGORYMESSAGE = DESCRIPTOR.message_types_by_name['AlloffCategoryMessage']
_LISTALLOFFCATEGORYREQUEST = DESCRIPTOR.message_types_by_name['ListAlloffCategoryRequest']
_LISTALLOFFCATEGORYRESPONSE = DESCRIPTOR.message_types_by_name['ListAlloffCategoryResponse']
AlloffCategoryMessage = _reflection.GeneratedProtocolMessageType('AlloffCategoryMessage', (_message.Message,), {
  'DESCRIPTOR' : _ALLOFFCATEGORYMESSAGE,
  '__module__' : 'alloffcategory_pb2'
  # @@protoc_insertion_point(class_scope:protos.AlloffCategoryMessage)
  })
_sym_db.RegisterMessage(AlloffCategoryMessage)

ListAlloffCategoryRequest = _reflection.GeneratedProtocolMessageType('ListAlloffCategoryRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTALLOFFCATEGORYREQUEST,
  '__module__' : 'alloffcategory_pb2'
  # @@protoc_insertion_point(class_scope:protos.ListAlloffCategoryRequest)
  })
_sym_db.RegisterMessage(ListAlloffCategoryRequest)

ListAlloffCategoryResponse = _reflection.GeneratedProtocolMessageType('ListAlloffCategoryResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTALLOFFCATEGORYRESPONSE,
  '__module__' : 'alloffcategory_pb2'
  # @@protoc_insertion_point(class_scope:protos.ListAlloffCategoryResponse)
  })
_sym_db.RegisterMessage(ListAlloffCategoryResponse)

_ALLOFFCATEGORY = DESCRIPTOR.services_by_name['AlloffCategory']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z6github.com/lessbutter/alloff-grpc-protos/gen/go/protos'
  _ALLOFFCATEGORYMESSAGE._serialized_start=32
  _ALLOFFCATEGORYMESSAGE._serialized_end=158
  _LISTALLOFFCATEGORYREQUEST._serialized_start=160
  _LISTALLOFFCATEGORYREQUEST._serialized_end=206
  _LISTALLOFFCATEGORYRESPONSE._serialized_start=208
  _LISTALLOFFCATEGORYRESPONSE._serialized_end=287
  _ALLOFFCATEGORY._serialized_start=289
  _ALLOFFCATEGORY._serialized_end=398
# @@protoc_insertion_point(module_scope)
