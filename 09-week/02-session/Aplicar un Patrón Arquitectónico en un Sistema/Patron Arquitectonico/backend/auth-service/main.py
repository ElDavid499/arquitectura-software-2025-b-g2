
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from typing import Dict
import uuid

app = FastAPI(title="Auth Service")

# Simple in-memory store (for demo). Replace with DB in production.
users: Dict[str, dict] = {}

class UserCreate(BaseModel):
    username: str
    password: str
    email: str = None

class TokenResponse(BaseModel):
    access_token: str
    token_type: str = "bearer"

@app.post("/register", status_code=201)
def register(u: UserCreate):
    if u.username in users:
        raise HTTPException(status_code=400, detail="User already exists")
    users[u.username] = {"username": u.username, "password": u.password, "email": u.email}
    return {"message": "User created"}

@app.post("/login", response_model=TokenResponse)
def login(u: UserCreate):
    user = users.get(u.username)
    if not user or user["password"] != u.password:
        raise HTTPException(status_code=401, detail="Invalid credentials")
    token = str(uuid.uuid4())
    return {"access_token": token, "token_type": "bearer"}

@app.get("/users")
def list_users():
    return list(users.values())
