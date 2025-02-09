### **What is GoLang?**  
Go (or GoLang) is an open-source, statically typed, compiled programming language developed by Google in 2007 and released in 2009. It was designed to improve productivity in software development, combining the performance of a compiled language with the simplicity and ease of use of dynamically typed languages.  

Go is particularly known for its efficiency in handling concurrent programming, making it ideal for building scalable and high-performance applications.  

---

### **Key Features of GoLang:**  

1. **Simple and Readable Syntax**  
   - Go has a clean and minimalistic syntax, making it easy to learn and read.  
   - No complex features like classes, inheritance, or explicit memory management.  

2. **Compiled and High Performance**  
   - Unlike interpreted languages (e.g., Python, JavaScript), Go is compiled, resulting in faster execution.  
   - Produces standalone binaries with minimal dependencies.  

3. **Garbage Collection (GC)**  
   - Go has an efficient garbage collector that manages memory automatically.  

4. **Concurrency Support (Goroutines)**  
   - Go has built-in concurrency using lightweight threads called **goroutines**.  
   - It uses **channels** for safe communication between goroutines, making it more efficient than traditional threading.  

5. **Static Typing with Type Inference**  
   - Statically typed, reducing runtime errors.  
   - Also supports type inference, making code more concise.  

6. **Cross-Platform**  
   - Go supports cross-platform development.  
   - A single Go program can be compiled to run on Windows, Linux, and macOS without changes.  

7. **Standard Library**  
   - Comes with a powerful standard library, eliminating the need for third-party packages for common tasks like HTTP handling, JSON parsing, and cryptography.  

8. **Built-in Testing**  
   - Provides an in-built testing framework (`testing` package) for unit testing and benchmarking.  

9. **Efficient Networking & Web Development**  
   - Optimized for building high-performance web services and APIs.  
   - Used in popular frameworks like Gin and Fiber.  

10. **Security & Simplicity**  
   - Avoids issues like dangling pointers and memory leaks.  
   - Enforces good coding practices and prevents unsafe operations.  

---

### **Why Should We Learn GoLang?**  

1. **High Demand for Go Developers**  
   - Companies like Google, Uber, Dropbox, Netflix, and PayPal use Go.  
   - Growing job opportunities in backend development, cloud computing, and DevOps.  

2. **Ideal for Backend Development**  
   - Perfect for building RESTful APIs, microservices, and backend systems.  
   - Frameworks like Gin, Fiber, and Echo make web development easy.  

3. **Great for Cloud Computing & DevOps**  
   - Used in cloud-native technologies like Docker, Kubernetes, and Terraform.  
   - Many DevOps tools are built using Go.  

4. **Concurrency & Scalability**  
   - Handles thousands of concurrent tasks efficiently.  
   - Suitable for applications requiring real-time data processing (e.g., chat apps, streaming services).  

5. **Ease of Learning**  
   - Simple syntax with fewer concepts to master compared to Java or C++.  
   - Good starting point for learning system programming.  

6. **Performance & Efficiency**  
   - Almost as fast as C/C++, with better memory management.  
   - Lower CPU and RAM usage, making it ideal for large-scale applications.  

7. **Community & Ecosystem**  
   - Active open-source community.  
   - Strong ecosystem with numerous libraries and frameworks.  

---

### **Should You Learn GoLang?**  
Yes, especially if you're interested in:  
✅ **Backend development** (APIs, microservices)  
✅ **Cloud computing & DevOps**  
✅ **Building high-performance, scalable applications**  
✅ **Concurrency & distributed systems**  
✅ **Learning a simple yet powerful language with fast execution**  

Here’s a **detailed explanation** of all the Go commands:  

---

### **1. `go bug`**  
- Opens a web page to report a bug in the Go project.  
- It automatically collects system information to help with debugging.  

📝 **Usage:**  
```sh
go bug
```
---

### **2. `go build`**  
- Compiles the Go source code and generates an executable binary.  
- It compiles all dependencies but does not install them.  

