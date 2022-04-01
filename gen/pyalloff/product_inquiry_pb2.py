# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: product_inquiry.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from . import product_inquiry_reply_pb2 as product__inquiry__reply__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x15product_inquiry.proto\x12\x0eproductinquiry\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1bproduct_inquiry_reply.proto\"\xf5\x01\n\x0eProductInquiry\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x12\n\nproduct_id\x18\x02 \x01(\t\x12\x15\n\rbrand_keyname\x18\x03 \x01(\t\x12\x17\n\x0f\x63ompany_keyname\x18\x04 \x01(\t\x12\x0f\n\x07user_id\x18\x05 \x01(\t\x12\x11\n\tuser_name\x18\x06 \x01(\t\x12\x0c\n\x04\x62ody\x18\x07 \x01(\t\x12\x12\n\ncreated_at\x18\x08 \x01(\t\x12\x39\n\x07replies\x18\t \x03(\x0b\x32(.productinquiryreply.ProductInquiryReply\x12\x12\n\nis_pending\x18\n \x01(\x08\"\xf2\x01\n\x19ProductInquiryListRequest\x12\x0c\n\x04page\x18\x01 \x01(\x03\x12\x0c\n\x04size\x18\x02 \x01(\x03\x12\x13\n\x06search\x18\x03 \x01(\tH\x00\x88\x01\x01\x12\x38\n\x06status\x18\x04 \x01(\x0e\x32(.productinquiry.ProductInquiryListStatus\x12\x17\n\x0f\x63ompany_keyname\x18\x05 \x01(\t\x12\x16\n\tdate_from\x18\x06 \x01(\tH\x01\x88\x01\x01\x12\x14\n\x07\x64\x61te_to\x18\x07 \x01(\tH\x02\x88\x01\x01\x42\t\n\x07_searchB\x0c\n\n_date_fromB\n\n\x08_date_to\"\x9c\x01\n\x1aProductInquiryListResponse\x12\r\n\x05\x63ount\x18\x01 \x01(\x03\x12\x11\n\x04next\x18\x02 \x01(\x03H\x00\x88\x01\x01\x12\x15\n\x08previous\x18\x03 \x01(\x03H\x01\x88\x01\x01\x12/\n\x07results\x18\x04 \x03(\x0b\x32\x1e.productinquiry.ProductInquiryB\x07\n\x05_nextB\x0b\n\t_previous\"D\n\x1dProductInquiryRetrieveRequest\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x17\n\x0f\x63ompany_keyname\x18\x02 \x01(\t\"\x7f\n ProductInquiryReplyCreateRequest\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x0c\n\x04\x62ody\x18\x02 \x01(\t\x12\x15\n\ruser_username\x18\x03 \x01(\t\x12\x11\n\tuser_uuid\x18\x04 \x01(\t\x12\x17\n\x0f\x63ompany_keyname\x18\x05 \x01(\t\"H\n!ProductInquiryReplyDestroyRequest\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x17\n\x0f\x63ompany_keyname\x18\x02 \x01(\t*=\n\x18ProductInquiryListStatus\x12\x07\n\x03\x41LL\x10\x00\x12\x0b\n\x07PENDING\x10\x01\x12\x0b\n\x07REPLIED\x10\x02\x32\x97\x03\n\x18ProductInquiryController\x12_\n\x04List\x12).productinquiry.ProductInquiryListRequest\x1a*.productinquiry.ProductInquiryListResponse\"\x00\x12[\n\x08Retrieve\x12-.productinquiry.ProductInquiryRetrieveRequest\x1a\x1e.productinquiry.ProductInquiry\"\x00\x12\x61\n\x0b\x43reateReply\x12\x30.productinquiry.ProductInquiryReplyCreateRequest\x1a\x1e.productinquiry.ProductInquiry\"\x00\x12Z\n\x0b\x44\x65leteReply\x12\x31.productinquiry.ProductInquiryReplyDestroyRequest\x1a\x16.google.protobuf.Empty\"\x00\x62\x06proto3')

_PRODUCTINQUIRYLISTSTATUS = DESCRIPTOR.enum_types_by_name['ProductInquiryListStatus']
ProductInquiryListStatus = enum_type_wrapper.EnumTypeWrapper(_PRODUCTINQUIRYLISTSTATUS)
ALL = 0
PENDING = 1
REPLIED = 2


