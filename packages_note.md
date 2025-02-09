# **📌 The `fmt` Package in GoLang (In-Depth Guide)**
The **`fmt`** package in Go (**Format Package**) is used for **formatted I/O (Input/Output)** operations, such as printing to the console and reading user input.

👉 **Import the `fmt` package** before using it:
```go
import "fmt"
```

---

## **🔹 1. Printing Functions in `fmt`**
### ✅ **`fmt.Print()`** - Prints without a newline
- It prints text **without** adding a newline (`\n`) at the end.
- If multiple arguments are passed, they are **concatenated** with a space.

✔️ **Example:**
```go
package main

import "fmt"

func main() {
    fmt.Print("Hello") 
    fmt.Print(" World!") 
}
```
✔️ **Output:**
```
Hello World!
```

---

### ✅ **`fmt.Println()`** - Prints with a newline
- Prints text and **automatically adds a newline (`\n`)** at the end.
- Multiple arguments are **separated by a space**.

✔️ **Example:**
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello")
    fmt.Println("GoLang")
}
```
✔️ **Output:**
```
Hello
GoLang
```

✔️ **With multiple arguments:**
```go
fmt.Println("Name:", "John", "Age:", 30)
```
✔️ **Output:**
```
Name: John Age: 30
```

---

### ✅ **`fmt.Printf()`** - Formatted output
- Allows **format specifiers** (`%d`, `%s`, `%f`, etc.).
- **Does NOT add a newline** automatically.

✔️ **Example:**
```go
package main

import "fmt"

func main() {
    name := "John"
    age := 25
    fmt.Printf("Name: %s, Age: %d", name, age)
}
```
✔️ **Output:**
```
Name: John, Age: 25
```

📌 **Common Format Specifiers:**
| **Specifier** | **Description** | **Example** |
|-------------|----------------|------------|
| `%s` | String | `"Hello"` → `fmt.Printf("%s", "Hello")` |
| `%d` | Integer (base 10) | `123` → `fmt.Printf("%d", 123)` |
| `%f` | Float | `3.14` → `fmt.Printf("%f", 3.14)` |
| `%.2f` | Float with 2 decimal places | `3.1415` → `3.14` |
| `%t` | Boolean | `true` → `fmt.Printf("%t", true)` |
| `%v` | Any value (default format) | `fmt.Printf("%v", value)` |
| `%T` | Type of variable | `fmt.Printf("%T", 42)` (prints `int`) |

✔️ **Example with multiple format specifiers:**
```go
package main

import "fmt"

func main() {
    price := 49.99
    available := true
    fmt.Printf("Price: %.2f, Available: %t\n", price, available)
}
```
✔️ **Output:**
```
Price: 49.99, Available: true
```

---

### ✅ **`fmt.Sprintf()`** - Returns formatted string
- Similar to `fmt.Printf()`, but **returns** the formatted string instead of printing.

✔️ **Example:**
```go
package main

import "fmt"

func main() {
    name := "Alice"
    age := 30
    message := fmt.Sprintf("Name: %s, Age: %d", name, age)
    fmt.Println(message) // Prints: Name: Alice, Age: 30
}
```

---

## **🔹 2. Reading Input in `fmt`**
### ✅ **`fmt.Scan()`** - Reads space-separated input
- Reads user input **until a space** is encountered.
- Stores values in the provided **memory addresses (`&variable`)**.

✔️ **Example:**
```go
package main

import "fmt"

func main() {
    var name string
    fmt.Print("Enter your name: ")
    fmt.Scan(&name) // Reads input until a space
    fmt.Println("Hello,", name)
}
```
✔️ **Input:**
```
John Doe
```
✔️ **Output:**
```
Hello, John
```
⚠ **Note:** Only "John" is captured because `fmt.Scan()` stops at a space.

---

### ✅ **`fmt.Scanln()`** - Reads entire line
- Reads input **until a newline (`\n`)**.
- Useful for multi-word input.

✔️ **Example:**
```go
package main

import "fmt"

