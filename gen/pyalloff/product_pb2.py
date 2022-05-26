# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: product.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from . import alloff_size_pb2 as alloff__size__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\rproduct.proto\x12\x08goalloff\x1a\x11\x61lloff_size.proto\".\n\x11GetProductRequest\x12\x19\n\x11\x61lloff_product_id\x18\x01 \x01(\t\"?\n\x12GetProductResponse\x12)\n\x07product\x18\x01 \x01(\x0b\x32\x18.goalloff.ProductMessage\"\x85\x01\n\x13ListProductsRequest\x12%\n\x05query\x18\x01 \x01(\x0b\x32\x16.goalloff.ProductQuery\x12\x0e\n\x06offset\x18\x02 \x01(\x05\x12\r\n\x05limit\x18\x03 \x01(\x05\x12\x18\n\x0bmodule_name\x18\x04 \x01(\tH\x00\x88\x01\x01\x42\x0e\n\x0c_module_name\"\xb8\x01\n\x14ListProductsResponse\x12*\n\x08products\x18\x01 \x03(\x0b\x32\x18.goalloff.ProductMessage\x12\x0e\n\x06offset\x18\x02 \x01(\x05\x12\r\n\x05limit\x18\x03 \x01(\x05\x12\x14\n\x0ctotal_counts\x18\x04 \x01(\x05\x12*\n\nlist_query\x18\x05 \x01(\x0b\x32\x16.goalloff.ProductQuery\x12\x13\n\x0bmodule_name\x18\x06 \x01(\t\"\xd9\x07\n\x14\x43reateProductRequest\x12\x13\n\x0b\x61lloff_name\x18\x01 \x01(\t\x12\x1b\n\x13is_foreign_delivery\x18\x02 \x01(\x08\x12\x17\n\nproduct_id\x18\x03 \x01(\tH\x00\x88\x01\x01\x12\x1b\n\x0eoriginal_price\x18\x04 \x01(\x05H\x01\x88\x01\x01\x12\x18\n\x10\x64iscounted_price\x18\x05 \x01(\x05\x12\x1a\n\rspecial_price\x18\x06 \x01(\x05H\x02\x88\x01\x01\x12\x16\n\x0e\x62rand_key_name\x18\x07 \x01(\t\x12\x34\n\tinventory\x18\x08 \x03(\x0b\x32!.goalloff.ProductInventoryMessage\x12\x13\n\x0b\x64\x65scription\x18\t \x03(\t\x12\x1a\n\x12is_refund_possible\x18\n \x01(\x08\x12\x0e\n\x06images\x18\x0b \x03(\t\x12\x1a\n\x12\x64\x65scription_images\x18\x0c \x03(\t\x12\x1e\n\x16\x65\x61rliest_delivery_days\x18\r \x01(\x05\x12\x1c\n\x14latest_delivery_days\x18\x0e \x01(\x05\x12\x12\n\nrefund_fee\x18\x0f \x01(\x05\x12\x18\n\x0bmodule_name\x18\x10 \x01(\tH\x03\x88\x01\x01\x12\x1f\n\x12\x61lloff_category_id\x18\x11 \x01(\tH\x04\x88\x01\x01\x12\x18\n\x0bproduct_url\x18\x12 \x01(\tH\x05\x88\x01\x01\x12O\n\x11\x64\x65scription_infos\x18\x13 \x03(\x0b\x32\x34.goalloff.CreateProductRequest.DescriptionInfosEntry\x12G\n\rproduct_infos\x18\x14 \x03(\x0b\x32\x30.goalloff.CreateProductRequest.ProductInfosEntry\x12\x1c\n\x0fthumbnail_image\x18\x15 \x01(\tH\x06\x88\x01\x01\x12,\n\rproduct_types\x18\x16 \x03(\x0e\x32\x15.goalloff.ProductType\x1a\x37\n\x15\x44\x65scriptionInfosEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x1a\x33\n\x11ProductInfosEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x42\r\n\x0b_product_idB\x11\n\x0f_original_priceB\x10\n\x0e_special_priceB\x0e\n\x0c_module_nameB\x15\n\x13_alloff_category_idB\x0e\n\x0c_product_urlB\x12\n\x10_thumbnail_image\"B\n\x15\x43reateProductResponse\x12)\n\x07product\x18\x01 \x01(\x0b\x32\x18.goalloff.ProductMessage\"\xfb\t\n\x12\x45\x64itProductRequest\x12\x18\n\x0b\x61lloff_name\x18\x01 \x01(\tH\x00\x88\x01\x01\x12 \n\x13is_foreign_delivery\x18\x02 \x01(\x08H\x01\x88\x01\x01\x12\x17\n\nproduct_id\x18\x03 \x01(\tH\x02\x88\x01\x01\x12\x1b\n\x0eoriginal_price\x18\x04 \x01(\x05H\x03\x88\x01\x01\x12\x1d\n\x10\x64iscounted_price\x18\x05 \x01(\x05H\x04\x88\x01\x01\x12\x1a\n\rspecial_price\x18\x06 \x01(\x05H\x05\x88\x01\x01\x12\x1b\n\x0e\x62rand_key_name\x18\x07 \x01(\tH\x06\x88\x01\x01\x12\x34\n\tinventory\x18\x08 \x03(\x0b\x32!.goalloff.ProductInventoryMessage\x12\x13\n\x0b\x64\x65scription\x18\t \x03(\t\x12\x1f\n\x12is_refund_possible\x18\n \x01(\x08H\x07\x88\x01\x01\x12\x0e\n\x06images\x18\x0b \x03(\t\x12\x1a\n\x12\x64\x65scription_images\x18\x0c \x03(\t\x12\x17\n\nis_removed\x18\r \x01(\x08H\x08\x88\x01\x01\x12#\n\x16\x65\x61rliest_delivery_days\x18\x0e \x01(\x05H\t\x88\x01\x01\x12!\n\x14latest_delivery_days\x18\x0f \x01(\x05H\n\x88\x01\x01\x12\x17\n\nrefund_fee\x18\x10 \x01(\x05H\x0b\x88\x01\x01\x12\x19\n\x11\x61lloff_product_id\x18\x11 \x01(\t\x12\x13\n\x0bmodule_name\x18\x12 \x01(\t\x12\x1f\n\x12\x61lloff_category_id\x18\x13 \x01(\tH\x0c\x88\x01\x01\x12\x17\n\nis_soldout\x18\x14 \x01(\x08H\r\x88\x01\x01\x12\x18\n\x0bproduct_url\x18\x15 \x01(\tH\x0e\x88\x01\x01\x12M\n\x11\x64\x65scription_infos\x18\x16 \x03(\x0b\x32\x32.goalloff.EditProductRequest.DescriptionInfosEntry\x12\x45\n\rproduct_infos\x18\x17 \x03(\x0b\x32..goalloff.EditProductRequest.ProductInfosEntry\x12\x1c\n\x0fthumbnail_image\x18\x18 \x01(\tH\x0f\x88\x01\x01\x12,\n\rproduct_types\x18\x19 \x03(\x0e\x32\x15.goalloff.ProductType\x1a\x37\n\x15\x44\x65scriptionInfosEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x1a\x33\n\x11ProductInfosEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x42\x0e\n\x0c_alloff_nameB\x16\n\x14_is_foreign_deliveryB\r\n\x0b_product_idB\x11\n\x0f_original_priceB\x13\n\x11_discounted_priceB\x10\n\x0e_special_priceB\x11\n\x0f_brand_key_nameB\x15\n\x13_is_refund_possibleB\r\n\x0b_is_removedB\x19\n\x17_earliest_delivery_daysB\x17\n\x15_latest_delivery_daysB\r\n\x0b_refund_feeB\x15\n\x13_alloff_category_idB\r\n\x0b_is_soldoutB\x0e\n\x0c_product_urlB\x12\n\x10_thumbnail_image\"@\n\x13\x45\x64itProductResponse\x12)\n\x07product\x18\x01 \x01(\x0b\x32\x18.goalloff.ProductMessage\"\xed\x07\n\x0eProductMessage\x12\x13\n\x0b\x61lloff_name\x18\x01 \x01(\t\x12\x1b\n\x13is_foreign_delivery\x18\x02 \x01(\x08\x12\x12\n\nproduct_id\x18\x03 \x01(\t\x12\x16\n\x0eoriginal_price\x18\x04 \x01(\x05\x12\x18\n\x10\x64iscounted_price\x18\x05 \x01(\x05\x12\x15\n\rspecial_price\x18\x06 \x01(\x05\x12\x16\n\x0e\x62rand_kor_name\x18\x07 \x01(\t\x12\x34\n\tinventory\x18\x08 \x03(\x0b\x32!.goalloff.ProductInventoryMessage\x12\x13\n\x0b\x64\x65scription\x18\t \x03(\t\x12\x1e\n\x16\x65\x61rliest_delivery_days\x18\n \x01(\x05\x12\x1c\n\x14latest_delivery_days\x18\x0b \x01(\x05\x12\x12\n\nrefund_fee\x18\x0c \x01(\x05\x12\x1a\n\x12is_refund_possible\x18\r \x01(\x08\x12\x0e\n\x06images\x18\x0e \x03(\t\x12\x1a\n\x12\x64\x65scription_images\x18\x0f \x03(\t\x12\x15\n\rcategory_name\x18\x10 \x01(\t\x12\x1c\n\x14\x61lloff_category_name\x18\x11 \x01(\t\x12\x12\n\nis_removed\x18\x12 \x01(\x08\x12\x12\n\nis_soldout\x18\x13 \x01(\x08\x12\x13\n\x0btotal_score\x18\x14 \x01(\x05\x12\x19\n\x11\x61lloff_product_id\x18\x15 \x01(\t\x12\x13\n\x0bmodule_name\x18\x16 \x01(\t\x12\x1a\n\x12\x61lloff_category_id\x18\x17 \x01(\t\x12\x13\n\x0bproduct_url\x18\x18 \x01(\t\x12\x1a\n\x12is_classified_done\x18\x19 \x01(\x08\x12\x1d\n\x15is_classified_touched\x18\x1a \x01(\x08\x12I\n\x11\x64\x65scription_infos\x18\x1b \x03(\x0b\x32..goalloff.ProductMessage.DescriptionInfosEntry\x12\x41\n\rproduct_infos\x18\x1c \x03(\x0b\x32*.goalloff.ProductMessage.ProductInfosEntry\x12\x17\n\x0fthumbnail_image\x18\x1d \x01(\t\x12,\n\rproduct_types\x18\x1e \x03(\x0e\x32\x15.goalloff.ProductType\x1a\x37\n\x15\x44\x65scriptionInfosEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x1a\x33\n\x11ProductInfosEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"l\n\x17ProductInventoryMessage\x12\x0c\n\x04size\x18\x01 \x01(\t\x12\x10\n\x08quantity\x18\x02 \x01(\x05\x12\x31\n\x0c\x61lloff_sizes\x18\x03 \x03(\x0b\x32\x1b.goalloff.AlloffSizeMessage\"\xa3\x02\n\x0cProductQuery\x12\x19\n\x0csearch_query\x18\x01 \x01(\tH\x00\x88\x01\x01\x12\x15\n\x08\x62rand_id\x18\x02 \x01(\tH\x01\x88\x01\x01\x12\x18\n\x0b\x63\x61tegory_id\x18\x03 \x01(\tH\x02\x88\x01\x01\x12\x1f\n\x12\x61lloff_category_id\x18\x04 \x01(\tH\x03\x88\x01\x01\x12)\n\x07options\x18\x05 \x03(\x0e\x32\x18.goalloff.SortingOptions\x12\x1f\n\x12is_classified_done\x18\x06 \x01(\x08H\x04\x88\x01\x01\x42\x0f\n\r_search_queryB\x0b\n\t_brand_idB\x0e\n\x0c_category_idB\x15\n\x13_alloff_category_idB\x15\n\x13_is_classified_done*9\n\x0bProductType\x12\n\n\x06\x46\x45MALE\x10\x00\x12\x08\n\x04MALE\x10\x01\x12\x08\n\x04KIDS\x10\x02\x12\n\n\x06SPORTS\x10\x03*\xc4\x01\n\x0eSortingOptions\x12\x13\n\x0fPRICE_ASCENDING\x10\x00\x12\x14\n\x10PRICE_DESCENDING\x10\x01\x12\x11\n\rDISCOUNT_0_30\x10\x02\x12\x12\n\x0e\x44ISCOUNT_30_50\x10\x03\x12\x12\n\x0e\x44ISCOUNT_50_70\x10\x04\x12\x13\n\x0f\x44ISCOUNT_70_100\x10\x05\x12\x1a\n\x16\x44ISCOUNTRATE_ASCENDING\x10\x06\x12\x1b\n\x17\x44ISCOUNTRATE_DESCENDING\x10\x07\x32\xbf\x02\n\x07Product\x12G\n\nGetProduct\x12\x1b.goalloff.GetProductRequest\x1a\x1c.goalloff.GetProductResponse\x12M\n\x0cListProducts\x12\x1d.goalloff.ListProductsRequest\x1a\x1e.goalloff.ListProductsResponse\x12P\n\rCreateProduct\x12\x1e.goalloff.CreateProductRequest\x1a\x1f.goalloff.CreateProductResponse\x12J\n\x0b\x45\x64itProduct\x12\x1c.goalloff.EditProductRequest\x1a\x1d.goalloff.EditProductResponseB7Z5github.com/lessbutter/alloff-grpc-protos/gen/goalloffb\x06proto3')

