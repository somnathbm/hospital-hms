from concurrent import futures
import grpc

import app.gen.billing_pb2 as billing_pb2
import app.gen.billing_pb2_grpc as billing_pb2_grpc

class BillingServiceServicer(billing_pb2_grpc.BillingServiceServicer):
    def __init__(self, core):
        self.core = core

    
    def CreateBill(self, request, context):
        billd_id, msg = self.core.create_bill(
            request.patient_id, request.visit_id, request.amount
        )
        return billing_pb2.CreateBillResponse(billd_id=billd_id, msg=msg)
    

def serve_grpc(core, grpc_port="50052"):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    billing_pb2_grpc.add_BillingServiceServicer_to_server(
        BillingServiceServicer(core), server)
    server.add_insecure_port(f"[::]:{grpc_port}")
    server.start()
    print(f"✅ gRPC server listening on {grpc_port}")
    server.wait_for_termination()