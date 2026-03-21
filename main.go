package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/llir/llvm/ir"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"proyecto3/checker"
	"proyecto3/encoder"
	"proyecto3/generated"
)

var errorMessage string

// To capture errors, implement the ANTLR ErrorListener interface.
type customErrorListener struct {
	*antlr.DiagnosticErrorListener
	errors []string
}

// NewCustomErrorListener creates a new instance of CustomErrorListener
func NewCustomErrorListener() *customErrorListener {
	return &customErrorListener{
		DiagnosticErrorListener: antlr.NewDiagnosticErrorListener(true),
	}
}

/*
SyntaxError is a method that implements the ErrorListener interface of ANTLR.
It is called when a syntax error is encountered during parsing.
*/
func (c *customErrorListener) SyntaxError(_ antlr.Recognizer,
	_ interface{}, line, column int, msg string,
	_ antlr.RecognitionException) {

	c.errors = append(c.errors, fmt.Sprintf("line %d:%d %s", line, column, msg))
}

func runModule(module *ir.Module) {
	// Escribir el módulo LLVM en un archivo
	f, err := os.Create("module.ll")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer f.Close()
	if _, err := module.WriteTo(f); err != nil {
		fmt.Println("Error al escribir el módulo:", err)
		return
	}

	// Compilar el módulo LLVM a código objeto
	cmd := exec.Command("clang", "", "module.ll", "-o", "module.exe")
	if err := cmd.Run(); err != nil {
		fmt.Println("Error al compilar el módulo:", err)
		return
	}

	fmt.Println("El archivo ejecutable .exe ha sido generado correctamente.")

	//ejecute el programa
	cmd = exec.Command("module.exe")
	var out bytes.Buffer
	cmd.Stdout = &out
	if errors.Is(cmd.Err, exec.ErrDot) {
		cmd.Err = nil
	}

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
		fmt.Println("Error al ejecutar el comando:", err)
		return
	}

	errorMessage = out.String()
	fmt.Println("Salida de la consola:", out.String())
}

func Index(w http.ResponseWriter, r *http.Request) {
	// Verificar si la solicitud es POST
	if r.Method == http.MethodPost {
		// Parsear el formulario
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error al parsear el formulario", http.StatusInternalServerError)
			return
		}

		// Obtener el valor del campo de texto (asumiendo que el campo tiene el nombre "campoTexto")
		campoTexto := r.Form.Get("campoTexto")

		//Instances of the necessary ANTLR variables.
		text := campoTexto
		input := antlr.NewInputStream(text)
		lexer := generated.NewMiniGoScanner(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		PM := generated.NewMiniGoParser(stream)

		errorListener := NewCustomErrorListener()
		PM.RemoveErrorListeners()
		PM.AddErrorListener(errorListener)

		//Calling the method of the main grammar.
		tree := PM.Root()

		// Error checking

		if len(errorListener.errors) > 0 {
			CheckerVar := checker.NewChecker()
			//rootCtx, _ := tree.(*generated.RootASTContext)
			CheckerVar.Visit(tree)
			//CheckerVar.VisitRootAST(rootCtx)

			// Construir el mensaje de error

			errorMessage = "Compilation errors:\n"
			for _, err := range errorListener.errors {
				errorMessage += fmt.Sprintf("%s\n", err)
			}

			// Enviar solo los elementos de la lista como respuesta
			fmt.Fprintf(w, "%s", errorMessage)
			return // Salir de la función para evitar enviar el HTML después
		} else {
			errorMessage = "Compilation successful!"
			encoder := encoder.NewEncoder()
			module := encoder.Visit(tree)
			fmt.Println(module)
			runModule(module.(*ir.Module))

			fmt.Fprintf(w, "%s", errorMessage)

			return // Salir de la función para evitar enviar el HTML después

		}
	}

	// Renderizar la plantilla HTML como antes
	template, err := template.ParseFiles("GUI/Index.html")
	if err != nil {
		http.Error(w, "Error al cargar la plantilla HTML", http.StatusInternalServerError)
		return
	}

	// Ejecutar la plantilla con los datos del usuario
	err = template.Execute(w, "")
	if err != nil {
		http.Error(w, "Error al renderizar la plantilla HTML", http.StatusInternalServerError)
		return
	}

}

func main() {

	http.HandleFunc("/", Index)
	http.ListenAndServe("localhost:8080", nil)

}
