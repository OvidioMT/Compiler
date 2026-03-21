# MiniGo Compiler (Proyecto)

## Descripción

Este repositorio contiene un compilador/transformador para un subconjunto de Go (MiniGo). Usa ANTLR para parsear la gramática, un análisis semántico básico en `checker` y genera un módulo LLVM IR con `llir/llvm` desde `encoder`.

La aplicación expone una interfaz web simple en `localhost:8080` que recibe código MiniGo, lo parsea y (si no hay errores) genera y compila un módulo LLVM usando `clang`.

## Estructura del proyecto

- `main.go` — servidor HTTP y flujo principal de compilación.
- `MiniGoParser.g4`, `MiniGoScanner.g4` — gramática ANTLR (fuente).
- `generated/` — código generado por ANTLR (parser/lexer/visitor).
- `checker/` — análisis semántico y tabla de símbolos.
- `encoder/` — generación de IR (usa `github.com/llir/llvm`).
- `GUI/Index.html` — interfaz web usada por el servidor.
- `resources/` — archivos de prueba.

## Requisitos

- Go 1.22+ (el proyecto usa `go.mod`).
- `clang` (para compilar `module.ll` a ejecutable). Asegúrate de tener `clang` disponible en `PATH` si quieres generar y ejecutar el binario.

## Dependencias

Las dependencias están en `go.mod`. Ejecuta:

```bash
go mod tidy
```

## Uso (desarrollo)

1. Instala dependencias:

```bash
go mod tidy
```

2. Arranca el servidor web:

```bash
go run main.go
```

3. Abre en el navegador: `http://localhost:8080` y pega código MiniGo en el formulario.

4. Si la compilación es exitosa, el servidor genera `module.ll`, intenta compilarlo con `clang` y ejecutar el binario resultante (nota: ejecutar binarios desde un servidor puede ser peligroso; usar sólo en entornos de confianza).

# MiniGo Compiler

## Description

This repository implements a small compiler/transformer for a MiniGo subset. It uses ANTLR for parsing, a basic semantic checker in the `checker` package, and generates LLVM IR using `llir/llvm` from the `encoder` package.

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

1. Install dependencies:

```bash
go mod tidy
```

2. Run the server:

```bash
go run main.go
```

3. Open your browser at `http://localhost:8080`, paste MiniGo code into the form and submit.

4. On successful compilation the server writes `module.ll`, runs `clang module.ll -o module.exe` and (optionally) executes the generated binary. Note: executing generated binaries from a server can be a security risk — use only in trusted environments.

## Important notes

- `main.go` currently invokes `clang` to build `module.ll`. If you don't want compilation/execution, modify the server to only write the IR file.
- `readmes.md` contains an example README from another project and does not describe this compiler.
- Many visitor methods in `checker` and `encoder` delegate to `VisitChildren`; those areas may need further implementation to support the full language.

## Recommended next steps

- Fix the `clang` invocation in `main.go` (there is an empty argument that should be removed).
- Add safety checks before executing any generated binary.
- Add unit tests for `checker` and `encoder` using inputs from `resources/`.

---

Author: Ovidio Martinez Taleno
