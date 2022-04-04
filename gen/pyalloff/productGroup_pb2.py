# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: productGroup.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from . import product_pb2 as product__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x12productGroup.proto\x12\x08goalloff\x1a\rproduct.proto\"2\n\x16GetProductGroupRequest\x12\x18\n\x10product_group_id\x18\x01 \x01(\t\"\xe8\x01\n\x19\x43reateProductGroupRequest\x12\r\n\x05title\x18\x01 \x01(\t\x12\x18\n\x0bshort_title\x18\x02 \x01(\tH\x00\x88\x01\x01\x12\x13\n\x0binstruction\x18\x03 \x03(\t\x12\x16\n\timage_url\x18\x04 \x01(\tH\x01\x88\x01\x01\x12\x12\n\nstart_time\x18\x05 \x01(\t\x12\x13\n\x0b\x66inish_time\x18\x06 \x01(\t\x12.\n\ngroup_type\x18\x07 \x01(\x0e\x32\x1a.goalloff.ProductGroupTypeB\x0e\n\x0c_short_titleB\x0c\n\n_image_url\"e\n\x18ListProductGroupsRequest\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.goalloff.ProductGroupQuery\x12\x0e\n\x06offset\x18\x02 \x01(\x05\x12\r\n\x05limit\x18\x03 \x01(\x05\"\xcc\x02\n\x17\x45\x64itProductGroupRequest\x12\x12\n\x05title\x18\x01 \x01(\tH\x00\x88\x01\x01\x12\x18\n\x0bshort_title\x18\x02 \x01(\tH\x01\x88\x01\x01\x12\x13\n\x0binstruction\x18\x03 \x03(\t\x12\x16\n\timage_url\x18\x04 \x01(\tH\x02\x88\x01\x01\x12\x17\n\nstart_time\x18\x05 \x01(\tH\x03\x88\x01\x01\x12\x18\n\x0b\x66inish_time\x18\x06 \x01(\tH\x04\x88\x01\x01\x12\x18\n\x10product_group_id\x18\x07 \x01(\t\x12\x33\n\ngroup_type\x18\t \x01(\x0e\x32\x1a.goalloff.ProductGroupTypeH\x05\x88\x01\x01\x42\x08\n\x06_titleB\x0e\n\x0c_short_titleB\x0c\n\n_image_urlB\r\n\x0b_start_timeB\x0e\n\x0c_finish_timeB\r\n\x0b_group_type\"m\n\x13ProductsInPgRequest\x12\x18\n\x10product_group_id\x18\x01 \x01(\t\x12<\n\x12product_priorities\x18\x02 \x03(\x0b\x32 .goalloff.ProductPriorityMessage\"H\n\x18RemoveProductInPgRequest\x12\x18\n\x10product_group_id\x18\x01 \x01(\t\x12\x12\n\nproduct_id\x18\x02 \x01(\t\"\x87\x02\n\x13ProductGroupMessage\x12\r\n\x05title\x18\x01 \x01(\t\x12\x13\n\x0bshort_title\x18\x02 \x01(\t\x12\x13\n\x0binstruction\x18\x03 \x03(\t\x12\x11\n\timage_url\x18\x04 \x01(\t\x12\x31\n\x08products\x18\x05 \x03(\x0b\x32\x1f.goalloff.ProductInGroupMessage\x12\x12\n\nstart_time\x18\x06 \x01(\t\x12\x13\n\x0b\x66inish_time\x18\x07 \x01(\t\x12\x18\n\x10product_group_id\x18\x08 \x01(\t\x12.\n\ngroup_type\x18\t \x01(\x0e\x32\x1a.goalloff.ProductGroupType\"|\n\x19ListProductGroupsResponse\x12*\n\x03pgs\x18\x01 \x03(\x0b\x32\x1d.goalloff.ProductGroupMessage\x12\x0e\n\x06offset\x18\x02 \x01(\x05\x12\r\n\x05limit\x18\x03 \x01(\x05\x12\x14\n\x0ctotal_counts\x18\x04 \x01(\x05\"T\n\x15ProductInGroupMessage\x12)\n\x07product\x18\x01 \x01(\x0b\x32\x18.goalloff.ProductMessage\x12\x10\n\x08priority\x18\x02 \x01(\x05\">\n\x16ProductPriorityMessage\x12\x12\n\nproduct_id\x18\x01 \x01(\t\x12\x10\n\x08priority\x18\x02 \x01(\x05\"\x83\x01\n\x11ProductGroupQuery\x12\x19\n\x0csearch_query\x18\x01 \x01(\tH\x00\x88\x01\x01\x12\x33\n\ngroup_type\x18\x02 \x01(\x0e\x32\x1a.goalloff.ProductGroupTypeH\x01\x88\x01\x01\x42\x0f\n\r_search_queryB\r\n\x0b_group_type*i\n\x10ProductGroupType\x12\x1a\n\x16PRODUCT_GROUP_TIMEDEAL\x10\x00\x12\x1c\n\x18PRODUCT_GROUP_EXHIBITION\x10\x01\x12\x1b\n\x17PRODUCT_GROUP_GROUPDEAL\x10\x02\x32\x8c\x05\n\x0cProductGroup\x12R\n\x0fGetProductGroup\x12 .goalloff.GetProductGroupRequest\x1a\x1d.goalloff.ProductGroupMessage\x12X\n\x12\x43reateProductGroup\x12#.goalloff.CreateProductGroupRequest\x1a\x1d.goalloff.ProductGroupMessage\x12\\\n\x11ListProductGroups\x12\".goalloff.ListProductGroupsRequest\x1a#.goalloff.ListProductGroupsResponse\x12T\n\x10\x45\x64itProductGroup\x12!.goalloff.EditProductGroupRequest\x1a\x1d.goalloff.ProductGroupMessage\x12Z\n\x1aPushProductsInProductGroup\x12\x1d.goalloff.ProductsInPgRequest\x1a\x1d.goalloff.ProductGroupMessage\x12\\\n\x1cUpdateProductsInProductGroup\x12\x1d.goalloff.ProductsInPgRequest\x1a\x1d.goalloff.ProductGroupMessage\x12`\n\x1bRemoveProductInProductGroup\x12\".goalloff.RemoveProductInPgRequest\x1a\x1d.goalloff.ProductGroupMessageB7Z5github.com/lessbutter/alloff-grpc-protos/gen/goalloffb\x06proto3')

