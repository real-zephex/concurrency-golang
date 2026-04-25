# Go API Backend (Mock API)

This is a Mock API for the project **Sequential-vs-Concurrent Data Fetching**. This API performs both the sequential and concurrent data fetching from dummyjson.com and feeds the aggregated results to the frontend.

Simple Go HTTP server that aggregates data from dummyjson.com. It exposes two endpoints that fetch posts, quotes, and todos either sequentially or concurrently, and returns a combined JSON payload with timings.

## Features

- Sequential and concurrent aggregation endpoints
- Basic CORS middleware
- Request logging with duration
- User-Agent filter that blocks curl on the concurrent endpoint

## Endpoints

- `GET /sequential` - Fetch posts, quotes, todos in order
- `GET /concurrent` - Fetch posts, quotes, todos in parallel

## Response Shape

Both endpoints return a JSON object that matches the `types.Combined` struct:

```json
{
  "duration": "123.456ms",
  "posts": {
    "posts": [/* ... */],
    "time": "45.678ms"
  },
  "quotes": {
    "quotes": [/* ... */],
    "time": "23.456ms"
  },
  "todos": {
    "todos": [/* ... */],
    "time": "54.321ms"
  }
}
```

## Project Structure

- `main.go` - HTTP server, routing, CORS, logging
- `scripts/` - Request handlers and helpers
- `routes/` - External API fetchers for posts, quotes, todos
- `types/` - Response structs used by handlers

## Setup

Requires Go 1.20+ (or a recent Go version).

```bash
go run main.go
```

Server starts on `http://localhost:8080`.

## Example Requests

```bash
curl -s http://localhost:8080/sequential
```

Note: `GET /concurrent` blocks curl-based User-Agent requests. Use a browser or set a custom User-Agent header:

```bash
curl -s -H "User-Agent: my-client" http://localhost:8080/concurrent
```

## External Dependencies

No third-party dependencies. The server calls these public APIs:

- `https://dummyjson.com/posts`
- `https://dummyjson.com/quotes`
- `https://dummyjson.com/todos`

## Notes

- The concurrent handler measures total duration across the three upstream requests.
- Each upstream response includes its own timing field in the payload.
