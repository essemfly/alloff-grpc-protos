# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from . import topbanner_pb2 as topbanner__pb2


class TopBannerStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetTopBanner = channel.unary_unary(
                '/goalloff.TopBanner/GetTopBanner',
                request_serializer=topbanner__pb2.GetTopBannerRequest.SerializeToString,
                response_deserializer=topbanner__pb2.GetTopBannerResponse.FromString,
                )
        self.ListTopBanners = channel.unary_unary(
                '/goalloff.TopBanner/ListTopBanners',
                request_serializer=topbanner__pb2.ListTopBannersRequest.SerializeToString,
                response_deserializer=topbanner__pb2.ListTopBannersResponse.FromString,
                )
        self.EditTopBanner = channel.unary_unary(
                '/goalloff.TopBanner/EditTopBanner',
                request_serializer=topbanner__pb2.EditTopBannerRequest.SerializeToString,
                response_deserializer=topbanner__pb2.EditTopBannerResponse.FromString,
                )
        self.CreateTopBanner = channel.unary_unary(
                '/goalloff.TopBanner/CreateTopBanner',
                request_serializer=topbanner__pb2.CreateTopBannerRequest.SerializeToString,
                response_deserializer=topbanner__pb2.CreateTopBannerResponse.FromString,
                )


class TopBannerServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetTopBanner(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListTopBanners(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def EditTopBanner(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreateTopBanner(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_TopBannerServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetTopBanner': grpc.unary_unary_rpc_method_handler(
                    servicer.GetTopBanner,
                    request_deserializer=topbanner__pb2.GetTopBannerRequest.FromString,
                    response_serializer=topbanner__pb2.GetTopBannerResponse.SerializeToString,
            ),
            'ListTopBanners': grpc.unary_unary_rpc_method_handler(
                    servicer.ListTopBanners,
                    request_deserializer=topbanner__pb2.ListTopBannersRequest.FromString,
                    response_serializer=topbanner__pb2.ListTopBannersResponse.SerializeToString,
            ),
            'EditTopBanner': grpc.unary_unary_rpc_method_handler(
                    servicer.EditTopBanner,
                    request_deserializer=topbanner__pb2.EditTopBannerRequest.FromString,
                    response_serializer=topbanner__pb2.EditTopBannerResponse.SerializeToString,
            ),
            'CreateTopBanner': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateTopBanner,
                    request_deserializer=topbanner__pb2.CreateTopBannerRequest.FromString,
                    response_serializer=topbanner__pb2.CreateTopBannerResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'goalloff.TopBanner', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class TopBanner(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetTopBanner(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/goalloff.TopBanner/GetTopBanner',
            topbanner__pb2.GetTopBannerRequest.SerializeToString,
            topbanner__pb2.GetTopBannerResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListTopBanners(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/goalloff.TopBanner/ListTopBanners',
            topbanner__pb2.ListTopBannersRequest.SerializeToString,
            topbanner__pb2.ListTopBannersResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def EditTopBanner(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/goalloff.TopBanner/EditTopBanner',
            topbanner__pb2.EditTopBannerRequest.SerializeToString,
            topbanner__pb2.EditTopBannerResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CreateTopBanner(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/goalloff.TopBanner/CreateTopBanner',
            topbanner__pb2.CreateTopBannerRequest.SerializeToString,
            topbanner__pb2.CreateTopBannerResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
