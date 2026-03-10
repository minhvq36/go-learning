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

## Next Tasks

The next exercises continue expanding Go fundamentals.

### Exercise 4 — Map Basics

Goal:
Use a map to store student scores.

Requirements:

* map[string]int
* at least two initial students
* add one additional student

Expected output:

map[Alice:90 Bob:85 Carol:78]

---

### Exercise 5 — Struct & Method

Goal:
Define a struct and attach a method.

Requirements:

Struct:

User {
Name string
Age  int
}

Method:

Greet() string

Expected output:

Hello, Alice

---

### Exercise 6 — Pointer

Goal:
Modify struct data using a pointer.

Requirements:

Function:

IncreaseAge(u *User)

Expected output:

Age before: 20
Age after: 21

---

### Exercise 7 — Basic Error Handling

Goal:
Return an error when dividing by zero.

Requirements:

func divide(a, b int) (int, error)

Expected output:

Result: 2
Cannot divide by zero

---

### Exercise 8 — Simple Goroutine

Goal:
Run code concurrently using goroutines.

Requirements:

* use go keyword
* main waits using time.Sleep

Expected output:

Hello from main
Hello from goroutine

(order may vary)

---

### Exercise 9 — Channel Basics

Goal:
Send data from a goroutine to main.

Requirements:

* create channel
* send value from goroutine
* receive in main

Expected output:

Received: 42

---

### Exercise 10 — Basic Context

Goal:
Cancel a goroutine after a timeout.

Requirements:

* context.WithTimeout (1 second)
* goroutine listens for ctx.Done()

Expected output:

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
	hello/
		main.go
internal/
	maps/
		exercise.go
	maths/
		calc.go
	models/
		user.go
	slice/
		exercise.go
```

---

## Notes

* Exercises should remain small and focused.
* Each exercise should be placed in its own folder under cmd/.
* Code clarity is more important than optimization at this stage.

Example structure:

cmd/

hello/
main.go

sum/
main.go

slice/
main.go

map/
main.go

---

## Current Status

Completed: 3 / 10 exercises

Next step:

Start implementing Exercise 4 — Map Basics.