_PRODUCTTYPE = DESCRIPTOR.enum_types_by_name['ProductType']
ProductType = enum_type_wrapper.EnumTypeWrapper(_PRODUCTTYPE)
_SORTINGOPTIONS = DESCRIPTOR.enum_types_by_name['SortingOptions']
SortingOptions = enum_type_wrapper.EnumTypeWrapper(_SORTINGOPTIONS)
FEMALE = 0
MALE = 1
KIDS = 2
SPORTS = 3
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
_CREATEPRODUCTREQUEST_DESCRIPTIONINFOSENTRY = _CREATEPRODUCTREQUEST.nested_types_by_name['DescriptionInfosEntry']
_CREATEPRODUCTREQUEST_PRODUCTINFOSENTRY = _CREATEPRODUCTREQUEST.nested_types_by_name['ProductInfosEntry']
_CREATEPRODUCTRESPONSE = DESCRIPTOR.message_types_by_name['CreateProductResponse']
_EDITPRODUCTREQUEST = DESCRIPTOR.message_types_by_name['EditProductRequest']
_EDITPRODUCTREQUEST_DESCRIPTIONINFOSENTRY = _EDITPRODUCTREQUEST.nested_types_by_name['DescriptionInfosEntry']
_EDITPRODUCTREQUEST_PRODUCTINFOSENTRY = _EDITPRODUCTREQUEST.nested_types_by_name['ProductInfosEntry']
_EDITPRODUCTRESPONSE = DESCRIPTOR.message_types_by_name['EditProductResponse']
_PRODUCTMESSAGE = DESCRIPTOR.message_types_by_name['ProductMessage']
_PRODUCTMESSAGE_DESCRIPTIONINFOSENTRY = _PRODUCTMESSAGE.nested_types_by_name['DescriptionInfosEntry']
_PRODUCTMESSAGE_PRODUCTINFOSENTRY = _PRODUCTMESSAGE.nested_types_by_name['ProductInfosEntry']
_PRODUCTINVENTORYMESSAGE = DESCRIPTOR.message_types_by_name['ProductInventoryMessage']
_PRODUCTQUERY = DESCRIPTOR.message_types_by_name['ProductQuery']
GetProductRequest = _reflection.GeneratedProtocolMessageType('GetProductRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETPRODUCTREQUEST,
  '__module__' : 'product_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.GetProductRequest)
  })
