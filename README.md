# Go Mastery: Senior Engineering & Architecture

> **Proof of Work**: Verification of Backend Engineering Depth & Production Readiness.

## 📌 Executive Summary
This repository is a **senior-level knowledge base** designed for Staff+ engineers and technical interviewers. It bypasses beginner tutorials to focus strictly on **correctness, low-level mechanics, specific concurrency patterns, and production reliability** in Go (Golang).

**Target Audience**: Senior Backend Engineers, Systems Architects, and Technical Leads.
**Focus**: Memory models, GC behavior, scheduler semantics, and distributed system patterns.
**Goal**: To demonstrate not just *how* to write Go, but *why* robust Go systems behave the way they do.

---

## 🧭 How to Use This Repository
Each concept is isolated in its own folder (e.g., [`20_goroutines/`](./20_goroutines)) containing:
*   `README.md`: Theory & production mental models.
*   `example.go`: Runnable, correct code.
*   `pitfalls.go`: Common bugs to avoid.
*   `interview.md`: Deep-dive questions.

### 🏃 Running Examples
```bash
cd 20_goroutines
go run example.go
go run pitfalls.go
```

---

## 📚 Complete Concept Index

| **1. Fundamentals** | **2. Control Flow** | **3. Data Structures** |
| :--- | :--- | :--- |
| [1_hello_world](./1_hello_world) | [5_for](./5_for) | [9_slices](./9_slices) |
| [2_simple_values](./2_simple_values) | [6_if_else](./6_if_else) | [10_maps](./10_maps) |
| [3_variables](./3_variables) | [7_switch](./7_switch) | [15_pointers](./15_pointers) |
| [4_constants](./4_constants) | [11_range](./11_range) | [16_structs](./16_structs) |
| [8_arrays](./8_arrays) | | |
| [18_enums](./18_enums) | | |

| **4. Abstractions** | **5. Concurrency** | **6. System & Web** |
| :--- | :--- | :--- |
| [12_functions](./12_functions) | [20_goroutines](./20_goroutines) | [23_files](./23_files) |
| [13_variadic_functions](./13_variadic_functions) | [21_channels](./21_channels) | [24_packages](./24_packages) |
| [14_closures](./14_closures) | [22_mutex](./22_mutex) | [27_testing](./27_testing) |
| [17_interfaces](./17_interfaces) | [25_context](./25_context) | [28_http_backend](./28_http_backend) |
| [19_generics](./19_generics) | | |
| [26_errors](./26_errors) | | |

---

## 2️⃣ Core Go Concepts

### A. Language Fundamentals (Senior Perspective)

#### 🔹 Value vs. Reference Semantics
*   **What**: Go passes everything by value. "Reference types" (slices, maps, channels) are just structs with internal pointers passed by value.
*   **Why**: Critical for performance/memory optimization. Misunderstanding leads to unintended data mutations or unnecessary heap allocations.
*   **Trap**: "Maps are passed by reference." -> **Correction**: Maps are passed as a pointer to a `hmap` struct. The pointer is copied.
*   **Senior View**: Visualize the backing array or struct header. Know when a copy stays on the stack vs. escapes to heap.
*   **Senior Code**:
    ```go
    type BigStruct struct { Data [1024]int }
    // Bad: Copies entire array (stack or heap alloc cost)
    func process(b BigStruct) { ... }
    // Good: Copies only the pointer (8 bytes)
    func process(b *BigStruct) { ... }
    ```

#### 🔹 Zero Values & Hidden Bugs
*   **What**: Variables declared without initialization get a "zero value" (`0`, `""`, `nil`).
*   **Why**: Go's "make the zero value useful" philosophy (e.g., `sync.Mutex`, `bytes.Buffer`) simplifies code but `nil` pointers/interfaces cause panics.
*   **Trap**: `var mu sync.Mutex` is ready to use. `var s []int` is usable (nil slice). `var m map[string]int` is **read-only** (panics on write).
*   **Senior View**: Always check if a struct needs a constructor or if the zero value is safe. A nil `slice` behaves like an empty slice, but a nil `map` is a timebomb.

