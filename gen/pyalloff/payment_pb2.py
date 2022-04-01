# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: payment.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\rpayment.proto\x12\x07payment\x1a\x1bgoogle/protobuf/empty.proto\"\xfb\x02\n\x07Payment\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x0f\n\x07imp_uid\x18\x02 \x01(\t\x12\x16\n\x0epayment_status\x18\x03 \x01(\t\x12\n\n\x02pg\x18\x04 \x01(\t\x12\x12\n\npay_method\x18\x05 \x01(\t\x12\x0c\n\x04name\x18\x06 \x01(\t\x12\x14\n\x0cmerchant_uid\x18\x07 \x01(\t\x12\x0e\n\x06\x61mount\x18\x08 \x01(\x05\x12\x12\n\nbuyer_name\x18\t \x01(\t\x12\x14\n\x0c\x62uyer_mobile\x18\n \x01(\t\x12\x15\n\rbuyer_address\x18\x0b \x01(\t\x12\x17\n\x0f\x62uyer_post_code\x18\x0c \x01(\t\x12\x0f\n\x07\x63ompany\x18\r \x01(\t\x12\x12\n\napp_scheme\x18\x0e \x01(\t\x12\x12\n\ncreated_at\x18\x0f \x01(\t\x12\x12\n\nupdated_at\x18\x10 \x01(\t\x12$\n\x17personal_customs_number\x18\x11 \x01(\tH\x00\x88\x01\x01\x42\x1a\n\x18_personal_customs_number\"\x14\n\x12PaymentListRequest\"$\n\x16PaymentRetrieveRequest\x12\n\n\x02id\x18\x01 \x01(\x03\x32\xa6\x02\n\x11PaymentController\x12\x39\n\x04List\x12\x1b.payment.PaymentListRequest\x1a\x10.payment.Payment\"\x00\x30\x01\x12.\n\x06\x43reate\x12\x10.payment.Payment\x1a\x10.payment.Payment\"\x00\x12?\n\x08Retrieve\x12\x1f.payment.PaymentRetrieveRequest\x1a\x10.payment.Payment\"\x00\x12.\n\x06Update\x12\x10.payment.Payment\x1a\x10.payment.Payment\"\x00\x12\x35\n\x07\x44\x65stroy\x12\x10.payment.Payment\x1a\x16.google.protobuf.Empty\"\x00\x62\x06proto3')



_PAYMENT = DESCRIPTOR.message_types_by_name['Payment']
_PAYMENTLISTREQUEST = DESCRIPTOR.message_types_by_name['PaymentListRequest']
_PAYMENTRETRIEVEREQUEST = DESCRIPTOR.message_types_by_name['PaymentRetrieveRequest']
Payment = _reflection.GeneratedProtocolMessageType('Payment', (_message.Message,), {
  'DESCRIPTOR' : _PAYMENT,
  '__module__' : 'payment_pb2'
  # @@protoc_insertion_point(class_scope:payment.Payment)
  })
_sym_db.RegisterMessage(Payment)

PaymentListRequest = _reflection.GeneratedProtocolMessageType('PaymentListRequest', (_message.Message,), {
  'DESCRIPTOR' : _PAYMENTLISTREQUEST,
  '__module__' : 'payment_pb2'
  # @@protoc_insertion_point(class_scope:payment.PaymentListRequest)
  })
_sym_db.RegisterMessage(PaymentListRequest)

PaymentRetrieveRequest = _reflection.GeneratedProtocolMessageType('PaymentRetrieveRequest', (_message.Message,), {
  'DESCRIPTOR' : _PAYMENTRETRIEVEREQUEST,
  '__module__' : 'payment_pb2'
  # @@protoc_insertion_point(class_scope:payment.PaymentRetrieveRequest)
  })
_sym_db.RegisterMessage(PaymentRetrieveRequest)

_PAYMENTCONTROLLER = DESCRIPTOR.services_by_name['PaymentController']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  _PAYMENT._serialized_start=56
  _PAYMENT._serialized_end=435
  _PAYMENTLISTREQUEST._serialized_start=437
  _PAYMENTLISTREQUEST._serialized_end=457
  _PAYMENTRETRIEVEREQUEST._serialized_start=459
  _PAYMENTRETRIEVEREQUEST._serialized_end=495
  _PAYMENTCONTROLLER._serialized_start=498
  _PAYMENTCONTROLLER._serialized_end=792
# @@protoc_insertion_point(module_scope)
