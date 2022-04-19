from django.db import models
from django_grpc_framework import proto_serializers

from ..pyalloff import (inventory_receipt_log_pb2, order_item_action_log_pb2,
                        order_item_alimtalk_log_pb2,
                        order_item_refund_update_log_pb2,
                        order_item_status_change_log_pb2,
                        received_item_generation_log_pb2)


class OrderItemActionType(models.TextChoices):
    STATUS_CHANGE = "STATUS_CHANGE"
    MEMO_ADD = "MEMO_ADD"
    MEMO_DELETE = "MEMO_DELETE"
    PAYMENT_ADJUSTMENT = "PAYMENT_ADJUSTMENT"
    REFUND_UPDATE = "REFUND_UPDATE"
    GENERATED_RECEIVED_ITEM = "GENERATED_RECEIVED_ITEM"
    FORCE_GENERATED_RECEIVED_ITEM = "FORCE_GENERATED_RECEIVED_ITEM"
    RECEIVED_INVENTORY = "RECEIVED_INVENTORY"


class OrderItemAlimtalkType(models.TextChoices):
    DELIVERY_STARTED = "DELIVERY_STARTED"
    CANCEL_FINISHED = "CANCEL_FINISHED"


class OrderItemAlimtalkLogProtoSerializer(proto_serializers.ProtoSerializer):
    class Meta:
        proto_class = order_item_alimtalk_log_pb2.OrderItemAlimtalkLog


class OrderItemRefundUpdateLogProtoSerializer(proto_serializers.ProtoSerializer):
    class Meta:
        proto_class = order_item_refund_update_log_pb2.OrderItemRefundUpdateLog


class OrderItemStatusChangeLogProtoSerializer(proto_serializers.ProtoSerializer):
    class Meta:
        proto_class = order_item_status_change_log_pb2.OrderItemStatusChangeLog


class ReceivedItemGenerationLogProtoSerializer(proto_serializers.ProtoSerializer):
    class Meta:
        proto_class = received_item_generation_log_pb2.ReceivedItemGenerationLog


class InventoryReceiptLogProtoSerializer(proto_serializers.ProtoSerializer):
    class Meta:
        proto_class = inventory_receipt_log_pb2.InventoryReceiptLog


class OrderItemActionLogProtoProtoSerializer(proto_serializers.ProtoSerializer):
    alimtalk = OrderItemAlimtalkLogProtoSerializer(allow_null=True)
    status_change = OrderItemStatusChangeLogProtoSerializer(allow_null=True)
    refund_update = OrderItemRefundUpdateLogProtoSerializer(allow_null=True)
    received_item = ReceivedItemGenerationLogProtoSerializer(allow_null=True)
    inventory = InventoryReceiptLogProtoSerializer(allow_null=True)

    class Meta:
        proto_class = order_item_action_log_pb2.OrderItemActionLog