📝 **Usage:**  
```sh
go build [package]
```
Example:  
```sh
go build main.go
```
- This compiles `main.go` into an executable file (`main.exe` on Windows, `main` on Linux/macOS).  

---

### **3. `go clean`**  
- Removes compiled object files and cached data from the Go build system.  
- Useful for cleaning up unnecessary files.  

📝 **Usage:**  
```sh
go clean
```
- Removes all build artifacts from the current module.  

```sh
go clean -cache -modcache -testcache
```
- Cleans the cache, module cache, and test cache.  

---

### **4. `go doc`**  
- Displays documentation for a Go package or function.  

📝 **Usage:**  
```sh
go doc fmt.Println
```
- Shows documentation for `fmt.Println`.  

```sh
go doc -all fmt
```
- Shows full documentation for the `fmt` package.  

---

### **5. `go env`**  
- Prints information about the Go environment variables (like `GOROOT`, `GOPATH`).  

📝 **Usage:**  
```sh
go env
```
- Shows all environment variables.  

```sh
go env GOPATH
```
- Prints only the `GOPATH` value.  

---

### **6. `go fix`**  
- Automatically updates Go code to use the latest API changes.  

📝 **Usage:**  
```sh
go fix ./...
```
- Fixes all Go files in the current directory and subdirectories.  

---

### **7. `go fmt`**  
- Formats Go source code to follow standard Go coding style.  

📝 **Usage:**  
```sh
go fmt main.go
```
- Formats `main.go`.  

```sh
go fmt ./...
```
- Formats all `.go` files in the current directory and subdirectories.  

---

### **8. `go generate`**  
- Runs code generation tools specified in Go files using `//go:generate` comments.  
- Often used for generating boilerplate code, mocks, or database models.  

📝 **Usage:**  
```sh
go generate
```
- Executes commands specified in `//go:generate` comments.  

Example in a Go file:  
```go
//go:generate go run mygenerator.go
```

---

### **9. `go get`**  
- Fetches and installs dependencies for a Go module.  

📝 **Usage:**  
```sh
go get github.com/gorilla/mux
```
- Installs the `gorilla/mux` package.  

```sh
go get -u
```
- Updates all dependencies in the project.  

---

### **10. `go install`**  
- Compiles and installs a package or program into `$GOBIN`.  
- Useful for installing CLI tools.  

📝 **Usage:**  
```sh
go install github.com/user/tool@latest
```
- Installs the latest version of `tool`.  

```sh
go install ./...
```
- Installs all executables from the current module.  

---

### **11. `go list`**  
- Lists installed Go packages or modules.  

📝 **Usage:**  
```sh
go list all
```
- Lists all available packages.  

```sh
go list -m
```
- Shows details about the current module.  

---

### **12. `go mod`**  
- Manages Go modules (Go’s dependency management system).  

📝 **Common Subcommands:**  
```sh
go mod init myproject
```
- Initializes a new module.  

```sh
go mod tidy
```
- Removes unused dependencies from `go.mod`.  

```sh
go mod verify
```
- Checks the integrity of dependencies.  

---

### **13. `go work`**  
- Manages workspaces containing multiple Go modules.  

📝 **Usage:**  
```sh
go work init
```
- Initializes a workspace file (`go.work`).  

```sh
go work use ./module1 ./module2
```
- Adds multiple modules to the workspace.  

---

### **14. `go run`**  
- Compiles and runs a Go program **without creating an executable file**.  

📝 **Usage:**  
```sh
go run main.go
```
- Runs `main.go` without generating an executable.  

```sh
go run .
```
- Runs all Go files in the current directory.  

---

### **15. `go telemetry`**  
- Manages telemetry data sent to Go developers.  

📝 **Usage:**  
```sh
go telemetry on
```
- Enables telemetry.  

```sh
go telemetry off
```
- Disables telemetry.  

---

### **16. `go test`**  
- Runs unit tests in the current package.  

📝 **Usage:**  
```sh
go test
```
- Runs all tests in the current directory.  

```sh
go test -v
```
- Runs tests in verbose mode.  

```sh
go test -cover
```
- Shows code coverage details.  

---

