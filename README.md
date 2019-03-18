# Zipkin instrumentation for sqlx

[![Build Status](https://travis-ci.org/jcchavezs/zipkin-instrumentation-sqlx.svg?branch=master)](https://travis-ci.org/jcchavezs/zipkin-instrumentation-sqlx)

**Deprecated:** use https://github.com/jcchavezs/zipkin-instrumentation-sql instead.

This package implements the interfaces from sqlx adding zipkin instrumentation

## Install

```bash
go get github.com/jmoiron/sqlx
```

## Setup

```go
tracer, _ := zipkin.NewTracer(...)
...
tracedDB := NewDb(db, tracer)
```

## Usage

This library does not add any functionality on top of sqlx. For more information, 
about `sqlx.DB` usage check https://github.com/jmoiron/sqlx 
