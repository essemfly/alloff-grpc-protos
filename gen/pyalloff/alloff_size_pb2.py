# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: alloff_size.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from . import alloffcategory_pb2 as alloffcategory__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x11\x61lloff_size.proto\x12\x08goalloff\x1a\x14\x61lloffcategory.proto\".\n\x14GetAlloffSizeRequest\x12\x16\n\x0e\x61lloff_size_id\x18\x01 \x01(\t\"\x17\n\x15ListAlloffSizeRequest\"K\n\x16ListAlloffSizeResponse\x12\x31\n\x0c\x61lloff_sizes\x18\x01 \x03(\x0b\x32\x1b.goalloff.AlloffSizeMessage\"~\n\x17\x43reateAlloffSizeRequest\x12\x11\n\tsize_name\x18\x01 \x01(\t\x12\x1a\n\x12\x61lloff_category_id\x18\x02 \x01(\t\x12\x18\n\x10\x61lloff_size_name\x18\x03 \x01(\t\x12\x1a\n\x12original_size_name\x18\x04 \x01(\t\"\xd3\x01\n\x15\x45\x64itAlloffSizeRequest\x12\x16\n\x0e\x61lloff_size_id\x18\x01 \x01(\t\x12\x1f\n\x12\x61lloff_category_id\x18\x02 \x01(\tH\x00\x88\x01\x01\x12\x1d\n\x10\x61lloff_size_name\x18\x03 \x01(\tH\x01\x88\x01\x01\x12\x1f\n\x12original_size_name\x18\x04 \x01(\tH\x02\x88\x01\x01\x42\x15\n\x13_alloff_category_idB\x13\n\x11_alloff_size_nameB\x15\n\x13_original_size_name\"\x9b\x01\n\x11\x41lloffSizeMessage\x12\x16\n\x0e\x61lloff_size_id\x18\x01 \x01(\t\x12\x38\n\x0f\x61lloff_category\x18\x02 \x01(\x0b\x32\x1f.goalloff.AlloffCategoryMessage\x12\x18\n\x10\x61lloff_size_name\x18\x03 \x01(\t\x12\x1a\n\x12original_size_name\x18\x04 \x01(\t2\xd3\x02\n\nAlloffSize\x12L\n\rGetAlloffSize\x12\x1e.goalloff.GetAlloffSizeRequest\x1a\x1b.goalloff.AlloffSizeMessage\x12S\n\x0eListAlloffSize\x12\x1f.goalloff.ListAlloffSizeRequest\x1a .goalloff.ListAlloffSizeResponse\x12N\n\x0e\x45\x64itAlloffSize\x12\x1f.goalloff.EditAlloffSizeRequest\x1a\x1b.goalloff.AlloffSizeMessage\x12R\n\x10\x43reateAlloffSize\x12!.goalloff.CreateAlloffSizeRequest\x1a\x1b.goalloff.AlloffSizeMessageB7Z5github.com/lessbutter/alloff-grpc-protos/gen/goalloffb\x06proto3')



_GETALLOFFSIZEREQUEST = DESCRIPTOR.message_types_by_name['GetAlloffSizeRequest']
_LISTALLOFFSIZEREQUEST = DESCRIPTOR.message_types_by_name['ListAlloffSizeRequest']
_LISTALLOFFSIZERESPONSE = DESCRIPTOR.message_types_by_name['ListAlloffSizeResponse']
_CREATEALLOFFSIZEREQUEST = DESCRIPTOR.message_types_by_name['CreateAlloffSizeRequest']
_EDITALLOFFSIZEREQUEST = DESCRIPTOR.message_types_by_name['EditAlloffSizeRequest']
_ALLOFFSIZEMESSAGE = DESCRIPTOR.message_types_by_name['AlloffSizeMessage']
GetAlloffSizeRequest = _reflection.GeneratedProtocolMessageType('GetAlloffSizeRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETALLOFFSIZEREQUEST,
  '__module__' : 'alloff_size_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.GetAlloffSizeRequest)
  })
_sym_db.RegisterMessage(GetAlloffSizeRequest)

ListAlloffSizeRequest = _reflection.GeneratedProtocolMessageType('ListAlloffSizeRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTALLOFFSIZEREQUEST,
  '__module__' : 'alloff_size_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.ListAlloffSizeRequest)
  })
_sym_db.RegisterMessage(ListAlloffSizeRequest)

ListAlloffSizeResponse = _reflection.GeneratedProtocolMessageType('ListAlloffSizeResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTALLOFFSIZERESPONSE,
  '__module__' : 'alloff_size_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.ListAlloffSizeResponse)
  })
_sym_db.RegisterMessage(ListAlloffSizeResponse)

CreateAlloffSizeRequest = _reflection.GeneratedProtocolMessageType('CreateAlloffSizeRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEALLOFFSIZEREQUEST,
  '__module__' : 'alloff_size_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.CreateAlloffSizeRequest)
  })
_sym_db.RegisterMessage(CreateAlloffSizeRequest)

EditAlloffSizeRequest = _reflection.GeneratedProtocolMessageType('EditAlloffSizeRequest', (_message.Message,), {
  'DESCRIPTOR' : _EDITALLOFFSIZEREQUEST,
  '__module__' : 'alloff_size_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.EditAlloffSizeRequest)
  })
_sym_db.RegisterMessage(EditAlloffSizeRequest)

AlloffSizeMessage = _reflection.GeneratedProtocolMessageType('AlloffSizeMessage', (_message.Message,), {
  'DESCRIPTOR' : _ALLOFFSIZEMESSAGE,
  '__module__' : 'alloff_size_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.AlloffSizeMessage)
  })
_sym_db.RegisterMessage(AlloffSizeMessage)

_ALLOFFSIZE = DESCRIPTOR.services_by_name['AlloffSize']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z5github.com/lessbutter/alloff-grpc-protos/gen/goalloff'
  _GETALLOFFSIZEREQUEST._serialized_start=53
  _GETALLOFFSIZEREQUEST._serialized_end=99
  _LISTALLOFFSIZEREQUEST._serialized_start=101
  _LISTALLOFFSIZEREQUEST._serialized_end=124
  _LISTALLOFFSIZERESPONSE._serialized_start=126
  _LISTALLOFFSIZERESPONSE._serialized_end=201
  _CREATEALLOFFSIZEREQUEST._serialized_start=203
  _CREATEALLOFFSIZEREQUEST._serialized_end=329
  _EDITALLOFFSIZEREQUEST._serialized_start=332
  _EDITALLOFFSIZEREQUEST._serialized_end=543
  _ALLOFFSIZEMESSAGE._serialized_start=546
  _ALLOFFSIZEMESSAGE._serialized_end=701
  _ALLOFFSIZE._serialized_start=704
  _ALLOFFSIZE._serialized_end=1043
# @@protoc_insertion_point(module_scope)