### **17. `go tool`**  
- Runs various Go tools like `pprof` (profiling), `vet` (error checking), etc.  

📝 **Usage:**  
```sh
go tool vet main.go
```
- Checks for possible mistakes in `main.go`.  

```sh
go tool pprof profile.out
```
- Analyzes a performance profile.  

---

### **18. `go version`**  
- Prints the installed Go version.  

📝 **Usage:**  
```sh
go version
```
- Example output:  
  ```
  go version go1.21.0 windows/amd64
  ```

---

### **19. `go vet`**  
- Detects potential bugs and incorrect code patterns.  

📝 **Usage:**  
```sh
go vet main.go
```
- Checks `main.go` for issues.  

```sh
go vet ./...
```
- Checks all files in the current project.  

---

### **Summary Table:**  

| Command | Description |
|---------|------------|
| `go bug` | Reports a Go bug. |
| `go build` | Compiles Go code into an executable. |
| `go clean` | Removes compiled files and cache. |
| `go doc` | Shows package documentation. |
| `go env` | Displays Go environment settings. |
| `go fix` | Updates Go code to new API versions. |
| `go fmt` | Formats Go source files. |
| `go generate` | Runs code generation tools. |
| `go get` | Fetches and installs dependencies. |
| `go install` | Installs executables in `$GOBIN`. |
| `go list` | Lists installed packages. |
| `go mod` | Manages Go modules. |
| `go work` | Manages workspaces. |
| `go run` | Compiles and runs a Go program. |
| `go telemetry` | Manages telemetry settings. |
| `go test` | Runs unit tests. |
| `go tool` | Runs additional Go tools. |
| `go version` | Prints the installed Go version. |
| `go vet` | Detects code mistakes. |

---
Great! Let's break down this **Go (Golang)** program **in great depth** so you understand every part of it. 🚀  

---

## **📌 Code Breakdown**
```go
package main
```
### **1️⃣ `package main` (Package Declaration)**
- Every Go program is organized into **packages**.
- The **`main` package** is **special**—it tells Go that this program is an **executable** (not a library).
- If a Go file is part of a **library**, it will have a different package name (e.g., `package mylib`).

---

```go
import "fmt"
```
### **2️⃣ `import "fmt"` (Import Statement)**
- `import` is used to include built-in **or external** packages.
- `"fmt"` (**format**) is a built-in **standard library** package for input/output operations.
- `"fmt"` provides functions to print text to the console, format output, and read input.

---

```go
func main() {
```
### **3️⃣ `func main()` (Main Function)**
- `func` **defines a function**.
- `main()` is **the entry point** of every Go program.
- When we run the program using `go run main.go`, **execution starts from `main()`**.

---

```go
    fmt.Println("Hello World!")
```
### **4️⃣ `fmt.Println("Hello World!")` (Printing to Console)**
- `fmt.Println()` prints output **with a newline (`\n`) at the end**.
- `"Hello World!"` is a **string literal** (a fixed sequence of characters).
- `Println()` automatically adds a **new line** at the end.
- Equivalent in JavaScript:
  ```js
  console.log("Hello World!");
  ```

---

## **📌 Complete Code Flow**
1. **Program starts execution** at `main()`.
2. **Imports** the `fmt` package.
3. Calls `fmt.Println("Hello World!")` to print the message.
4. The **program exits** after execution.

---

## **🔥 Experimenting with Modifications**
Try these changes to learn more:  

### **1️⃣ Without `fmt.Println`**
```go
package main

import "fmt"

func main() {
    // Nothing here!
}
```
💡 **Runs successfully but prints nothing!**  

---

### **2️⃣ Using `fmt.Print()` Instead of `fmt.Println()`**
```go
package main

import "fmt"

func main() {
    fmt.Print("Hello ")
    fmt.Print("World!")
}
```
💡 **Output:**  
```
Hello World!
```
👉 `fmt.Print()` does **not** add a newline (`\n`) automatically.

---

