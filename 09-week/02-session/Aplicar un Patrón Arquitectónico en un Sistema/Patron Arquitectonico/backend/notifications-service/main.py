
from fastapi import FastAPI
from pydantic import BaseModel
from typing import Dict
import time

app = FastAPI(title="Notifications Service")

class Notification(BaseModel):
    to: str
    subject: str
    body: str

sent = []

@app.post("/notify")
def send_notification(n: Notification):
    # Simulate sending (in production integrate with SMTP, SMS gateway, push service)
    record = {"to": n.to, "subject": n.subject, "body": n.body, "sent_at": time.time()}
    sent.append(record)
    return {"message": "Notification queued", "record": record}

@app.get("/sent")
def list_sent():
    return sent