_PRODUCTINQUIRY = DESCRIPTOR.message_types_by_name['ProductInquiry']
_PRODUCTINQUIRYLISTREQUEST = DESCRIPTOR.message_types_by_name['ProductInquiryListRequest']
_PRODUCTINQUIRYLISTRESPONSE = DESCRIPTOR.message_types_by_name['ProductInquiryListResponse']
_PRODUCTINQUIRYRETRIEVEREQUEST = DESCRIPTOR.message_types_by_name['ProductInquiryRetrieveRequest']
_PRODUCTINQUIRYREPLYCREATEREQUEST = DESCRIPTOR.message_types_by_name['ProductInquiryReplyCreateRequest']
_PRODUCTINQUIRYREPLYDESTROYREQUEST = DESCRIPTOR.message_types_by_name['ProductInquiryReplyDestroyRequest']
ProductInquiry = _reflection.GeneratedProtocolMessageType('ProductInquiry', (_message.Message,), {
  'DESCRIPTOR' : _PRODUCTINQUIRY,
  '__module__' : 'product_inquiry_pb2'
  # @@protoc_insertion_point(class_scope:productinquiry.ProductInquiry)
  })
_sym_db.RegisterMessage(ProductInquiry)

ProductInquiryListRequest = _reflection.GeneratedProtocolMessageType('ProductInquiryListRequest', (_message.Message,), {
  'DESCRIPTOR' : _PRODUCTINQUIRYLISTREQUEST,
  '__module__' : 'product_inquiry_pb2'
  # @@protoc_insertion_point(class_scope:productinquiry.ProductInquiryListRequest)
  })
_sym_db.RegisterMessage(ProductInquiryListRequest)

ProductInquiryListResponse = _reflection.GeneratedProtocolMessageType('ProductInquiryListResponse', (_message.Message,), {
  'DESCRIPTOR' : _PRODUCTINQUIRYLISTRESPONSE,
  '__module__' : 'product_inquiry_pb2'
  # @@protoc_insertion_point(class_scope:productinquiry.ProductInquiryListResponse)
  })
_sym_db.RegisterMessage(ProductInquiryListResponse)

ProductInquiryRetrieveRequest = _reflection.GeneratedProtocolMessageType('ProductInquiryRetrieveRequest', (_message.Message,), {
  'DESCRIPTOR' : _PRODUCTINQUIRYRETRIEVEREQUEST,
  '__module__' : 'product_inquiry_pb2'
  # @@protoc_insertion_point(class_scope:productinquiry.ProductInquiryRetrieveRequest)
  })
_sym_db.RegisterMessage(ProductInquiryRetrieveRequest)

ProductInquiryReplyCreateRequest = _reflection.GeneratedProtocolMessageType('ProductInquiryReplyCreateRequest', (_message.Message,), {
  'DESCRIPTOR' : _PRODUCTINQUIRYREPLYCREATEREQUEST,
  '__module__' : 'product_inquiry_pb2'
  # @@protoc_insertion_point(class_scope:productinquiry.ProductInquiryReplyCreateRequest)
  })
_sym_db.RegisterMessage(ProductInquiryReplyCreateRequest)

ProductInquiryReplyDestroyRequest = _reflection.GeneratedProtocolMessageType('ProductInquiryReplyDestroyRequest', (_message.Message,), {
  'DESCRIPTOR' : _PRODUCTINQUIRYREPLYDESTROYREQUEST,
  '__module__' : 'product_inquiry_pb2'
  # @@protoc_insertion_point(class_scope:productinquiry.ProductInquiryReplyDestroyRequest)
  })
_sym_db.RegisterMessage(ProductInquiryReplyDestroyRequest)

_PRODUCTINQUIRYCONTROLLER = DESCRIPTOR.services_by_name['ProductInquiryController']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  _PRODUCTINQUIRYLISTSTATUS._serialized_start=1024
  _PRODUCTINQUIRYLISTSTATUS._serialized_end=1085
  _PRODUCTINQUIRY._serialized_start=100
  _PRODUCTINQUIRY._serialized_end=345
  _PRODUCTINQUIRYLISTREQUEST._serialized_start=348
  _PRODUCTINQUIRYLISTREQUEST._serialized_end=590
  _PRODUCTINQUIRYLISTRESPONSE._serialized_start=593
  _PRODUCTINQUIRYLISTRESPONSE._serialized_end=749
  _PRODUCTINQUIRYRETRIEVEREQUEST._serialized_start=751
  _PRODUCTINQUIRYRETRIEVEREQUEST._serialized_end=819
  _PRODUCTINQUIRYREPLYCREATEREQUEST._serialized_start=821
  _PRODUCTINQUIRYREPLYCREATEREQUEST._serialized_end=948
  _PRODUCTINQUIRYREPLYDESTROYREQUEST._serialized_start=950
  _PRODUCTINQUIRYREPLYDESTROYREQUEST._serialized_end=1022
  _PRODUCTINQUIRYCONTROLLER._serialized_start=1088
  _PRODUCTINQUIRYCONTROLLER._serialized_end=1495
# @@protoc_insertion_point(module_scope)