### **3️⃣ Printing Multiple Values**
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello", "World!", 2025)
}
```
💡 **Output:**  
```
Hello World! 2025
```
👉 `fmt.Println()` automatically adds spaces between arguments.

---

## **🎯 Summary**
✔️ **`package main`** → Defines an **executable** Go program.  
✔️ **`import "fmt"`** → Imports the **formatting library** for input/output.  
✔️ **`func main()`** → The **entry point** of execution.  
✔️ **`fmt.Println()`** → Prints text to the console **with a newline**.  

## **📌 Packages and Modules in GoLang (Detailed Explanation)**  

Go (Golang) is designed for **scalability and modularity**, and it organizes code into **packages** and **modules** to make projects more manageable and reusable.

---

# **🔹 1. Packages in GoLang**
### **📌 What is a Package?**
- A **package** is a collection of **Go source files** that are compiled together.
- Every Go file **must belong** to a package.
- A package is defined using the `package` keyword at the beginning of a Go file.

### **📌 Types of Packages**
There are **two types** of packages in Go:

1. **Executable Packages** (Standalone programs)
2. **Library Packages** (Reusable code)

---

## **📌 1.1 Executable Packages (`package main`)**
- Every Go program **must have a `main` package**.
- The **`main` package** tells the Go compiler to create an executable program.
- The `main()` function inside `package main` is the entry point for execution.

### **Example:**
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```
✔️ **Explanation:**  
- `package main` → Defines an **executable package**.
- `import "fmt"` → Imports the standard `fmt` package.
- `func main()` → The **entry point** of the program.

✅ **Command to Run the Program:**
```sh
go run main.go
```

---

## **📌 1.2 Library Packages (Reusable Code)**
- These packages **do not contain** a `main()` function.
- They are designed to be **imported and used by other packages**.
- Example: The standard Go **`fmt`**, `math`, and `strings` packages.

### **Creating a Custom Library Package**
Let's create a custom **math** package.

### **Step 1: Create a Folder**
Inside your project directory:
```
/project
  /mathutil
    math.go
  main.go
```

### **Step 2: Create `mathutil/math.go`**
```go
package mathutil

// Add adds two numbers
func Add(a int, b int) int {
    return a + b
}
```
✔️ **Explanation:**  
- `package mathutil` → Defines a reusable package named `mathutil`.
- `Add()` function is exported because it **starts with an uppercase letter**.

---

### **Step 3: Import and Use the Package in `main.go`**
```go
package main

import (
    "fmt"
    "project/mathutil"  // Importing our custom package
)

func main() {
    result := mathutil.Add(5, 7)
    fmt.Println("Result:", result)
}
```

✅ **Command to Run the Program:**
```sh
go run main.go
```
✔️ **Output:**
```
Result: 12
```

---

# **🔹 2. Modules in GoLang**
### **📌 What is a Module?**
- A **module** is a collection of related **Go packages**.
- It defines a **versioned unit of reusable code**.
- Modules are managed using the `go.mod` file.

---

## **📌 2.1 Creating a Go Module**
### **Step 1: Initialize a Module**
Run the following command in the project directory:
```sh
go mod init project
```
✔️ This creates a `go.mod` file with:
```
module project

go 1.19
```

### **Step 2: Verify the Module**
To check the dependencies:
```sh
go list -m all
```

---

## **📌 2.2 Understanding `go.mod` and `go.sum`**
- **`go.mod`** → Tracks **module name** and dependencies.
- **`go.sum`** → Stores **checksums** of dependencies.

---

# **🔹 Difference Between Packages and Modules**
| Feature     | Package | Module |
|------------|---------|--------|
| Definition | Collection of Go source files | Collection of related packages |
| Scope | Local to a project | Can be shared and versioned |
| Example | `fmt`, `mathutil` | `github.com/user/mymodule` |
| Managed By | `import` statement | `go.mod` file |

---

## **🎯 Summary**
✔️ **Packages** are used to **organize code** into reusable units.  
✔️ **Modules** are used to **manage dependencies** and package versioning.  
✔️ **`go.mod`** is required for module-based dependency management.  
✔️ Go **automatically fetches dependencies** from online repositories when using modules.  
---

## **📌 Data Types in GoLang (Detailed Explanation)**  

