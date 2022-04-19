# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: exhibition.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from . import productGroup_pb2 as productGroup__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x10\x65xhibition.proto\x12\x08goalloff\x1a\x12productGroup.proto\"-\n\x14GetExhibitionRequest\x12\x15\n\rexhibition_id\x18\x01 \x01(\t\"\x94\x01\n\x16ListExhibitionsRequest\x12\x0e\n\x06offset\x18\x01 \x01(\x05\x12\r\n\x05limit\x18\x02 \x01(\x05\x12,\n\ngroup_type\x18\x03 \x01(\x0e\x32\x18.goalloff.ExhibitionType\x12\x0f\n\x07is_live\x18\x04 \x01(\x08\x12\x12\n\x05query\x18\x05 \x01(\tH\x00\x88\x01\x01\x42\x08\n\x06_query\"\xdc\x03\n\x15\x45\x64itExhibitionRequest\x12\x15\n\rexhibition_id\x18\x01 \x01(\t\x12\x19\n\x0c\x62\x61nner_image\x18\x02 \x01(\tH\x00\x88\x01\x01\x12\x1c\n\x0fthumbnail_image\x18\x03 \x01(\tH\x01\x88\x01\x01\x12\x12\n\x05title\x18\x04 \x01(\tH\x02\x88\x01\x01\x12\x15\n\x08subtitle\x18\x05 \x01(\tH\x03\x88\x01\x01\x12\x18\n\x0b\x64\x65scription\x18\x06 \x01(\tH\x04\x88\x01\x01\x12\x17\n\nstart_time\x18\x07 \x01(\tH\x05\x88\x01\x01\x12\x18\n\x0b\x66inish_time\x18\x08 \x01(\tH\x06\x88\x01\x01\x12\x0e\n\x06pg_ids\x18\t \x03(\t\x12\x14\n\x07is_live\x18\n \x01(\x08H\x07\x88\x01\x01\x12\x19\n\x0ctarget_sales\x18\x0b \x01(\x05H\x08\x88\x01\x01\x12\x32\n\x07\x62\x61nners\x18\x0c \x03(\x0b\x32!.goalloff.ExhibitionBannerMessageB\x0f\n\r_banner_imageB\x12\n\x10_thumbnail_imageB\x08\n\x06_titleB\x0b\n\t_subtitleB\x0e\n\x0c_descriptionB\r\n\x0b_start_timeB\x0e\n\x0c_finish_timeB\n\n\x08_is_liveB\x0f\n\r_target_sales\"\xb4\x02\n\x17\x43reateExhibitionRequest\x12\x14\n\x0c\x62\x61nner_image\x18\x01 \x01(\t\x12\x17\n\x0fthumbnail_image\x18\x02 \x01(\t\x12\r\n\x05title\x18\x03 \x01(\t\x12\x10\n\x08subtitle\x18\x04 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x05 \x01(\t\x12\x12\n\nstart_time\x18\x06 \x01(\t\x12\x13\n\x0b\x66inish_time\x18\x07 \x01(\t\x12\x0e\n\x06pg_ids\x18\x08 \x03(\t\x12\x31\n\x0f\x65xhibition_type\x18\t \x01(\x0e\x32\x18.goalloff.ExhibitionType\x12\x14\n\x0ctarget_sales\x18\n \x01(\x05\x12\x32\n\x07\x62\x61nners\x18\x0b \x03(\x0b\x32!.goalloff.ExhibitionBannerMessage\"H\n\x15GetExhibitionResponse\x12/\n\nexhibition\x18\x01 \x01(\x0b\x32\x1b.goalloff.ExhibitionMessage\"\xce\x01\n\x17ListExhibitionsResponse\x12\x30\n\x0b\x65xhibitions\x18\x01 \x03(\x0b\x32\x1b.goalloff.ExhibitionMessage\x12\x0e\n\x06offset\x18\x02 \x01(\x05\x12\r\n\x05limit\x18\x03 \x01(\x05\x12\x14\n\x0ctotal_counts\x18\x04 \x01(\x05\x12,\n\ngroup_type\x18\x05 \x01(\x0e\x32\x18.goalloff.ExhibitionType\x12\x0f\n\x07is_live\x18\x06 \x01(\x08\x12\r\n\x05query\x18\x07 \x01(\t\"I\n\x16\x45\x64itExhibitionResponse\x12/\n\nexhibition\x18\x01 \x01(\x0b\x32\x1b.goalloff.ExhibitionMessage\"K\n\x18\x43reateExhibitionResponse\x12/\n\nexhibition\x18\x01 \x01(\x0b\x32\x1b.goalloff.ExhibitionMessage\"\x89\x03\n\x11\x45xhibitionMessage\x12\x15\n\rexhibition_id\x18\x01 \x01(\t\x12\x14\n\x0c\x62\x61nner_image\x18\x02 \x01(\t\x12\x17\n\x0fthumbnail_image\x18\x03 \x01(\t\x12\r\n\x05title\x18\x04 \x01(\t\x12\x10\n\x08subtitle\x18\x05 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x06 \x01(\t\x12\x12\n\nstart_time\x18\x07 \x01(\t\x12\x13\n\x0b\x66inish_time\x18\x08 \x01(\t\x12*\n\x03pgs\x18\t \x03(\x0b\x32\x1d.goalloff.ProductGroupMessage\x12\x0f\n\x07is_live\x18\n \x01(\x08\x12\x31\n\x0f\x65xhibition_type\x18\x0b \x01(\x0e\x32\x18.goalloff.ExhibitionType\x12\x14\n\x0ctarget_sales\x18\x0c \x01(\x05\x12\x15\n\rcurrent_sales\x18\r \x01(\x05\x12\x32\n\x07\x62\x61nners\x18\x0e \x03(\x0b\x32!.goalloff.ExhibitionBannerMessage\"e\n\x17\x45xhibitionBannerMessage\x12\x0f\n\x07img_url\x18\x01 \x01(\t\x12\r\n\x05title\x18\x02 \x01(\t\x12\x10\n\x08subtitle\x18\x03 \x01(\t\x12\x18\n\x10product_group_id\x18\x04 \x01(\t*Z\n\x0e\x45xhibitionType\x12\x15\n\x11\x45XHIBITION_NORMAL\x10\x00\x12\x17\n\x13\x45XHIBITION_TIMEDEAL\x10\x01\x12\x18\n\x14\x45XHIBITION_GROUPDEAL\x10\x02\x32\xe6\x02\n\nExhibition\x12P\n\rGetExhibition\x12\x1e.goalloff.GetExhibitionRequest\x1a\x1f.goalloff.GetExhibitionResponse\x12V\n\x0fListExhibitions\x12 .goalloff.ListExhibitionsRequest\x1a!.goalloff.ListExhibitionsResponse\x12S\n\x0e\x45\x64itExhibition\x12\x1f.goalloff.EditExhibitionRequest\x1a .goalloff.EditExhibitionResponse\x12Y\n\x10\x43reateExhibition\x12!.goalloff.CreateExhibitionRequest\x1a\".goalloff.CreateExhibitionResponseB7Z5github.com/lessbutter/alloff-grpc-protos/gen/goalloffb\x06proto3')

