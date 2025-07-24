# 🟦 **Go Collections**

## 🔹 **Session Breakdown**
___________________________________________________________________________________
| Module                            | Description                                 |
| --------------------------------- | ------------------------------------------- |
| Introduction to Collections in Go | Overview, memory model, value vs reference  |
| Arrays                            | Syntax, indexing, iteration, limitations    |
| Slices                            | Creation, slicing, make, nil slices, growth |
| Slice Manipulation & Internals    | Capacity, append, copy, memory sharing      |
| Maps                              | Key-value stores, existence checks, delete  |
| Built-in Functions                | `append`, `copy`, `delete`, len, cap        |
| Practical Examples & Use Cases    | Small real-world scenarios                  |
| Q\&A + Wrap-Up                    | Summary, resources, questions               |
-----------------------------------------------------------------------------------
---

## 🔸Introduction to Collections in Go

### ✦ Topics:

* What are collections? 
    Collections are data structures used to group and manage multiple values under a single variable. They allow developers to store, retrieve, iterate, and manipulate groups of values efficiently.
* Homogeneous types
    Collections in Go are homogeneous—meaning all elements in a collection must be of the same type.
* Go memory model: stack vs heap
    | Feature        | Stack                         | Heap                           |
    | -------------- | ----------------------------- | ------------------------------ |
    | **Scope**      | Function-local                | Global or shared between funcs |
    | **Lifetime**   | Short (tied to function call) | Long (until GC collects it)    |
    | **Speed**      | Very fast (LIFO)              | Slower (requires GC)           |
    | **Managed by** | Compiler                      | Garbage Collector (GC)         |
    | **Use case**   | Local variables, small data   | Data escaping function scope   |

* Value types vs reference types
In Go, data types are broadly categorized into **value types** and **reference types** based on how they are stored in memory and how they behave when copied or passed to functions.

---

## 🔹 Value Types

### ✅ Definition:

Value types **directly store their data**. When a value type is assigned to another variable or passed to a function, a **complete copy** of the data is made.
➡️ **Changes to the copy do not affect the original.**

### 💡 Examples:

* **Primitive types**: `int`, `float64`, `bool`, `string`

  > (Note: Strings have an underlying pointer to immutable data, but behave like value types.)
* **Arrays**: The entire array is copied on assignment or function call.
* **Structs**: All fields are copied when assigned or passed.

### 🧠 Memory Location:

* Typically allocated on the **stack**.

---

## 🔹 Reference Types

### ✅ Definition:

Reference types store a **memory address (reference)** to the actual data, not the data itself.
➡️ **Copying a reference type means both variables point to the same data** — changes via one reflect in the other.

### 💡 Examples:

* **Slices**: A view over an array with shared memory.
* **Maps**: References to a hash table.
* **Channels**: References to communication conduits.
* **Pointers**: Explicit references to memory addresses.
* **Functions**: Behave as reference types when passed or assigned.

### 🧠 Memory Location:

* The underlying data is typically stored on the **heap**, managed by Go’s garbage collector.

---

## 🔁 Key Differences: Summary Table

| Feature                 | Value Types                           | Reference Types                     |
| ----------------------- | ------------------------------------- | ----------------------------------- |
| **Data Storage**        | Directly stores data                  | Stores a reference (memory address) |
| **Copying Behavior**    | Full copy of the data                 | Only the reference is copied        |
| **Modification Impact** | Changes to copy don’t affect original | Changes affect original             |
| **Typical Memory**      | Stack                                 | Heap (for underlying data)          |

---

> 🧠 Tip: Use Go's `go build -gcflags="-m"` to inspect where variables are stored (stack vs heap).

---

✅ Primary Collections in Go
| Collection | Description                                 | Key Characteristics                         |
| ---------- | ------------------------------------------- | ------------------------------------------- |
| **Array**  | Fixed-length sequence of elements           | Size known at compile time, value type      |
| **Slice**  | Dynamic-length, flexible view over an array | Most common collection type, reference type |
| **Map**    | Key-value store                             | Dynamic size, fast lookup by key            |

### 🧪 Example for variables of primitive data types vs collections :

```go
x := 10
y := x
y = 20
fmt.Println(x) // Output: 10
```

Compare with:

```go
s1 := []int{1, 2, 3}
s2 := s1
s2[0] = 100
fmt.Println(s1) // Output: [100 2 3] — because slices share underlying array (slices are reference type)
```

