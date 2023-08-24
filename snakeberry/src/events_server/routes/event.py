"""
    Main router for Events
"""
from datetime import date
from typing import List

from fastapi import APIRouter, HTTPException, Query, status
from models.error import Error
from models.mopc_event import MOPCEvent

router = APIRouter()


@router.get(
    "/events",
    responses={
        200: {"model": List[MOPCEvent], "description": "OK"},
        400: {"description": "Incorrect timespan passed"},
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
) -> List[MOPCEvent]:
    """Return events list for given dates span"""

    left_date = date.fromisoformat(start_date)
    right_date = date.fromisoformat(end_date)
    if left_date >= right_date:
        raise HTTPException(
            status_code=status.HTTP_400_BAD_REQUEST,
            detail="Incorrect timespan passed",
        )
    events = await MOPCEvent.find(
        MOPCEvent.start_date < right_date, MOPCEvent.end_date > left_date
    ).to_list()
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


# @router.get(
#     "/events/{id}",
#     responses={
#         200: {"model": Event, "description": "Fetched event"},
#         404: {"model": Error, "description": "The specified resource was not found"},
#         500: {"model": Error, "description": "Unexpected error"},
#     },
#     tags=["default"],
#     response_model_by_alias=True,
# )
# async def events_id_get(
#     id: int = Path(description="ID of event to fetch"),
# ) -> Event:
#     """Returns an event by its ID"""
#     ...


# @router.post(
#     "/events",
#     responses={
#         201: {"model": Event, "description": "Event created"},
#         400: {"model": Error, "description": "Incorrect value"},
#         409: {"model": Error, "description": "Name already exists"},
#         500: {"model": Error, "description": "Unexpected error"},
#     },
#     tags=["default"],
#     response_model_by_alias=True,
# )
# async def events_post(
#     new_event: NewEvent = Body(None, description="Event to add"),
# ) -> Event:
#     """Creates a new event. Duplicate names are not allowed"""
#     ...