Go is a statically typed language, meaning variables must have a **specific type** at compile time. Go provides various **primitive** and **composite** data types.

---

# **🔹 1. Basic (Primitive) Data Types in GoLang**
These are the fundamental data types in Go:

| Type       | Description                          | Example |
|------------|--------------------------------------|---------|
| `bool`     | Boolean (true/false)               | `true`, `false` |
| `string`   | Sequence of characters             | `"Hello, Go!"` |
| `int`      | Signed integer (default `int64`)   | `10`, `-42` |
| `float64`  | Floating-point numbers             | `3.14`, `-2.71` |
| `complex`  | Complex numbers                    | `2 + 3i` |
| `byte`     | Alias for `uint8`                   | `var ch byte = 'A'` |
| `rune`     | Alias for `int32` (Unicode char)   | `var ch rune = '✓'` |

---

## **📌 1.1 Boolean (`bool`)**
A **boolean** represents **true** or **false**.

### **Example:**
```go
package main

import "fmt"

func main() {
    var isGoFun bool = true
    fmt.Println("Is Go fun?", isGoFun)
}
```
✔️ **Output:**  
```
Is Go fun? true
```

---

## **📌 1.2 String (`string`)**
A **string** is a sequence of **UTF-8** characters.

### **Example:**
```go
package main

import "fmt"

func main() {
    var message string = "Hello, Go!"
    fmt.Println(message)
}
```
✔️ **Output:**  
```
Hello, Go!
```

✔️ **Multi-line String:**
```go
str := `This is
a multi-line
string`
```

---

## **📌 1.3 Integer (`int`, `uint`)**
Go supports various integer sizes.

### **Signed Integers (Can store negative values)**
| Type   | Size  | Range |
|--------|------|-----------------------------|
| `int8` | 8-bit  | -128 to 127 |
| `int16` | 16-bit | -32,768 to 32,767 |
| `int32` | 32-bit | -2,147,483,648 to 2,147,483,647 |
| `int64` | 64-bit | -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807 |
| `int`   | Depends on OS (32-bit or 64-bit) |

### **Unsigned Integers (Only positive values)**
| Type   | Size  | Range |
|--------|------|-----------------------------|
| `uint8` | 8-bit  | 0 to 255 |
| `uint16` | 16-bit | 0 to 65,535 |
| `uint32` | 32-bit | 0 to 4,294,967,295 |
| `uint64` | 64-bit | 0 to 18,446,744,073,709,551,615 |
| `uint`   | Depends on OS (32-bit or 64-bit) |

✔️ **Example:**
```go
package main

import "fmt"

func main() {
    var age int = 29
    var distance uint = 500

    fmt.Println("Age:", age)
    fmt.Println("Distance:", distance)
}
```
✔️ **Output:**
```
Age: 29
Distance: 500
```

---

## **📌 1.4 Floating Point (`float32`, `float64`)**
| Type     | Size  | Precision |
|----------|------|-----------|
| `float32` | 32-bit | 6-7 decimal places |
| `float64` | 64-bit | 15-16 decimal places |

✔️ **Example:**
```go
package main

import "fmt"

func main() {
    var pi float64 = 3.14159
    fmt.Println("Pi:", pi)
}
```
✔️ **Output:**
```
Pi: 3.14159
```

---

## **📌 1.5 Complex Numbers (`complex64`, `complex128`)**
Go supports **complex numbers** with real and imaginary parts.

✔️ **Example:**
```go
package main

import "fmt"

func main() {
    var num complex64 = 2 + 3i
    fmt.Println("Complex Number:", num)
}
```
✔️ **Output:**
```
Complex Number: (2+3i)
```

---

## **📌 1.6 Byte and Rune**
- **`byte`** is an alias for **`uint8`** (used for ASCII characters).
- **`rune`** is an alias for **`int32`** (used for Unicode characters).

✔️ **Example:**
```go
package main

import "fmt"

func main() {
    var ch byte = 'A'
    var symbol rune = '✓'

    fmt.Println("Byte:", ch)       // Output: 65 (ASCII code of 'A')
    fmt.Println("Rune:", symbol)    // Output: 10003 (Unicode code point)
}
```

