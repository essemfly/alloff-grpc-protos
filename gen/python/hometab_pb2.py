# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: hometab.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import exhibition_pb2 as exhibition__pb2
import product_pb2 as product__pb2
import brand_pb2 as brand__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\rhometab.proto\x12\x06golang\x1a\x10\x65xhibition.proto\x1a\rproduct.proto\x1a\x0b\x62rand.proto\"(\n\x15GetHomeTabItemRequest\x12\x0f\n\x07item_id\x18\x01 \x01(\t\"8\n\x17ListHomeTabItemsRequest\x12\x0e\n\x06offset\x18\x01 \x01(\x05\x12\r\n\x05limit\x18\x02 \x01(\x05\"\xef\x02\n\x16\x45\x64itHomeTabItemRequest\x12\x12\n\nhometab_id\x18\x01 \x01(\t\x12\x12\n\x05title\x18\x02 \x01(\tH\x00\x88\x01\x01\x12\x18\n\x0b\x64\x65scription\x18\x03 \x01(\tH\x01\x88\x01\x01\x12\x0c\n\x04tags\x18\x04 \x03(\t\x12\x1b\n\x0e\x62\x61\x63k_image_url\x18\x05 \x01(\tH\x02\x88\x01\x01\x12\x17\n\nstart_time\x18\x06 \x01(\tH\x03\x88\x01\x01\x12\x18\n\x0b\x66inish_time\x18\x07 \x01(\tH\x04\x88\x01\x01\x12\'\n\x08\x63ontents\x18\x08 \x01(\x0b\x32\x15.golang.ItemRequester\x12\x13\n\x06weight\x18\t \x01(\x05H\x05\x88\x01\x01\x12\x14\n\x07is_live\x18\n \x01(\x08H\x06\x88\x01\x01\x42\x08\n\x06_titleB\x0e\n\x0c_descriptionB\x11\n\x0f_back_image_urlB\r\n\x0b_start_timeB\x0e\n\x0c_finish_timeB\t\n\x07_weightB\n\n\x08_is_live\"\xde\x01\n\x18\x43reateHomeTabItemRequest\x12\r\n\x05title\x18\x01 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x02 \x01(\t\x12\x0c\n\x04tags\x18\x03 \x03(\t\x12\x1b\n\x0e\x62\x61\x63k_image_url\x18\x04 \x01(\tH\x00\x88\x01\x01\x12\x12\n\nstart_time\x18\x05 \x01(\t\x12\x13\n\x0b\x66inish_time\x18\x06 \x01(\t\x12\'\n\x08\x63ontents\x18\x07 \x01(\x0b\x32\x15.golang.ItemRequester\x12\x0e\n\x06weight\x18\x08 \x01(\x05\x42\x11\n\x0f_back_image_url\"B\n\x16GetHomeTabItemResponse\x12(\n\x04item\x18\x01 \x01(\x0b\x32\x1a.golang.HomeTabItemMessage\"z\n\x18ListHomeTabItemsResponse\x12)\n\x05items\x18\x01 \x03(\x0b\x32\x1a.golang.HomeTabItemMessage\x12\x0e\n\x06offset\x18\x02 \x01(\x05\x12\r\n\x05limit\x18\x03 \x01(\x05\x12\x14\n\x0ctotal_counts\x18\x04 \x01(\x05\"C\n\x17\x45\x64itHomeTabItemResponse\x12(\n\x04item\x18\x01 \x01(\x0b\x32\x1a.golang.HomeTabItemMessage\"E\n\x19\x43reateHomeTabItemResponse\x12(\n\x04item\x18\x01 \x01(\x0b\x32\x1a.golang.HomeTabItemMessage\"\x96\x03\n\x12HomeTabItemMessage\x12\x0f\n\x07item_id\x18\x01 \x01(\t\x12\r\n\x05title\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12\x0c\n\x04tags\x18\x04 \x03(\t\x12\x16\n\x0e\x62\x61\x63k_image_url\x18\x05 \x01(\t\x12#\n\titem_type\x18\x06 \x01(\x0e\x32\x10.golang.ItemType\x12\x12\n\nstart_time\x18\x07 \x01(\t\x12\x13\n\x0b\x66inish_time\x18\x08 \x01(\t\x12(\n\x08products\x18\t \x03(\x0b\x32\x16.golang.ProductMessage\x12$\n\x06\x62rands\x18\n \x03(\x0b\x32\x14.golang.BrandMessage\x12.\n\x0b\x65xhibitions\x18\x0b \x03(\x0b\x32\x19.golang.ExhibitionMessage\x12\x36\n\treference\x18\x0c \x01(\x0b\x32#.golang.HomeTabItemReferenceMessage\x12\x0e\n\x06weight\x18\r \x01(\x05\x12\x0f\n\x07is_live\x18\x0e \x01(\x08\"d\n\x1bHomeTabItemReferenceMessage\x12\x0c\n\x04path\x18\x01 \x01(\t\x12\x0e\n\x06params\x18\x02 \x01(\t\x12\'\n\x07options\x18\x03 \x03(\x0e\x32\x16.golang.SortingOptions\"\xd8\x01\n\rItemRequester\x12#\n\titem_type\x18\x01 \x01(\x0e\x32\x10.golang.ItemType\x12\x16\n\x0e\x62rand_keynames\x18\x02 \x03(\t\x12\x16\n\x0e\x65xhibition_ids\x18\x03 \x03(\t\x12\x1e\n\x11\x61lloffcategory_id\x18\x04 \x01(\tH\x00\x88\x01\x01\x12\'\n\x07options\x18\x05 \x03(\x0e\x32\x16.golang.SortingOptions\x12\x13\n\x0bproduct_ids\x18\x06 \x03(\tB\x14\n\x12_alloffcategory_id*\xc5\x01\n\x08ItemType\x12\x17\n\x13HOMETAB_ITEM_BRANDS\x10\x00\x12\x1d\n\x19HOMETAB_ITEM_EXHIBITION_A\x10\x01\x12\x1c\n\x18HOMETAB_ITEM_EXHIBITIONS\x10\x02\x12\x1b\n\x17HOMETAB_ITEM_EXHIBITION\x10\x03\x12 \n\x1cHOMETAB_ITEM_PRODUCTS_BRANDS\x10\x04\x12$\n HOMETAB_ITEM_PRODUCTS_CATEGORIES\x10\x05\x32\xe3\x02\n\x0bHomeTabItem\x12O\n\x0eGetHomeTabItem\x12\x1d.golang.GetHomeTabItemRequest\x1a\x1e.golang.GetHomeTabItemResponse\x12U\n\x10ListHomeTabItems\x12\x1f.golang.ListHomeTabItemsRequest\x1a .golang.ListHomeTabItemsResponse\x12R\n\x0f\x45\x64itHomeTabItem\x12\x1e.golang.EditHomeTabItemRequest\x1a\x1f.golang.EditHomeTabItemResponse\x12X\n\x11\x43reateHomeTabItem\x12 .golang.CreateHomeTabItemRequest\x1a!.golang.CreateHomeTabItemResponseB5Z3github.com/lessbutter/alloff-grpc-protos/gen/golangb\x06proto3')

