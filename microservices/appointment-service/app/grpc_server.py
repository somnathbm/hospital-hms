from concurrent import futures
import grpc

import app.gen.appointment_pb2 as appointment_pb2
import app.gen.appointment_pb2_grpc as appointment_pb2_grpc

class AppointmentServiceServicer(appointment_pb2_grpc.AppointmentServiceServicer):
    def __init__(self, core):
        self.core = core

    def CheckAppointment(self, request, context):
        valid, msg = self.core.check_appointment(request.patient_id, request.appointment_id)
        return appointment_pb2.CheckAppointmentResponse(valid=valid, msg=msg)


    def BookAppointment(self, request, context):
        appt_id, msg = self.core.book_appointment(
            request.patient_id, 
            request.doctor_id, 
            request.slot)
        return appointment_pb2.BookAppointmentResponse(appointment_id=appt_id, msg=msg)
    
    
def serve_grpc(core, grpc_port="50052"):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    appointment_pb2_grpc.add_AppointmentServiceServicer_to_server(
        AppointmentServiceServicer(core), server)
    server.add_insecure_port(f"[::]:{grpc_port}")
    server.start()
    print(f"✅ gRPC server listening on {grpc_port}")
    server.wait_for_termination()
