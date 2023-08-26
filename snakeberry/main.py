from fastapi import FastAPI
import uvicorn

from dotenv import load_dotenv


# from auth.jwt_bearer import JWTBearer
# from config.config import initiate_database

# from routes.admin import router as AdminRouter

load_dotenv()

from routes.event import router as EventRouter
from utils.database import init_database
from utils.dev_utils import populate_test_events, clear_events


load_dotenv()

app = FastAPI(title="MOPC project events API")

# token_listener = JWTBearer()


@app.on_event("startup")
async def start_database():
    """Starts databases"""
    await init_database()
    await clear_events()
    await populate_test_events()


# app.include_router(AdminRouter, tags=["Administrator"], prefix="/admin")
app.include_router(
    EventRouter,
    tags=["Events"],
    prefix="/events",
    # dependencies=[Depends(token_listener)],
)

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)
