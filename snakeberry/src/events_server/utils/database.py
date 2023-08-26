"""
    Various functions that works with DB
"""

import os

from beanie import init_beanie
from motor.motor_asyncio import AsyncIOMotorClient
from motor.core import AgnosticClient

import models


async def init_database() -> AgnosticClient:
    """Connect to MongoDB. Init beanie models"""
    client: AgnosticClient = AsyncIOMotorClient(os.environ["MONGODB_CONNECTION_STRING"])
    await init_beanie(database=client["events"], document_models=models.__all__)
    return client
