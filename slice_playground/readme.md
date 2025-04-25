

# 🧩 Slice Playground (Go CLI Tool)

A simple command-line playground to practice slice operations in Go.  
This tool is perfect for beginners who want to understand how slices work under the hood.

---

## ✨ Features

- ✅ Append elements to a slice
- ✅ Remove elements by index
- ✅ Reverse the entire slice
- ✅ Search for a specific element
- ✅ View current state of the slice

---

## 🛠 Getting Started

### Prerequisites

- Go installed (version 1.16+ recommended)

### Run the app

```bash
go run main.go
```

---

## 📸 Sample Output

```
==== Slice Playground ====
Current Slice: []
1. Append element
2. Remove element by index
3. Reverse slice
4. Search for an element
5. Exit
Choose an option:
```

---

## 📚 Concepts Practiced

- Slice operations (`append`, slicing, index access)
- Buffered input using `bufio.NewReader`
- Basic CLI menus
- String and integer conversion
- Error handling

---

## 🤔 Why use `bufio`?

We use `bufio.NewReader(os.Stdin)` to read full lines of input (including spaces) from the user. It more flexible than `fmt.Scan` and handles user input in a cleaner way.

---

## 📂 Project Structure

```
slice-playground/
│
├── main.go         # Entry point with menu and slice logic
├── README.md       # You're reading it!
```

---

## 🧠 Next Steps (Optional Ideas)

- Add save/load functionality using files
- Add support for other slice types (e.g. strings)
- Build a web version using Gin or Fiber
- Create unit tests for each function

---

