FROM python:3.10-slim AS build

WORKDIR /app

COPY . .

RUN pip install --no-cache-dir -r requirements.txt

EXPOSE 8000 50052

CMD ["python", "-m", "app.main"]