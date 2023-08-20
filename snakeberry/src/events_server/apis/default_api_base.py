# coding: utf-8

from typing import ClassVar, Dict, List, Tuple  # noqa: F401

from events_server.models.error import Error
from events_server.models.event import Event
from events_server.models.new_event import NewEvent


class BaseDefaultApi:
    subclasses: ClassVar[Tuple] = ()

    def __init_subclass__(cls, **kwargs):
        super().__init_subclass__(**kwargs)
        BaseDefaultApi.subclasses = BaseDefaultApi.subclasses + (cls,)
        print(BaseDefaultApi.subclasses)

    def events_get(
        self,
        start_date: str,
        end_date: str,
    ) -> List[Event]:
        """Return events list for given dates span"""
        ...

    def events_id_delete(
        self,
        id: int,
    ) -> None:
        """Deletes a single event based on the ID supplied"""
        ...

    def events_id_get(
        self,
        id: int,
    ) -> Event:
        """Returns an event by its ID"""
        ...

    def events_post(
        self,
        new_event: NewEvent,
    ) -> Event:
        """Creates a new event. Duplicate names are not allowed"""
        ...