func main() {
    var name string
    fmt.Print("Enter your full name: ")
    fmt.Scanln(&name) // Reads input until newline
    fmt.Println("Hello,", name)
}
```

✔️ **Input:**
```
John Doe
```
✔️ **Output:**
```
Hello, John Doe
```

⚠ **Issue:** `fmt.Scanln()` still only captures **one word**.  
✅ **Solution:** Use `bufio.Reader` instead.

---

### ✅ **`fmt.Scanf()`** - Reads formatted input
- Similar to `fmt.Scan()`, but allows **format specifiers**.

✔️ **Example:**
```go
package main

import "fmt"

func main() {
    var name string
    var age int
    fmt.Print("Enter name and age: ")
    fmt.Scanf("%s %d", &name, &age) // Read formatted input
    fmt.Printf("Hello %s, you are %d years old.\n", name, age)
}
```

✔️ **Input:**
```
Alice 25
```
✔️ **Output:**
```
Hello Alice, you are 25 years old.
```

---

## **🔹 3. `fmt.Errorf()` - Custom Error Formatting**
- Used to create formatted error messages.

✔️ **Example:**
```go
package main

import (
    "fmt"
    "errors"
)

func main() {
    err := fmt.Errorf("error: %s", "file not found")
    fmt.Println(err)
}
```
✔️ **Output:**
```
error: file not found
```

---

## **🚀 Summary of `fmt` Package Functions**
| **Function** | **Description** |
|-------------|----------------|
| `fmt.Print()` | Print without newline |
| `fmt.Println()` | Print with newline |
| `fmt.Printf()` | Print with format specifiers |
| `fmt.Sprintf()` | Format string and return it |
| `fmt.Scan()` | Read space-separated input |
| `fmt.Scanln()` | Read entire line |
| `fmt.Scanf()` | Read formatted input |
| `fmt.Errorf()` | Create formatted errors |

---
# **📌 Functions in GoLang (Detailed Explanation)**  

A **function** in Go (**GoLang**) is a reusable block of code that performs a specific task. Functions improve **code organization, reusability, and maintainability**.

---

# **🔹 1. Defining a Function in Go**
A function in Go is declared using the **`func`** keyword, followed by:  
✅ Function name  
✅ Parameters (optional)  
✅ Return type (optional)  
✅ Function body enclosed in `{}`  

### ✅ **Basic Function Example**
```go
package main

import "fmt"

// Function declaration
func greet() {
    fmt.Println("Hello, GoLang!")
}

func main() {
    greet() // Function call
}
```
✔️ **Output:**
```
Hello, GoLang!
```

---

# **🔹 2. Function Parameters in Go**
Functions can accept **parameters** to perform operations based on user input.

### ✅ **Function with One Parameter**
```go
package main

import "fmt"

// Function with a parameter
func greet(name string) {
    fmt.Println("Hello,", name)
}

func main() {
    greet("Alice") // Function call with argument
}
```
✔️ **Output:**
```
Hello, Alice
```

---

### ✅ **Function with Multiple Parameters**
```go
package main

import "fmt"

// Function with multiple parameters
func add(a int, b int) {
    sum := a + b
    fmt.Println("Sum:", sum)
}

func main() {
    add(5, 10) // Function call with arguments
}
```
✔️ **Output:**
```
Sum: 15
```

🛑 **Note:** Instead of specifying `int` twice, we can write:
```go
func add(a, b int) { ... }
```

---

# **🔹 3. Return Values in Functions**
Functions can return **values** using the `return` keyword.

### ✅ **Function with Return Value**
```go
package main

import "fmt"

// Function that returns a value
func square(n int) int {
    return n * n
}

func main() {
    result := square(4)
    fmt.Println("Square:", result)
}
```
✔️ **Output:**
```
Square: 16
```

---

### ✅ **Function with Multiple Return Values**
Go allows **multiple return values**, useful for returning multiple results.

✔️ **Example:**
```go
package main

import "fmt"

// Function returning multiple values
func divide(a, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}

func main() {
    q, r := divide(10, 3)
    fmt.Println("Quotient:", q, "Remainder:", r)
}
```
✔️ **Output:**
```
Quotient: 3 Remainder: 1
```

---

# **🔹 4. Named Return Values**
Instead of returning values explicitly, we can use **named return variables**.

✔️ **Example:**
```go
package main