_sym_db.RegisterMessage(GetProductRequest)

GetProductResponse = _reflection.GeneratedProtocolMessageType('GetProductResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETPRODUCTRESPONSE,
  '__module__' : 'product_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.GetProductResponse)
  })
_sym_db.RegisterMessage(GetProductResponse)

ListProductsRequest = _reflection.GeneratedProtocolMessageType('ListProductsRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTPRODUCTSREQUEST,
  '__module__' : 'product_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.ListProductsRequest)
  })
_sym_db.RegisterMessage(ListProductsRequest)

ListProductsResponse = _reflection.GeneratedProtocolMessageType('ListProductsResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTPRODUCTSRESPONSE,
  '__module__' : 'product_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.ListProductsResponse)
  })
_sym_db.RegisterMessage(ListProductsResponse)

CreateProductRequest = _reflection.GeneratedProtocolMessageType('CreateProductRequest', (_message.Message,), {

  'DescriptionInfosEntry' : _reflection.GeneratedProtocolMessageType('DescriptionInfosEntry', (_message.Message,), {
    'DESCRIPTOR' : _CREATEPRODUCTREQUEST_DESCRIPTIONINFOSENTRY,
    '__module__' : 'product_pb2'
    # @@protoc_insertion_point(class_scope:goalloff.CreateProductRequest.DescriptionInfosEntry)
    })
  ,

  'ProductInfosEntry' : _reflection.GeneratedProtocolMessageType('ProductInfosEntry', (_message.Message,), {
    'DESCRIPTOR' : _CREATEPRODUCTREQUEST_PRODUCTINFOSENTRY,
    '__module__' : 'product_pb2'
    # @@protoc_insertion_point(class_scope:goalloff.CreateProductRequest.ProductInfosEntry)
    })
  ,
  'DESCRIPTOR' : _CREATEPRODUCTREQUEST,
  '__module__' : 'product_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.CreateProductRequest)
  })
