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


import product_pb2 as product__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x12productGroup.proto\x12\x06protos\x1a\rproduct.proto\"2\n\x16GetProductGroupRequest\x12\x18\n\x10product_group_id\x18\x01 \x01(\t\"B\n\x17GetProductGroupResponse\x12\'\n\x02pg\x18\x01 \x01(\x0b\x32\x1b.protos.ProductGroupMessage\"\xe6\x01\n\x19\x43reateProductGroupRequest\x12\r\n\x05title\x18\x01 \x01(\t\x12\x18\n\x0bshort_title\x18\x02 \x01(\tH\x00\x88\x01\x01\x12\x13\n\x0binstruction\x18\x03 \x03(\t\x12\x16\n\timage_url\x18\x04 \x01(\tH\x01\x88\x01\x01\x12\x12\n\nstart_time\x18\x05 \x01(\t\x12\x13\n\x0b\x66inish_time\x18\x06 \x01(\t\x12,\n\ngroup_type\x18\x07 \x01(\x0e\x32\x18.protos.ProductGroupTypeB\x0e\n\x0c_short_titleB\x0c\n\n_image_url\"E\n\x1a\x43reateProductGroupResponse\x12\'\n\x02pg\x18\x01 \x01(\x0b\x32\x1b.protos.ProductGroupMessage\"\xca\x02\n\x17\x45\x64itProductGroupRequest\x12\x12\n\x05title\x18\x01 \x01(\tH\x00\x88\x01\x01\x12\x18\n\x0bshort_title\x18\x02 \x01(\tH\x01\x88\x01\x01\x12\x13\n\x0binstruction\x18\x03 \x03(\t\x12\x16\n\timage_url\x18\x04 \x01(\tH\x02\x88\x01\x01\x12\x17\n\nstart_time\x18\x05 \x01(\tH\x03\x88\x01\x01\x12\x18\n\x0b\x66inish_time\x18\x06 \x01(\tH\x04\x88\x01\x01\x12\x18\n\x10product_group_id\x18\x07 \x01(\t\x12\x31\n\ngroup_type\x18\t \x01(\x0e\x32\x18.protos.ProductGroupTypeH\x05\x88\x01\x01\x42\x08\n\x06_titleB\x0e\n\x0c_short_titleB\x0c\n\n_image_urlB\r\n\x0b_start_timeB\x0e\n\x0c_finish_timeB\r\n\x0b_group_type\"C\n\x18\x45\x64itProductGroupResponse\x12\'\n\x02pg\x18\x01 \x01(\x0b\x32\x1b.protos.ProductGroupMessage\"c\n\x18ListProductGroupsRequest\x12(\n\x05query\x18\x01 \x01(\x0b\x32\x19.protos.ProductGroupQuery\x12\x0e\n\x06offset\x18\x02 \x01(\x05\x12\r\n\x05limit\x18\x03 \x01(\x05\"z\n\x19ListProductGroupsResponse\x12(\n\x03pgs\x18\x01 \x03(\x0b\x32\x1b.protos.ProductGroupMessage\x12\x0e\n\x06offset\x18\x02 \x01(\x05\x12\r\n\x05limit\x18\x03 \x01(\x05\x12\x14\n\x0ctotal_counts\x18\x04 \x01(\x05\"m\n\x17PushProductsInPgRequest\x12\x18\n\x10product_group_id\x18\x01 \x01(\t\x12\x38\n\x10product_priority\x18\x02 \x03(\x0b\x32\x1e.protos.ProductPriorityMessage\"C\n\x18PushProductsInPgResponse\x12\'\n\x02pg\x18\x01 \x01(\x0b\x32\x1b.protos.ProductGroupMessage\"H\n\x18RemoveProductInPgRequest\x12\x18\n\x10product_group_id\x18\x01 \x01(\t\x12\x12\n\nproduct_id\x18\x02 \x01(\t\"D\n\x19RemoveProductInPgResponse\x12\'\n\x02pg\x18\x01 \x01(\x0b\x32\x1b.protos.ProductGroupMessage\"R\n\x15ProductInGroupMessage\x12\'\n\x07product\x18\x01 \x01(\x0b\x32\x16.protos.ProductMessage\x12\x10\n\x08priority\x18\x02 \x01(\x05\"\x83\x02\n\x13ProductGroupMessage\x12\r\n\x05title\x18\x01 \x01(\t\x12\x13\n\x0bshort_title\x18\x02 \x01(\t\x12\x13\n\x0binstruction\x18\x03 \x03(\t\x12\x11\n\timage_url\x18\x04 \x01(\t\x12/\n\x08products\x18\x05 \x03(\x0b\x32\x1d.protos.ProductInGroupMessage\x12\x12\n\nstart_time\x18\x06 \x01(\t\x12\x13\n\x0b\x66inish_time\x18\x07 \x01(\t\x12\x18\n\x10product_group_id\x18\x08 \x01(\t\x12,\n\ngroup_type\x18\t \x01(\x0e\x32\x18.protos.ProductGroupType\">\n\x16ProductPriorityMessage\x12\x12\n\nproduct_id\x18\x01 \x01(\t\x12\x10\n\x08priority\x18\x02 \x01(\x05\"\x81\x01\n\x11ProductGroupQuery\x12\x19\n\x0csearch_query\x18\x01 \x01(\tH\x00\x88\x01\x01\x12\x31\n\ngroup_type\x18\x02 \x01(\x0e\x32\x18.protos.ProductGroupTypeH\x01\x88\x01\x01\x42\x0f\n\r_search_queryB\r\n\x0b_group_type*L\n\x10ProductGroupType\x12\x1a\n\x16PRODUCT_GROUP_TIMEDEAL\x10\x00\x12\x1c\n\x18PRODUCT_GROUP_EXHIBITION\x10\x01\x32\xb5\x04\n\x0cProductGroup\x12R\n\x0fGetProductGroup\x12\x1e.protos.GetProductGroupRequest\x1a\x1f.protos.GetProductGroupResponse\x12[\n\x12\x43reateProductGroup\x12!.protos.CreateProductGroupRequest\x1a\".protos.CreateProductGroupResponse\x12X\n\x11ListProductGroups\x12 .protos.ListProductGroupsRequest\x1a!.protos.ListProductGroupsResponse\x12U\n\x10\x45\x64itProductGroup\x12\x1f.protos.EditProductGroupRequest\x1a .protos.EditProductGroupResponse\x12_\n\x1aPushProductsInProductGroup\x12\x1f.protos.PushProductsInPgRequest\x1a .protos.PushProductsInPgResponse\x12\x62\n\x1bRemoveProductInProductGroup\x12 .protos.RemoveProductInPgRequest\x1a!.protos.RemoveProductInPgResponseB8Z6github.com/lessbutter/alloff-grpc-protos/gen/go/protosb\x06proto3')

