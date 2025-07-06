## 🎯 **Understanding Data Types and Operators in Golang**

### 📌 **Session Objectives**

By the end of this session, you will be able to:

* Understand the basic and composite data types in Go.
* Use type conversion and inference effectively.
* Apply various operators in Go (arithmetic, relational, logical, etc.)
* Write basic Go programs using appropriate data types and operators.

---

## 🗂️ **Session Outline**

### 1. **Introduction**

* Quick overview of Go's philosophy: simplicity, performance, concurrency.
* Data types and operators: building blocks of any Go program.

---

### 2. **Basic Data Types in Go**

#### a. **Numeric Types**

```go
int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64
float32, float64
complex64, complex128
byte // alias for uint8
rune // alias for int32, used to represent a Unicode code point
```

✅ *Example:*

```go
var age int = 30
var pi float64 = 3.14
```

```go
package main

import "fmt"

func main() {
    var r rune = '♥'
    fmt.Printf("Rune: %c, Unicode: %U, Decimal: %d\n", r, r, r)
}
```

## 🧠 Runes in Strings

In Go, strings are UTF-8 encoded by default. If you want to iterate over actual Unicode characters (not just bytes), use a loop over `[]rune` or `range`:

```go
s := "Go♥"
for i, r := range s {
    fmt.Printf("Index: %d, Rune: %c, Unicode: %U\n", i, r, r)
}
```

## ⚠️ Runes vs Bytes

* A **byte** in Go is an alias for `uint8` and represents a single **byte** (0–255).
* A **rune** is an alias for `int32` and can represent a full Unicode character (including emojis, non-English characters).

---

## 🧪 Real-World Use Cases for runes

* Handling non-English input (e.g., Chinese, Arabic, emojis)
* Validating characters in strings
* Implementing custom parsers or lexers

---

#### b. **Boolean Type**

```go
var isActive bool = true
```

#### c. **String Type**

```go
var name string = "Gopher"
```

🔹 **String operations:**

```go
len(name), name[0], name + "s"
```

#### d. **Zero Values**

Each type has a default zero value:

* `int → 0`, `float → 0.0`, `bool → false`, `string → ""`

---

### 3. **Type Inference and Conversion **

#### a. **Short Variable Declaration**

```go
x := 100 // kind of a syntactical sugar for `var x int = 5`
```

#### b. **Type Conversion**

```go
var a int = 42
var b float64 = float64(a)
```

🔸 Go is **strict** about conversions—no implicit casting.

---

### 4. **Composite Data Types **

Brief overview (detailed in future sessions):

* **Arrays**
* **Slices**
* **Maps**
* **Structs**

✅ *Example:*

```go
var nums = [3]int{1, 2, 3}
var cities = map[string]string{"IN": "Delhi", "US": "New York"}
```

---
#### Constants and iota:

In Go (Golang), `iota` is a **predeclared identifier** used to simplify the definition of **incrementing constants**. It's most commonly used within `const` blocks to generate a sequence of related values.

##### How `iota` works:

* It starts at **0** in each `const` block.
* It increments by **1** for each new line in the block.
* You can use it with expressions to generate patterns (e.g., bit shifting, enums, masks).

---

##### 🔹 Basic Example:

```go
const (
    A = iota // 0
    B        // 1
    C        // 2
)
```

This is equivalent to:

```go
const (
    A = 0
    B = 1
    C = 2
)
```

---

##### 🔹 Using Expressions:

```go
const (
    _  = iota             // ignore first value (0)
    KB = 1 << (10 * iota) // 1 << (10*1)
    MB = 1 << (10 * iota) // 1 << (10*2)
    GB = 1 << (10 * iota) // 1 << (10*3)
)
```

Output values:

* `KB = 1024`
* `MB = 1048576`
* `GB = 1073741824`

---

##### 🔹 Skipping Values:

You can use `_` to skip an `iota` value:

```go
const (
    Red = iota
    _
    Blue
)
```

