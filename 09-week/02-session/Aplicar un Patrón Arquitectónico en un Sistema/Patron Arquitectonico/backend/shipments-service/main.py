
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from typing import Dict, List
import uuid
from datetime import datetime

app = FastAPI(title="Shipments Service")

shipments: Dict[str, dict] = {}

class ShipmentCreate(BaseModel):
    origin: str
    destination: str
    weight: float
    customer_id: str

class Shipment(BaseModel):
    id: str
    origin: str
    destination: str
    weight: float
    customer_id: str
    status: str
    created_at: datetime

@app.post("/shipments", response_model=Shipment, status_code=201)
def create_shipment(s: ShipmentCreate):
    sid = str(uuid.uuid4())
    shipment = {
        "id": sid,
        "origin": s.origin,
        "destination": s.destination,
        "weight": s.weight,
        "customer_id": s.customer_id,
        "status": "created",
        "created_at": datetime.utcnow().isoformat()
    }
    shipments[sid] = shipment
    return shipment

@app.get("/shipments", response_model=List[Shipment])
def list_shipments():
    return list(shipments.values())

@app.get("/shipments/{shipment_id}", response_model=Shipment)
def get_shipment(shipment_id: str):
    sh = shipments.get(shipment_id)
    if not sh:
        raise HTTPException(status_code=404, detail="Shipment not found")
    return sh
