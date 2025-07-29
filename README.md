# Go Learning Repository

Welcome to my Go programming learning journey! This repository contains exercises, practice projects, and notes as I learn the Go programming language.

## About This Repository

This repo is my personal collection of Go code written while learning the language fundamentals. It includes basic exercises, experiments, and small projects to help solidify my understanding of Go concepts.

## What's Inside

### 📝 Basic Exercises
- **FizzBuzz** - Classic programming exercise
- **Palindrome Checker** - String manipulation and Unicode handling
- **Factorial Calculator** - Recursion and mathematical operations
- **Prime Number Checker** - Algorithm implementation
- **String Operations** - Working with strings, runes, and Unicode
- **Slice Manipulation** - Arrays, slices, and data structures

### 🔧 Key Concepts Covered
- Variables and data types (`int`, `string`, `bool`, slices, maps)
- Control structures (`for`, `if/else`, `switch`)
- Functions with parameters and return values
- String manipulation and the `strings` package
- Working with runes and Unicode characters
- Error handling basics
- Go idioms and best practices

## Getting Started

### Prerequisites
- Go 1.19 or later installed on your system
- Basic understanding of programming concepts

### Running the Code

1. Clone this repository:
```bash
git clone <your-repo-url>
cd go-learning
```

2. Run individual Go files:
```bash
go run filename.go
```

3. Or run with go build:
```bash
go build filename.go
./filename
```

## Learning Resources

Here are some resources I'm using to learn Go:

- [Official Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [A Tour of Go](https://tour.golang.org/)
- [Effective Go](https://golang.org/doc/effective_go.html)

## Progress Tracking

### ✅ Completed Topics
- [x] Basic syntax and variables
- [x] Control structures (loops, conditionals)
- [x] Functions and parameters
- [x] Strings and runes
- [x] Slices and arrays
- [x] Maps
- [x] Range loops and the blank identifier

### 🚧 Currently Learning
- [ ] Structs and methods
- [ ] Interfaces
- [ ] Error handling patterns
- [ ] Goroutines and concurrency
- [ ] Packages and modules

### 📚 Future Topics
- [ ] HTTP servers and web development
- [ ] Database connectivity
- [ ] Testing in Go
- [ ] Advanced concurrency patterns
- [ ] Performance optimization

## Notes and Observations

### Key Go Concepts I've Learned
- **The blank identifier (`_`)** - Used to ignore values you don't need
- **Runes vs bytes** - Important for proper Unicode string handling
- **Multiple return values** - Functions can return multiple values
- **Zero values** - Every type has a default zero value
- **Explicit is better** - Go favors explicit code over implicit

### Common Gotchas
- `len(string)` returns bytes, not characters - use `[]rune(string)` for character count
- Unused variables cause compilation errors - use `_` to ignore them
- Go is statically typed - no implicit type conversions

## File Structure

```
├── README.md
├── exercises/
│   ├── basic/
│   │   ├── fizzbuzz.go
│   │   ├── palindrome.go
│   │   └── factorial.go
│   └── intermediate/
│       └── (future exercises)
├── projects/
│   └── (small projects)
└── notes/
    └── concepts.md
```

## Contributing

This is a personal learning repository, but if you spot any issues or have suggestions for improvements, feel free to open an issue or submit a pull request!

## License

This project is open source and available under the [MIT License](LICENSE).

---

**Happy Coding!** 🐹

*Last updated: July 29, 2025*
