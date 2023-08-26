from datetime import date

from motor.motor_asyncio import AsyncIOMotorClient

from models.mopc_event import MOPCEvent, EventType


async def populate_test_events():
    """Saves some tests events for testing purposes"""
    await MOPCEvent(
        name="Event_1",
        type=EventType.ONLINE,
        description="Some new event",
        startDate=date(2023, 8, 28),
        endDate=date(2023, 9, 10),
    ).create()


async def clear_events():
    """Clears events collection"""
    await MOPCEvent.find_all().delete()
    # async for e in events:
    #     e.delete()
