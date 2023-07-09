from datetime import date, datetime
from typing import Any
import json

import requests
from pydantic import BaseModel

from config import config


class Ticket(BaseModel):
    total_amount: float
    total_currency: str
    departure: datetime
    arrival: datetime
    airline: str


class TicketServiceError(Exception):
    pass


def get_tickets(origin: str, destination: str, date: date) -> list[Ticket]:
    res = requests.get(
        url='http://localhost:'  + str(config.api_port) + '/tickets',
        data=json.dumps({'origin': origin, 'destination': destination, 'date': str(date)}),
        headers={'Content-Type': 'application/json'}
        )
    if not res.ok:
        raise TicketServiceError(
            f'Ticket service response {res.status_code}: {res.text}')
    return _parse_response(res)


def _parse_response(response: dict[str, Any]) -> list[Ticket]:
    return [Ticket(**t) for t in response.json()['data']]