_EXHIBITIONTYPE = DESCRIPTOR.enum_types_by_name['ExhibitionType']
ExhibitionType = enum_type_wrapper.EnumTypeWrapper(_EXHIBITIONTYPE)
EXHIBITION_NORMAL = 0
EXHIBITION_TIMEDEAL = 1
EXHIBITION_GROUPDEAL = 2


_GETEXHIBITIONREQUEST = DESCRIPTOR.message_types_by_name['GetExhibitionRequest']
_LISTEXHIBITIONSREQUEST = DESCRIPTOR.message_types_by_name['ListExhibitionsRequest']
_EDITEXHIBITIONREQUEST = DESCRIPTOR.message_types_by_name['EditExhibitionRequest']
_CREATEEXHIBITIONREQUEST = DESCRIPTOR.message_types_by_name['CreateExhibitionRequest']
_GETEXHIBITIONRESPONSE = DESCRIPTOR.message_types_by_name['GetExhibitionResponse']
_LISTEXHIBITIONSRESPONSE = DESCRIPTOR.message_types_by_name['ListExhibitionsResponse']
_EDITEXHIBITIONRESPONSE = DESCRIPTOR.message_types_by_name['EditExhibitionResponse']
_CREATEEXHIBITIONRESPONSE = DESCRIPTOR.message_types_by_name['CreateExhibitionResponse']
_EXHIBITIONMESSAGE = DESCRIPTOR.message_types_by_name['ExhibitionMessage']
_EXHIBITIONBANNERMESSAGE = DESCRIPTOR.message_types_by_name['ExhibitionBannerMessage']
GetExhibitionRequest = _reflection.GeneratedProtocolMessageType('GetExhibitionRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETEXHIBITIONREQUEST,
  '__module__' : 'exhibition_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.GetExhibitionRequest)
  })
_sym_db.RegisterMessage(GetExhibitionRequest)

ListExhibitionsRequest = _reflection.GeneratedProtocolMessageType('ListExhibitionsRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTEXHIBITIONSREQUEST,
  '__module__' : 'exhibition_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.ListExhibitionsRequest)
  })
