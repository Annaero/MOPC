"""
    Main router for Events
"""
from datetime import date
from typing import List
from beanie import PydanticObjectId

from fastapi import APIRouter, Body, HTTPException, Path, Query, status
from models.mopc_event import MOPCEvent
from models.http_error import HTTPError
from utils.dev_utils import debug_only

router = APIRouter()


@router.get(
    "/events",
    responses={
        200: {"model": List[MOPCEvent], "description": "OK"},
        400: {"description": "Incorrect timespan passed"},
        500: {"model": HTTPError, "description": "Unexpected error"},
    },
)
async def events_get(
    start_date: str = Query(
        None, description="the date FROM which the events will be returned (inclusive)"
    ),
    end_date: str = Query(
        None, description="the date BY which the events will be returned (inclusive)"
    ),
) -> List[MOPCEvent]:
    """Return events list for given dates span"""

    left_date = date.fromisoformat(start_date)
    right_date = date.fromisoformat(end_date)
    if left_date > right_date:
        raise HTTPException(
            status_code=status.HTTP_400_BAD_REQUEST,
            detail="Incorrect timespan passed",
        )
    events = await MOPCEvent.find(
        MOPCEvent.start_date < right_date, MOPCEvent.end_date > left_date
    ).to_list()
    return events


@router.post(
    "/events",
    responses={
        201: {"model": MOPCEvent, "description": "Event created"},
        400: {"model": HTTPError, "description": "Incorrect value"},
        409: {"model": HTTPError, "description": "Name already exists"},
        500: {"model": HTTPError, "description": "Unexpected error"},
    },
)
async def events_post(
    new_event: MOPCEvent = Body(None, description="Event to add"),
) -> MOPCEvent:
    """Creates a new event. Duplicate names are not allowed"""
    name_exists = await MOPCEvent.find(MOPCEvent.name == new_event.name).count() > 0
    if name_exists:
        raise HTTPException(
            status_code=status.HTTP_409_CONFLICT,
            detail="Event with this name already exists",
        )
    created_event = await new_event.create()
    return created_event


@router.get(
    "/events2",
    tags=["default"],
    response_model_by_alias=True,
)
@debug_only
async def events_get2() -> List[MOPCEvent]:
    """Return events list for given dates span"""
    events = await MOPCEvent.find({}).to_list()
    return events


# @router.delete(
#     "/events/{id}",
#     responses={
#         204: {"description": "Event deleted"},
#         404: {"model": Error, "description": "The specified resource was not found"},
#         500: {"model": Error, "description": "Unexpected error"},
#     },
#     tags=["default"],
#     response_model_by_alias=True,
# )
# async def events_id_delete(
#     id: int = Path(description="ID of event to delete"),
# ) -> None:
#     """Deletes a single event based on the ID supplied"""
#     ...


@router.get(
    "/events/{event_id}",
    responses={
        200: {"model": MOPCEvent, "description": "Fetched event"},
        404: {
            "model": HTTPError,
            "description": "The specified resource was not found",
        },
        500: {"model": HTTPError, "description": "Unexpected error"},
    },
)
async def events_id_get(
    event_id: PydanticObjectId = Path(description="ID of event to fetch"),
) -> MOPCEvent:
    """Returns an event by its ID"""
    event = await MOPCEvent.get(event_id)
    if not event:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND,
            detail="Event with this id does not exists",
        )
    return event
