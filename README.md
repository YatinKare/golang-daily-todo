# Golang Daily To-Do CLI

A cli go daily to-do manager in ~100 lines. Features include:

- Creating Daily Notes with a template to a directory
- Migrating Un-checked to-do's from the previous day into today's daily note upon creation

Motivation: Learning Golang syntax

## Usage

```bash
go build to-do.go
./to-do --version
./to-do --help
./to-do new
```

## Learnings

I learned...
- How to capture command line arguments with `os.Args`
- How to do switch statements
- The different ways of printing (`fmt.Println, fmt.Printf)
- How to format a string with `fmt.Sprintf()`
- How to declare a function that returns a value 
- How to check for errors
- How to use the `time` package to use dates
- How to handle file paths with the `path/filepath` package.
- How to read a file with `os.ReadFile`
- How to convert bytes to strings
- how to use the `strings` package to split strings, to check if a string contains a substring
- How to initialize a slice (not an array)
- How to itterate over a slice using the `_, val := range {slice}` syntax
- How to append to a slice
- How to write buffers
- How to open a file, use file descriptors, check if files exists and write a buffer to a file.
- How to handle errors gracefully

---
> Author: Yatin Kare
