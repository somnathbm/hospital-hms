import threading
import uvicorn
from prometheus_fastapi_instrumentator import Instrumentator

from app.service import AppointmentServiceCore
from app.rest import create_rest_api
from app.grpc_server import serve_grpc

def main():
    core = AppointmentServiceCore()

    # start gRPC in a thread
    grpc_thread = threading.Thread(target=serve_grpc, args=(core,))
    grpc_thread.start()

    # start REST API with FastAPI
    app = create_rest_api(core)

    # Instrument the app server
    Instrumentator().instrument(app).expose(app)
    
    uvicorn.run(app, host="0.0.0.0", port=8000)


if __name__ == "__main__":
    main()