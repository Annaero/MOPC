from pydantic import BaseModel


class HTTPError(BaseModel):
    """Model for http error"""

    detail: str

    class Config:  # pylint: disable=missing-class-docstring
        json_schema_extra = {
            "example": {"detail": "HTTPException raised."},
        }
