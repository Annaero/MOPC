from fastapi.testclient import TestClient
import pytest

from server import app
from models.mopc_event import MOPCEvent


@pytest.fixture
def client():
    with TestClient(app) as client:
        yield client


def test_events_get(client):
    """Test case for events_get"""
    params = [("start_date", "2013-10-20"), ("end_date", "2013-10-20")]
    headers = {}
    response = client.request(
        "GET",
        "/events/events",
        headers=headers,
        params=params,
    )

    assert response.status_code == 200, response.text


def test_events_post(client):
    """Test case for events_post"""
    new_event = {
        "end_date": "2000-02-25",
        "name": "name",
        "description": "description",
        "type": "online",
        "start_date": "2000-01-23",
    }

    headers = {}
    response = client.request(
        "POST",
        "/events/events",
        headers=headers,
        json=new_event,
    )

    assert response.status_code == 200, response.text
