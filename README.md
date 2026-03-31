# MiniGo Compiler

## Description

This repository implements a small compiler/transformer for a MiniGo subset. It uses ANTLR for parsing, a basic semantic checker in the `checker` package, and generates LLVM IR using `llvm` from the `encoder` package.

The application exposes a simple web UI at `http://localhost:8080` where you can submit MiniGo source; if parsing and checking succeed, the project writes `module.ll` and attempts to compile it with `clang`.

## Project structure

- `main.go` — HTTP server and main compile pipeline.
- `MiniGoParser.g4`, `MiniGoScanner.g4` — ANTLR grammar sources.
- `generated/` — ANTLR-generated parser/lexer/visitor code.
- `checker/` — semantic analysis and symbol table.
- `encoder/` — LLVM IR generation using `github.com/llir/llvm`.
- `GUI/Index.html` — simple web interface used by the server.
- `resources/` — example test inputs.

## Requirements

- Go 1.22 or newer.
- `clang` (required only if you want the server to compile `module.ll` into an executable). Make sure `clang` is available in your `PATH`.

## Dependencies

Dependencies are declared in `go.mod`. Prepare the environment with:

```bash
go mod tidy
```

## Development / Usage

1. Run the server:

```bash
go run main.go
```

2. Open your browser at `http://localhost:8080`, paste MiniGo code into the form and submit.

3. On successful compilation the server writes `module.ll`, runs `clang module.ll -o module.exe` and (optionally) executes the generated binary. Note: executing generated binaries from a server can be a security risk — use only in trusted environments.

## Important notes

- `main.go` currently invokes `clang` to build `module.ll`. If you don't want compilation/execution, modify the server to only write the IR file.
- Many visitor methods in `checker` and `encoder` delegate to `VisitChildren`; those areas may need further implementation to support the full language.

## Recommended next steps

- Fix the `clang` invocation in `main.go` (there is an empty argument that should be removed).
- Add safety checks before executing any generated binary.
- Add unit tests for `checker` and `encoder` using inputs from `resources/`.

---

Author: Ovidio Martinez Taleno
