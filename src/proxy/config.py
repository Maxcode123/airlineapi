from typing import Optional
import os

from pydantic import BaseModel


class Config(BaseModel):
    api_port: Optional[int] = None
    proxy_port: Optional[int] = None

    def init(self) -> None:
        self.api_port = int(os.getenv("API_PORT"))
        self.proxy_port = int(os.getenv("PROXY_PORT"))

config = Config()
