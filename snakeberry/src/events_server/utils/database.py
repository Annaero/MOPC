"""
    Various functions that works with DB
"""

import os

from beanie import init_beanie
from motor.motor_asyncio import AsyncIOMotorClient

import models


async def initiate_database():
    """Connect to MongoDB. Init beanie models"""
    client = AsyncIOMotorClient(os.environ["MONGODB_CONNECTION_STRING"])
    await init_beanie(database=client["events"], document_models=models.__all__)
