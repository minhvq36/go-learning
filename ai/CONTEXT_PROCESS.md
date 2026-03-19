# CONTEXT_PROCESS.md

## Overview

This document tracks the current learning process focused on the Go programming language. The objective is to gradually build a solid understanding of Go fundamentals and progress toward mastering the core features of Go that are commonly used in real-world backend systems, concurrency models, and scalable services.

The approach follows a progressive exercise-based method. Each task introduces a specific Go concept while keeping the scope small and clear. By completing these exercises sequentially, the learning path builds from basic syntax to more advanced constructs such as concurrency and context management.

The exercises are intentionally minimalistic so that the focus remains on understanding Go's core behavior rather than building large applications prematurely.

---

## Learning Goal

Primary goal:

Become comfortable with Go core concepts including:

* basic syntax
* functions
* slices and maps
* structs and methods
* pointers
* error handling
* goroutines
* channels
* context management

Secondary goal:

Develop intuition for how Go programs are structured in real production codebases, especially regarding:

* concurrency
* safe data flow
* cancellation patterns

Long term direction:

Move toward deeper Go core topics such as:

* goroutine lifecycle
* worker pools
* pipeline patterns
* context propagation
* Go memory model basics

---

## What Has Been Done

Completed exercises:

### Exercise 1 — Hello World

Goal:
Print a simple message using Go.

Concepts learned:

* basic Go project structure
* using the fmt package
* main package and entry point

---

### Exercise 2 — Sum Two Numbers

Goal:
Create a function that adds two integers.

Concepts learned:

* defining functions
* parameters
* return values

---

### Exercise 3 — Basic Slice

Goal:
Work with slices and append.

Concepts learned:

* slice creation
* append operation
* printing slices

---

### Exercise 4 — Map Basics

Goal:
Use a map to store student scores.

Concepts learned:

* map literal initialization
* inserting values into a map
* iterating or printing map content

---

### Exercise 5 — Struct & Method

Goal:
Define a struct and attach a method.

Concepts learned:

* struct definition
* method receivers
* basic object-like behavior in Go

---

### Exercise 6 — Pointer

Goal:
Modify struct data using a pointer.

Concepts learned:

* pointer receiver
* struct mutation
* Go method call auto-addressing

---

### Exercise 7 — Basic Error Handling

Goal:
Return an error when dividing by zero.

Concepts learned:

* multiple return values
* error interface
* basic error handling pattern

---

### Exercise 8 — Simple Goroutine

Goal:
Run code concurrently using goroutines.

Concepts learned:

* go keyword
* goroutine execution model
* basic synchronization using time.Sleep

---

### Exercise 9 — Channel Basics

Goal:
Send data from a goroutine to main.

Concepts learned:

* channel creation
* send / receive operations
* blocking synchronization
* directional channels

---

### Exercise 10 — Basic Context

Goal:
Cancel a goroutine after a timeout.

Concepts learned:

* context.WithTimeout
* cancellation propagation
* select with multiple signals

---

### Exercise 11 — Interface

Concepts learned:

* interface design
* implicit implementation
* polymorphism in Go

---

### Exercise 12 — JSON Encode/Decode

Concepts learned:

* json.Marshal / json.Unmarshal
* struct tags
* API data format handling

---

### Exercise 13 — File I/O

Concepts learned:

* os.ReadFile / os.WriteFile
* []byte <-> string conversion
* basic persistence pattern

---

### Exercise 14 — HTTP Server

Concepts learned:

* net/http
* http.Handler
* basic server setup

---

### Exercise 15 — HTTP JSON API

Concepts learned:

* JSON response handling
* API design basics
* response structuring

---

### Exercise 16 — Simple Router

Concepts learned:

* route handling
* method checking
* request parsing

---

### Exercise 17 — Worker Pool

Concepts learned:

* goroutine pool
* job channel / result channel
* fan-out / fan-in pattern
* sync.WaitGroup coordination

---

### Exercise 18 — Rate Limiter

Concepts learned:

* time.Ticker
* channel-based throttling
* basic rate limiting pattern (local)

---

### Exercise 19 — Context Propagation

Concepts learned:

* passing context through layers
* cancellation handling
* request lifecycle awareness

---

### Exercise 20 — Graceful Shutdown

Concepts learned:

* signal.NotifyContext
* http.Server.Shutdown
* graceful shutdown lifecycle
* context-driven cancellation in handlers

---

## Project Architecture (Core Files and Folders)

```
go.mod
README.md
ai/
	CONTEXT_PROCESS.md
	SKILLS.md
cmd/
	context_demo/
		main.go
	file_io/
		main.go
	gocore/
		main.go
	goroutine/
		main.go
	hello/
		main.go
	http_server/
		main.go
	json_exercise/
		main.go
	limiter_demo/
		main.go
	worker_demo/
		main.go
internal/
	api/
		note_handler.go
		router.go
		routes.go
		util.go
	dto/
			note.go
			response.go
	limiter/
		litmiter.go
	maps/
		exercise.go
	maths/
		calc.go
	models/
		user.go
	repository/
		note_repo.go
	service/
		note_service.go
	shapes/
		interface.go
	slice/
		exercise.go
	storage/
		file.go
	worker/
		pool.go
```

---

## Key Concepts Learned

Through the exercises, the following Go core primitives have been covered:

* Go module structure
* functions and return values
* slices and maps
* structs and methods
* pointer receivers
* error handling patterns
* goroutines
* channels
* context cancellation
* JSON encoding/decoding
* file I/O operations
* HTTP server and handlers
* API response design
* worker pool pattern
* rate limiting (local)
* graceful shutdown

These concepts form the foundation of most Go backend systems.

---

## Next Learning Direction

Move from exercises to real backend systems.

Planned:

* build mini backend projects (Notes API, Worker Queue)
* apply concurrency patterns in real scenarios
* introduce database layer (SQLite)
* integrate Redis for caching / rate limiting
* design handler → service → repository structure

---

## Current Status

Completed: 20 / 20 exercises

Current level:

* solid Go fundamentals
* intermediate concurrency understanding
* ready for backend project implementation

Next step:

Start building mini backend projects (Stage 2)
