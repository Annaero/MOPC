# coding: utf-8

from __future__ import annotations
from datetime import date, datetime
from enum import Enum  # noqa: F401

import re  # noqa: F401
from typing import Any, Dict, List, Optional  # noqa: F401

from pydantic import AnyUrl, BaseModel, EmailStr, Field, validator  # noqa: F401

from beanie import Document, Indexed
import pymongo


class EventType(str, Enum):
    ONLINE = "online"
    OFFLINE = "offline"


class MOPCEvent(Document):
    """Event
    # id: The id of this Event.
    name: The name of this Event.
    type: The type of this Event. One of "online" | "offline"
    description: The description of this Event.
    start_date: The start_date of this Event.
    end_date: The end_date of this Event.
    """

    # id: int = Field(alias="id")
    name: str = Field(alias="name")
    type: EventType = Field(alias="type")
    description: Optional[str] = Field(alias="description")
    start_date: Indexed(date) = Field(alias="startDate")
    end_date: Indexed(date) = Field(alias="endDate")

    class Settings:
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
        schema_extra = {
            "example": {
                # "id": 1,
                "name": "My crazy event",
                "type": "online",
                "description": "My crazy event",
                "start_date": date(2023, 9, 10),
                "end_date": date(2023, 9, 15),
            }
        }


class NewMOPCEvent(Document):
    """Event
    id: The id of this Event.
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

    class Settings:
        name = "mopc_event"
        indexes = [
            [
                ("start_date", pymongo.DESCENDING),
                ("end_date", pymongo.DESCENDING),
            ],
        ]

    class Config:
        schema_extra = {
            "example": {
                "name": "My crazy event",
                "type": "online",
                "description": "My crazy event",
                "start_date": date(2023, 9, 10),
                "end_date": date(2023, 9, 15),
            }
        }
