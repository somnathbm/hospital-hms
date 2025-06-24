import threading
import uvicorn

from app.service import BillingServiceCore
from app.rest import create_rest_api
from app.grpc_server import serve_grpc

def main():
    core = BillingServiceCore()

    # start gRPC in thread
    grpc_thread = threading.Thread(target=serve_grpc, args=(core,))
    grpc_thread.start()

    # start FastAPI
    app = create_rest_api(core)
    uvicorn.run(app, host="0.0.0.0", port=8000)

if __name__ == "__main__":
    main()
