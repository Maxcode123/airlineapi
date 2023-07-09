FROM python:3.10-slim-buster

WORKDIR /app

COPY requirements.txt .env ./

RUN pip install -r requirements.txt

COPY ./*.py ./

CMD ["python", "main.py"]