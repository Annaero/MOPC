# coding: utf-8

from typing import Dict, List  # noqa: F401

from events_server.apis.events_api import EventsAPI
from events_server.apis.default_api_base import BaseDefaultApi

from fastapi import (  # noqa: F401
    APIRouter,
    Body,
    Cookie,
    Depends,
    Form,
    Header,
    Path,
    Query,
    Response,
    Security,
    status,
)

from events_server.models.extra_models import TokenModel  # noqa: F401
from events_server.models.error import Error
from events_server.models.event import Event
from events_server.models.new_event import NewEvent


router = APIRouter()


@router.get(
    "/events",
    responses={
        200: {"model": List[Event], "description": "OK"},
        400: {"model": Error, "description": "Wrong timespan passed"},
        500: {"model": Error, "description": "Unexpected error"},
    },
    tags=["default"],
    response_model_by_alias=True,
)
async def events_get(
    start_date: str = Query(
        None, description="the date FROM which the events will be returned (inclusive)"
    ),
    end_date: str = Query(
        None, description="the date BY which the events will be returned (inclusive)"
    ),
) -> List[Event]:
    """Return events list for given dates span"""
    return BaseDefaultApi.subclasses[0]().events_get(start_date, end_date)


@router.delete(
    "/events/{id}",
    responses={
        204: {"description": "Event deleted"},
        404: {"model": Error, "description": "The specified resource was not found"},
        500: {"model": Error, "description": "Unexpected error"},
    },
    tags=["default"],
    response_model_by_alias=True,
)
async def events_id_delete(
    id: int = Path(description="ID of event to delete"),
) -> None:
    """Deletes a single event based on the ID supplied"""
    return BaseDefaultApi.subclasses[0]().events_id_delete(id)


@router.get(
    "/events/{id}",
    responses={
        200: {"model": Event, "description": "Fetched event"},
        404: {"model": Error, "description": "The specified resource was not found"},
        500: {"model": Error, "description": "Unexpected error"},
    },
    tags=["default"],
    response_model_by_alias=True,
)
async def events_id_get(
    id: int = Path(description="ID of event to fetch"),
) -> Event:
    """Returns an event by its ID"""
    return BaseDefaultApi.subclasses[0]().events_id_get(id)


@router.post(
    "/events",
    responses={
        201: {"model": Event, "description": "Event created"},
        400: {"model": Error, "description": "Incorrect value"},
        409: {"model": Error, "description": "Name already exists"},
        500: {"model": Error, "description": "Unexpected error"},
    },
    tags=["default"],
    response_model_by_alias=True,
)
async def events_post(
    new_event: NewEvent = Body(None, description="Event to add"),
) -> Event:
    """Creates a new event. Duplicate names are not allowed"""
    return BaseDefaultApi.subclasses[0]().events_post(new_event)
