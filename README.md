# IDATG2005 — Assignment 1

A small Go HTTP service that have a few endpoints which fetch data from external APIs (currency rates and country info).

## How to use

### 1) Start the server
Run the project from the root folder:

```bash
go run .
```
---
## Features

- Root endpoint (`/`) shows a simple HTML message with links to available endpoints.
- Currency exchange endpoint (`/exchange/{code}`) for requesting exchange related data by country code.
- Country information endpoint (`/info/{code}`) for requesting country related data by country code.
- Status endpoint (`/status`) for basic service status/health info.

---

## API Endpoints

### `GET /`
Returns a short HTML page describing available endpoints and links to:
- `/status`
- `/exchange/{code}`
- `/info/{code}`

### `GET /status`
Status endpoint.

### `GET /exchange/{code}`
Exchange endpoint where `{code}` is a country code (example: `no`, `se`, `us`).

Example:
- `/exchange/us`

### `GET /info/{code}`
Information endpoint where `{code}` is a country code (example: `no`, `se`, `us`).

Example:
- `/info/no` 

---

## External APIs Used
- Currency API base:
  - `http://129.241.150.113:9090/currency/`
- Countries API base:
  - `http://129.241.150.113:8080/v3.1/alpha/`
