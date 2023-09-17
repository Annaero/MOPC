from __future__ import annotations
from datetime import date, datetime
from enum import Enum

from typing import Any, Optional

from pydantic import (
    Field,
    model_validator,
)

from beanie import Document, Indexed
import pymongo


class EventType(str, Enum):
    ONLINE = "online"
    OFFLINE = "offline"


class MOPCEvent(Document):
    """
    MOPS paltform event

    name: The name of this Event.
    type: The type of this Event. One of "online" | "offline"
    description: The description of this Event.
    start_date: The start_date of this Event.
    end_date: The end_date of this Event.
    """

    name: str = Field(alias="name")
    type: EventType = Field(alias="type")
    description: Optional[str] = Field(alias="description")
    start_date: Indexed(date) = Field(alias="startDate")
    end_date: Indexed(date) = Field(alias="endDate")

    @model_validator(mode="after")
    @classmethod
    def validate_dates(cls, event: MOPCEvent) -> MOPCEvent:
        """Validate that start date not exceed end_date"""
        assert (
            event.start_date <= event.end_date
        ), f"start_date mast be prior or same as end_date, but got {event.start_date} > {event.end_date}"
        return event

    class Settings:
        """Setting"""

        name = "mopc_event"
        indexes = [
            [
                ("start_date", pymongo.DESCENDING),
                ("end_date", pymongo.DESCENDING),
            ],
        ]
        bson_encoders = {
            date: lambda dt: datetime(
                year=dt.year, month=dt.month, day=dt.day, hour=0, minute=0, second=0
            )
        }

    class Config:
        """Config"""

        json_schema_extra = {
            "example": {
                "name": "My crazy event",
                "type": "online",
                "description": "My crazy event",
                "start_date": date(2023, 9, 10),
                "end_date": date(2023, 9, 15),
            }
        }