---

## 🔸 Arrays in Go

### ✦ Topics:

* Fixed size: Arrays in Go have a fixed length defined at compile time and cannot be resized.
* Declaration and initialization: Arrays can be declared with a size and initialized using literals or default values.
* Iteration: You can iterate over arrays using for loops or for range to access index and value.
* Multidimensional arrays: Go supports arrays with multiple dimensions (e.g., 2D arrays), useful for grids or matrices.

### 🧪 Code Example:

```go
var arr [3]int // creates array with zero value of int [0, 0, 0]
arr[0] = 10
arr[1] = 20
arr[2] = 30

for i, v := range arr {
    fmt.Printf("Index %d: %d\n", i, v)
}

var strArr [3]string // creates array with zero value of string ["", "", ""]
for i, v := range strArr {
    fmt.Printf("Index %d: %d\n", i, v)
}
```

### 📝 Exercise:

* Create a 2D array and print it in matrix format.
* Write a program to sum of squares of all elements of an integer array.

---

## 🔸Slices

### ✦ Topics:

Here’s an expanded explanation for each key point about **slices in Go**, structured clearly for learning or inclusion in documentation:

---

## 🔹 Slices in Go

### 1. **What are slices?**

Slices are a flexible, dynamic data structures that provide a view over an array. Unlike arrays, slices can grow and shrink in size, making them the most commonly used collection type in Go.

---

### 2. **Relationship to arrays**

Internally, a slice consists of:

* A pointer to an underlying array,
* A length (number of elements),
* A capacity (maximum size before reallocation is needed).

Modifying elements in a slice affects the underlying array, and multiple slices can share the same array.

```go
arr := [5]int{10, 20, 30, 40, 50}
s := arr[1:4] // slice of elements: [20, 30, 40] , len = 3, cap = 4
s[0] = 99 // modifies arr[1], arr: [10, 99, 30, 40, 50]
s = append(s, 60) // OK, within array capacity, s: [20, 30, 40, 60], arr: [10, 20, 30, 40, 60]
s = append(s, 999) // triggers reallocation!! (cap exceeded)
fmt.Println(arr) // still [10, 99, 30, 40, 60], not affected
fmt.Println(s) // [20, 30, 40, 60, 999]
```

---

### 3. **`make` function**

The `make` function is used to create slices with a specified length and capacity. It allocates an underlying array and returns a slice referencing it.

```go
s := make([]int, 3, 5) // length: 3, capacity: 5
```

This is especially useful when you want to preallocate space for performance.

---

### 4. **Nil vs empty slice**

| Type        | Definition                     | `len()` | `cap()` | `== nil` |
| ----------- | ------------------------------ | ------- | ------- | -------- |
| Nil slice   | Declared but not initialized   | 0       | 0       | `true`   |
| Empty slice | Initialized with zero elements | 0       | ≥0      | `false`  |

```go
var a []int         // nil slice
b := []int{}        // empty slice

// note: append() can work with nil slice too.
a = append(a, 1, 2, 3) // works even if a is nil slice
```

Both are safe to use in most operations, but they behave differently when checked against `nil`.

```go
// Different ways to create a slice
var a []int              // nil slice
var b []int{1,2,3,4}     // slice initialized with elements
b := []int{}             // empty slice - commonly used
c := []int{1,2,3}        // slice initialized with elements
d := make([]int, 3, 5)   // empty slice with predefined length and capacity
```
---

### 5. **Iterating slices**

You can iterate over slices using a traditional `for` loop or a `for range` loop:

```go
nums := []int{1, 2, 3}
for i := 0; i < len(nums); i++ {
    fmt.Println(nums[i])
}

for index, value := range nums {
    fmt.Println(index, value)
}
```

---

### 6. **Slice literals**

Slice literals are a shorthand way to declare and initialize slices without explicitly specifying length or capacity.

```go
names := []string{"Alice", "Bob", "Charlie"}
```

Behind the scenes, this creates an array and a slice pointing to it.

### 🧪 Code Example:

```go
s := []string{"Go", "Rust", "Python"}
fmt.Println(len(s), cap(s)) // Output: 3, 3

s = append(s, "JavaScript")
fmt.Println(s)              // Output: ["Go" "Rust" "Python" "JavaScript"]
```

### 📝 Exercise:

* Use slicing to extract every second element from a list.
* Write a program to reverses a slice.

---

