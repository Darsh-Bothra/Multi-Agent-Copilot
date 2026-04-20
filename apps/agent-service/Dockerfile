FROM python:3.11-slim AS base
ENV PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1
WORKDIR /app
RUN pip install --no-cache-dir uv
COPY apps/agent-service/pyproject.toml apps/agent-service/uv.lock ./
RUN uv sync --frozen --no-dev --no-install-project

FROM base AS dev
# Copy source after dependency install to maximize Docker cache hits.
COPY apps/agent-service/ ./
EXPOSE 8000
CMD ["uv", "run", "uvicorn", "app.main:app", "--reload", "--host", "0.0.0.0", "--port", "8000"]

FROM base AS prod
COPY apps/agent-service/ ./
EXPOSE 8000
CMD ["uv", "run", "uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8000"]
