
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from typing import Dict, List
import uuid

app = FastAPI(title="Fleet Service")

drivers: Dict[str, dict] = {}

class DriverCreate(BaseModel):
    name: str
    license_number: str

class Driver(BaseModel):
    id: str
    name: str
    license_number: str
    status: str

@app.post("/drivers", response_model=Driver)
def create_driver(d: DriverCreate):
    did = str(uuid.uuid4())
    driver = {"id": did, "name": d.name, "license_number": d.license_number, "status": "available"}
    drivers[did] = driver
    return driver

@app.get("/drivers", response_model=List[Driver])
def list_drivers():
    return list(drivers.values())

@app.get("/drivers/{driver_id}", response_model=Driver)
def get_driver(driver_id: str):
    drv = drivers.get(driver_id)
    if not drv:
        raise HTTPException(status_code=404, detail="Driver not found")
    return drv