Here, `Red = 0`, `Blue = 2`

---

##### 🔹 Mixing with Other Constants:

```go
const (
    FlagNone = 0
    FlagA    = 1 << iota // 1 << 0 = 1
    FlagB                // 1 << 1 = 2
    FlagC                // 1 << 2 = 4
)
```

This is useful for creating **bit masks**.

#### Creating Enums with Custom Types:

```go
package main

import "fmt"

// Custom type for LogLevel
type LogLevel int

const (
    Debug LogLevel = iota // 0
    Info                  // 1
    Warning               // 2
    Error                 // 3
    Fatal                 // 4
)

func (l LogLevel) String() string {
    return [...]string{"Debug", "Info", "Warning", "Error", "Fatal"}[l]
}

func main() {
    level := Warning
    fmt.Println("Log Level:", level)         // Output: Log Level: 2
    fmt.Println("Log Level:", level.String()) // Output: Log Level: Warning
}
```

---

##### Summary:

* `iota` is reset to `0` in every new `const` block.
* It’s great for enums, flags, and patterns.
* It's evaluated at **compile time**, so no runtime cost.

---

### 5. **Operators in Go**

#### a. **Arithmetic Operators**

```go
+ - * / %
```

✅ *Example:*

```go
a, b := 10, 3
fmt.Println(a + b, a - b, a * b, a / b, a % b)
```

#### b. **Relational Operators**

```go
== != > < >= <=
```

#### c. **Logical Operators**

```go
&& || !
```

✅ *Example:*

```go
x := 5
fmt.Println(x > 3 && x < 10)
```

#### d. **Bitwise Operators**

Go supports these bitwise operators on integer types:
```go
& | ^ &^ << >>
| Operator | Description         | Example  |     |     |
| -------- | ------------------- | -------- | --- | --- |
| `&`      | AND                 | `a & b`  |     |     |
| `|`      | OR                  | `a | b`  |     |     |
| `^`      | XOR (exclusive OR)  | `a ^ b`  |     |     |
| `&^`     | AND NOT (bit clear) | `a &^ b` |     |     |
| `<<`     | Left shift          | `a << 2` |     |     |
| `>>`     | Right shift         | `a >> 2` |     |     |

```

✅ *Example:*

```go
a, b := 5, 3 // binary: 101, 011
fmt.Println(a & b)  // 1
```

---

## 1️⃣ Bitwise AND `&`

```go
package main

import "fmt"

func main() {
    a := 12  // binary 1100
    b := 10  // binary 1010

    fmt.Printf("%b & %b = %b\n", a, b, a&b)
}
```

**Output:**

```
1100 & 1010 = 1000
```

Explanation: Only bits set in **both** `a` and `b` remain 1 (`1000` = 8).

---

## 2️⃣ Bitwise OR `|`

```go
a := 12  // 1100
b := 10  // 1010

fmt.Printf("%b | %b = %b\n", a, b, a|b)
```

**Output:**

```
1100 | 1010 = 1110
```

Explanation: Bits set in **either** `a` or `b` become 1 (`1110` = 14).

---

## 3️⃣ Bitwise XOR `^`

```go
a := 12  // 1100
b := 10  // 1010

fmt.Printf("%b ^ %b = %b\n", a, b, a^b)
```

**Output:**

```
1100 ^ 1010 = 110
```

Explanation: Bits set in **only one** operand become 1 (`0110` = 6).

---

## 4️⃣ AND NOT `&^` (Clear bits)

```go
a := 12  // 1100
b := 10  // 1010

fmt.Printf("%b &^ %b = %b\n", a, b, a&^b)
```

**Output:**

```
1100 &^ 1010 = 100
```

Explanation: Clears bits of `a` where `b` has bits set.

* `a` = 1100
* `b` = 1010
* Result: Keep bits from `a` that are **not set** in `b` → `0100` (decimal 4)

---

## 5️⃣ Left Shift `<<`

