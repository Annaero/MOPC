from datetime import date, datetime
import logging
import os
import random
from fastapi import HTTPException, status

from motor.motor_asyncio import AsyncIOMotorClient

from models.mopc_event import MOPCEvent, EventType

logger = logging.getLogger(__name__)


def debug_only(rout):
    """Allows rout only for debug env"""
    if os.environ["MOPC_ENV"] == "debug":
        return rout

    async def __not_available_in_none_debug__(*args, **kwargs):
        raise HTTPException(
            status_code=status.HTTP_405_METHOD_NOT_ALLOWED,
            detail="Don't call me, don't come by my house",
        )

    return __not_available_in_none_debug__


async def populate_test_events(n_events: int = 3):
    """Saves some tests events for testing purposes"""

    events = []
    for i in range(n_events):
        current_month = datetime.now().month
        current_year = datetime.now().year

        month1 = random.randint(current_month - 1, current_month + 2)
        month2 = random.randint(month1, month1 + 2)

        day1 = random.randint(1, 28)
        day2 = random.randint(1, 28)
        if month1 == month2:
            while day2 < day1:
                day2 = random.randint(1, 28)

        events.append(
            MOPCEvent(
                name=f"Event {i}",
                type=EventType.ONLINE,
                description="Some new event",
                startDate=date(current_year, month1, day1),
                endDate=date(current_year, month2, day2),
            )
        )

    logger.info(f"Populate db with events: {events}")
    await MOPCEvent.insert_many(events)


async def clear_events():
    """Clears events collection"""
    await MOPCEvent.find_all().delete()
