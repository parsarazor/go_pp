# Standard Streams & fmt Package

## Deep Dive in Go

---

# Agenda

1.  **Standard Streams** (Stdin, Stdout, Stderr)
2.  **The `fmt` Package**
3.  **Advanced Interfaces** (Stringer, Formatter)
4.  **Best Practices**

---

# 1. Standard Streams

## The Unix Philosophy

- **Stdin (0):** Standard Input (Keyboard)
- **Stdout (1):** Standard Output (Terminal)
- **Stderr (2):** Standard Error (Logs/Errors)

---

# Standard Streams in Go

## `os` Package

Type: `*os.File` (Implements `io.Reader` & `io.Writer`)

```go
package main
import "os"

func main() {
    // Writing to Stdout
    os.Stdout.WriteString("Hello Stdout\n")

    // Writing to Stderr
    os.Stderr.WriteString("Error Message\n")
}
```

---

# 2. The `fmt` Package

## Formatted I/O

Three main categories:

1.  **Print/Scan:** `Stdout` / `Stdin`
    - `Println`, `Printf`, `Scanln`
2.  **Fprint/Fscan:** `io.Writer` / `io.Reader`
    - `Fprintf`, `Fscan`
3.  **Sprint/Sscan:** `string`
    - `Sprintf`, `Sscan`

---

# Formatting Verbs (1/2)

## General & Integers

| Verb  | Description    | Example              |
| :---- | :------------- | :------------------- |
| `%v`  | Default format | `Hello`              |
| `%+v` | Struct fields  | `{Name:Ali}`         |
| `%#v` | Go Syntax      | `main.P{Name:"Ali"}` |
| `%T`  | Type           | `main.Person`        |
| `%d`  | Base 10        | `123`                |
| `%x`  | Hex            | `7b`                 |

---

# Formatting Verbs (2/2)

## Strings & Others

- **String:** `%s`
- **Quoted String:** `%q` (Escapes control chars)
- **Boolean:** `%t`
- **Pointer:** `%p` (Hex address)
- **Float:** `%.2f` (Precision control)

---

# 3. Advanced Interfaces

## Customizing Output

Control how your types are printed.

1.  **`Stringer`**: For `%v` and `%s`
2.  **`GoStringer`**: For `%#v` (Debug)
3.  **`Formatter`**: Full control over all verbs

---

# `Stringer` Example

```go
type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%s (%d)", p.Name, p.Age)
}

// fmt.Println(p) -> "Ali (20)"
```

---

# 4. Best Practices

- **Use `log` for Errors:**
  - Prefer `log.Println` over `fmt.Println` for system errors (writes to Stderr + Timestamp).
- **Buffering:**
  - `fmt` is unbuffered. For high-volume I/O, use `bufio.Writer`.
- **Input Safety:**
  - Always check errors from `Scan` functions.

---

# Thank You

## Q&A

- Go Documentation: `pkg.go.dev/fmt`
