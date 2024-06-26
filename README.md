# Skeleton Go Web Server

![Build (Go)](<https://github.com/scharissis/go-server-skeleton/workflows/Build%20(Go)/badge.svg?branch=master>)
[![Go Report Card](https://goreportcard.com/badge/github.com/scharissis/go-server-skeleton)](https://goreportcard.com/report/github.com/scharissis/go-server-skeleton)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/scharissis/go-server-skeleton?tab=doc)

A very simple template for a golang web service.

It is opinionated and intended to act as an example of best practices.

## Building & Running

- Build

  `./build.sh` (recommended), OR

  `go build ./...`

- Run

  `go run main/skeleton.go`

## What's here?

1. A web server, which:

   - Exposes one method (/api/answer), which allows GET and POST and responds with JSON

     ```
     GET  /api/answer               // 'Hello!'
     POST /api/answer 'Stefano'     // 'Hello, Stefano!'
     ```

   - Has middleware to limit the type of HTTP request a method will allow
   - Shuts down gracefully

1. An example 3rd party service client library (skeleton/numbers)

1. Tests

### Design choices at a glance

- Server shuts down gracefully
- Routes go in their own file
- API unit tests use **testify**: github.com/stretchr/testify/assert

## Contributing

Contributions are very welcome!

Submit a pull request with:

1. New functionality or edits.
2. Tests covering the above.
3. Reasoning as to why it should be included.
