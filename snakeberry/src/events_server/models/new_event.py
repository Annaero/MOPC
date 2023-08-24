# coding: utf-8

from __future__ import annotations
from datetime import date, datetime  # noqa: F401

import re  # noqa: F401
from typing import Any, Dict, List, Optional
from beanie import Indexed  # noqa: F401

from pydantic import AnyUrl, BaseModel, EmailStr, Field, validator

from .mopc_event import EventType  # noqa: F401


class NewMOPCEvent(BaseModel):
    """New MOPC Event

    name: The name of this NewEvent.
    type: The type of this NewEvent.
    description: The description of this NewEvent.
    start_date: The start_date of this NewEvent.
    end_date: The end_date of this NewEvent.
    """

    id: int = Field(alias="id")
    name: str = Field(alias="name")
    type: EventType = Field(alias="type")
    description: Optional[str] = Field(alias="description")
    start_date: Indexed(date) = Field(alias="startDate")
    end_date: Indexed(date) = Field(alias="endDate")