---

# **🔹 2. Composite Data Types in GoLang**
Composite types **group multiple values** together.

| Type        | Description |
|------------|-------------|
| `array`    | Fixed-size collection of elements |
| `slice`    | Dynamic-size collection of elements |
| `map`      | Key-value pairs (like dictionaries) |
| `struct`   | Custom data type with multiple fields |
| `pointer`  | Stores memory address of a variable |
| `interface`| Defines behavior without implementation |

---

## **📌 2.1 Arrays**
- **Fixed-size** collection of the **same data type**.

✔️ **Example:**
```go
package main

import "fmt"

func main() {
    var arr [3]int = [3]int{10, 20, 30}
    fmt.Println("Array:", arr)
}
```
✔️ **Output:**
```
Array: [10 20 30]
```

---

## **📌 2.2 Slices**
- **Dynamic size** (can grow and shrink).
- More commonly used than arrays.

✔️ **Example:**
```go
package main

import "fmt"

func main() {
    nums := []int{1, 2, 3, 4}
    fmt.Println("Slice:", nums)
}
```
✔️ **Output:**
```
Slice: [1 2 3 4]
```

---

## **📌 2.3 Maps**
- **Key-value pairs** (like Python dictionaries).

✔️ **Example:**
```go
package main

import "fmt"

func main() {
    capitals := map[string]string{
        "India": "New Delhi",
        "USA":   "Washington D.C.",
    }
    fmt.Println(capitals)
}
```
✔️ **Output:**
```
map[India:New Delhi USA:Washington D.C.]
```

---

## **📌 2.4 Structs**
- Used to define **custom data types**.

✔️ **Example:**
```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    person := Person{Name: "Alice", Age: 25}
    fmt.Println(person)
}
```
✔️ **Output:**
```
{Name: Alice Age: 25}
```

---

# **🎯 Summary**
✔️ **Primitive Types**: `int`, `float`, `bool`, `string`, `byte`, `rune`.  
✔️ **Composite Types**: `array`, `slice`, `map`, `struct`.  
✔️ **Pointers & Interfaces**: Advanced concepts in Go.  

---
# **📌 Type Conversion in GoLang (With Examples)**  

Go is a **statically typed** language, meaning **explicit type conversion** is required when changing from one type to another. Unlike dynamically typed languages like JavaScript, Go **does not allow implicit type conversions** between different types.

---

# **🔹 1. Convert Between Integer Types**
In Go, you must **explicitly** convert between integer types.

✔️ **Example: Converting `int` to `int32`, `int64`, etc.**  
```go
package main

import "fmt"

func main() {
    var num int = 100
    var num32 int32 = int32(num) // Convert int to int32
    var num64 int64 = int64(num) // Convert int to int64

    fmt.Println("Original (int):", num)
    fmt.Println("Converted to int32:", num32)
    fmt.Println("Converted to int64:", num64)
}
```
✔️ **Output:**
```
Original (int): 100
Converted to int32: 100
Converted to int64: 100
```

---

# **🔹 2. Convert Between Float and Integer**
- Converting **float to int** **truncates** the decimal part (rounding is not automatic).  

✔️ **Example:**
```go
package main

import "fmt"

func main() {
    var floatNum float64 = 9.8
    var intNum int = int(floatNum) // Truncate decimal part

    fmt.Println("Original (float64):", floatNum)
    fmt.Println("Converted to int:", intNum)
}
```
✔️ **Output:**
```
Original (float64): 9.8
Converted to int: 9
```

⚠ **Note:** No automatic rounding. `9.8 → 9` (truncated).

If we want **rounding**, we must use `math.Round()`:
```go
import "math"
var rounded int = int(math.Round(floatNum))
```

---

# **🔹 3. Convert String to Integer**
Use **`strconv.Atoi()`** to convert a string to an integer.

✔️ **Example:**
```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    str := "123"
    num, err := strconv.Atoi(str) // Convert string to int

    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Converted to int:", num)
    }
}
```
✔️ **Output:**
```
Converted to int: 123
```

