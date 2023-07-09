import uvicorn

from app import app
from config import config


if __name__ == '__main__':
    config.init()
    uvicorn.run(app, host='0.0.0.0', port=config.proxy_port)
