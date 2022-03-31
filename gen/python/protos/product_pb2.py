# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protos/product.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x14protos/product.proto\x12\x06protos\".\n\x11GetProductRequest\x12\x19\n\x11\x61lloff_product_id\x18\x01 \x01(\t\"=\n\x12GetProductResponse\x12\'\n\x07product\x18\x01 \x01(\x0b\x32\x16.protos.ProductMessage\"\x83\x01\n\x13ListProductsRequest\x12#\n\x05query\x18\x01 \x01(\x0b\x32\x14.protos.ProductQuery\x12\x0e\n\x06offset\x18\x02 \x01(\x05\x12\r\n\x05limit\x18\x03 \x01(\x05\x12\x18\n\x0bmodule_name\x18\x04 \x01(\tH\x00\x88\x01\x01\x42\x0e\n\x0c_module_name\"\xb4\x01\n\x14ListProductsResponse\x12(\n\x08products\x18\x01 \x03(\x0b\x32\x16.protos.ProductMessage\x12\x0e\n\x06offset\x18\x02 \x01(\x05\x12\r\n\x05limit\x18\x03 \x01(\x05\x12\x14\n\x0ctotal_counts\x18\x04 \x01(\x05\x12(\n\nlist_query\x18\x05 \x01(\x0b\x32\x14.protos.ProductQuery\x12\x13\n\x0bmodule_name\x18\x06 \x01(\t\"\xef\x04\n\x14\x43reateProductRequest\x12\x13\n\x0b\x61lloff_name\x18\x01 \x01(\t\x12\x1b\n\x13is_foreign_delivery\x18\x02 \x01(\x08\x12\x17\n\nproduct_id\x18\x03 \x01(\tH\x00\x88\x01\x01\x12\x1b\n\x0eoriginal_price\x18\x04 \x01(\x05H\x01\x88\x01\x01\x12\x18\n\x10\x64iscounted_price\x18\x05 \x01(\x05\x12\x1a\n\rspecial_price\x18\x06 \x01(\x05H\x02\x88\x01\x01\x12\x16\n\x0e\x62rand_key_name\x18\x07 \x01(\t\x12\x32\n\tinventory\x18\x08 \x03(\x0b\x32\x1f.protos.ProductInventoryMessage\x12\x13\n\x0b\x64\x65scription\x18\t \x03(\t\x12\x1a\n\x12is_refund_possible\x18\n \x01(\x08\x12\x0e\n\x06images\x18\x0b \x03(\t\x12\x1a\n\x12\x64\x65scription_images\x18\x0c \x03(\t\x12\x1e\n\x16\x65\x61rliest_delivery_days\x18\r \x01(\x05\x12\x1c\n\x14latest_delivery_days\x18\x0e \x01(\x05\x12\x12\n\nrefund_fee\x18\x0f \x01(\x05\x12\x18\n\x0bmodule_name\x18\x10 \x01(\tH\x03\x88\x01\x01\x12\x1f\n\x12\x61lloff_category_id\x18\x11 \x01(\tH\x04\x88\x01\x01\x12\x18\n\x0bproduct_url\x18\x12 \x01(\tH\x05\x88\x01\x01\x42\r\n\x0b_product_idB\x11\n\x0f_original_priceB\x10\n\x0e_special_priceB\x0e\n\x0c_module_nameB\x15\n\x13_alloff_category_idB\x0e\n\x0c_product_url\"@\n\x15\x43reateProductResponse\x12\'\n\x07product\x18\x01 \x01(\x0b\x32\x16.protos.ProductMessage\"\x95\x07\n\x12\x45\x64itProductRequest\x12\x18\n\x0b\x61lloff_name\x18\x01 \x01(\tH\x00\x88\x01\x01\x12 \n\x13is_foreign_delivery\x18\x02 \x01(\x08H\x01\x88\x01\x01\x12\x17\n\nproduct_id\x18\x03 \x01(\tH\x02\x88\x01\x01\x12\x1b\n\x0eoriginal_price\x18\x04 \x01(\x05H\x03\x88\x01\x01\x12\x1d\n\x10\x64iscounted_price\x18\x05 \x01(\x05H\x04\x88\x01\x01\x12\x1a\n\rspecial_price\x18\x06 \x01(\x05H\x05\x88\x01\x01\x12\x1b\n\x0e\x62rand_key_name\x18\x07 \x01(\tH\x06\x88\x01\x01\x12\x32\n\tinventory\x18\x08 \x03(\x0b\x32\x1f.protos.ProductInventoryMessage\x12\x13\n\x0b\x64\x65scription\x18\t \x03(\t\x12\x1f\n\x12is_refund_possible\x18\n \x01(\x08H\x07\x88\x01\x01\x12\x0e\n\x06images\x18\x0b \x03(\t\x12\x1a\n\x12\x64\x65scription_images\x18\x0c \x03(\t\x12\x17\n\nis_removed\x18\r \x01(\x08H\x08\x88\x01\x01\x12#\n\x16\x65\x61rliest_delivery_days\x18\x0e \x01(\x05H\t\x88\x01\x01\x12!\n\x14latest_delivery_days\x18\x0f \x01(\x05H\n\x88\x01\x01\x12\x17\n\nrefund_fee\x18\x10 \x01(\x05H\x0b\x88\x01\x01\x12\x19\n\x11\x61lloff_product_id\x18\x11 \x01(\t\x12\x13\n\x0bmodule_name\x18\x12 \x01(\t\x12\x1f\n\x12\x61lloff_category_id\x18\x13 \x01(\tH\x0c\x88\x01\x01\x12\x17\n\nis_soldout\x18\x14 \x01(\x08H\r\x88\x01\x01\x12\x18\n\x0bproduct_url\x18\x15 \x01(\tH\x0e\x88\x01\x01\x42\x0e\n\x0c_alloff_nameB\x16\n\x14_is_foreign_deliveryB\r\n\x0b_product_idB\x11\n\x0f_original_priceB\x13\n\x11_discounted_priceB\x10\n\x0e_special_priceB\x11\n\x0f_brand_key_nameB\x15\n\x13_is_refund_possibleB\r\n\x0b_is_removedB\x19\n\x17_earliest_delivery_daysB\x17\n\x15_latest_delivery_daysB\r\n\x0b_refund_feeB\x15\n\x13_alloff_category_idB\r\n\x0b_is_soldoutB\x0e\n\x0c_product_url\">\n\x13\x45\x64itProductResponse\x12\'\n\x07product\x18\x01 \x01(\x0b\x32\x16.protos.ProductMessage\"\xa8\x05\n\x0eProductMessage\x12\x13\n\x0b\x61lloff_name\x18\x01 \x01(\t\x12\x1b\n\x13is_foreign_delivery\x18\x02 \x01(\x08\x12\x12\n\nproduct_id\x18\x03 \x01(\t\x12\x16\n\x0eoriginal_price\x18\x04 \x01(\x05\x12\x18\n\x10\x64iscounted_price\x18\x05 \x01(\x05\x12\x15\n\rspecial_price\x18\x06 \x01(\x05\x12\x16\n\x0e\x62rand_kor_name\x18\x07 \x01(\t\x12\x32\n\tinventory\x18\x08 \x03(\x0b\x32\x1f.protos.ProductInventoryMessage\x12\x13\n\x0b\x64\x65scription\x18\t \x03(\t\x12\x1e\n\x16\x65\x61rliest_delivery_days\x18\n \x01(\x05\x12\x1c\n\x14latest_delivery_days\x18\x0b \x01(\x05\x12\x12\n\nrefund_fee\x18\x0c \x01(\x05\x12\x1a\n\x12is_refund_possible\x18\r \x01(\x08\x12\x0e\n\x06images\x18\x0e \x03(\t\x12\x1a\n\x12\x64\x65scription_images\x18\x0f \x03(\t\x12\x15\n\rcategory_name\x18\x10 \x01(\t\x12\x1c\n\x14\x61lloff_category_name\x18\x11 \x01(\t\x12\x12\n\nis_removed\x18\x12 \x01(\x08\x12\x12\n\nis_soldout\x18\x13 \x01(\x08\x12\x13\n\x0btotal_score\x18\x14 \x01(\x05\x12\x19\n\x11\x61lloff_product_id\x18\x15 \x01(\t\x12\x13\n\x0bmodule_name\x18\x16 \x01(\t\x12\x1a\n\x12\x61lloff_category_id\x18\x17 \x01(\t\x12\x13\n\x0bproduct_url\x18\x18 \x01(\t\x12\x1a\n\x12is_classified_done\x18\x19 \x01(\x08\x12\x1d\n\x15is_classified_touched\x18\x1a \x01(\x08\"9\n\x17ProductInventoryMessage\x12\x0c\n\x04size\x18\x01 \x01(\t\x12\x10\n\x08quantity\x18\x02 \x01(\x05\"\xa1\x02\n\x0cProductQuery\x12\x19\n\x0csearch_query\x18\x01 \x01(\tH\x00\x88\x01\x01\x12\x15\n\x08\x62rand_id\x18\x02 \x01(\tH\x01\x88\x01\x01\x12\x18\n\x0b\x63\x61tegory_id\x18\x03 \x01(\tH\x02\x88\x01\x01\x12\x1f\n\x12\x61lloff_category_id\x18\x04 \x01(\tH\x03\x88\x01\x01\x12\'\n\x07options\x18\x05 \x03(\x0e\x32\x16.protos.SortingOptions\x12\x1f\n\x12is_classified_done\x18\x06 \x01(\x08H\x04\x88\x01\x01\x42\x0f\n\r_search_queryB\x0b\n\t_brand_idB\x0e\n\x0c_category_idB\x15\n\x13_alloff_category_idB\x15\n\x13_is_classified_done*\xc4\x01\n\x0eSortingOptions\x12\x13\n\x0fPRICE_ASCENDING\x10\x00\x12\x14\n\x10PRICE_DESCENDING\x10\x01\x12\x11\n\rDISCOUNT_0_30\x10\x02\x12\x12\n\x0e\x44ISCOUNT_30_50\x10\x03\x12\x12\n\x0e\x44ISCOUNT_50_70\x10\x04\x12\x13\n\x0f\x44ISCOUNT_70_100\x10\x05\x12\x1a\n\x16\x44ISCOUNTRATE_ASCENDING\x10\x06\x12\x1b\n\x17\x44ISCOUNTRATE_DESCENDING\x10\x07\x32\xaf\x02\n\x07Product\x12\x43\n\nGetProduct\x12\x19.protos.GetProductRequest\x1a\x1a.protos.GetProductResponse\x12I\n\x0cListProducts\x12\x1b.protos.ListProductsRequest\x1a\x1c.protos.ListProductsResponse\x12L\n\rCreateProduct\x12\x1c.protos.CreateProductRequest\x1a\x1d.protos.CreateProductResponse\x12\x46\n\x0b\x45\x64itProduct\x12\x1a.protos.EditProductRequest\x1a\x1b.protos.EditProductResponseB<Z:github.com/lessbutter/alloff-grpc-protos/gen/golang/protosb\x06proto3')