🔹 **Alternative:** Use `strconv.ParseInt()` for `int64`:
```go
num64, err := strconv.ParseInt(str, 10, 64) // Convert to int64
```

---

# **🔹 4. Convert Integer to String**
Use **`strconv.Itoa()`** to convert an integer to a string.

✔️ **Example:**
```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    num := 456
    str := strconv.Itoa(num) // Convert int to string

    fmt.Println("Converted to string:", str)
}
```
✔️ **Output:**
```
Converted to string: 456
```

🔹 **Alternative:** Use `fmt.Sprintf()`
```go
str := fmt.Sprintf("%d", num)
```

---

# **🔹 5. Convert String to Float**
Use **`strconv.ParseFloat()`**.

✔️ **Example:**
```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    str := "3.1415"
    floatNum, err := strconv.ParseFloat(str, 64) // Convert string to float64

    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Converted to float64:", floatNum)
    }
}
```
✔️ **Output:**
```
Converted to float64: 3.1415
```

---

# **🔹 6. Convert Float to String**
Use **`strconv.FormatFloat()`**.

✔️ **Example:**
```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    pi := 3.1415
    str := strconv.FormatFloat(pi, 'f', 2, 64) // Keep 2 decimal places

    fmt.Println("Converted to string:", str)
}
```
✔️ **Output:**
```
Converted to string: 3.14
```

🔹 **Alternative:** Use `fmt.Sprintf()`
```go
str := fmt.Sprintf("%.2f", pi)
```

---

# **🔹 7. Convert Boolean to String**
Use **`strconv.FormatBool()`**.

✔️ **Example:**
```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    flag := true
    str := strconv.FormatBool(flag) // Convert bool to string

    fmt.Println("Converted to string:", str)
}
```
✔️ **Output:**
```
Converted to string: true
```

---

# **🔹 8. Convert String to Boolean**
Use **`strconv.ParseBool()`**.

✔️ **Example:**
```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    str := "true"
    flag, err := strconv.ParseBool(str) // Convert string to bool

    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Converted to bool:", flag)
    }
}
```
✔️ **Output:**
```
Converted to bool: true
```

✅ Acceptable string values for **true:** `"1"`, `"t"`, `"T"`, `"true"`, `"TRUE"`, `"True"`  
✅ Acceptable string values for **false:** `"0"`, `"f"`, `"F"`, `"false"`, `"FALSE"`, `"False"`

---

# **🔹 9. Convert Byte to String**
Use **`string()`** or `[]byte()`.

✔️ **Example:**
```go
package main

import "fmt"

func main() {
    bytes := []byte{72, 101, 108, 108, 111} // ASCII for "Hello"
    str := string(bytes) // Convert byte slice to string

    fmt.Println("Converted to string:", str)
}
```
✔️ **Output:**
```
Converted to string: Hello
```

---

# **🔹 10. Convert String to Byte**
Use **`[]byte()`**.

✔️ **Example:**
```go
package main

import "fmt"

func main() {
    str := "Hello"
    bytes := []byte(str) // Convert string to byte slice

    fmt.Println("Converted to byte slice:", bytes)
}
```
✔️ **Output:**
```
Converted to byte slice: [72 101 108 108 111]
```

---

# **🎯 Summary of Type Conversions in GoLang**
| **Conversion** | **Function** |
|--------------|----------------|
| **Int → Float** | `float64(intVar)` |
| **Float → Int** | `int(floatVar)` (Truncates) |
| **String → Int** | `strconv.Atoi(str)` |
| **Int → String** | `strconv.Itoa(intVar)` |
| **String → Float** | `strconv.ParseFloat(str, 64)` |
| **Float → String** | `strconv.FormatFloat(floatVar, 'f', 2, 64)` |
| **String → Bool** | `strconv.ParseBool(str)` |
| **Bool → String** | `strconv.FormatBool(boolVar)` |
| **Byte → String** | `string([]byteVar)` |
| **String → Byte** | `[]byte(stringVar)` |

---