_sym_db.RegisterMessage(ListExhibitionsRequest)

EditExhibitionRequest = _reflection.GeneratedProtocolMessageType('EditExhibitionRequest', (_message.Message,), {
  'DESCRIPTOR' : _EDITEXHIBITIONREQUEST,
  '__module__' : 'exhibition_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.EditExhibitionRequest)
  })
_sym_db.RegisterMessage(EditExhibitionRequest)

CreateExhibitionRequest = _reflection.GeneratedProtocolMessageType('CreateExhibitionRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEEXHIBITIONREQUEST,
  '__module__' : 'exhibition_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.CreateExhibitionRequest)
  })
_sym_db.RegisterMessage(CreateExhibitionRequest)

GetExhibitionResponse = _reflection.GeneratedProtocolMessageType('GetExhibitionResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETEXHIBITIONRESPONSE,
  '__module__' : 'exhibition_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.GetExhibitionResponse)
  })
_sym_db.RegisterMessage(GetExhibitionResponse)

ListExhibitionsResponse = _reflection.GeneratedProtocolMessageType('ListExhibitionsResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTEXHIBITIONSRESPONSE,
  '__module__' : 'exhibition_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.ListExhibitionsResponse)
  })
_sym_db.RegisterMessage(ListExhibitionsResponse)

EditExhibitionResponse = _reflection.GeneratedProtocolMessageType('EditExhibitionResponse', (_message.Message,), {
  'DESCRIPTOR' : _EDITEXHIBITIONRESPONSE,
  '__module__' : 'exhibition_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.EditExhibitionResponse)
  })
_sym_db.RegisterMessage(EditExhibitionResponse)

CreateExhibitionResponse = _reflection.GeneratedProtocolMessageType('CreateExhibitionResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATEEXHIBITIONRESPONSE,
  '__module__' : 'exhibition_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.CreateExhibitionResponse)
  })
_sym_db.RegisterMessage(CreateExhibitionResponse)

ExhibitionMessage = _reflection.GeneratedProtocolMessageType('ExhibitionMessage', (_message.Message,), {
  'DESCRIPTOR' : _EXHIBITIONMESSAGE,
  '__module__' : 'exhibition_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.ExhibitionMessage)
  })
_sym_db.RegisterMessage(ExhibitionMessage)

ExhibitionBannerMessage = _reflection.GeneratedProtocolMessageType('ExhibitionBannerMessage', (_message.Message,), {
  'DESCRIPTOR' : _EXHIBITIONBANNERMESSAGE,
  '__module__' : 'exhibition_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.ExhibitionBannerMessage)
  })
_sym_db.RegisterMessage(ExhibitionBannerMessage)

_EXHIBITION = DESCRIPTOR.services_by_name['Exhibition']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z5github.com/lessbutter/alloff-grpc-protos/gen/goalloff'
  _EXHIBITIONTYPE._serialized_start=1972
  _EXHIBITIONTYPE._serialized_end=2062
  _GETEXHIBITIONREQUEST._serialized_start=50
  _GETEXHIBITIONREQUEST._serialized_end=95
  _LISTEXHIBITIONSREQUEST._serialized_start=98
  _LISTEXHIBITIONSREQUEST._serialized_end=246
  _EDITEXHIBITIONREQUEST._serialized_start=249
  _EDITEXHIBITIONREQUEST._serialized_end=725
  _CREATEEXHIBITIONREQUEST._serialized_start=728
  _CREATEEXHIBITIONREQUEST._serialized_end=1036
  _GETEXHIBITIONRESPONSE._serialized_start=1038
  _GETEXHIBITIONRESPONSE._serialized_end=1110
  _LISTEXHIBITIONSRESPONSE._serialized_start=1113
  _LISTEXHIBITIONSRESPONSE._serialized_end=1319
  _EDITEXHIBITIONRESPONSE._serialized_start=1321
  _EDITEXHIBITIONRESPONSE._serialized_end=1394
  _CREATEEXHIBITIONRESPONSE._serialized_start=1396
  _CREATEEXHIBITIONRESPONSE._serialized_end=1471
  _EXHIBITIONMESSAGE._serialized_start=1474
  _EXHIBITIONMESSAGE._serialized_end=1867
  _EXHIBITIONBANNERMESSAGE._serialized_start=1869
  _EXHIBITIONBANNERMESSAGE._serialized_end=1970
  _EXHIBITION._serialized_start=2065
  _EXHIBITION._serialized_end=2423
# @@protoc_insertion_point(module_scope)
