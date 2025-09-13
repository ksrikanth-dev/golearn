# Control Structures

## Objective:

* Introduce the types of control structures in Go and their significance in program flow.
* Understanding and hands-on experience with Go's control structures. These include conditional statements, loops, switch cases, and control keywords like break, continue, and goto.

## Topics Covered:

* What are control structures?
* Why are they important?
* Types of control structures in Go

## What are control structures:

Control structures dictate the flow of execution within a Go program. Understanding them is crucial to writing logical, efficient, and bug-free code.

### Types:

* Conditional: `if`, `if-else`, `if-else if`
* Multi-branch: `switch`, `type switch`
* Looping: `for`, `for-range`
* Control flow: `break`, `fallthrough`, `continue`, `goto`
* Error handling: `defer`, `panic`, `recover` (will be discussed separately in future session)

---

## The `if`, `else if`, and `else` Statements

### Topics Covered:

* Basic `if` statement
* `else if` and `else`
* Short variable declaration in `if`
* Nesting `if` statements

### Code Example:

```go
age := 20
failedDlTest := true
if age >= 18 {
    fmt.Println("Eligible to vote")
} else if age >= 16 {
    fmt.Println("Can apply for learner's license")
    if failedDlTest {
        fmt.Println("Wait for 6 weeks before redoing the test")
    }

} else {
    fmt.Println("Too young")
}

scores := map[string]int{"Alex": 50, "John": 20}
student := "Alex"
if score, ok := scores[student]; ok {
    fmt.Println(student, " scored ", score)
} else {
    fmt.Printl(student, " score is not available")
}
```

### Exercise:

* Write a Go program to categorize a number as positive, negative, or zero.
* Write a program to categorize a group of numbers based on divisibility by 2, 3, 4, 5.
---

## `switch` Statements

Use `switch` statements to simplify complex `if-else` chains.

### Topics Covered:

* Simple switch syntax
* Multiple expressions in a single case
* Switch without condition (like `if-else if` ladder)
* Fallthrough behavior

### Code Example:

```go
switch ccType {
case "Exclusive":
    fmt.Println("credit card has exclusive benefits")
    break // optional, enforced by default, doesn't fall through like C, C++
case "Premium":
    fmt.Println("credit card has premium benefits")
    fallthrough // requesting for fallthrough explicitely
case "Traveller":
    fmt.Println("credit card has traveller benefits")
case "Shopper":
    fmt.Println("credit card has shopper benefits")
default:
    fmt.Println("credit card has 1% cash back")
}
```

### Hands-on:

Create a program that uses a switch to display weekdays based on input.
---

In Go, the expressions used in `switch` cases must be **comparable types** — meaning they can be used with the `==` operator.

---

### ✅ Suitable Data Types for `switch case`

You can use **any data type that supports equality comparison** as a case expression. These include:

#### 1. **Basic Types**

* `int`, `int8`, `int16`, `int32`, `int64`
* `uint`, `uint8`, `uint16`, `uint32`, `uint64`
* `float32`, `float64` *(allowed but uncommon)*
* `string`
* `bool`
* `rune` and `byte` (aliases for `int32` and `uint8`)

✅ Example:

```go
var grade string = "B"
switch grade {
case "A":
    fmt.Println("Excellent")
case "B":
    fmt.Println("Good")
case "C":
    fmt.Println("Average")
default:
    fmt.Println("Invalid grade")
}
```

---

#### 2. **Constants (typed or untyped)**

You can use constant expressions as case values:

```go
const (
    Red   = 1
    Blue  = 2
    Green = 3
)

color := 2
switch color {
case Red:
    fmt.Println("Red")
case Blue:
    fmt.Println("Blue")
}
```

---

#### 3. **Enums using iota**

Go doesn't have enums like C/C++, but `iota` lets you simulate them.

```go
type Day int

const (
    Sunday Day = iota
    Monday
    Tuesday
)

var d Day = Monday
switch d {
case Sunday:
    fmt.Println("Weekend")
case Monday:
    fmt.Println("Weekday")
}
```

---

#### 4. **Type switches (special case for interfaces)**

If you're using a **type switch**, the value must be an **interface**.

```go
var x interface{} = 10

switch v := x.(type) {
case int:
    fmt.Println("int:", v)
case string:
    fmt.Println("string:", v)
}
```

---

#### ❌ Unsupported Types

You **cannot** use types that are not comparable in regular switch cases:

| ❌ Unsupported | Reason                                    |
| ------------- | ----------------------------------------- |
| `slice`       | Not comparable with `==`                  |
| `map`         | Not comparable                            |
| `func`        | Not comparable                            |
| `struct`      | Only if it contains non-comparable fields |

---

#### 🧪 Summary Table

| Type                    | Can be used in case?                    |
| ----------------------- | --------------------------------------- |
| `int`, `string`, `bool` | ✅ Yes                                   |
| `slice`, `map`, `func`  | ❌ No                                    |
| `interface{}`           | ✅ In type switch                        |
| `struct`                | ✅ Only if all fields are comparable     |
| `float64`               | ✅ But use cautiously (due to precision) |

---


## Loops (`for`, `range`)

### Topics Covered:

* Traditional `for` loop with init, condition, post
* `while` style loop (`for condition {}`)
* Infinite loop with `for {}`
* Looping over arrays, slices, maps, strings with `range`
* Accessing index and value

### Code Example:

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

```go
nums := []int{1, 2, 3}
i := 0
for i < len(nums) {
    fmt.Println(i, nums[i])
    i++
}
```

```go
for {
    readFromSocket() // blocking function that reads data from socket
}
```


```go
nums := []int{1, 2, 3}
fmt.Println("index and value: ")
for index, value := range nums {
    fmt.Println(index, value)
}
fmt.Println("Only values: ")
for _, value := range nums {
    fmt.Println(value)
}
fmt.Println("Only indices: ")
for index, _ := range nums {
    fmt.Println(index)
}
for index := range nums { // similar operation for maps returns only keys
    fmt.Println(index)
}
fmt.Println("Looping over a certain number of times: ")
for _, _ := range nums {
    fmt.Println("Do some task")
}

```

### Exercise:

* Sum all elements of a slice using `range`.
* Count the number of vowels in a string.
* Write a program to loop over arrays, slices, maps, strings with `range`

---

## Control Keywords – `break`, `fallthrough`, `continue`, `goto`

Manipulate loop execution using control statements.

### Topics Covered:

* `break` to exit loops early
* `continue` to skip to the next iteration
* Using labels
* Introduction to `goto` (used rarely)

### Code Example:

```go
for i := 1; i <= 10; i++ {
    if i == 5 {
        continue
    }
    if i == 8 {
        break
    }
    fmt.Println(i)
}
```

### Goto Example:

```go
count := 0
Here: // this a label
    fmt.Println(count)
    count++
    if count < 3 {
        goto Here
    }
```

### Hands-on:

* Write a program to skip printing multiples of 3.
* Write a loop that exits using `break` when a random number is divisible by 7.

---

## Real-world Scenarios and Best Practices

### Objective:

Apply all learned control structures in real-world inspired examples.

### Activities:

* Simulate a login system (conditions, loops, break/continue)
* Build a command-line menu (switch, loops, labels)
* Create a mini calculator using `switch` and loop

### Discussion:

* When to use which control structure
* Readability and maintainability considerations

### * Resources:

  * [Go Tour](https://tour.golang.org/flowcontrol/1)
  * [Effective Go - Control Structures](https://golang.org/doc/effective_go.html#control-structures)

---
