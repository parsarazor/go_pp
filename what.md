# What I Learned Today 🎓

## 1. Zed Tasks System 📋

### File Requirements
- Tasks must be named `.zed/tasks.json` (NOT `.jsonc`)
- Zed's JSON parser accepts JSONC syntax (comments with `//`)
- **No trailing commas allowed** - this was the key issue!

### Task Structure
```json
[
  {
    "label": "Task Name",
    "command": "command to run",
    "cwd": "${ZED_WORKTREE_ROOT}",
    "tags": ["category"]
  }
]
```

### Using Tasks
- Open with `cmd-shift-p` → `task: spawn` (or `opt-shift-t`)
- Rerun last task: `task: rerun` (or `opt-t`)
- Tasks are filtered by available variables

---

## 2. Unix Philosophy & Standard Streams 🔀

### The Power of stdout
Instead of complicated HTTP servers and curl commands, the Unix way is simpler:

```bash
go run main.go > output.gif
```

### How It Works
```go
func main() {
    lissajous(os.Stdout)  // Write to standard output
}
```

- `os.Stdout` = standard output stream
- Shell's `>` operator redirects it to a file
- Without `>`, binary data appears as garbage in terminal

### Why This is Better
✅ **Flexible**: Users decide where output goes
```bash
./program > file.gif           # Save to file
./program | display -          # Pipe to image viewer
./program > /dev/null          # Discard output
```

✅ **Simple**: No need for file handling in code
✅ **Composable**: Works with other Unix tools

---

## 3. Go Binary Compilation 🔨

### What `go build` Does
1. Compiles Go source code to machine code
2. Links all dependencies
3. Creates a **standalone executable**
4. No Go runtime needed to run it

### Example
```bash
go build -o bin/lissajous main.go
```

**Result:**
- Input: `main.go` (~4KB source code)
- Output: `bin/lissajous` (~2.4MB executable)
- Format: Mach-O 64-bit executable arm64 (for Apple Silicon)

### Benefits
- ⚡ **Fast startup** - no compilation at runtime
- 📦 **Portable** - copy to another machine and it works
- 🚀 **No dependencies** - users don't need Go installed
- 💨 **Native performance** - runs at full CPU speed

### Running the Binary
```bash
# Instead of:
go run main.go > output.gif

# You can:
./bin/lissajous > output.gif
```

---

## 4. Cross-Compilation 🌍

Go can build binaries for different platforms from a single machine:

```bash
# For Linux
GOOS=linux GOARCH=amd64 go build -o bin/lissajous-linux main.go

# For Windows
GOOS=windows GOARCH=amd64 go build -o bin/lissajous.exe main.go

# For macOS Intel
GOOS=darwin GOARCH=amd64 go build -o bin/lissajous-mac-intel main.go

# For macOS ARM (M1/M2/M3)
GOOS=darwin GOARCH=arm64 go build -o bin/lissajous-mac-arm main.go
```

This means you can distribute your program to users on any platform! 🎉

---

## 5. Code Organization Best Practices 📁

### Commenting vs Deleting
Instead of deleting old code, **comment it out** with explanations:

```go
// Uncomment below to run as HTTP server:
// handler := func(w http.ResponseWriter, r *http.Request){
//     lissajous(w)
// }
// http.HandleFunc("/", handler)
// log.Fatal(http.ListenAndServe("localhost:8000", nil))

// Output GIF to stdout (use with: go run main.go > output.gif)
lissajous(os.Stdout)
```

**Why?**
- Easy to switch between modes
- Preserves working code
- Documents alternative approaches
- Helps future you remember options

---

## Key Takeaways 💡

1. **Keep it simple** - `go run main.go > output.gif` beats complex HTTP + curl
2. **Write to stdout** - Let the shell handle redirection
3. **Binaries are powerful** - Distribute standalone executables
4. **Comment, don't delete** - Preserve alternative implementations
5. **Unix philosophy** - Do one thing well, compose with others

---

## Project Structure

```
go_pp/
├── .zed/
│   └── tasks.json          # Zed task definitions
├── shape/
│   └── main.go             # Shape animations server
├── main.go                 # Lissajous GIF generator
├── go.mod                  # Go module file
└── *.gif                   # Generated animations
```

---

## Useful Commands

```bash
# Generate GIF
go run main.go > output.gif

# Build binary
go build -o bin/lissajous main.go

# Run binary
./bin/lissajous > output.gif

# Format and vet code
go fmt ./... && go vet ./...

# Open GIF (macOS)
open output.gif
```

---

*Last updated: 2024*