_PRODUCTGROUPTYPE = DESCRIPTOR.enum_types_by_name['ProductGroupType']
ProductGroupType = enum_type_wrapper.EnumTypeWrapper(_PRODUCTGROUPTYPE)
PRODUCT_GROUP_TIMEDEAL = 0
PRODUCT_GROUP_EXHIBITION = 1
PRODUCT_GROUP_GROUPDEAL = 2


_GETPRODUCTGROUPREQUEST = DESCRIPTOR.message_types_by_name['GetProductGroupRequest']
_CREATEPRODUCTGROUPREQUEST = DESCRIPTOR.message_types_by_name['CreateProductGroupRequest']
_LISTPRODUCTGROUPSREQUEST = DESCRIPTOR.message_types_by_name['ListProductGroupsRequest']
_EDITPRODUCTGROUPREQUEST = DESCRIPTOR.message_types_by_name['EditProductGroupRequest']
_PRODUCTSINPGREQUEST = DESCRIPTOR.message_types_by_name['ProductsInPgRequest']
_REMOVEPRODUCTINPGREQUEST = DESCRIPTOR.message_types_by_name['RemoveProductInPgRequest']
_PRODUCTGROUPMESSAGE = DESCRIPTOR.message_types_by_name['ProductGroupMessage']
_LISTPRODUCTGROUPSRESPONSE = DESCRIPTOR.message_types_by_name['ListProductGroupsResponse']
_PRODUCTINGROUPMESSAGE = DESCRIPTOR.message_types_by_name['ProductInGroupMessage']
_PRODUCTPRIORITYMESSAGE = DESCRIPTOR.message_types_by_name['ProductPriorityMessage']
_PRODUCTGROUPQUERY = DESCRIPTOR.message_types_by_name['ProductGroupQuery']
GetProductGroupRequest = _reflection.GeneratedProtocolMessageType('GetProductGroupRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETPRODUCTGROUPREQUEST,
  '__module__' : 'productGroup_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.GetProductGroupRequest)
  })
_sym_db.RegisterMessage(GetProductGroupRequest)

CreateProductGroupRequest = _reflection.GeneratedProtocolMessageType('CreateProductGroupRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEPRODUCTGROUPREQUEST,
  '__module__' : 'productGroup_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.CreateProductGroupRequest)
  })
_sym_db.RegisterMessage(CreateProductGroupRequest)

ListProductGroupsRequest = _reflection.GeneratedProtocolMessageType('ListProductGroupsRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTPRODUCTGROUPSREQUEST,
  '__module__' : 'productGroup_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.ListProductGroupsRequest)
  })
_sym_db.RegisterMessage(ListProductGroupsRequest)

EditProductGroupRequest = _reflection.GeneratedProtocolMessageType('EditProductGroupRequest', (_message.Message,), {
  'DESCRIPTOR' : _EDITPRODUCTGROUPREQUEST,
  '__module__' : 'productGroup_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.EditProductGroupRequest)
  })