_sym_db.RegisterMessage(CreateProductRequest)
_sym_db.RegisterMessage(CreateProductRequest.DescriptionInfosEntry)
_sym_db.RegisterMessage(CreateProductRequest.ProductInfosEntry)

CreateProductResponse = _reflection.GeneratedProtocolMessageType('CreateProductResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATEPRODUCTRESPONSE,
  '__module__' : 'product_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.CreateProductResponse)
  })
_sym_db.RegisterMessage(CreateProductResponse)

EditProductRequest = _reflection.GeneratedProtocolMessageType('EditProductRequest', (_message.Message,), {

  'DescriptionInfosEntry' : _reflection.GeneratedProtocolMessageType('DescriptionInfosEntry', (_message.Message,), {
    'DESCRIPTOR' : _EDITPRODUCTREQUEST_DESCRIPTIONINFOSENTRY,
    '__module__' : 'product_pb2'
    # @@protoc_insertion_point(class_scope:goalloff.EditProductRequest.DescriptionInfosEntry)
    })
  ,

  'ProductInfosEntry' : _reflection.GeneratedProtocolMessageType('ProductInfosEntry', (_message.Message,), {
    'DESCRIPTOR' : _EDITPRODUCTREQUEST_PRODUCTINFOSENTRY,
    '__module__' : 'product_pb2'
    # @@protoc_insertion_point(class_scope:goalloff.EditProductRequest.ProductInfosEntry)
    })
  ,
  'DESCRIPTOR' : _EDITPRODUCTREQUEST,
  '__module__' : 'product_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.EditProductRequest)
  })
