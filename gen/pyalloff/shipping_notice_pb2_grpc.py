# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from . import shipping_candidate_pb2 as shipping__candidate__pb2
from . import shipping_notice_pb2 as shipping__notice__pb2


class ShippingNoticeControllerStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.List = channel.unary_unary(
                '/shipping_notice.ShippingNoticeController/List',
                request_serializer=shipping__notice__pb2.ShippingNoticeListRequest.SerializeToString,
                response_deserializer=shipping__notice__pb2.ShippingNoticeListResponse.FromString,
                )
        self.Create = channel.unary_unary(
                '/shipping_notice.ShippingNoticeController/Create',
                request_serializer=shipping__notice__pb2.ShippingNotice.SerializeToString,
                response_deserializer=shipping__notice__pb2.ShippingNotice.FromString,
                )
        self.Retrieve = channel.unary_unary(
                '/shipping_notice.ShippingNoticeController/Retrieve',
                request_serializer=shipping__notice__pb2.ShippingNoticeRetrieveRequest.SerializeToString,
                response_deserializer=shipping__notice__pb2.ShippingNotice.FromString,
                )
        self.Update = channel.unary_unary(
                '/shipping_notice.ShippingNoticeController/Update',
                request_serializer=shipping__notice__pb2.ShippingNotice.SerializeToString,
                response_deserializer=shipping__notice__pb2.ShippingNotice.FromString,
                )
        self.Destroy = channel.unary_unary(
                '/shipping_notice.ShippingNoticeController/Destroy',
                request_serializer=shipping__notice__pb2.ShippingNotice.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.RemoveItem = channel.unary_unary(
                '/shipping_notice.ShippingNoticeController/RemoveItem',
                request_serializer=shipping__notice__pb2.ShippingNoticeRemoveItemRequest.SerializeToString,
                response_deserializer=shipping__notice__pb2.ShippingNotice.FromString,
                )
        self.MoveItem = channel.unary_unary(
                '/shipping_notice.ShippingNoticeController/MoveItem',
                request_serializer=shipping__notice__pb2.ShippingNoticeMoveItemRequest.SerializeToString,
                response_deserializer=shipping__notice__pb2.ShippingNotice.FromString,
                )
        self.LockAndPackage = channel.unary_unary(
                '/shipping_notice.ShippingNoticeController/LockAndPackage',
                request_serializer=shipping__notice__pb2.ShippingNoticeRetrieveRequest.SerializeToString,
                response_deserializer=shipping__notice__pb2.ShippingNotice.FromString,
                )
        self.RecordShippingTemplate = channel.unary_unary(
                '/shipping_notice.ShippingNoticeController/RecordShippingTemplate',
                request_serializer=shipping__notice__pb2.RecordShippingTemplateRequest.SerializeToString,
                response_deserializer=shipping__notice__pb2.ShippingNotice.FromString,
                )
        self.SubmitTrackingExcel = channel.unary_unary(
                '/shipping_notice.ShippingNoticeController/SubmitTrackingExcel',
                request_serializer=shipping__notice__pb2.SubmitTrackingExcelRequest.SerializeToString,
                response_deserializer=shipping__notice__pb2.ShippingNotice.FromString,
                )
        self.GetCandidates = channel.unary_stream(
                '/shipping_notice.ShippingNoticeController/GetCandidates',
                request_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
                response_deserializer=shipping__candidate__pb2.ShippingCandidate.FromString,
                )
        self.SubmitCandidates = channel.unary_unary(
                '/shipping_notice.ShippingNoticeController/SubmitCandidates',
                request_serializer=shipping__notice__pb2.ShippingNoticeCandidateSubmitRequest.SerializeToString,
                response_deserializer=shipping__notice__pb2.ShippingNotice.FromString,
                )


class ShippingNoticeControllerServicer(object):
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

    def RemoveItem(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def MoveItem(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def LockAndPackage(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RecordShippingTemplate(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SubmitTrackingExcel(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetCandidates(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SubmitCandidates(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ShippingNoticeControllerServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'List': grpc.unary_unary_rpc_method_handler(
                    servicer.List,
                    request_deserializer=shipping__notice__pb2.ShippingNoticeListRequest.FromString,
                    response_serializer=shipping__notice__pb2.ShippingNoticeListResponse.SerializeToString,
            ),
            'Create': grpc.unary_unary_rpc_method_handler(
                    servicer.Create,
                    request_deserializer=shipping__notice__pb2.ShippingNotice.FromString,
                    response_serializer=shipping__notice__pb2.ShippingNotice.SerializeToString,
            ),
            'Retrieve': grpc.unary_unary_rpc_method_handler(
                    servicer.Retrieve,
                    request_deserializer=shipping__notice__pb2.ShippingNoticeRetrieveRequest.FromString,
                    response_serializer=shipping__notice__pb2.ShippingNotice.SerializeToString,
            ),
            'Update': grpc.unary_unary_rpc_method_handler(
                    servicer.Update,
                    request_deserializer=shipping__notice__pb2.ShippingNotice.FromString,
                    response_serializer=shipping__notice__pb2.ShippingNotice.SerializeToString,
            ),
            'Destroy': grpc.unary_unary_rpc_method_handler(
                    servicer.Destroy,
                    request_deserializer=shipping__notice__pb2.ShippingNotice.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'RemoveItem': grpc.unary_unary_rpc_method_handler(
                    servicer.RemoveItem,
                    request_deserializer=shipping__notice__pb2.ShippingNoticeRemoveItemRequest.FromString,
                    response_serializer=shipping__notice__pb2.ShippingNotice.SerializeToString,
            ),
            'MoveItem': grpc.unary_unary_rpc_method_handler(
                    servicer.MoveItem,
                    request_deserializer=shipping__notice__pb2.ShippingNoticeMoveItemRequest.FromString,
                    response_serializer=shipping__notice__pb2.ShippingNotice.SerializeToString,
            ),
            'LockAndPackage': grpc.unary_unary_rpc_method_handler(
                    servicer.LockAndPackage,
                    request_deserializer=shipping__notice__pb2.ShippingNoticeRetrieveRequest.FromString,
                    response_serializer=shipping__notice__pb2.ShippingNotice.SerializeToString,
            ),
            'RecordShippingTemplate': grpc.unary_unary_rpc_method_handler(
                    servicer.RecordShippingTemplate,
                    request_deserializer=shipping__notice__pb2.RecordShippingTemplateRequest.FromString,
                    response_serializer=shipping__notice__pb2.ShippingNotice.SerializeToString,
            ),
            'SubmitTrackingExcel': grpc.unary_unary_rpc_method_handler(
                    servicer.SubmitTrackingExcel,
                    request_deserializer=shipping__notice__pb2.SubmitTrackingExcelRequest.FromString,
                    response_serializer=shipping__notice__pb2.ShippingNotice.SerializeToString,
            ),
            'GetCandidates': grpc.unary_stream_rpc_method_handler(
                    servicer.GetCandidates,
                    request_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                    response_serializer=shipping__candidate__pb2.ShippingCandidate.SerializeToString,
            ),
            'SubmitCandidates': grpc.unary_unary_rpc_method_handler(
                    servicer.SubmitCandidates,
                    request_deserializer=shipping__notice__pb2.ShippingNoticeCandidateSubmitRequest.FromString,
                    response_serializer=shipping__notice__pb2.ShippingNotice.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'shipping_notice.ShippingNoticeController', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ShippingNoticeController(object):
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
        return grpc.experimental.unary_unary(request, target, '/shipping_notice.ShippingNoticeController/List',
            shipping__notice__pb2.ShippingNoticeListRequest.SerializeToString,
            shipping__notice__pb2.ShippingNoticeListResponse.FromString,
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
        return grpc.experimental.unary_unary(request, target, '/shipping_notice.ShippingNoticeController/Create',
            shipping__notice__pb2.ShippingNotice.SerializeToString,
            shipping__notice__pb2.ShippingNotice.FromString,
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
        return grpc.experimental.unary_unary(request, target, '/shipping_notice.ShippingNoticeController/Retrieve',
            shipping__notice__pb2.ShippingNoticeRetrieveRequest.SerializeToString,
            shipping__notice__pb2.ShippingNotice.FromString,
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
        return grpc.experimental.unary_unary(request, target, '/shipping_notice.ShippingNoticeController/Update',
            shipping__notice__pb2.ShippingNotice.SerializeToString,
            shipping__notice__pb2.ShippingNotice.FromString,
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
        return grpc.experimental.unary_unary(request, target, '/shipping_notice.ShippingNoticeController/Destroy',
            shipping__notice__pb2.ShippingNotice.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def RemoveItem(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/shipping_notice.ShippingNoticeController/RemoveItem',
            shipping__notice__pb2.ShippingNoticeRemoveItemRequest.SerializeToString,
            shipping__notice__pb2.ShippingNotice.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def MoveItem(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/shipping_notice.ShippingNoticeController/MoveItem',
            shipping__notice__pb2.ShippingNoticeMoveItemRequest.SerializeToString,
            shipping__notice__pb2.ShippingNotice.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def LockAndPackage(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/shipping_notice.ShippingNoticeController/LockAndPackage',
            shipping__notice__pb2.ShippingNoticeRetrieveRequest.SerializeToString,
            shipping__notice__pb2.ShippingNotice.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def RecordShippingTemplate(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/shipping_notice.ShippingNoticeController/RecordShippingTemplate',
            shipping__notice__pb2.RecordShippingTemplateRequest.SerializeToString,
            shipping__notice__pb2.ShippingNotice.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SubmitTrackingExcel(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/shipping_notice.ShippingNoticeController/SubmitTrackingExcel',
            shipping__notice__pb2.SubmitTrackingExcelRequest.SerializeToString,
            shipping__notice__pb2.ShippingNotice.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetCandidates(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/shipping_notice.ShippingNoticeController/GetCandidates',
            google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            shipping__candidate__pb2.ShippingCandidate.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SubmitCandidates(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/shipping_notice.ShippingNoticeController/SubmitCandidates',
            shipping__notice__pb2.ShippingNoticeCandidateSubmitRequest.SerializeToString,
            shipping__notice__pb2.ShippingNotice.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