```go
a := 3   // 11 in binary

fmt.Printf("%b << 2 = %b\n", a, a<<2)
```

**Output:**

```
11 << 2 = 1100
```

Explanation: Shifts bits left by 2, equivalent to multiplying by 2² = 4.

* `3 << 2` = `3 * 4 = 12` (binary 1100)

---

## 6️⃣ Right Shift `>>`

```go
a := 12  // 1100

fmt.Printf("%b >> 2 = %b\n", a, a>>2)
```

**Output:**

```
1100 >> 2 = 11
```

Explanation: Shifts bits right by 2, equivalent to dividing by 2² = 4.

* `12 >> 2` = `12 / 4 = 3` (binary 11)

---

# Bonus: Using bitwise operators for flags

```go
const (
    FlagRead  = 1 << 0 // 0001
    FlagWrite = 1 << 1 // 0010
    FlagExec  = 1 << 2 // 0100
)

func main() {
    var permissions byte = FlagRead | FlagWrite

    fmt.Printf("Permissions: %b\n", permissions) // 0011

    // Check if write permission is set
    if permissions&FlagWrite != 0 {
        fmt.Println("Write permission granted")
    }

    // Remove write permission
    permissions &^= FlagWrite
    fmt.Printf("Permissions after removing write: %b\n", permissions) // 0001
}
```

---

#### e. **Assignment & Composite Assignment Operators**

```go
= += -= *= /= %= <<= >>= &= ^= |=
```

## 1️⃣ Basic Assignment `=`

Assigns the value on the right to the variable on the left.

```go
package main

import "fmt"

func main() {
    var x int
    x = 10
    fmt.Println("x =", x) // Output: x = 10
}
```

---

## 2️⃣ Addition Assignment `+=`

Adds the right operand to the left operand and assigns the result to the left operand.

```go
x := 5
x += 3  // Equivalent to: x = x + 3
fmt.Println("x after += 3:", x) // Output: 8
```
Other operators usage can be infered from this example...

### Bitwise Composite Assignment Example:

```go
x := 12  // 1100 in binary
y := 10  // 1010 in binary

x &= y   // x = x & y => 1100 & 1010 = 1000 (8)
fmt.Println("x after &= y:", x)

x |= 2   // x = x | 0010 => 1000 | 0010 = 1010 (10)
fmt.Println("x after |= 2:", x)

x <<= 1  // x = x << 1 => 1010 << 1 = 10100 (20)
fmt.Println("x after <<= 1:", x)
```

#### f. **Other Operators**

* `&` (address-of)
* `*` (pointer dereference)
* `:=` (short declaration)

---

### Summary Table of Composite Assignments

| Operator | Meaning                    | Equivalent To         |         |     |
| -------- | -------------------------- | --------------------- | ------- | --- |
| `+=`     | Add and assign             | `x = x + y`           |         |     |
| `-=`     | Subtract and assign        | `x = x - y`           |         |     |
| `*=`     | Multiply and assign        | `x = x * y`           |         |     |
| `/=`     | Divide and assign          | `x = x / y`           |         |     |
| `%=`     | Modulus and assign         | `x = x % y`           |         |     |
| `&=`     | Bitwise AND and assign     | `x = x & y`           |         |     |
| `|=`     | Bitwise OR and assign      | `x = x | y`           |         |     |
| `^=`     | Bitwise XOR and assign     | `x = x ^ y`           |         |     |
| `&^=`    | Bitwise AND NOT and assign | `x = x &^ y`          |         |     |
| `<<=`    | Left shift and assign      | `x = x << y`          |         |     |
| `>>=`    | Right shift and assign     | `x = x >> y`          |         |     |

---

### 6. **Hands-on Exercises**

1. Declare variables of all basic types and print their zero values.
2. Write a program to calculate area and perimeter of a rectangle.
3. Use logical operators to check if a number is in a range.
4. Understand bitwise operations with examples.

---
