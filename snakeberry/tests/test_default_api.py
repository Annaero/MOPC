# coding: utf-8

from fastapi.testclient import TestClient

from events_server.models.error import Error  # noqa: F401
from events_server.models.event import Event  # noqa: F401
from events_server.models.new_event import NewEvent  # noqa: F401


def test_events_get(client: TestClient):
    """Test case for events_get"""
    params = [("start_date", "2013-10-20"), ("end_date", "2013-10-20")]
    headers = {}
    response = client.request(
        "GET",
        "/events",
        headers=headers,
        params=params,
    )

    # uncomment below to assert the status code of the HTTP response
    # assert response.status_code == 200


def test_events_id_delete(client: TestClient):
    """Test case for events_id_delete"""

    headers = {}
    response = client.request(
        "DELETE",
        "/events/{id}".format(id=56),
        headers=headers,
    )

    # uncomment below to assert the status code of the HTTP response
    # assert response.status_code == 200


def test_events_id_get(client: TestClient):
    """Test case for events_id_get"""

    headers = {}
    response = client.request(
        "GET",
        "/events/{id}".format(id=56),
        headers=headers,
    )

    # uncomment below to assert the status code of the HTTP response
    # assert response.status_code == 200


def test_events_post(client: TestClient):
    """Test case for events_post"""
    new_event = {
        "end_date": "2000-01-23",
        "name": "name",
        "description": "description",
        "type": "type",
        "start_date": "2000-01-23",
    }

    headers = {}
    response = client.request(
        "POST",
        "/events",
        headers=headers,
        json=new_event,
    )

    # uncomment below to assert the status code of the HTTP response
    # assert response.status_code == 200
