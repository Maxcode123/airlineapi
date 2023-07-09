from typing import Optional
import os

from pydantic import BaseModel

from dotenv import load_dotenv


class Config(BaseModel):
    """
    Configuration object holding project settings.
    """
    api_port: Optional[int] = None
    proxy_port: Optional[int] = None

    def init(self) -> None:
        """
        Execute this function to load configuration settings from environment.
        """
        load_dotenv()
        self.api_port = int(os.getenv("API_PORT"))
        self.proxy_port = int(os.getenv("PROXY_PORT"))

config = Config()
