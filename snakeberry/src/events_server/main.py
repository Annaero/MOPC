from fastapi import FastAPI

# from auth.jwt_bearer import JWTBearer
# from config.config import initiate_database

# from routes.admin import router as AdminRouter
from routes.event import router as EventRouter
from utils.database import initiate_database

app = FastAPI(title="MOPC project events API")

# token_listener = JWTBearer()


@app.on_event("startup")
async def start_database():
    """Starts databases"""
    await initiate_database()


# app.include_router(AdminRouter, tags=["Administrator"], prefix="/admin")
app.include_router(
    EventRouter,
    tags=["Events"],
    prefix="/events",
    # dependencies=[Depends(token_listener)],
)
