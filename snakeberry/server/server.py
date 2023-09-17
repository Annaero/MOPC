from fastapi import FastAPI

from routes.events import router as EventRouter
from utils.database import init_database
from utils.dev_utils import populate_test_events, clear_events

app = FastAPI(title="MOPC project events API")


@app.on_event("startup")
async def start_database():
    """Starts databases"""
    await init_database()
    await clear_events()
    await populate_test_events()


app.include_router(
    EventRouter,
    tags=["Events"],
    prefix="/events",
)