## 🔸Slice Manipulation & Internals

### ✦ Topics:

* `append()` and how Go handles resizing: The append() function adds elements to a slice and automatically allocates a new underlying array if capacity is exceeded.
* Underlying array sharing: Multiple slices can share the same underlying array, so changes in one slice may affect others.
* `copy()` function: copies elements from one slice to another and is useful for making independent copies of data.
* Capacity and reallocation: When a slice grows beyond its capacity, Go reallocates a new, larger array and moves the contents to it.
* Pitfalls: modifying shared slices - Mutating one slice can unintentionally change data in another slice if they share the same backing array.

### 🧪 Code Example:

```go
a := []int{1, 2, 3}
b := a[:2]        // Shares backing array
b[0] = 100
fmt.Println(a)    // Output: [100 2 3]
```

### 🧪 Deep Copy:

```go
original := []int{1, 2, 3}
copySlice := make([]int, len(original))
copy(copySlice, original)
```

### 📝 Exercise:

* Write a program that appends items to a slice and prints len/cap at each step.
* Implement a logic to extract even numbers from an array using slice and append.

---

## 🔸 Maps in Go

## 🔹 1. **What is a Map in Go?**

### ✅ Definition:

A **map** is a built-in data structure in Go that associates **keys with values** (like dictionaries in Python or hash maps in Java).

```go
ages := map[string]int{"Alice": 25, "Bob": 30}
```

### ✅ Characteristics:

* Unordered
* Dynamically sized
* Fast lookups, inserts, deletes (average O(1))
* Keys must be **comparable** (can’t be slices, maps, or functions)
| Field Type in Struct           | Can Be Map Key?    |
|-------------------------------|--------------------|
| `int`, `float64`              | ✅ Yes             |
| `string`, `bool`              | ✅ Yes             |
| `array` (of comparable types) | ✅ Yes             |
| `slice`, `map`, `func`        | ❌ No              |
| another struct (comparable)   | ✅ Yes             |
| another struct                | ❌ No *(if contains uncomparable fields)* |

---

## 🔹 2. **Declaring and Initializing Maps**

### a) **Using a map literal**:

```go
colors := map[string]string{
    "red":   "#FF0000",
    "green": "#00FF00",
}
```

### b) **Using `make()`**:
`In Go, you can create a map with an initial size (capacity) using the built-in `make` function:

---

#### ✅ Syntax

```go
m := make(map[KeyType]ValueType, initialCapacity)
```

* `KeyType` and `ValueType` are the types for the map.
* `initialCapacity` is **optional**, but if provided, Go will preallocate space to reduce allocations as the map grows.
* `initialCapacity` avoids repeated rehashing and resizing (until the capacity is reached), improving performance.

---

#### ✅ Example

```go
m := make(map[string]int, 100) // Preallocate for ~100 map/hash buckets

m["one"] = 1
m["two"] = 2

fmt.Println(len(m)) // Output: 2

scores := make(map[string]int) // without preallocating the map/hash buckets
scores["John"] = 85