_PRODUCTGROUPTYPE = DESCRIPTOR.enum_types_by_name['ProductGroupType']
ProductGroupType = enum_type_wrapper.EnumTypeWrapper(_PRODUCTGROUPTYPE)
PRODUCT_GROUP_TIMEDEAL = 0
PRODUCT_GROUP_EXHIBITION = 1


_GETPRODUCTGROUPREQUEST = DESCRIPTOR.message_types_by_name['GetProductGroupRequest']
_GETPRODUCTGROUPRESPONSE = DESCRIPTOR.message_types_by_name['GetProductGroupResponse']
_CREATEPRODUCTGROUPREQUEST = DESCRIPTOR.message_types_by_name['CreateProductGroupRequest']
_CREATEPRODUCTGROUPRESPONSE = DESCRIPTOR.message_types_by_name['CreateProductGroupResponse']
_EDITPRODUCTGROUPREQUEST = DESCRIPTOR.message_types_by_name['EditProductGroupRequest']
_EDITPRODUCTGROUPRESPONSE = DESCRIPTOR.message_types_by_name['EditProductGroupResponse']
_LISTPRODUCTGROUPSREQUEST = DESCRIPTOR.message_types_by_name['ListProductGroupsRequest']
_LISTPRODUCTGROUPSRESPONSE = DESCRIPTOR.message_types_by_name['ListProductGroupsResponse']
_PUSHPRODUCTSINPGREQUEST = DESCRIPTOR.message_types_by_name['PushProductsInPgRequest']
_PUSHPRODUCTSINPGRESPONSE = DESCRIPTOR.message_types_by_name['PushProductsInPgResponse']
_REMOVEPRODUCTINPGREQUEST = DESCRIPTOR.message_types_by_name['RemoveProductInPgRequest']
_REMOVEPRODUCTINPGRESPONSE = DESCRIPTOR.message_types_by_name['RemoveProductInPgResponse']
_PRODUCTINGROUPMESSAGE = DESCRIPTOR.message_types_by_name['ProductInGroupMessage']
_PRODUCTGROUPMESSAGE = DESCRIPTOR.message_types_by_name['ProductGroupMessage']
_PRODUCTPRIORITYMESSAGE = DESCRIPTOR.message_types_by_name['ProductPriorityMessage']
_PRODUCTGROUPQUERY = DESCRIPTOR.message_types_by_name['ProductGroupQuery']
GetProductGroupRequest = _reflection.GeneratedProtocolMessageType('GetProductGroupRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETPRODUCTGROUPREQUEST,
  '__module__' : 'productGroup_pb2'
  # @@protoc_insertion_point(class_scope:protos.GetProductGroupRequest)
  })