_ITEMTYPE = DESCRIPTOR.enum_types_by_name['ItemType']
ItemType = enum_type_wrapper.EnumTypeWrapper(_ITEMTYPE)
HOMETAB_ITEM_BRANDS = 0
HOMETAB_ITEM_EXHIBITION_A = 1
HOMETAB_ITEM_EXHIBITIONS = 2
HOMETAB_ITEM_EXHIBITION = 3
HOMETAB_ITEM_PRODUCTS_BRANDS = 4
HOMETAB_ITEM_PRODUCTS_CATEGORIES = 5


_GETHOMETABITEMREQUEST = DESCRIPTOR.message_types_by_name['GetHomeTabItemRequest']
_LISTHOMETABITEMSREQUEST = DESCRIPTOR.message_types_by_name['ListHomeTabItemsRequest']
_EDITHOMETABITEMREQUEST = DESCRIPTOR.message_types_by_name['EditHomeTabItemRequest']
_CREATEHOMETABITEMREQUEST = DESCRIPTOR.message_types_by_name['CreateHomeTabItemRequest']
_GETHOMETABITEMRESPONSE = DESCRIPTOR.message_types_by_name['GetHomeTabItemResponse']
_LISTHOMETABITEMSRESPONSE = DESCRIPTOR.message_types_by_name['ListHomeTabItemsResponse']
_EDITHOMETABITEMRESPONSE = DESCRIPTOR.message_types_by_name['EditHomeTabItemResponse']
_CREATEHOMETABITEMRESPONSE = DESCRIPTOR.message_types_by_name['CreateHomeTabItemResponse']
_HOMETABITEMMESSAGE = DESCRIPTOR.message_types_by_name['HomeTabItemMessage']
_HOMETABITEMREFERENCEMESSAGE = DESCRIPTOR.message_types_by_name['HomeTabItemReferenceMessage']
_ITEMREQUESTER = DESCRIPTOR.message_types_by_name['ItemRequester']
GetHomeTabItemRequest = _reflection.GeneratedProtocolMessageType('GetHomeTabItemRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETHOMETABITEMREQUEST,
  '__module__' : 'hometab_pb2'
  # @@protoc_insertion_point(class_scope:golang.GetHomeTabItemRequest)
  })
_sym_db.RegisterMessage(GetHomeTabItemRequest)

ListHomeTabItemsRequest = _reflection.GeneratedProtocolMessageType('ListHomeTabItemsRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTHOMETABITEMSREQUEST,
  '__module__' : 'hometab_pb2'
  # @@protoc_insertion_point(class_scope:golang.ListHomeTabItemsRequest)
  })
_sym_db.RegisterMessage(ListHomeTabItemsRequest)

EditHomeTabItemRequest = _reflection.GeneratedProtocolMessageType('EditHomeTabItemRequest', (_message.Message,), {
  'DESCRIPTOR' : _EDITHOMETABITEMREQUEST,
  '__module__' : 'hometab_pb2'
  # @@protoc_insertion_point(class_scope:golang.EditHomeTabItemRequest)
  })
