# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

import app.gen.appointment_pb2 as appointment__pb2

GRPC_GENERATED_VERSION = '1.73.1'
GRPC_VERSION = grpc.__version__
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    raise RuntimeError(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in appointment_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
    )


class AppointmentServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CheckAppointment = channel.unary_unary(
                '/appointment.AppointmentService/CheckAppointment',
                request_serializer=appointment__pb2.CheckAppointmentRequest.SerializeToString,
                response_deserializer=appointment__pb2.CheckAppointmentResponse.FromString,
                _registered_method=True)
        self.BookAppointment = channel.unary_unary(
                '/appointment.AppointmentService/BookAppointment',
                request_serializer=appointment__pb2.BookAppointmentRequest.SerializeToString,
                response_deserializer=appointment__pb2.BookAppointmentResponse.FromString,
                _registered_method=True)


class AppointmentServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def CheckAppointment(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def BookAppointment(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_AppointmentServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CheckAppointment': grpc.unary_unary_rpc_method_handler(
                    servicer.CheckAppointment,
                    request_deserializer=appointment__pb2.CheckAppointmentRequest.FromString,
                    response_serializer=appointment__pb2.CheckAppointmentResponse.SerializeToString,
            ),
            'BookAppointment': grpc.unary_unary_rpc_method_handler(
                    servicer.BookAppointment,
                    request_deserializer=appointment__pb2.BookAppointmentRequest.FromString,
                    response_serializer=appointment__pb2.BookAppointmentResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'appointment.AppointmentService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('appointment.AppointmentService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class AppointmentService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def CheckAppointment(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/appointment.AppointmentService/CheckAppointment',
            appointment__pb2.CheckAppointmentRequest.SerializeToString,
            appointment__pb2.CheckAppointmentResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def BookAppointment(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/appointment.AppointmentService/BookAppointment',
            appointment__pb2.BookAppointmentRequest.SerializeToString,
            appointment__pb2.BookAppointmentResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
