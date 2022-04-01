# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from . import order_item_status_change_log_pb2 as order__item__status__change__log__pb2


class OrderItemStatusChangeLogControllerStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.List = channel.unary_stream(
                '/orderitemstatuschangelog.OrderItemStatusChangeLogController/List',
                request_serializer=order__item__status__change__log__pb2.OrderItemStatusChangeLogListRequest.SerializeToString,
                response_deserializer=order__item__status__change__log__pb2.OrderItemStatusChangeLog.FromString,
                )
        self.Create = channel.unary_unary(
                '/orderitemstatuschangelog.OrderItemStatusChangeLogController/Create',
                request_serializer=order__item__status__change__log__pb2.OrderItemStatusChangeLog.SerializeToString,
                response_deserializer=order__item__status__change__log__pb2.OrderItemStatusChangeLog.FromString,
                )
        self.Retrieve = channel.unary_unary(
                '/orderitemstatuschangelog.OrderItemStatusChangeLogController/Retrieve',
                request_serializer=order__item__status__change__log__pb2.OrderItemStatusChangeLogRetrieveRequest.SerializeToString,
                response_deserializer=order__item__status__change__log__pb2.OrderItemStatusChangeLog.FromString,
                )
        self.Update = channel.unary_unary(
                '/orderitemstatuschangelog.OrderItemStatusChangeLogController/Update',
                request_serializer=order__item__status__change__log__pb2.OrderItemStatusChangeLog.SerializeToString,
                response_deserializer=order__item__status__change__log__pb2.OrderItemStatusChangeLog.FromString,
                )
        self.Destroy = channel.unary_unary(
                '/orderitemstatuschangelog.OrderItemStatusChangeLogController/Destroy',
                request_serializer=order__item__status__change__log__pb2.OrderItemStatusChangeLog.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )


class OrderItemStatusChangeLogControllerServicer(object):
    """Missing associated documentation comment in .proto file."""

    def List(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Create(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Retrieve(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Update(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Destroy(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_OrderItemStatusChangeLogControllerServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'List': grpc.unary_stream_rpc_method_handler(
                    servicer.List,
                    request_deserializer=order__item__status__change__log__pb2.OrderItemStatusChangeLogListRequest.FromString,
                    response_serializer=order__item__status__change__log__pb2.OrderItemStatusChangeLog.SerializeToString,
            ),
            'Create': grpc.unary_unary_rpc_method_handler(
                    servicer.Create,
                    request_deserializer=order__item__status__change__log__pb2.OrderItemStatusChangeLog.FromString,
                    response_serializer=order__item__status__change__log__pb2.OrderItemStatusChangeLog.SerializeToString,
            ),
            'Retrieve': grpc.unary_unary_rpc_method_handler(
                    servicer.Retrieve,
                    request_deserializer=order__item__status__change__log__pb2.OrderItemStatusChangeLogRetrieveRequest.FromString,
                    response_serializer=order__item__status__change__log__pb2.OrderItemStatusChangeLog.SerializeToString,
            ),
            'Update': grpc.unary_unary_rpc_method_handler(
                    servicer.Update,
                    request_deserializer=order__item__status__change__log__pb2.OrderItemStatusChangeLog.FromString,
                    response_serializer=order__item__status__change__log__pb2.OrderItemStatusChangeLog.SerializeToString,
            ),
            'Destroy': grpc.unary_unary_rpc_method_handler(
                    servicer.Destroy,
                    request_deserializer=order__item__status__change__log__pb2.OrderItemStatusChangeLog.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'orderitemstatuschangelog.OrderItemStatusChangeLogController', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class OrderItemStatusChangeLogController(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def List(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/orderitemstatuschangelog.OrderItemStatusChangeLogController/List',
            order__item__status__change__log__pb2.OrderItemStatusChangeLogListRequest.SerializeToString,
            order__item__status__change__log__pb2.OrderItemStatusChangeLog.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Create(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/orderitemstatuschangelog.OrderItemStatusChangeLogController/Create',
            order__item__status__change__log__pb2.OrderItemStatusChangeLog.SerializeToString,
            order__item__status__change__log__pb2.OrderItemStatusChangeLog.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Retrieve(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/orderitemstatuschangelog.OrderItemStatusChangeLogController/Retrieve',
            order__item__status__change__log__pb2.OrderItemStatusChangeLogRetrieveRequest.SerializeToString,
            order__item__status__change__log__pb2.OrderItemStatusChangeLog.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Update(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/orderitemstatuschangelog.OrderItemStatusChangeLogController/Update',
            order__item__status__change__log__pb2.OrderItemStatusChangeLog.SerializeToString,
            order__item__status__change__log__pb2.OrderItemStatusChangeLog.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Destroy(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/orderitemstatuschangelog.OrderItemStatusChangeLogController/Destroy',
            order__item__status__change__log__pb2.OrderItemStatusChangeLog.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