_SORTINGOPTIONS = DESCRIPTOR.enum_types_by_name['SortingOptions']
SortingOptions = enum_type_wrapper.EnumTypeWrapper(_SORTINGOPTIONS)
PRICE_ASCENDING = 0
PRICE_DESCENDING = 1
DISCOUNT_0_30 = 2
DISCOUNT_30_50 = 3
DISCOUNT_50_70 = 4
DISCOUNT_70_100 = 5
DISCOUNTRATE_ASCENDING = 6
DISCOUNTRATE_DESCENDING = 7


_GETPRODUCTREQUEST = DESCRIPTOR.message_types_by_name['GetProductRequest']
_GETPRODUCTRESPONSE = DESCRIPTOR.message_types_by_name['GetProductResponse']
_LISTPRODUCTSREQUEST = DESCRIPTOR.message_types_by_name['ListProductsRequest']
_LISTPRODUCTSRESPONSE = DESCRIPTOR.message_types_by_name['ListProductsResponse']
_CREATEPRODUCTREQUEST = DESCRIPTOR.message_types_by_name['CreateProductRequest']
_CREATEPRODUCTRESPONSE = DESCRIPTOR.message_types_by_name['CreateProductResponse']
_EDITPRODUCTREQUEST = DESCRIPTOR.message_types_by_name['EditProductRequest']
_EDITPRODUCTRESPONSE = DESCRIPTOR.message_types_by_name['EditProductResponse']
_PRODUCTMESSAGE = DESCRIPTOR.message_types_by_name['ProductMessage']
_PRODUCTINVENTORYMESSAGE = DESCRIPTOR.message_types_by_name['ProductInventoryMessage']
_PRODUCTQUERY = DESCRIPTOR.message_types_by_name['ProductQuery']
GetProductRequest = _reflection.GeneratedProtocolMessageType('GetProductRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETPRODUCTREQUEST,
  '__module__' : 'protos.product_pb2'
  # @@protoc_insertion_point(class_scope:protos.GetProductRequest)
  })
