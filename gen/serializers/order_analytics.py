from django_grpc_framework import proto_serializers
from rest_framework import fields

from ..pyalloff import order_analytics_pb2


class OrderStatFigureSerializer(proto_serializers.ProtoSerializer):
    count = fields.IntegerField()
    amount = fields.IntegerField()
    quantity = fields.IntegerField()
    class Meta:
        proto_class = order_analytics_pb2.OrderStatFigure

class OrderStatSerializer(proto_serializers.ProtoSerializer):
    timestamp = fields.CharField()
    total = OrderStatFigureSerializer()
    paid = OrderStatFigureSerializer()
    canceled = OrderStatFigureSerializer()
    class Meta:
        proto_class = order_analytics_pb2.OrderStat

class OrderStatResponseSerializer(proto_serializers.ProtoSerializer):
    summary = OrderStatSerializer()
    order_stats = OrderStatSerializer(many=True)
    class Meta:
        proto_class = order_analytics_pb2.OrderStatResponse