_sym_db.RegisterMessage(GetProductGroupRequest)

GetProductGroupResponse = _reflection.GeneratedProtocolMessageType('GetProductGroupResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETPRODUCTGROUPRESPONSE,
  '__module__' : 'productGroup_pb2'
  # @@protoc_insertion_point(class_scope:protos.GetProductGroupResponse)
  })
_sym_db.RegisterMessage(GetProductGroupResponse)

CreateProductGroupRequest = _reflection.GeneratedProtocolMessageType('CreateProductGroupRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEPRODUCTGROUPREQUEST,
  '__module__' : 'productGroup_pb2'
  # @@protoc_insertion_point(class_scope:protos.CreateProductGroupRequest)
  })
_sym_db.RegisterMessage(CreateProductGroupRequest)

CreateProductGroupResponse = _reflection.GeneratedProtocolMessageType('CreateProductGroupResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATEPRODUCTGROUPRESPONSE,
  '__module__' : 'productGroup_pb2'
  # @@protoc_insertion_point(class_scope:protos.CreateProductGroupResponse)
  })
_sym_db.RegisterMessage(CreateProductGroupResponse)

EditProductGroupRequest = _reflection.GeneratedProtocolMessageType('EditProductGroupRequest', (_message.Message,), {
  'DESCRIPTOR' : _EDITPRODUCTGROUPREQUEST,
  '__module__' : 'productGroup_pb2'
  # @@protoc_insertion_point(class_scope:protos.EditProductGroupRequest)
  })
_sym_db.RegisterMessage(EditProductGroupRequest)

EditProductGroupResponse = _reflection.GeneratedProtocolMessageType('EditProductGroupResponse', (_message.Message,), {
  'DESCRIPTOR' : _EDITPRODUCTGROUPRESPONSE,
  '__module__' : 'productGroup_pb2'
  # @@protoc_insertion_point(class_scope:protos.EditProductGroupResponse)
  })
_sym_db.RegisterMessage(EditProductGroupResponse)

ListProductGroupsRequest = _reflection.GeneratedProtocolMessageType('ListProductGroupsRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTPRODUCTGROUPSREQUEST,
  '__module__' : 'productGroup_pb2'
  # @@protoc_insertion_point(class_scope:protos.ListProductGroupsRequest)
  })
_sym_db.RegisterMessage(ListProductGroupsRequest)

ListProductGroupsResponse = _reflection.GeneratedProtocolMessageType('ListProductGroupsResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTPRODUCTGROUPSRESPONSE,
  '__module__' : 'productGroup_pb2'
  # @@protoc_insertion_point(class_scope:protos.ListProductGroupsResponse)
  })
_sym_db.RegisterMessage(ListProductGroupsResponse)

PushProductsInPgRequest = _reflection.GeneratedProtocolMessageType('PushProductsInPgRequest', (_message.Message,), {
  'DESCRIPTOR' : _PUSHPRODUCTSINPGREQUEST,
  '__module__' : 'productGroup_pb2'
  # @@protoc_insertion_point(class_scope:protos.PushProductsInPgRequest)
  })
_sym_db.RegisterMessage(PushProductsInPgRequest)

PushProductsInPgResponse = _reflection.GeneratedProtocolMessageType('PushProductsInPgResponse', (_message.Message,), {
  'DESCRIPTOR' : _PUSHPRODUCTSINPGRESPONSE,
  '__module__' : 'productGroup_pb2'
  # @@protoc_insertion_point(class_scope:protos.PushProductsInPgResponse)
  })
_sym_db.RegisterMessage(PushProductsInPgResponse)

RemoveProductInPgRequest = _reflection.GeneratedProtocolMessageType('RemoveProductInPgRequest', (_message.Message,), {
  'DESCRIPTOR' : _REMOVEPRODUCTINPGREQUEST,
  '__module__' : 'productGroup_pb2'
  # @@protoc_insertion_point(class_scope:protos.RemoveProductInPgRequest)
  })
_sym_db.RegisterMessage(RemoveProductInPgRequest)

RemoveProductInPgResponse = _reflection.GeneratedProtocolMessageType('RemoveProductInPgResponse', (_message.Message,), {
  'DESCRIPTOR' : _REMOVEPRODUCTINPGRESPONSE,
  '__module__' : 'productGroup_pb2'
  # @@protoc_insertion_point(class_scope:protos.RemoveProductInPgResponse)
  })
