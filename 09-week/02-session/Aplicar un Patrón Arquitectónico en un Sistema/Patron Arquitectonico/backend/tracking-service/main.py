
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from typing import Dict, List
import uuid
from datetime import datetime

app = FastAPI(title="Tracking Service")

locations: Dict[str, List[dict]] = {}

class LocationUpdate(BaseModel):
    shipment_id: str
    latitude: float
    longitude: float
    timestamp: datetime = None

@app.post("/track")
def update_location(loc: LocationUpdate):
    loc.timestamp = loc.timestamp or datetime.utcnow()
    rec = {"latitude": loc.latitude, "longitude": loc.longitude, "timestamp": loc.timestamp.isoformat()}
    locations.setdefault(loc.shipment_id, []).append(rec)
    return {"message": "Location recorded"}

@app.get("/track/{shipment_id}")
def get_locations(shipment_id: str):
    return locations.get(shipment_id, [])