_sym_db.RegisterMessage(EditProductRequest)
_sym_db.RegisterMessage(EditProductRequest.DescriptionInfosEntry)
_sym_db.RegisterMessage(EditProductRequest.ProductInfosEntry)

EditProductResponse = _reflection.GeneratedProtocolMessageType('EditProductResponse', (_message.Message,), {
  'DESCRIPTOR' : _EDITPRODUCTRESPONSE,
  '__module__' : 'product_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.EditProductResponse)
  })
_sym_db.RegisterMessage(EditProductResponse)

ProductMessage = _reflection.GeneratedProtocolMessageType('ProductMessage', (_message.Message,), {

  'DescriptionInfosEntry' : _reflection.GeneratedProtocolMessageType('DescriptionInfosEntry', (_message.Message,), {
    'DESCRIPTOR' : _PRODUCTMESSAGE_DESCRIPTIONINFOSENTRY,
    '__module__' : 'product_pb2'
    # @@protoc_insertion_point(class_scope:goalloff.ProductMessage.DescriptionInfosEntry)
    })
  ,

  'ProductInfosEntry' : _reflection.GeneratedProtocolMessageType('ProductInfosEntry', (_message.Message,), {
    'DESCRIPTOR' : _PRODUCTMESSAGE_PRODUCTINFOSENTRY,
    '__module__' : 'product_pb2'
    # @@protoc_insertion_point(class_scope:goalloff.ProductMessage.ProductInfosEntry)
    })
  ,
  'DESCRIPTOR' : _PRODUCTMESSAGE,
  '__module__' : 'product_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.ProductMessage)
  })
_sym_db.RegisterMessage(ProductMessage)
_sym_db.RegisterMessage(ProductMessage.DescriptionInfosEntry)
_sym_db.RegisterMessage(ProductMessage.ProductInfosEntry)

ProductInventoryMessage = _reflection.GeneratedProtocolMessageType('ProductInventoryMessage', (_message.Message,), {
  'DESCRIPTOR' : _PRODUCTINVENTORYMESSAGE,
  '__module__' : 'product_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.ProductInventoryMessage)
  })
_sym_db.RegisterMessage(ProductInventoryMessage)

ProductQuery = _reflection.GeneratedProtocolMessageType('ProductQuery', (_message.Message,), {
  'DESCRIPTOR' : _PRODUCTQUERY,
  '__module__' : 'product_pb2'
  # @@protoc_insertion_point(class_scope:goalloff.ProductQuery)
  })
_sym_db.RegisterMessage(ProductQuery)

