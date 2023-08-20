from datetime import date
from typing import Dict, List, Tuple  # noqa: F401

from events_server.models.event import Event
from events_server.models.new_event import NewEvent
from events_server.apis.default_api_base import BaseDefaultApi


events__ = [
    Event(
        id=1,
        name="My fancy event",
        description="My fancy event description",
        type="online",
        startDate=date(2023, 8, 5),
        endDate=date(2023, 8, 6),
    ),
    Event(
        id=2,
        name="Another event",
        description="Another event description",
        type="online",
        startDate=date(2023, 8, 6),
        endDate=date(2023, 9, 10),
    ),
    Event(
        id=3,
        name="mega cool XXXL so loooong",
        description="",
        type="online",
        startDate=date(2023, 8, 23),
        endDate=date(2023, 9, 1),
    ),
    Event(
        id=4,
        name="small event",
        description="",
        type="online",
        startDate=date(2023, 7, 10),
        endDate=date(2023, 8, 3),
    ),
]


class EventsAPI(BaseDefaultApi):
    """Main realisation of Event API"""

    def events_get(
        self,
        start_date: str,
        end_date: str,
    ) -> List[Event]:
        """Return events list for given dates span"""
        return events__

    def events_id_delete(
        self,
        id: int,
    ) -> None:
        """Deletes a single event based on the ID supplied"""
        return 200

    def events_id_get(
        self,
        id: int,
    ) -> Event:
        """Returns an event by its ID"""
        return events__[0]

    def events_post(
        self,
        new_event: NewEvent,
    ) -> Event:
        """Creates a new event. Duplicate names are not allowed"""
        return 200