_sym_db.RegisterMessage(GetProductRequest)

GetProductResponse = _reflection.GeneratedProtocolMessageType('GetProductResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETPRODUCTRESPONSE,
  '__module__' : 'protos.product_pb2'
  # @@protoc_insertion_point(class_scope:protos.GetProductResponse)
  })
_sym_db.RegisterMessage(GetProductResponse)

ListProductsRequest = _reflection.GeneratedProtocolMessageType('ListProductsRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTPRODUCTSREQUEST,
  '__module__' : 'protos.product_pb2'
  # @@protoc_insertion_point(class_scope:protos.ListProductsRequest)
  })
_sym_db.RegisterMessage(ListProductsRequest)

ListProductsResponse = _reflection.GeneratedProtocolMessageType('ListProductsResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTPRODUCTSRESPONSE,
  '__module__' : 'protos.product_pb2'
  # @@protoc_insertion_point(class_scope:protos.ListProductsResponse)
  })
_sym_db.RegisterMessage(ListProductsResponse)

CreateProductRequest = _reflection.GeneratedProtocolMessageType('CreateProductRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEPRODUCTREQUEST,
  '__module__' : 'protos.product_pb2'
  # @@protoc_insertion_point(class_scope:protos.CreateProductRequest)
  })
_sym_db.RegisterMessage(CreateProductRequest)

