
import requests

def test_create_and_get_shipment():
    # Assumes shipments service running on localhost:8002
    base = "http://localhost:8002"
    payload = {"origin":"Bogota","destination":"Medellin","weight":2.5,"customer_id":"cust-1"}
    r = requests.post(base + "/shipments", json=payload)
    assert r.status_code == 201
    data = r.json()
    sid = data["id"]
    r2 = requests.get(base + f"/shipments/{sid}")
    assert r2.status_code == 200
    assert r2.json()["origin"] == "Bogota"