_PRODUCT = DESCRIPTOR.services_by_name['Product']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z5github.com/lessbutter/alloff-grpc-protos/gen/goalloff'
  _CREATEPRODUCTREQUEST_DESCRIPTIONINFOSENTRY._options = None
  _CREATEPRODUCTREQUEST_DESCRIPTIONINFOSENTRY._serialized_options = b'8\001'
  _CREATEPRODUCTREQUEST_PRODUCTINFOSENTRY._options = None
  _CREATEPRODUCTREQUEST_PRODUCTINFOSENTRY._serialized_options = b'8\001'
  _EDITPRODUCTREQUEST_DESCRIPTIONINFOSENTRY._options = None
  _EDITPRODUCTREQUEST_DESCRIPTIONINFOSENTRY._serialized_options = b'8\001'
  _EDITPRODUCTREQUEST_PRODUCTINFOSENTRY._options = None
  _EDITPRODUCTREQUEST_PRODUCTINFOSENTRY._serialized_options = b'8\001'
  _PRODUCTMESSAGE_DESCRIPTIONINFOSENTRY._options = None
  _PRODUCTMESSAGE_DESCRIPTIONINFOSENTRY._serialized_options = b'8\001'
  _PRODUCTMESSAGE_PRODUCTINFOSENTRY._options = None
  _PRODUCTMESSAGE_PRODUCTINFOSENTRY._serialized_options = b'8\001'
  _PRODUCTTYPE._serialized_start=4294
  _PRODUCTTYPE._serialized_end=4351
  _SORTINGOPTIONS._serialized_start=4354
  _SORTINGOPTIONS._serialized_end=4550
  _GETPRODUCTREQUEST._serialized_start=46
  _GETPRODUCTREQUEST._serialized_end=92
  _GETPRODUCTRESPONSE._serialized_start=94
  _GETPRODUCTRESPONSE._serialized_end=157
  _LISTPRODUCTSREQUEST._serialized_start=160
  _LISTPRODUCTSREQUEST._serialized_end=293
  _LISTPRODUCTSRESPONSE._serialized_start=296
  _LISTPRODUCTSRESPONSE._serialized_end=480
  _CREATEPRODUCTREQUEST._serialized_start=483
  _CREATEPRODUCTREQUEST._serialized_end=1468
  _CREATEPRODUCTREQUEST_DESCRIPTIONINFOSENTRY._serialized_start=1233
  _CREATEPRODUCTREQUEST_DESCRIPTIONINFOSENTRY._serialized_end=1288
  _CREATEPRODUCTREQUEST_PRODUCTINFOSENTRY._serialized_start=1290
  _CREATEPRODUCTREQUEST_PRODUCTINFOSENTRY._serialized_end=1341
  _CREATEPRODUCTRESPONSE._serialized_start=1470
  _CREATEPRODUCTRESPONSE._serialized_end=1536
  _EDITPRODUCTREQUEST._serialized_start=1539
  _EDITPRODUCTREQUEST._serialized_end=2814
  _EDITPRODUCTREQUEST_DESCRIPTIONINFOSENTRY._serialized_start=1233
  _EDITPRODUCTREQUEST_DESCRIPTIONINFOSENTRY._serialized_end=1288
  _EDITPRODUCTREQUEST_PRODUCTINFOSENTRY._serialized_start=1290
  _EDITPRODUCTREQUEST_PRODUCTINFOSENTRY._serialized_end=1341
  _EDITPRODUCTRESPONSE._serialized_start=2816
  _EDITPRODUCTRESPONSE._serialized_end=2880
  _PRODUCTMESSAGE._serialized_start=2883
  _PRODUCTMESSAGE._serialized_end=3888
  _PRODUCTMESSAGE_DESCRIPTIONINFOSENTRY._serialized_start=1233
  _PRODUCTMESSAGE_DESCRIPTIONINFOSENTRY._serialized_end=1288
  _PRODUCTMESSAGE_PRODUCTINFOSENTRY._serialized_start=1290
  _PRODUCTMESSAGE_PRODUCTINFOSENTRY._serialized_end=1341
  _PRODUCTINVENTORYMESSAGE._serialized_start=3890
  _PRODUCTINVENTORYMESSAGE._serialized_end=3998
  _PRODUCTQUERY._serialized_start=4001
  _PRODUCTQUERY._serialized_end=4292
  _PRODUCT._serialized_start=4553
  _PRODUCT._serialized_end=4872
# @@protoc_insertion_point(module_scope)