CreateProductResponse = _reflection.GeneratedProtocolMessageType('CreateProductResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATEPRODUCTRESPONSE,
  '__module__' : 'protos.product_pb2'
  # @@protoc_insertion_point(class_scope:protos.CreateProductResponse)
  })
_sym_db.RegisterMessage(CreateProductResponse)

EditProductRequest = _reflection.GeneratedProtocolMessageType('EditProductRequest', (_message.Message,), {
  'DESCRIPTOR' : _EDITPRODUCTREQUEST,
  '__module__' : 'protos.product_pb2'
  # @@protoc_insertion_point(class_scope:protos.EditProductRequest)
  })
_sym_db.RegisterMessage(EditProductRequest)

EditProductResponse = _reflection.GeneratedProtocolMessageType('EditProductResponse', (_message.Message,), {
  'DESCRIPTOR' : _EDITPRODUCTRESPONSE,
  '__module__' : 'protos.product_pb2'
  # @@protoc_insertion_point(class_scope:protos.EditProductResponse)
  })
_sym_db.RegisterMessage(EditProductResponse)

ProductMessage = _reflection.GeneratedProtocolMessageType('ProductMessage', (_message.Message,), {
  'DESCRIPTOR' : _PRODUCTMESSAGE,
  '__module__' : 'protos.product_pb2'
  # @@protoc_insertion_point(class_scope:protos.ProductMessage)
  })
_sym_db.RegisterMessage(ProductMessage)

ProductInventoryMessage = _reflection.GeneratedProtocolMessageType('ProductInventoryMessage', (_message.Message,), {
  'DESCRIPTOR' : _PRODUCTINVENTORYMESSAGE,
  '__module__' : 'protos.product_pb2'
  # @@protoc_insertion_point(class_scope:protos.ProductInventoryMessage)
  })
_sym_db.RegisterMessage(ProductInventoryMessage)

ProductQuery = _reflection.GeneratedProtocolMessageType('ProductQuery', (_message.Message,), {
  'DESCRIPTOR' : _PRODUCTQUERY,
  '__module__' : 'protos.product_pb2'
  # @@protoc_insertion_point(class_scope:protos.ProductQuery)
  })
_sym_db.RegisterMessage(ProductQuery)

_PRODUCT = DESCRIPTOR.services_by_name['Product']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z:github.com/lessbutter/alloff-grpc-protos/gen/golang/protos'
  _SORTINGOPTIONS._serialized_start=3171
  _SORTINGOPTIONS._serialized_end=3367
  _GETPRODUCTREQUEST._serialized_start=32
  _GETPRODUCTREQUEST._serialized_end=78
  _GETPRODUCTRESPONSE._serialized_start=80
  _GETPRODUCTRESPONSE._serialized_end=141
  _LISTPRODUCTSREQUEST._serialized_start=144
  _LISTPRODUCTSREQUEST._serialized_end=275
  _LISTPRODUCTSRESPONSE._serialized_start=278
  _LISTPRODUCTSRESPONSE._serialized_end=458
  _CREATEPRODUCTREQUEST._serialized_start=461
  _CREATEPRODUCTREQUEST._serialized_end=1084
  _CREATEPRODUCTRESPONSE._serialized_start=1086
  _CREATEPRODUCTRESPONSE._serialized_end=1150
  _EDITPRODUCTREQUEST._serialized_start=1153
  _EDITPRODUCTREQUEST._serialized_end=2070
  _EDITPRODUCTRESPONSE._serialized_start=2072
  _EDITPRODUCTRESPONSE._serialized_end=2134
  _PRODUCTMESSAGE._serialized_start=2137
  _PRODUCTMESSAGE._serialized_end=2817
  _PRODUCTINVENTORYMESSAGE._serialized_start=2819
  _PRODUCTINVENTORYMESSAGE._serialized_end=2876
  _PRODUCTQUERY._serialized_start=2879
  _PRODUCTQUERY._serialized_end=3168
  _PRODUCT._serialized_start=3370
  _PRODUCT._serialized_end=3673
# @@protoc_insertion_point(module_scope)