```
---

### 📝 Notes

* The capacity is **not a hard limit**; the map can grow beyond the initial size.
* Unlike slices, we can't retrieve the capacity of a map using `cap()`.
---

### c) **Zero value is nil map**:

```go
var data map[string]int // nil map, cannot insert until initialized!!
```

```go
// Different ways to create a map
var a map[string]int                                  // nil map
var b map[string]int{"admin": 30, "employee": 20}     // map initialized with key-value pairs
b := map[string]int{}                                 // empty map - commonly used
c := map[string]int{"admin": 30, "employee": 20}      // map initialized key-value pairs
d := make(map[string]int, 5)   // empty map with predefined capacity
e := make(map[string]int)      // empty map without predefined capacity
```

---

## 🔹 3. **Basic Map Operations**

### ✅ Inserting / Updating:

```go
m := make(map[string]int)
m["apple"] = 3       // insert a kv pair
m["apple"] = 10      // update
```

### ✅ Retrieving:

```go
val := m["apple"]
fmt.Println(val) // prints 10
```

### ✅ Deleting:

```go
delete(m, "apple")
```

### ✅ Checking for Existence (`comma ok` idiom):

```go
val, ok := m["banana"]
if ok {
    fmt.Println("Found:", val)
} else {
    fmt.Println("Not found")
}
```

---

## 🔹 4. **Iterating Over Maps**

```go
for key, value := range m {
    fmt.Printf("%s: %d\n", key, value)
}
```

* Iteration order is **randomized** on each run (to prevent relying on order).

---

## 🔹 5. **Map Key and Value Types**

* **Valid key types**: strings, integers, floats, booleans, structs (comparable types).
* **Invalid key types**: slices, maps, functions (not comparable).

```go
type Point struct {
    X, Y int
}
locations := map[Point]string{
    {1, 2}: "origin",
}
```

---

## 🔹 6. **Maps of Slices, Slices of Maps**

### Map of slices:

```go
groups := map[string][]string{
    "devs": {"Alice", "Bob"},
}
```

### Slice of maps:

```go
records := []map[string]int{
    {"math": 90},
    {"science": 85},
}
```

---

## 🔹 7. **Nil Maps vs Empty Maps**

| Type      | Initialized? | Can store values? |
| --------- | ------------ | ----------------- |
| `nil` map | No           | ❌ panic on write  |
| Empty map | Yes          | ✅ usable          |

```go
var m1 map[string]int      // nil
m2 := make(map[string]int) // empty
m3 := map[string]int{} // empty
```

---

## 🔹 8. **Pitfalls & Gotchas**

* Modifying a map during iteration is allowed but can lead to unpredictable behavior.
* Use caution when **sharing maps between goroutines**—they are **not thread-safe**. Use `sync.Mutex` or `sync.Map` for concurrency. (will be discussed in future sessions along with concurrency primitives)

---

## 🔹 9. **Useful Patterns**

### a) **Counting Frequencies**:

```go
words := []string{"go", "go", "lang"}
freq := make(map[string]int)

for _, word := range words {
    freq[word]++
}
```

### b) **Grouping by Key**:

```go
names := []string{"alice", "bob", "amy", "brian"}
groups := make(map[byte][]string)

for _, name := range names {
    first := name[0]
    groups[first] = append(groups[first], name) // grouping names based on starting letter
}
```

---

## 🔹 10. **Advanced: `sync.Map` (Concurrent Maps)**

For concurrent access without locking:

```go
var m sync.Map

m.Store("foo", 42)
val, ok := m.Load("foo")
```

Note: `sync.Map` has different semantics and is optimized for specific use cases. This will be covered in future sessions in detail.

---

## 🧪 Hands-on Exercises

1. Write a program to count character frequencies in a string.
2. Group numbers into categores: +ve&even, -ve&even, +ve&odd, -ve&odd.
3. Create a map of slices representing employees in departments.
4. Implement a program that merges two maps (overwriting common keys).
5. Write a phone book program (name → phone number).
6. Group students by grade level using `map[int][]string`.
7. Write a program that uses `delete` to remove inactive users from a map (inactivity refers to users who didn't login for > 30 days).

---

## ✅ Summary

| Feature       | Description                           |
| ------------- | ------------------------------------- |
| Declaration   | Use `map[KeyType]ValueType`           |
| Zero value    | `nil`, must initialize before writing |
| Insert/Update | `m[key] = value`                      |
| Delete        | `delete(m, key)`                      |
| Lookup        | `val, ok := m[key]`                   |
| Iteration     | `for k, v := range m`                 |
| Concurrency   | Use `sync.Mutex` or `sync.Map`        |

---

❓ Why No Set, List, Queue in golang?
Go keeps things simple. If you need structures like:

* Set → Use map[T]bool
* Queue/Stack → Use slices with append and custom logic
* LinkedList, Tree → Implement manually or use third-party libs

## 🔸 Real-World Practical Examples

### ✅ Mini Projects:

#### 1. **To-Do List Manager**

* Use slices for tasks
* Use `append`, `delete` (manual via slice tricks)
* Show completed vs pending

#### 2. **Shopping Cart**

* Map of item → quantity
* Add, remove, update quantity
* Total price calculation

#### 3. **Tag-Based Note Organizer**

* Map\[string]\[]string — tag → notes
* Add, search, delete notes

---

## 🔸 Wrap-Up & Q\&A

### ✅ Recap:

* Arrays vs Slices
* Shared memory issues
* Dynamic map behavior
* Built-in tools to manipulate collections

### 🎁 Resources:

* [Go Slices: Usage and Internals](https://go.dev/blog/slices-intro)
* [Effective Go](https://golang.org/doc/effective_go.html)
* [Go by Example](https://gobyexample.com/)

---
