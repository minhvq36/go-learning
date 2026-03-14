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

Expected output:

Hello, Go

---

### Exercise 2 — Sum Two Numbers

Goal:
Create a function that adds two integers.

Concepts learned:

* defining functions
* parameters
* return values

Function signature:

sum(a, b int) int

Expected output:

5

---

### Exercise 3 — Basic Slice

Goal:
Work with slices and append.

Concepts learned:

* slice creation
* append operation
* printing slices

Example slice:

[]int{1,2,3}

Expected output:

[1 2 3 4]

---

### Exercise 4 — Map Basics

Goal:
Use a map to store student scores.

Concepts learned:

* map literal initialization
* inserting values into a map
* iterating or printing map content

Example:

map[string]int

Expected output:

map[Alice:90 Bob:85 Carol:78]

---

### Exercise 5 — Struct & Method

Goal:
Define a struct and attach a method.

Concepts learned:

* struct definition
* method receivers
* basic object-like behavior in Go

Example:

User struct with method:

Greet()

Expected output:

Hello, Alice

---

### Exercise 6 — Pointer

Goal:
Modify struct data using a pointer.

Concepts learned:

* pointer receiver
* struct mutation
* Go method call auto-addressing

Example:

IncreaseAge()

Expected output:

Age before: 20
Age after: 21

---

### Exercise 7 — Basic Error Handling

Goal:
Return an error when dividing by zero.

Concepts learned:

* multiple return values
* error interface
* basic error handling pattern

Function:

func Divide(a, b int) (int, error)

---

### Exercise 8 — Simple Goroutine

Goal:
Run code concurrently using goroutines.

Concepts learned:

* go keyword
* goroutine execution model
* basic synchronization using time.Sleep

Example output:

Hello from main
Hello from goroutine

(order may vary)

---

### Exercise 9 — Channel Basics

Goal:
Send data from a goroutine to main.

Concepts learned:

* channel creation
* send / receive operations
* blocking synchronization
* directional channels

Example output:

Received: 42

---

### Exercise 10 — Basic Context

Goal:
Cancel a goroutine after a timeout.

Concepts learned:

* context.WithTimeout
* cancellation propagation
* select with multiple signals

Example output:

Cancelled

---

## Project Architecture (Core Files and Folders)

```
go.mod
README.md
ai/
	CONTEXT_PROCESS.md
	SKILLS.md
cmd/
	gocore/
		main.go
	goroutine/
		main.go
	hello/
		main.go
internal/
	maps/
		exercise.go
	maths/
		calc.go
	models/
		user.go
	shapes/
		interface.go
	slice/
		exercise.go
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

These concepts form the foundation of most Go backend systems.

---

## Next Learning Direction

After completing the first 10 exercises, the next step is to explore deeper Go runtime and concurrency behavior.

Planned topics:

* goroutine scheduling
* blocking channel behavior
* context propagation across layers
* memory escape analysis

These topics will help understand how Go behaves internally and how to write efficient concurrent systems.

---

## Current Status

Completed: 10 / 10 exercises

Next step:

Move from basic language exercises to deeper Go runtime concepts and concurrency behavior.