import "fmt"

// Named return values
func rectangleDimensions(length, width int) (area int, perimeter int) {
    area = length * width
    perimeter = 2 * (length + width)
    return // No explicit return needed!
}

func main() {
    a, p := rectangleDimensions(5, 3)
    fmt.Println("Area:", a, "Perimeter:", p)
}
```
✔️ **Output:**
```
Area: 15 Perimeter: 16
```

---

# **🔹 5. Variadic Functions (Multiple Arguments)**
A function can accept a **variable number of arguments** using `...` (ellipsis).

### ✅ **Example: Sum of Multiple Numbers**
```go
package main

import "fmt"

// Variadic function to sum numbers
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3))       // 6
    fmt.Println(sum(10, 20, 30, 40)) // 100
}
```
✔️ **Output:**
```
6
100
```

🛑 **Note:** The function accepts **any number of arguments**.

---

# **🔹 6. Anonymous Functions**
Anonymous functions are functions **without a name**.  
They can be assigned to variables or executed **immediately**.

✔️ **Example: Assigning to a Variable**
```go
package main

import "fmt"

func main() {
    // Anonymous function assigned to a variable
    greet := func(name string) {
        fmt.Println("Hello,", name)
    }
    
    greet("GoLang") // Function call
}
```
✔️ **Output:**
```
Hello, GoLang
```

✔️ **Example: Immediately Invoked Function**
```go
package main

import "fmt"

func main() {
    // Immediately executing anonymous function
    func(msg string) {
        fmt.Println(msg)
    }("Welcome to Go!")
}
```
✔️ **Output:**
```
Welcome to Go!
```

---

# **🔹 7. Higher-Order Functions (Passing Functions as Arguments)**
Functions can be passed as **arguments** to other functions.

✔️ **Example: Passing a Function**
```go
package main

import "fmt"

// Function that takes another function as a parameter
func applyOperation(a, b int, operation func(int, int) int) int {
    return operation(a, b)
}

func main() {
    add := func(x, y int) int {
        return x + y
    }
    
    result := applyOperation(5, 3, add)
    fmt.Println("Result:", result)
}
```
✔️ **Output:**
```
Result: 8
```

---

# **🔹 8. Recursive Functions**
A function can call **itself** (recursion).

✔️ **Example: Factorial Using Recursion**
```go
package main

import "fmt"

// Recursive function
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    fmt.Println("Factorial of 5:", factorial(5))
}
```
✔️ **Output:**
```
Factorial of 5: 120
```

---

# **🔹 9. First-Class Functions in Go**
Go treats functions as **first-class citizens**, meaning:
✅ Functions can be assigned to variables.  
✅ Functions can be passed as arguments.  
✅ Functions can return other functions.

✔️ **Example: Returning a Function**
```go
package main

import "fmt"

// Function returning another function
func multiplier(x int) func(int) int {
    return func(y int) int {
        return x * y
    }
}

func main() {
    double := multiplier(2) // Function that doubles numbers
    fmt.Println(double(5))  // 10
}
```
✔️ **Output:**
```
10
```

---

# **🎯 Summary of Functions in Go**
| **Concept** | **Example** |
|------------|------------|
| **Basic Function** | `func greet() { fmt.Println("Hello") }` |
| **With Parameters** | `func add(a, b int) { fmt.Println(a + b) }` |
| **With Return Value** | `func square(n int) int { return n * n }` |
| **Multiple Returns** | `func divide(a, b int) (int, int) { return a / b, a % b }` |
| **Named Return Values** | `func rectangle(length, width int) (area int, perimeter int) { ... }` |
| **Variadic Function** | `func sum(nums ...int) int { ... }` |
| **Anonymous Function** | `greet := func(name string) { fmt.Println(name) }` |
| **Passing Functions** | `func applyOperation(a, b int, op func(int, int) int) int { return op(a, b) }` |
| **Recursive Function** | `func factorial(n int) int { if n == 0 { return 1 } return n * factorial(n-1) }` |

---
