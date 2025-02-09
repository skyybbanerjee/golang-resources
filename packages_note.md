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