#### 🔹 Struct Embedding != Inheritance
*   **What**: Syntactic sugar for automatic field/method delegation. *Not* polymorphism.
*   **Why**: Composition over inheritance. Prevents fragile base class problems.
*   **Trap**: Expecting the embedded type's methods to handle the outer type's state (dynamic dispatch doesn't exist like in Java/C++).
*   **Senior View**: The embedded type knows *nothing* about the embedding struct.
*   **Code**:
    ```go
    type Base struct{}
    func (b Base) Hello() { fmt.Println("Base") }
    func (b Base) CallHello() { b.Hello() } // Calls Base.Hello, not Child.Hello

    type Child struct { Base }
    func (c Child) Hello() { fmt.Println("Child") }
    
    // c.CallHello() prints "Base", because Base.CallHello receives Base!
    ```

#### 🔹 Interfaces: Implicit & Nil Pitfalls
*   **What**: Defined by method sets. Implicit satisfaction.
*   **Why**: Decoupling. "Accept interfaces, return structs."
*   **Trap**: A non-nil interface containing a nil concrete value is **not nil**.
*   **Senior View**: An interface is a tuple `(type, value)`. It is nil only if BOTH are nil.
*   **Code**:
    ```go
    func returnsError() error {
        var err *MyCustomError = nil // pointer is nil
        return err // Interface {Type: *MyCustomError, Value: nil} != nil
    }
    // Result: The caller thinks an error occurred!
    // Fix: explicitly return nil
    ```

#### 🔹 Method Sets (Pointer vs. Value Receivers)
*   **What**: Rules governing which methods attach to a type.
*   **Why**: affect interface satisfaction.
*   **Trap**: A value can call a pointer method (via auto-referencing), but a value *stored in an interface* cannot call a pointer method (needs addressability).
*   **Senior View**: Use pointer receivers if you mutate state OR if the struct is large. Consistency is key.

#### 🔹 Error Handling Philosophy
*   **What**: Errors are values, not control flow exceptions.
*   **Why**: Forces explicit handling. No "GOTO catch" magic.
*   **Trap**: Ignoring errors. Using `panic` for normal logic.
*   **Senior View**: Wrap errors for context (`fmt.Errorf("doing X: %w", err)`). Use `errors.Is` / `errors.As` for sentinel unwrapping.

### B. Concurrency & Parallelism (The "Killer" Section)

#### 🔹 Goroutines: Lifecycle & Scheduler
*   **What**: User-space threads managed by the Go runtime (not OS threads). Start with ~2KB stack.
*   **Why**: Cheap to spawn thousands (vs ~1MB for OS threads).
*   **Trap**: "Goroutines are parallel." -> **Correction**: They are *concurrent*. Parallelism depends on `GOMAXPROCS`.
*   **Senior View**: Understand the **M:P:G model**.
    *   **M** (Machine): OS Thread.
    *   **P** (Processor): Context for execution (limit = GOMAXPROCS).
    *   **G** (Goroutine): Indestructible until function returns.
    *   **Trap**: A blocking syscall (not network IO) can consume an M, causing runtime to spawn new Ms.
    *   **Leak**: A goroutine blocked on a nil channel or waiting for a lock that never opens *never dies* (until program exit).

#### 🔹 Channels: Unbuffered vs. Buffered
*   **What**: Typed conduits for synchronization and data transfer.
*   **Trap**: Treating buffered channels like a queue system.
*   **Crucial Mechanic**:
    *   **Unbuffered**: Synchronization point. Sender blocks until Receiver behaves. "I need you to take this *now*."
    *   **Buffered**: Decoupling. "I'm leaving this here for you."
    *   **Senior View**: **Don't use buffered channels** unless you have a specific performance reason or need to break a deadlock cycle. Unbuffered = stronger guarantees.

#### 🔹 Select: Fairness & Randomness
*   **What**: Control structure for waiting on multiple channel operations.
*   **Trap**: Expecting deterministic order. `select` chooses a ready case *at random* (pseudo-random) to prevent starvation.
*   **Senior View**: Use `default` carefully—it turns a blocking operation into a non-blocking spin-loop if inside a `for`.

#### 🔹 Sync Primitives (Mutex, RWMutex, Once)
*   **What**: Low-level lock-based synchronization.
*   **Why**: Channels are for passing ownership; Mutexes are for checking/caching state.
*   **Common Fatal Bug**: **Copying a Mutex**.
    *   `func check(mu sync.Mutex) { ... }` -> This copies the lock state! The function gets a *new* lock, rendering it useless.
    *   **Fix**: Always pass pointers: `func check(mu *sync.Mutex)`.
*   **RWMutex**: Use only when *many* readers and *few* writers. If writers are frequent, RWMutex performs worse than Mutex due to writer starvation and overhead.

#### 🔹 Context: Propagation & Cancellation
*   **What**: Standards for deadlines, cancellation signals, and request-scoped values.
*   **Why**: Preventing "goroutine leaks" in abandoned requests.
*   **Rule**: context.Background() is the root.
*   **Trap**: Storing Context in a struct. **Don't do it** (few exceptions). Pass it explicitly as the first argument (`ctx context.Context`).
*   **Senior View**: Use `context.WithoutCancel` (Go 1.21+) if you need to detach a cleanup task from a cancelled parent context.

#### 🔹 Fan-Out / Fan-In & Worker Pools
*   **Prop Pattern**: Constrain concurrency. Don't just `go func()` in a loop of 100k items.
*   **Senior Code (Worker Pool)**:
    ```go
    func worker(id int, jobs <-chan int, results chan<- int) {
        for j := range jobs {
            results <- process(j)
        }
    }
    // Main
    for w := 1; w <= 3; w++ { go worker(w, jobs, results) }
    ```

### C. Memory & Performance

#### 🔹 Stack vs. Heap (Escape Analysis)
*   **What**: Determining where "variables" live.
*   **Mechanic**:
    *   **Stack**: Cheap, self-cleaning (popping stack frame).
    *   **Heap**: Expensive allocation, requires GC.
*   **Rule**: If the compiler cannot prove a variable is unused after the function returns, it **escapes to the heap**.
*   **Common Causes**:
    *   Returning a pointer to a local variable.
    *   Storing a value in an interface (often causes escape).
    *   Passing values to `fmt.Println` (uses `interface{}`).
*   **Command**: `go build -gcflags="-m"` to see escape analysis decisions.

#### 🔹 Garbage Collector (GC) Reality
*   **What**: Concurrent, Tri-color Mark-and-Sweep.
*   **Goal**: Low latency (pauses < 1ms usually), not max throughput.
*   **Senior View**:
    *   **GOGC**: Knob for aggressiveness (default 100 = triggers GC when heap doubles).
    *   **Ballast**: (Legacy trick) Allocating a large byte slice to reduce GC frequency.
    *   **Write Barrier**: The overhead incurred during concurrent marking to maintain consistency.

#### 🔹 Slices, Arrays, and Maps (Internals)
*   **Slice**: A struct `Header { Data *Type, Len, Cap }`.
    *   **Trap**: `append` works on the copy of the header. If the backing array needs to grow, the *new* pointer is not seen by the caller unless you return the slice.
    *   **Memory Leak**: Slicing a small chunk of a huge array (`tiny := huge[0:1]`) keeps the *entire* huge array in memory.
    *   **Fix**: `copy()` the data to a new slice.
*   **Map**: Hash map with buckets.
    *   **Trap**: Maps *never shrink* (memory wise) even if you delete keys. They only grow or empty (if you rebuild).

#### 🔹 Profiling & Benchmarking (pprof)
*   **What**: The tool for finding bottlenecks.
*   **Commands**:
    *   `go test -bench=. -benchmem` (shows allocs/op).
    *   `go tool pprof -http=:8080 cpu.prof`.
*   **Senior View**: Don't guess. Use pprof. If `runtime.mallocgc` is high, you have an allocation problem. If `syscall.Read` is high, you have I/O latency.

### D. Packages & Build System

#### 🔹 go mod Internals
*   **What**: Minimal Version Selection (MVS).
*   **Trap**: Thinking Go resolves to the "newest" version.
*   **Reality**: Go picks the *oldest* version that satisfies all requirements (to ensure stability).
*   **go.sum**: Not a lockfile. It's a checksum DB to verify integrity.

#### 🔹 internal Packages
*   **What**: Directories named `internal/`.
*   **Why**: Enforced privacy. Only the parent directory (and siblings) can import it.
*   **Senior View**: Use this aggressively. If a package isn't meant for public consumption, hiding it prevents API leakage and maintenance burden.

#### 🔹 Build Tags
*   **What**: Conditional compilation. `//go:build linux || darwin`.
*   **Why**: Platform-specific logic without runtime checks.
*   **Senior Pattern**: Separate files `storage_posix.go`, `storage_windows.go` rather than `if runtime.GOOS == "windows"`.

### E. HTTP, APIs & Backend Engineering

#### 🔹 net/http Internals
*   **What**: Standard library for building HTTP servers/clients.
*   **Model**: Spawns a **new goroutine** for every incoming request.
*   **Implication**: Handlers must be thread-safe.

#### 🔹 The "No Timeouts" Killer
*   **Trap**: `http.Get(url)` or `&http.Client{}` has **NO timeout**. It will hang forever if the server server stalls.
*   **Senior View**: ALWAYS use explicit timeouts.
    ```go
    // Bad
    c := &http.Client{}
    // Good
    c := &http.Client{Timeout: 10 * time.Second}
    ```
*   **Server Side**: Set `ReadTimeout` and `WriteTimeout` on your `http.Server`. Default is 0 (infinite), which exposes you to Slowloris attacks.

#### 🔹 Context in HTTP
*   **What**: `req.Context()`.
*   **Rule**: Propagate this context to databases and outbound calls. If the client disconnects, the context cancels, and you should abort work to save resources.

#### 🔹 Graceful Shutdown
*   **What**: Stopping the server without killing active requests.
*   **Pattern**: Listen for `SIGINT`/`SIGTERM`. Call `server.Shutdown(ctx)` (waits for handlers to finish).

### F. Testing & Reliability

#### 🔹 Table-Driven Tests
*   **What**: The "Go way" to test.
*   **Why**: Separates test data from test logic. Easy to add cases.
*   **Code**:
    ```go
    tests := []struct{ name, input, want string }{...}
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) { ... })
    }
    ```

#### 🔹 The Race Detector
*   **What**: `go test -race`.
*   **Why**: **Data races are undefined behavior**. A program with a race is theoretically broken even if it "works".
*   **Rule**: CI must run with `-race`.

#### 🔹 Mocks vs. Fakes
*   **Senior View**: Prefer **Fakes** (in-memory implementations) over mocking frameworks (gomock, mockery) when possible.
*   **Why**: Mocks couple tests to implementation details ("Assert method X called 3 times"). Fakes test behavior ("Assert state is Y").

#### 🔹 Parallel Testing
*   **Trap**: `t.Parallel()` requires care with closure variables (prior to Go 1.22 loop fix) and shared state.
*   **Senior View**: Subtests sharing a database must run serially or use isolated transactions.

## 3️⃣ Go Interview Traps & Killer Questions

### 1. The "Nil Interface" Trap
**Q**: What does this print?
```go
func main() {
    var err *os.PathError = nil
    var i interface{} = err
    fmt.Println(i == nil)
}
```
*   **Tricky because**: `err` is typed nil.
*   **Answer**: `false`.
*   **Why**: Interface `i` contains `(type=*os.PathError, value=nil)`. An interface is only nil if both type and value are nil.

### 2. The "Defer Evaluation" Trap
**Q**: What is the output?
```go
func main() {
    i := 0
    defer fmt.Println(i)
    i++
    return
}
```
*   **Tricky because**: Users confuse execution time (end of function) with evaluation time.
*   **Answer**: `0`.
*   **Why**: Function arguments to `defer` are evaluated **immediately** when the `defer` statement is reached.

### 3. The "Slice Append" Side Effect
**Q**: If I pass a slice to a function and append to it, does the caller see the new elements?
*   **Answer**: No (usually).
*   **Why**: `append` returns a *new* slice header (new length). The caller still holds the *old* slice header with the old length. The backing array might be modified, but the caller won't "see" the new elements unless they re-slice their view or the function returns the new slice.

### 4. The "Closure Capture" (Go < 1.22)
**Q**: Why does `for i := 0; i < 3; i++ { go func() { fmt.Println(i) }() }` print `3 3 3`?
*   **Answer**: The variable `i` is shared across all goroutines. By the time they run, the loop has finished.
*   **Note**: Fixed in Go 1.22! But legacy code is full of `i := i` shadowing fixes.

---

## 4️⃣ Learning Roadmap (Zero to Senior)

### Phase 1: Correctness (Weeks 1-2)
*   **Study**: Go Tour, "Effective Go" (Official Docs).
*   **Practice**: Write CLI tools. internalize `if err != nil`.
*   **Goal**: Write code that creates no lint errors.

### Phase 2: Concurrency Mastery (Weeks 3-4)
*   **Study**: "The Go Memory Model", Rob Pike's "Concurrency is not Parallelism".
*   **Practice**: Build a worker pool. Implement a graceful shutdown HTTP server.
*   **Goal**: Understand why `sync.Mutex` is sometimes better than channels.

### Phase 3: Performance & Internals (Weeks 5-6)
*   **Study**: Escape Analysis, GC Pacer, `pprof` deep dive.
*   **Practice**: Optimize a hot path allocation to zero. Benchmarking.
*   **Goal**: Know precisely where your variables live (Stack vs Heap).

### Phase 4: Distributed Systems (Ongoing)
*   **Study**: Distributed tracing, Circuit breakers in Go, gRPC.
*   **Practice**: Build a microservice with standard observability (metrics, logs, traces).

---

> *Built for high-performance engineering excellence.*



