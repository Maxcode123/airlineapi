
from fastapi import FastAPI, APIRouter

from proxy import get_tickets
from config import config

app = FastAPI()

router = APIRouter()
router.get('/tickets')(get_tickets)

app.include_router(router)


@app.on_event('startup')
def on_startup():
    config.init()
