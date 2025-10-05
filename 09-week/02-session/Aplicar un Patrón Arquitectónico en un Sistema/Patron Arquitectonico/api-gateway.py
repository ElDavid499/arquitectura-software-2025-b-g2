
from fastapi import FastAPI, Request
import httpx
app = FastAPI(title="API Gateway")

# Simple route proxying to services (hardcoded for demo)
SERVICE_MAP = {
    "auth": "http://auth-service:8000",
    "shipments": "http://shipments-service:8000",
    "tracking": "http://tracking-service:8000",
    "notifications": "http://notifications-service:8000",
    "fleet": "http://fleet-service:8000"
}

@app.get("/health")
def health():
    return {"status": "ok"}

@app.api_route("/proxy/{service}/{path:path}", methods=["GET","POST","PUT","DELETE"])
async def proxy(service: str, path: str, request: Request):
    base = SERVICE_MAP.get(service)
    if not base:
        return {"error": "Unknown service"}
    url = f"{base}/{path}"
    async with httpx.AsyncClient() as client:
        if request.method == "GET":
            r = await client.get(url, params=dict(request.query_params))
            return r.json()
        else:
            body = await request.json()
            r = await client.request(request.method, url, json=body)
            return r.json()