_sym_db.RegisterMessage(RemoveProductInPgResponse)

ProductInGroupMessage = _reflection.GeneratedProtocolMessageType('ProductInGroupMessage', (_message.Message,), {
  'DESCRIPTOR' : _PRODUCTINGROUPMESSAGE,
  '__module__' : 'productGroup_pb2'
  # @@protoc_insertion_point(class_scope:protos.ProductInGroupMessage)
  })
_sym_db.RegisterMessage(ProductInGroupMessage)

ProductGroupMessage = _reflection.GeneratedProtocolMessageType('ProductGroupMessage', (_message.Message,), {
  'DESCRIPTOR' : _PRODUCTGROUPMESSAGE,
  '__module__' : 'productGroup_pb2'
  # @@protoc_insertion_point(class_scope:protos.ProductGroupMessage)
  })
_sym_db.RegisterMessage(ProductGroupMessage)

ProductPriorityMessage = _reflection.GeneratedProtocolMessageType('ProductPriorityMessage', (_message.Message,), {
  'DESCRIPTOR' : _PRODUCTPRIORITYMESSAGE,
  '__module__' : 'productGroup_pb2'
  # @@protoc_insertion_point(class_scope:protos.ProductPriorityMessage)
  })
_sym_db.RegisterMessage(ProductPriorityMessage)

ProductGroupQuery = _reflection.GeneratedProtocolMessageType('ProductGroupQuery', (_message.Message,), {
  'DESCRIPTOR' : _PRODUCTGROUPQUERY,
  '__module__' : 'productGroup_pb2'
  # @@protoc_insertion_point(class_scope:protos.ProductGroupQuery)
  })
_sym_db.RegisterMessage(ProductGroupQuery)

_PRODUCTGROUP = DESCRIPTOR.services_by_name['ProductGroup']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z6github.com/lessbutter/alloff-grpc-protos/gen/go/protos'
  _PRODUCTGROUPTYPE._serialized_start=1962
  _PRODUCTGROUPTYPE._serialized_end=2038
  _GETPRODUCTGROUPREQUEST._serialized_start=45
  _GETPRODUCTGROUPREQUEST._serialized_end=95
  _GETPRODUCTGROUPRESPONSE._serialized_start=97
  _GETPRODUCTGROUPRESPONSE._serialized_end=163
  _CREATEPRODUCTGROUPREQUEST._serialized_start=166
  _CREATEPRODUCTGROUPREQUEST._serialized_end=396
  _CREATEPRODUCTGROUPRESPONSE._serialized_start=398
  _CREATEPRODUCTGROUPRESPONSE._serialized_end=467
  _EDITPRODUCTGROUPREQUEST._serialized_start=470
  _EDITPRODUCTGROUPREQUEST._serialized_end=800
  _EDITPRODUCTGROUPRESPONSE._serialized_start=802
  _EDITPRODUCTGROUPRESPONSE._serialized_end=869
  _LISTPRODUCTGROUPSREQUEST._serialized_start=871
  _LISTPRODUCTGROUPSREQUEST._serialized_end=970
  _LISTPRODUCTGROUPSRESPONSE._serialized_start=972
  _LISTPRODUCTGROUPSRESPONSE._serialized_end=1094
  _PUSHPRODUCTSINPGREQUEST._serialized_start=1096
  _PUSHPRODUCTSINPGREQUEST._serialized_end=1205
  _PUSHPRODUCTSINPGRESPONSE._serialized_start=1207
  _PUSHPRODUCTSINPGRESPONSE._serialized_end=1274
  _REMOVEPRODUCTINPGREQUEST._serialized_start=1276
  _REMOVEPRODUCTINPGREQUEST._serialized_end=1348
  _REMOVEPRODUCTINPGRESPONSE._serialized_start=1350
  _REMOVEPRODUCTINPGRESPONSE._serialized_end=1418
  _PRODUCTINGROUPMESSAGE._serialized_start=1420
  _PRODUCTINGROUPMESSAGE._serialized_end=1502
  _PRODUCTGROUPMESSAGE._serialized_start=1505
  _PRODUCTGROUPMESSAGE._serialized_end=1764
  _PRODUCTPRIORITYMESSAGE._serialized_start=1766
  _PRODUCTPRIORITYMESSAGE._serialized_end=1828
  _PRODUCTGROUPQUERY._serialized_start=1831
  _PRODUCTGROUPQUERY._serialized_end=1960
  _PRODUCTGROUP._serialized_start=2041
  _PRODUCTGROUP._serialized_end=2606
# @@protoc_insertion_point(module_scope)
