# 🤖 Multi-Agent Copilot

> **Split bills, track spends, and let agents help** — a playful monorepo where a **Go** REST API meets a **Python** brain powered by [LangGraph](https://github.com/langchain-ai/langgraph).

A small monorepo that combines a **Go REST API** for shared expenses and transactions with a **Python agent service** built on LangGraph. The API persists data in PostgreSQL; the agent layer orchestrates natural-language style flows and can call the API for group balances and settlement suggestions.

This repository is suitable as a **starting point** for production systems: clear service boundaries, explicit configuration via environment variables, and layered handlers in the Go service. Hardening for real deployments (authentication, rate limits, observability, CI, and a declared open-source license) is left for you to add as requirements grow.

## 📑 Table of contents

- [🏗️ Architecture](#architecture)
- [🗂️ Repository layout](#repository-layout)
- [✅ Prerequisites](#prerequisites)
- [⚙️ Configuration](#configuration)
- [🚀 Run the app (step by step)](#run-the-app-step-by-step)
- [🌐 REST API (Go)](#rest-api-go)
- [🐍 Agent service (Python)](#agent-service-python)
- [🗄️ Database migrations](#database-migrations)
- [🛠️ Development](#development)
- [🔐 Operations and security](#operations-and-security)
- [🤝 Contributing](#contributing)
- [📜 License](#license)

## 🏗️ Architecture

```mermaid
flowchart LR
  Client[Client / UI] --> Agent[agent-service\nFastAPI + LangGraph]
  Client --> API[api\nGo + Gin]
  Agent -->|HTTP| API
  API --> PG[(PostgreSQL)]
```

- **`apps/api`** 🐹: HTTP API for groups, expenses, per-user balances, simplified settlements, and standalone transactions.
- **`apps/agent-service`** 🐍: FastAPI app with a LangGraph workflow (parse → tool execution → format) and tools that call the Go API using `BACKEND_URL`.

## 🗂️ Repository layout

| Path | Role |
|------|------|
| `apps/api/cmd/server` | 🚪 Application entrypoint |
| `apps/api/routes` | 🛤️ Gin router and HTTP routes |
| `apps/api/handlers` | 📥 Request binding and HTTP responses |
| `apps/api/service` | 🧠 Business logic |
| `apps/api/repository` | 💾 Database access |
| `apps/api/internal/config` | 🔧 Environment-based configuration |
| `apps/api/internal/db` | 🐘 PostgreSQL client and SQL migrations |
| `apps/agent-service/app` | ✨ FastAPI app, LangGraph graph, tools, LLM helpers |

## ✅ Prerequisites

Gather these before you dive in:

- **Go** 1.26+ (see `apps/api/go.mod`) 🐹
- **Python** 3.11+ 🐍
- **PostgreSQL** 14+ recommended 🐘
- An **OpenAI API key** if you use the default `ChatOpenAI` configuration in the agent service 🔑

Dependency management for Python uses **`uv`** (see [Astral uv](https://docs.astral.sh/uv/)) or a standard virtual environment with `pip` and `pyproject.toml`. Pick what makes you happiest. ☀️

## ⚙️ Configuration

### Go API (`apps/api`) 🐹

Loaded from the process environment. You may place an `apps/api/.env` file (see `.gitignore`); if it is missing, only the environment is used.

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | HTTP listen port |
| `DB_HOST` | `localhost` | PostgreSQL host |
| `DB_PORT` | `5432` | PostgreSQL port |
| `DB_USER` | `postgres` | Database user |
| `DB_PASSWORD` | `password` | Database password |
| `DB_NAME` | `upi_app` | Database name |
| `DB_SSLMODE` | `disable` | PostgreSQL `sslmode` |

### Agent service (`apps/agent-service`) 🐍

| Variable | Default | Description |
|----------|---------|-------------|
| `OPENAI_API_KEY` | _(unset)_ | API key for OpenAI-compatible chat models 🔑 |
| `BACKEND_URL` | `http://localhost:8080` | Base URL of the Go API for HTTP tools 🔗 |

## 🚀 Run the app (step by step)

You will use **two terminal windows** 🪟🪟: one for the Go API, one for the Python agent. PostgreSQL must be running before the API starts. All paths below assume your shell’s current directory is the **repository root** (the folder that contains `apps/`).

### Step 1 — 🧰 Install prerequisites on your machine

1. Install **Go** 1.26 or newer ([downloads](https://go.dev/dl/)).
2. Install **Python** 3.11 or newer.
3. Install **PostgreSQL** 14+ and start the server (service name varies by OS; ensure it listens on the host/port you will put in `DB_HOST` / `DB_PORT`, usually `localhost:5432`).
4. Install **`uv`** for Python (optional but recommended): see the [uv install guide](https://docs.astral.sh/uv/getting-started/installation/).

Sanity check — you should see version numbers, not errors:

```bash
go version
python3 --version
psql --version
```

### Step 2 — 🐘 Create the PostgreSQL database

Connect as a superuser (often `postgres`) and create a database. The API defaults to database name **`upi_app`**; use that name unless you plan to override `DB_NAME`.

```bash
psql -U postgres -h localhost -c "CREATE DATABASE upi_app;"
```

If your local user is already a superuser, you can use `psql -c "CREATE DATABASE upi_app;"` instead.

### Step 3 — ✏️ (Optional) Configure the Go API with a `.env` file

From the repository root:

```bash
cd apps/api
```

Create `apps/api/.env` if you want non-default credentials (otherwise the API uses the defaults in the **Go API** table under [Configuration](#configuration)):

```bash
# Example only — adjust to match your PostgreSQL setup
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=upi_app
DB_SSLMODE=disable
```

Return to the repository root when finished:

```bash
cd ../..
```

You can skip this step if the default user `postgres`, password `password`, and database `upi_app` match your local PostgreSQL. Easy mode. 😎

### Step 4 — 📜 Apply database migrations

Still from the **repository root**, run the SQL files **in order** (transactions first, then group expense schema). Adjust the connection URL to match `DB_USER`, `DB_PASSWORD`, `DB_HOST`, `DB_PORT`, `DB_NAME`, and `DB_SSLMODE`.

```bash
psql "postgres://postgres:password@localhost:5432/upi_app?sslmode=disable" \
  -f apps/api/internal/db/migrations/20260415172225_create_transactions.up.sql

psql "postgres://postgres:password@localhost:5432/upi_app?sslmode=disable" \
  -f apps/api/internal/db/migrations/20260417170710_create_group_expense.up.sql
```

You should see no errors from `psql`. If authentication fails, fix the URL or `.env` and rerun this step — you’ve got this. 💪

### Step 5 — 🐹 Start the Go API (terminal 1)

```bash
cd apps/api
go run ./cmd/server
```

Wait until the process logs that the server is starting and Gin is listening on **`PORT`** (default **8080**). Leave this terminal open. First engine: online. ✅

### Step 6 — 🧪 Smoke-test the API (optional)

In a **new** shell from the repository root:

```bash
curl -s -X POST http://localhost:8080/transactions \
  -H "Content-Type: application/json" \
  -d '{"amount": 10.5, "merchant": "Test"}'
```

You should get a `201` response with a success message. List transactions:

```bash
curl -s http://localhost:8080/transactions
```

If you see JSON, the API is awake and waving hello. 👋

### Step 7 — 📦 Install Python dependencies (agent service)

Open **terminal 2** (keep terminal 1 running). From the repository root:

**Option A — using `uv` (recommended) ⚡**

```bash
cd apps/agent-service
uv sync
```

**Option B — using `venv` and `pip` 🐢**

```bash
cd apps/agent-service
python3 -m venv .venv
source .venv/bin/activate   # On Windows CMD: .venv\Scripts\activate.bat
pip install -e .
```

### Step 8 — 🔑 Set environment variables for the agent

The default LLM path uses OpenAI. Export a valid key (do not commit it to git):

```bash
export OPENAI_API_KEY="sk-..."   # use your real key
export BACKEND_URL="http://localhost:8080"
```

If the Go API runs on another host or port, set `BACKEND_URL` to match (no trailing slash).

### Step 9 — 🐍 Start the agent service (terminal 2)

From `apps/agent-service` (same directory as Step 7):

**With `uv`:**

```bash
uv run uvicorn app.main:app --reload --host 0.0.0.0 --port 8000
```

**With an activated `venv`:**

```bash
uvicorn app.main:app --reload --host 0.0.0.0 --port 8000
```

You should see Uvicorn report that it is listening on **port 8000**. Second engine: also online. 🎉

### Step 10 — 💬 Call the chat endpoint

The `/chat` route takes `query` as a **query string** (not JSON body). Example:

```bash
curl -X POST "http://localhost:8000/chat?query=show%20balance"
```

Open **interactive API docs** in a browser: [http://localhost:8000/docs](http://localhost:8000/docs) — click around, it’s allowed. 🖱️✨

> **💡 Heads-up:** The sample LangGraph nodes use a **placeholder group id** and simple keyword routing (`balance` / `settlement` in the query). For meaningful balance or settlement data from PostgreSQL, create a group and expenses via the [REST API](#rest-api-go), then align the graph or tools with a real `group_id` (UUID) as you evolve the project.

### 🎯 Summary

| Step | What you run | Default URL |
|------|----------------|-------------|
| 1–4 | PostgreSQL + migrations | `postgres://…/upi_app` 🐘 |
| 5 | `go run ./cmd/server` in `apps/api` | API: `http://localhost:8080` 🐹 |
| 7–9 | `uvicorn app.main:app` in `apps/agent-service` | Agent: `http://localhost:8000` 🐍 |

## 🌐 REST API (Go)

Base URL: `http://localhost:${PORT}` (default port **8080**).

| Method | Path | Description |
|--------|------|-------------|
| `POST` | `/groups` | Create a group with a creator and member list 👥 |
| `POST` | `/expenses` | Record an expense and how it is split across members 💸 |
| `GET` | `/groups/:group_id/balances` | Per-user balances for the group ⚖️ |
| `GET` | `/groups/:group_id/settlements` | Suggested transfers to settle balances 🤝 |
| `GET` | `/transactions` | List transactions 📋 |
| `POST` | `/transactions` | Create a transaction ➕ |

**Create group** — JSON body:

```json
{
  "name": "Trip",
  "created_by": "user-alice",
  "members": ["user-alice", "user-bob"]
}
```

**Create expense** — JSON body:

```json
{
  "group_id": "<uuid-from-create-group>",
  "paid_by": "user-alice",
  "amount": 120.0,
  "description": "Dinner",
  "splits": [
    { "user_id": "user-alice", "amount": 60.0 },
    { "user_id": "user-bob", "amount": 60.0 }
  ]
}
```

**Create transaction** — JSON body:

```json
{
  "amount": 42.5,
  "merchant": "Coffee Shop"
}
```

The API uses structured logging middleware and a request ID middleware; extend these hooks for production tracing. 📡

## 🐍 Agent service (Python)

- **Stack**: FastAPI, LangGraph, LangChain (see `apps/agent-service/pyproject.toml`) 📚
- **Graph**: `parse` → `execute` → `format` (`app/graph/graph_builder.py`), with tool calls wired to the Go API for balances and settlements 🔁
- **LLM**: Default model is `gpt-4o-mini` via `ChatOpenAI` (`app/services/llm_service.py`) 🧠

The agent graph and tools are intentionally minimal so you can replace parsing with an LLM planner, add authentication, streaming, persistence, and evaluation without restructuring the whole service. Sky’s the limit. ☁️

## 🗄️ Database migrations

SQL files live under `apps/api/internal/db/migrations/`. Apply them in chronological order with `psql` or your migration runner of choice.

Example with `psql` (adjust credentials and database name):

```bash
psql "postgres://postgres:password@localhost:5432/upi_app?sslmode=disable" \
  -f apps/api/internal/db/migrations/20260415172225_create_transactions.up.sql

psql "postgres://postgres:password@localhost:5432/upi_app?sslmode=disable" \
  -f apps/api/internal/db/migrations/20260417170710_create_group_expense.up.sql
```

## 🛠️ Development

**Go API** 🐹

```bash
cd apps/api
go test ./...
go build -o bin/server ./cmd/server
```

**Python** 🐍

```bash
cd apps/agent-service
uv run ruff check .    # if you add Ruff or another linter
uv run pytest        # if you add tests
```

There is no root `Makefile` or Docker Compose file yet; adding them is a natural next step for reproducible local and CI environments. (Pull requests welcome!) 🎁

## 🔐 Operations and security

- **Secrets**: Never commit `.env` files or API keys. Use a secret manager or sealed environment variables in production. 🤫
- **Database**: Use TLS (`DB_SSLMODE=require` or equivalent) when PostgreSQL is not on localhost. 🔒
- **Agent → API**: Today the agent calls the API over plain HTTP; place both services behind your network controls, mutual TLS, or an API gateway as needed. 🛡️
- **Public exposure**: Add authentication, authorization, input limits, and abuse protection before exposing either service to the internet. 🌍
- **Vulnerability reports**: If you fork this for public use, add a `SECURITY.md` with contact details for responsible disclosure. 📬

## 🤝 Contributing

Issues and pull requests are welcome. For changes that touch both services, describe the API contract and any migration steps in the PR. Keep commits focused and match existing naming and layout conventions in each app. Thanks for making this repo better. 🙌

## 📜 License

There is **no `LICENSE` file** in this repository yet. Before distributing or packaging releases, add an explicit license (for example MIT or Apache-2.0) at the repository root and update this section to match. Then celebrate with your favorite beverage. 🥤
