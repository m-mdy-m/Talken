# Contributing to Talken

First of all, **thank you** for taking the time to contribute to Talken! This project exists because of amazing contributors like you.

## 🛠️ How to Contribute

### 1. Fork & Clone

```bash
git clone https://github.com/your-username/Talken.git
cd Talken
```

### 2. Install Go and Dependencies

Make sure you have **Go 1.22+** installed. Run:

```bash
go mod tidy
go build ./...
```

### 3. Create a Feature Branch

```bash
git checkout -b feature/your-amazing-idea
```

### 4. Write Your Code

- Follow [Effective Go](https://golang.org/doc/effective_go.html) practices
- Format with `go fmt ./...`
- Use `go vet ./...`
- Write unit tests in `_test.go` files
- Run all tests with `go test ./...`

### 5. Commit with Meaningful Messages

```bash
git commit -m "feat(core): add identity struct and ID generation"
```

### 6. Push and Create a Pull Request

```bash
git push origin feature/your-amazing-idea
```

Then open a pull request to the `main` branch. Fill out the PR template.

## Communication

If you have questions or ideas, open an [Issue](https://github.com/m-mdy-m/Talken/issues) or start a [Discussion](https://github.com/m-mdy-m/Talken/discussions).

## Security

Please refer to [SECURITY.md](./SECURITY.md) to report vulnerabilities responsibly.

Thank you for being part of Talken 🚀