_sym_db.RegisterMessage(EditProductGroupRequest)

ProductsInPgRequest = _reflection.GeneratedProtocolMessageType('ProductsInPgRequest', (_message.Message,), {
  'DESCRIPTOR' : _PRODUCTSINPGREQUEST,
  '__module__' : 'productGroup_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.ProductsInPgRequest)
  })
_sym_db.RegisterMessage(ProductsInPgRequest)

RemoveProductInPgRequest = _reflection.GeneratedProtocolMessageType('RemoveProductInPgRequest', (_message.Message,), {
  'DESCRIPTOR' : _REMOVEPRODUCTINPGREQUEST,
  '__module__' : 'productGroup_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.RemoveProductInPgRequest)
  })
_sym_db.RegisterMessage(RemoveProductInPgRequest)

ProductGroupMessage = _reflection.GeneratedProtocolMessageType('ProductGroupMessage', (_message.Message,), {
  'DESCRIPTOR' : _PRODUCTGROUPMESSAGE,
  '__module__' : 'productGroup_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.ProductGroupMessage)
  })
_sym_db.RegisterMessage(ProductGroupMessage)

ListProductGroupsResponse = _reflection.GeneratedProtocolMessageType('ListProductGroupsResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTPRODUCTGROUPSRESPONSE,
  '__module__' : 'productGroup_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.ListProductGroupsResponse)
  })
_sym_db.RegisterMessage(ListProductGroupsResponse)

ProductInGroupMessage = _reflection.GeneratedProtocolMessageType('ProductInGroupMessage', (_message.Message,), {
  'DESCRIPTOR' : _PRODUCTINGROUPMESSAGE,
  '__module__' : 'productGroup_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.ProductInGroupMessage)
  })
_sym_db.RegisterMessage(ProductInGroupMessage)

ProductPriorityMessage = _reflection.GeneratedProtocolMessageType('ProductPriorityMessage', (_message.Message,), {
  'DESCRIPTOR' : _PRODUCTPRIORITYMESSAGE,
  '__module__' : 'productGroup_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.ProductPriorityMessage)
  })
_sym_db.RegisterMessage(ProductPriorityMessage)

ProductGroupQuery = _reflection.GeneratedProtocolMessageType('ProductGroupQuery', (_message.Message,), {
  'DESCRIPTOR' : _PRODUCTGROUPQUERY,
  '__module__' : 'productGroup_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.ProductGroupQuery)
  })
_sym_db.RegisterMessage(ProductGroupQuery)

_PRODUCTGROUP = DESCRIPTOR.services_by_name['ProductGroup']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z5github.com/lessbutter/alloff-grpc-protos/gen/goalloff'
  _PRODUCTGROUPTYPE._serialized_start=1633
  _PRODUCTGROUPTYPE._serialized_end=1738
  _GETPRODUCTGROUPREQUEST._serialized_start=47
  _GETPRODUCTGROUPREQUEST._serialized_end=97
  _CREATEPRODUCTGROUPREQUEST._serialized_start=100
  _CREATEPRODUCTGROUPREQUEST._serialized_end=332
  _LISTPRODUCTGROUPSREQUEST._serialized_start=334
  _LISTPRODUCTGROUPSREQUEST._serialized_end=435
  _EDITPRODUCTGROUPREQUEST._serialized_start=438
  _EDITPRODUCTGROUPREQUEST._serialized_end=770
  _PRODUCTSINPGREQUEST._serialized_start=772
  _PRODUCTSINPGREQUEST._serialized_end=881
  _REMOVEPRODUCTINPGREQUEST._serialized_start=883
  _REMOVEPRODUCTINPGREQUEST._serialized_end=955
  _PRODUCTGROUPMESSAGE._serialized_start=958
  _PRODUCTGROUPMESSAGE._serialized_end=1221
  _LISTPRODUCTGROUPSRESPONSE._serialized_start=1223
  _LISTPRODUCTGROUPSRESPONSE._serialized_end=1347
  _PRODUCTINGROUPMESSAGE._serialized_start=1349
  _PRODUCTINGROUPMESSAGE._serialized_end=1433
  _PRODUCTPRIORITYMESSAGE._serialized_start=1435
  _PRODUCTPRIORITYMESSAGE._serialized_end=1497
  _PRODUCTGROUPQUERY._serialized_start=1500
  _PRODUCTGROUPQUERY._serialized_end=1631
  _PRODUCTGROUP._serialized_start=1741
  _PRODUCTGROUP._serialized_end=2393
# @@protoc_insertion_point(module_scope)