_sym_db.RegisterMessage(EditHomeTabItemRequest)

CreateHomeTabItemRequest = _reflection.GeneratedProtocolMessageType('CreateHomeTabItemRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEHOMETABITEMREQUEST,
  '__module__' : 'hometab_pb2'
  # @@protoc_insertion_point(class_scope:golang.CreateHomeTabItemRequest)
  })
_sym_db.RegisterMessage(CreateHomeTabItemRequest)

GetHomeTabItemResponse = _reflection.GeneratedProtocolMessageType('GetHomeTabItemResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETHOMETABITEMRESPONSE,
  '__module__' : 'hometab_pb2'
  # @@protoc_insertion_point(class_scope:golang.GetHomeTabItemResponse)
  })
_sym_db.RegisterMessage(GetHomeTabItemResponse)

ListHomeTabItemsResponse = _reflection.GeneratedProtocolMessageType('ListHomeTabItemsResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTHOMETABITEMSRESPONSE,
  '__module__' : 'hometab_pb2'
  # @@protoc_insertion_point(class_scope:golang.ListHomeTabItemsResponse)
  })
_sym_db.RegisterMessage(ListHomeTabItemsResponse)

EditHomeTabItemResponse = _reflection.GeneratedProtocolMessageType('EditHomeTabItemResponse', (_message.Message,), {
  'DESCRIPTOR' : _EDITHOMETABITEMRESPONSE,
  '__module__' : 'hometab_pb2'
  # @@protoc_insertion_point(class_scope:golang.EditHomeTabItemResponse)
  })
_sym_db.RegisterMessage(EditHomeTabItemResponse)

CreateHomeTabItemResponse = _reflection.GeneratedProtocolMessageType('CreateHomeTabItemResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATEHOMETABITEMRESPONSE,
  '__module__' : 'hometab_pb2'
  # @@protoc_insertion_point(class_scope:golang.CreateHomeTabItemResponse)
  })
_sym_db.RegisterMessage(CreateHomeTabItemResponse)

HomeTabItemMessage = _reflection.GeneratedProtocolMessageType('HomeTabItemMessage', (_message.Message,), {
  'DESCRIPTOR' : _HOMETABITEMMESSAGE,
  '__module__' : 'hometab_pb2'
  # @@protoc_insertion_point(class_scope:golang.HomeTabItemMessage)
  })
_sym_db.RegisterMessage(HomeTabItemMessage)

HomeTabItemReferenceMessage = _reflection.GeneratedProtocolMessageType('HomeTabItemReferenceMessage', (_message.Message,), {
  'DESCRIPTOR' : _HOMETABITEMREFERENCEMESSAGE,
  '__module__' : 'hometab_pb2'
  # @@protoc_insertion_point(class_scope:golang.HomeTabItemReferenceMessage)
  })
_sym_db.RegisterMessage(HomeTabItemReferenceMessage)

ItemRequester = _reflection.GeneratedProtocolMessageType('ItemRequester', (_message.Message,), {
  'DESCRIPTOR' : _ITEMREQUESTER,
  '__module__' : 'hometab_pb2'
  # @@protoc_insertion_point(class_scope:golang.ItemRequester)
  })
_sym_db.RegisterMessage(ItemRequester)

_HOMETABITEM = DESCRIPTOR.services_by_name['HomeTabItem']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z3github.com/lessbutter/alloff-grpc-protos/gen/golang'
  _ITEMTYPE._serialized_start=1829
  _ITEMTYPE._serialized_end=2026
  _GETHOMETABITEMREQUEST._serialized_start=71
  _GETHOMETABITEMREQUEST._serialized_end=111
  _LISTHOMETABITEMSREQUEST._serialized_start=113
  _LISTHOMETABITEMSREQUEST._serialized_end=169
  _EDITHOMETABITEMREQUEST._serialized_start=172
  _EDITHOMETABITEMREQUEST._serialized_end=539
  _CREATEHOMETABITEMREQUEST._serialized_start=542
  _CREATEHOMETABITEMREQUEST._serialized_end=764
  _GETHOMETABITEMRESPONSE._serialized_start=766
  _GETHOMETABITEMRESPONSE._serialized_end=832
  _LISTHOMETABITEMSRESPONSE._serialized_start=834
  _LISTHOMETABITEMSRESPONSE._serialized_end=956
  _EDITHOMETABITEMRESPONSE._serialized_start=958
  _EDITHOMETABITEMRESPONSE._serialized_end=1025
  _CREATEHOMETABITEMRESPONSE._serialized_start=1027
  _CREATEHOMETABITEMRESPONSE._serialized_end=1096
  _HOMETABITEMMESSAGE._serialized_start=1099
  _HOMETABITEMMESSAGE._serialized_end=1505
  _HOMETABITEMREFERENCEMESSAGE._serialized_start=1507
  _HOMETABITEMREFERENCEMESSAGE._serialized_end=1607
  _ITEMREQUESTER._serialized_start=1610
  _ITEMREQUESTER._serialized_end=1826
  _HOMETABITEM._serialized_start=2029
  _HOMETABITEM._serialized_end=2384
# @@protoc_insertion_point(module_scope)
