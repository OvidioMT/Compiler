package encoder

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"proyecto3/generated"
	"strconv"
	"strings"
)

type Encoder struct {
	*generated.BaseMiniGoParserVisitor
	mainModule *ir.Module
	aquiStr    *ir.Global
}

func NewEncoder() *Encoder {
	return &Encoder{
		BaseMiniGoParserVisitor: &generated.BaseMiniGoParserVisitor{},
		mainModule:              ir.NewModule(),
		aquiStr:                 nil,
	}
}

func (v *Encoder) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *Encoder) VisitChildren(tree antlr.RuleNode) any {
	var result any
	n := tree.GetChildCount()
	for i := 0; i < n; i++ {
		c := tree.GetChild(i)
		val := c.(antlr.ParseTree)
		result2 := result
		result = v.Visit(val)
		if result == nil {
			result = result2
		}
	}
	return result
}

var arreglosL ListaArreglos
var general listaVariablesGeneral
var realVariables []string

var funcActual *ir.Func //se usa para la función que se esté visitando de turno... solo una a la vez

var blocksFunActual Stack               //se usa almacenar los bloques de la función que se esté visitando de turno
var variablesLocales BlockVariableTable //se usa para almacenar las variables locales del bloque de turno. Variables X Bloque
//esta última tabla de variables locales tiene que pensarse mejor porque funciona bien cuando hablamos de identificadores declarados en el mismo bloque donde se usan
//si estamos hablando de identificadores más globales al bloque donde se usan, entonces no creo que funcione tan bien o al menos debería de pensarse mejor y cambiarse

// Para los append
var isAppend bool = false
var appendVar string = ""

// para los index o len
var varToBeAssing string = ""
var varNameArray string = ""
var isIndex bool = false

//segundaPruebaIndex

var buscandoIndex bool = false
var positionIndex int = 0
var finalVarNameStack string = ""
var countingProcess int = 0

//var tempIndex int = 0

func (v *Encoder) VisitRootAST(ctx *generated.RootASTContext) interface{} {
	//se inicializa la estructura de variables locales
	variablesLocales = *NewBlockVariableTable()
	v.VisitChildren(ctx)
	return v.mainModule
}

func (v *Encoder) VisitTopDeclarationListAST(ctx *generated.TopDeclarationListASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitVariableDeclSingAST(ctx *generated.VariableDeclSingASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitVariableDeclPaRAST(ctx *generated.VariableDeclPaRASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitVariableDeclNullAST(ctx *generated.VariableDeclNullASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitInnerVarDeclsAST(ctx *generated.InnerVarDeclsASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitSingleVarDeclIdenAST(ctx *generated.SingleVarDeclIdenASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitSingleVarDeclIAST(ctx *generated.SingleVarDeclIASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitSingleVarDeclNoExpsSVDAST(ctx *generated.SingleVarDeclNoExpsSVDASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitSingleVarDeclNoExpsAST(ctx *generated.SingleVarDeclNoExpsASTContext) interface{} {
	//el manejo de declaraciones se debe hacer en el bloque actual en el que se esté trabajando
	// por lo que se obtiene de la pila de bloques, al actual
	blockactual, _ := blocksFunActual.Peek()

	//se obtienen todos los identificadores de la declaración para declararlos todos
	ids := ctx.IdentifierList().(*generated.IdentifierListASTContext).AllID()
	var ids2 string

	ids2 = ids[0].GetText()
	// se recorre la lista para generar el código para solicitar memoria para el identificador
	// se usa quemado el tipo entero... debe obenerse del contexto

	for i := 0; i < len(ids); i++ {
		if ctx.DeclType().GetText() == "int" {

			vActual := blockactual.NewAlloca(types.I32) //tipo de dato int
			blockactual.NewStore(constant.NewInt(types.I32, 0), vActual)
			//se agrega a una tabla de variablesXbloque la nueva variable creada
			variablesLocales.AddVariable(blockactual, ids[i].GetText(), vActual)

		} else if ctx.DeclType().GetText() == "float" {
			vActual := blockactual.NewAlloca(types.Float) //tipo de dato float
			blockactual.NewStore(constant.NewFloat(types.Float, 0.0), vActual)
			//se agrega a una tabla de variablesXbloque la nueva variable creada
			variablesLocales.AddVariable(blockactual, ids[i].GetText(), vActual)

		} else if ctx.DeclType().GetText() == "rune" {
			vActual := blockactual.NewAlloca(types.I32)                           // Crear una variable 'char' (int8)
			blockactual.NewStore(constant.NewInt(types.I32, int64('0')), vActual) // Inicializar con valor 0
			// Agregar la nueva variable a la tabla de variables por bloque
			variablesLocales.AddVariable(blockactual, ids[i].GetText(), vActual)

		} else { //si no es nignuno de los anteriores, la variable que se esta guardando se trata de un arreglo
			vActual := blockactual.NewAlloca(types.I32) //tipo de dato int
			blockactual.NewStore(constant.NewInt(types.I32, 0), vActual)
			//se agrega a una tabla de variablesXbloque la nueva variable creada
			variablesLocales.AddVariable(blockactual, ids2, vActual)
		}
	}
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitTypeDeclTypeAST(ctx *generated.TypeDeclTypeASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitTypeDeclPARAST(ctx *generated.TypeDeclPARASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitTypeDeclNULPARAST(ctx *generated.TypeDeclNULPARASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitInnerTypeDeclsAST(ctx *generated.InnerTypeDeclsASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitSingleTypeDeclAST(ctx *generated.SingleTypeDeclASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitFuncDeclAST(ctx *generated.FuncDeclASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitFuncFrontDeclAST(ctx *generated.FuncFrontDeclASTContext) interface{} {
	// se crea la función actual global con el nombfre que venga en el nodo del ctx
	// no se incluyen los parámetros porque no se hace esta visita... debe hacerse
	// solo se obtiene el tipo de retorno para int... debe ampliarse para el resto

	if v.Visit(ctx.DeclType()) == "int" {

		funcActual = v.mainModule.NewFunc(ctx.ID().GetText(), types.I32)

	} else if v.Visit(ctx.DeclType()) == "float" {

		funcActual = v.mainModule.NewFunc(ctx.ID().GetText(), types.Float)

	}

	//se visitan sus hijos de esta forma para abarcar el block, pero puede que sea necesario hacerse de manera individual y luego retornar nil
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitFuncArgDeclsAST(ctx *generated.FuncArgDeclsASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitDeclTypeParAST(ctx *generated.DeclTypeParASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitDeclTypeIDAST(ctx *generated.DeclTypeIDASTContext) interface{} {
	return ctx.ID().GetText()
}

func (v *Encoder) VisitSliceDeclTypeDeclAST(ctx *generated.SliceDeclTypeDeclASTContext) interface{} {
	variable, err := general.obtenerPrimerElemento()

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		arregloTemp := nuevoArreglo(variable.nombre, []int{})
		arreglosL.AgregarAlInicio(arregloTemp)
		general.eliminarPorNombre(variable.nombre)

	}
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitArrayDeclTypeDeclAST(ctx *generated.ArrayDeclTypeDeclASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitStructDeclTypeDeckAST(ctx *generated.StructDeclTypeDeckASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitSliceDeclTypeAST(ctx *generated.SliceDeclTypeASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitArrayDeclTypeAST(ctx *generated.ArrayDeclTypeASTContext) interface{} {

	//todo: arraydeclaration
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitStructDeclTypeAST(ctx *generated.StructDeclTypeASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitStructMemDeclsAST(ctx *generated.StructMemDeclsASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitIdentifierListAST(ctx *generated.IdentifierListASTContext) interface{} {
	//se devuelve toda la lista de tokens con los ids que vengan
	//especialmente util pasar este contenido al nodo anterior en las declaraciones de variables
	return ctx.AllID()
}

func (v *Encoder) VisitSumaExpresionAST(ctx *generated.SumaExpresionASTContext) interface{} {
	var1, _ := variablesLocales.FindValue(ctx.Expression(0).GetText())
	var2, _ := variablesLocales.FindValue(ctx.Expression(1).GetText())

	//se obtiene el bloque actual de la pila de bloques
	blockActual, _ := blocksFunActual.Peek()
	//se crea la instrucción de comparación en el bloque actual esperando que sea el correcto y se devuelve
	auxiliar := blockActual.NewAdd(var2, var1)
	numero, _ := strconv.Atoi(auxiliar.Y.Ident())
	numero2, _ := strconv.Atoi(auxiliar.X.Ident())
	valor := numero + numero2

	valReturn := constant.NewInt(types.I32, int64(valor))

	return valReturn
}

func (v *Encoder) VisitExpressionAST(ctx *generated.ExpressionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitMayorExpresionAST(ctx *generated.MayorExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitMenorIgExpresionAST(ctx *generated.MenorIgExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitRestaExExpresionAST(ctx *generated.RestaExExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitAnddExpresionAST(ctx *generated.AnddExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitMayorMExpresionAST(ctx *generated.MayorMExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitMenorExpresionAST(ctx *generated.MenorExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitSumaExExpresionAST(ctx *generated.SumaExExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitBitExpresionAST(ctx *generated.BitExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitDiferenteExpresionAST(ctx *generated.DiferenteExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitMultExpresionAST(ctx *generated.MultExpresionASTContext) interface{} {
	var1, _ := variablesLocales.FindValue(ctx.Expression(0).GetText())
	var2, _ := variablesLocales.FindValue(ctx.Expression(1).GetText())

	//se obtiene el bloque actual de la pila de bloques
	blockActual, _ := blocksFunActual.Peek()
	//se crea la instrucción de comparación en el bloque actual esperando que sea el correcto y se devuelve
	auxiliar := blockActual.NewMul(var2, var1)
	numero, _ := strconv.Atoi(auxiliar.Y.Ident())
	numero2, _ := strconv.Atoi(auxiliar.X.Ident())
	valor := numero * numero2

	valReturn := constant.NewInt(types.I32, int64(valor))

	return valReturn
}

func (v *Encoder) VisitDivExpresionAST(ctx *generated.DivExpresionASTContext) interface{} {
	var1, _ := variablesLocales.FindValue(ctx.Expression(0).GetText())
	var2, _ := variablesLocales.FindValue(ctx.Expression(1).GetText())

	//se obtiene el bloque actual de la pila de bloques
	blockActual, _ := blocksFunActual.Peek()
	//se crea la instrucción de comparación en el bloque actual esperando que sea el correcto y se devuelve
	auxiliar := blockActual.NewSDiv(var1, var2)
	numero, _ := strconv.Atoi(auxiliar.X.Ident())
	numero2, _ := strconv.Atoi(auxiliar.Y.Ident())
	valor := numero / numero2

	valReturn := constant.NewInt(types.I32, int64(valor))

	return valReturn
}

func (v *Encoder) VisitOrExpresionAST(ctx *generated.OrExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitXorExpresionAST(ctx *generated.XorExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitIdenticoExpresionAST(ctx *generated.IdenticoExpresionASTContext) interface{} {
	//una comparación tiene dos expresiones que se deben visitar para obtener los Values
	valExpression1 := v.Visit(ctx.Expression(0)).(value.Value)
	valExpression2 := v.Visit(ctx.Expression(1)).(value.Value)
	//se obtiene el bloque actual de la pila de bloques
	blockActual, _ := blocksFunActual.Peek()
	//se crea la instrucción de comparación en el bloque actual esperando que sea el correcto y se devuelve
	valReturn := blockActual.NewICmp(enum.IPredEQ, valExpression1, valExpression2)

	return valReturn
}

func (v *Encoder) VisitMayorIgExpresionAST(ctx *generated.MayorIgExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitExclExpresionAST(ctx *generated.ExclExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitAndExpresionAST(ctx *generated.AndExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitOrdExpresionAST(ctx *generated.OrdExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitRestaExpresionAST(ctx *generated.RestaExpresionASTContext) interface{} {
	var1, _ := variablesLocales.FindValue(ctx.Expression(0).GetText())
	var2, _ := variablesLocales.FindValue(ctx.Expression(1).GetText())

	//se obtiene el bloque actual de la pila de bloques
	blockActual, _ := blocksFunActual.Peek()
	//se crea la instrucción de comparación en el bloque actual esperando que sea el correcto y se devuelve
	auxiliar := blockActual.NewSub(var1, var2)
	numero, _ := strconv.Atoi(auxiliar.X.Ident())
	numero2, _ := strconv.Atoi(auxiliar.Y.Ident())
	valor := numero - numero2

	valReturn := constant.NewInt(types.I32, int64(valor))

	return valReturn
}

func (v *Encoder) VisitXorExExpresionAST(ctx *generated.XorExExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitPorcentExpresionAST(ctx *generated.PorcentExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitMenorMExpresionAST(ctx *generated.MenorMExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitExpressionListAST(ctx *generated.ExpressionListASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitPrimaryExpressionInAST(ctx *generated.PrimaryExpressionInASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitPrimaryExpressionArAST(ctx *generated.PrimaryExpressionArASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitPrimaryExpressionAppendAST(ctx *generated.PrimaryExpressionAppendASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitPrimaryExpressionSeAST(ctx *generated.PrimaryExpressionSeASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitPrimaryExpressionCapEAST(ctx *generated.PrimaryExpressionCapEASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitPrimaryExpressionOpAST(ctx *generated.PrimaryExpressionOpASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitPrimaryExpressionLengAST(ctx *generated.PrimaryExpressionLengASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitOperandLitAST(ctx *generated.OperandLitASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitOperandIDAST(ctx *generated.OperandIDASTContext) interface{} {
	//se busca en el almacen de variables locales del bloque actual la variable del contexto para devolverla
	//TODO: se debe considerar qué pasa si la variable no fue declarado en el bloque del tope, por ejemplo variables globales al contexto o parámetros
	blockActual, _ := blocksFunActual.Peek()
	val, _ := variablesLocales.GetVariable(blockActual, ctx.ID().GetText())
	valReturn := blockActual.NewLoad(types.I32, val)

	if isAppend == true {
		appendVar = ctx.ID().GetText()
	} else { //si no es para un append, podría ser para un index
		varNameArray = ctx.ID().GetText()
	}
	return valReturn
}

func (v *Encoder) VisitOperandPARAST(ctx *generated.OperandPARASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitLiteralIntAST(ctx *generated.LiteralIntASTContext) interface{} {
	valReturn, _ := constant.NewIntFromString(types.I32, ctx.INT().GetText())
	if isAppend == true {
		valueT := ctx.INT().GetText()

		entero, _ := strconv.Atoi(valueT)

		arreglosL.AgregarNumeroPorTexto(appendVar, entero)

		isAppend = false

	} else if buscandoIndex == true {

		numero := ctx.INT().GetText()
		positionIndex, _ = strconv.Atoi(numero)
		countingProcess++
	}
	return valReturn
}

func (v *Encoder) VisitLiteralFloatAST(ctx *generated.LiteralFloatASTContext) interface{} {
	valReturn, _ := constant.NewFloatFromString(types.Float, ctx.FLOAT().GetText())
	return valReturn
}

func (v *Encoder) VisitLiteralRuneAST(ctx *generated.LiteralRuneASTContext) interface{} {
	runeLiteral := ctx.RUNE().GetText()
	runeStr := strings.Trim(runeLiteral, "'")

	// Intentar convertir el string a un número entero
	if runeValue, err := strconv.Atoi(runeStr); err == nil {
		// El string representa un número, retornar el valor entero correspondiente
		valReturn := constant.NewInt(types.I32, int64(runeValue))
		return valReturn
	}

	// Si no se pudo convertir a un número, asumir que es una letra y obtener su valor Unicode
	if len(runeStr) != 1 {
		panic("Invalid rune literal")
	}
	runeValue := int64(runeStr[0])
	valReturn := constant.NewInt(types.I32, runeValue)
	return valReturn
}

func (v *Encoder) VisitLiteralRawsAST(ctx *generated.LiteralRawsASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitLiteralInterAST(ctx *generated.LiteralInterASTContext) interface{} {
	var b string = ctx.INTERPRETEDSTRING().GetText()
	result := strings.ReplaceAll(b, `"`, "")
	SaquiStr := v.mainModule.NewGlobalDef("aquiStr", constant.NewCharArrayFromString(result+"\n\x00"))

	puts := v.mainModule.NewFunc("puts", types.I32, ir.NewParam("", types.I8Ptr))

	blockactual, _ := blocksFunActual.Peek()

	blockactual.NewCall(puts, SaquiStr)

	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitIndexAST(ctx *generated.IndexASTContext) interface{} {
	numIndex := ctx.Expression().GetText()

	varNameInStack := varNameArray + numIndex

	finalVarNameStack = varNameInStack

	blockActual, _ := blocksFunActual.Peek()

	newValue, _ := variablesLocales.FindValue(varNameInStack)

	variablesLocales.AddVariable(blockActual, varToBeAssing, newValue)

	numero, _ := strconv.Atoi(numIndex)

	positionIndex = numero

	isIndex = true
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitArgumentsAST(ctx *generated.ArgumentsASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitSelectorAST(ctx *generated.SelectorASTContext) interface{} {
	return v.VisitChildren(ctx)
}

var tempIndex int = 0

func (v *Encoder) VisitAppendExpressionAST(ctx *generated.AppendExpressionASTContext) interface{} {
	blockActual, _ := blocksFunActual.Peek()

	varNameOriginal := ctx.Expression(0).GetText()
	valorAssign := v.Visit(ctx.Expression(1)).(value.Value)
	//valorAssignText := ctx.Expression(1).GetText()
	//valorAssignInt, _ := strconv.Atoi(valorAssignText)
	cadena := strconv.Itoa(tempIndex)

	varNameForStack := varNameOriginal + cadena

	variablesLocales.AddVariable(blockActual, varNameForStack, valorAssign)

	tempIndex++
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitLengthExpressionAST(ctx *generated.LengthExpressionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitCapExpressionAST(ctx *generated.CapExpressionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitStatementListAST(ctx *generated.StatementListASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitBlockAST(ctx *generated.BlockASTContext) interface{} {
	//se agrega el bloque actual creado a la lista de bloques.. sin nombre
	blocksFunActual.Push(funcActual.NewBlock(""))

	//se visita el contenido del Block
	v.VisitChildren(ctx)

	//que se devuelva el bloque creado para conocerlo a la vuelta del visit
	returnBlock, _ := blocksFunActual.Peek()
	return returnBlock
}

func (v *Encoder) VisitStatementAST(ctx *generated.StatementASTContext) interface{} {
	if strings.HasPrefix(ctx.ExpressionList().GetText(), `"`) {

		return v.VisitChildren(ctx)
	} else {
		val1, _ := variablesLocales.FindValue(ctx.ExpressionList().GetText())

		SaquiStr := v.mainModule.NewGlobalDef("aquiStr", constant.NewCharArrayFromString(val1.Ident()+"\x00"))

		puts := v.mainModule.NewFunc("puts", types.I32, ir.NewParam("", types.I8Ptr))
		//block, _ := variablesLocales.GetBlockNumberWithVariable(ctx.ExpressionList().GetText())
		blockactual, _ := blocksFunActual.Peek()

		blockactual.NewCall(puts, SaquiStr)

	}

	return nil
}

func (v *Encoder) VisitPrintlStAST(ctx *generated.PrintlStASTContext) interface{} {
	if strings.HasPrefix(ctx.ExpressionList().GetText(), `"`) {

		return v.VisitChildren(ctx)
	} else {
		val1, _ := variablesLocales.FindValue(ctx.ExpressionList().GetText())

		SaquiStr := v.mainModule.NewGlobalDef("aquiStr", constant.NewCharArrayFromString(val1.Ident()+"\n\x00"))

		puts := v.mainModule.NewFunc("puts", types.I32, ir.NewParam("", types.I8Ptr))
		block, _ := variablesLocales.GetBlockNumberWithVariable(ctx.ExpressionList().GetText())
		blockactual, _ := blocksFunActual.Peek2(block)

		blockactual.NewCall(puts, SaquiStr)

	}

	return nil

}

func (v *Encoder) VisitReturnStAST(ctx *generated.ReturnStASTContext) interface{} {
	//se obtiene el bloque actual de la pila de bloques
	blockActual, _ := blocksFunActual.Peek()
	//se escribe la instrucción en el bloque correspondiente...
	//el tipo de retorno es quemado a nil..
	//debe cambiarse para que retorne el valor que se obtenga del ctx

	retValue := v.Visit(ctx.Expression())
	return blockActual.NewRet(retValue.(value.Value))
}

func (v *Encoder) VisitBreakStAST(ctx *generated.BreakStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitContinueStAST(ctx *generated.ContinueStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitSimpleStatementStAST(ctx *generated.SimpleStatementStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitBlockStAST(ctx *generated.BlockStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitSwitchStAST(ctx *generated.SwitchStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitIfStatementStAST(ctx *generated.IfStatementStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitLoopStAST(ctx *generated.LoopStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitTypeDeclStAST(ctx *generated.TypeDeclStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitVariableDeclStAST(ctx *generated.VariableDeclStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitSimpleStatementAST(ctx *generated.SimpleStatementASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitExpressionSimpStAST(ctx *generated.ExpressionSimpStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitAssignmentStatementSimpStAST(ctx *generated.AssignmentStatementSimpStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitAssignSimpStAST(ctx *generated.AssignSimpStASTContext) interface{} {
	blocksFunActual, _ := blocksFunActual.Peek()
	vExpre := v.Visit(ctx.ExpressionList(1)).(value.Value)

	variablesLocales.AddVariable(blocksFunActual, ctx.ExpressionList(0).GetText(), vExpre)
	if isIndex == true {

		realVariables = append(realVariables, ctx.ExpressionList(0).GetText())

		b := realVariables[0]

		getter, _ := variablesLocales.FindValue(finalVarNameStack)

		println(getter)

		variablesLocales.AddVariable(blocksFunActual, b, getter)
	}
	return nil
}

func (v *Encoder) VisitAssignmentStatementAST(ctx *generated.AssignmentStatementASTContext) interface{} {
	blocksFunActual, _ := blocksFunActual.Peek()
	//vIds := v.Visit(ctx.ExpressionList(0)).(value.Value)
	vIds, _ := variablesLocales.GetVariable(blocksFunActual, ctx.ExpressionList(0).GetText())
	vExpre := v.Visit(ctx.ExpressionList(1)).(value.Value)

	variablesLocales.SetVariableValue(blocksFunActual, ctx.ExpressionList(0).GetText(), vExpre)

	blocksFunActual.NewStore(vExpre, vIds)

	return nil
}

func (v *Encoder) VisitIncremAsSTAST(ctx *generated.IncremAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitAndCompAsSTAST(ctx *generated.AndCompAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitDecremeAsSTAST(ctx *generated.DecremeAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitAssignNorAsSTAST(ctx *generated.AssignNorAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitAssignMultAsSTAST(ctx *generated.AssignMultAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitAssingnXorAsSTAST(ctx *generated.AssingnXorAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitAssignMaAsSTAST(ctx *generated.AssignMaAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitAssignMeAsSTAST(ctx *generated.AssignMeAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitAssignBitAsSTAST(ctx *generated.AssignBitAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitAssignPorAsSTAST(ctx *generated.AssignPorAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitAssignDivAsSTAST(ctx *generated.AssignDivAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitIfStatementAST(ctx *generated.IfStatementASTContext) interface{} {
	//se respada el bloque que precede al IF
	blockAnteriorIf, _ := blocksFunActual.Peek()
	//se visita la expresión para generar el código de la comparación
	valCond := v.Visit(ctx.Expression()).(value.Value)

	//blockAnteriorIf.NewLoad(types.I32, vActual)
	//valExpression2 := constant.NewInt(types.I32, 10)

	//se crea el cloque para el cuerpo del if...
	//luego se debe volver a poner las instrucciones del comparación y saldo al bloque Padre
	v.Visit(ctx.Block())

	//se respalda el bloque del cuerpo del IF
	blockCuerpoIf, _ := blocksFunActual.Peek()

	//se crea el bloque para el cierre del if
	blocksFunActual.Push(funcActual.NewBlock(""))
	blockEndIf, _ := blocksFunActual.Peek()
	//una vez creado el bloque de cierre, se debe agregar la instrucción de salto al final del bloque del cuerpo del if para terminar dicho bloque
	blockCuerpoIf.NewBr(blockEndIf)

	//se debe agregar la instrucción de salto al final del bloque anterior al if para que se ejecute el bloque de cierre
	blockAnteriorIf.NewCondBr(valCond, blockCuerpoIf, blockEndIf)

	return nil
}

func (v *Encoder) VisitIfElseStAST(ctx *generated.IfElseStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitIfElseBlockStAST(ctx *generated.IfElseBlockStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitIfElseExpressionBlAST(ctx *generated.IfElseExpressionBlASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitIfElseExpIfstAST(ctx *generated.IfElseExpIfstASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitIfElseExpBlockBAST(ctx *generated.IfElseExpBlockBASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitForBlLoopAST(ctx *generated.ForBlLoopASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitForExpBlLoopAST(ctx *generated.ForExpBlLoopASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitForSimpStEXLoopAST(ctx *generated.ForSimpStEXLoopASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitForSimpStSimpStLoopAST(ctx *generated.ForSimpStSimpStLoopASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitSwitchAST(ctx *generated.SwitchASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitSwitchExprAST(ctx *generated.SwitchExprASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitSwitchSimpStAST(ctx *generated.SwitchSimpStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitSwitchExCaseAST(ctx *generated.SwitchExCaseASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitExpressionCaseClauseListAST(ctx *generated.ExpressionCaseClauseListASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitExpressionCaseClauseExpAST(ctx *generated.ExpressionCaseClauseExpASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitExpressionCaseClauseAST(ctx *generated.ExpressionCaseClauseASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitExpressionSwitchCaseAST(ctx *generated.ExpressionSwitchCaseASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitDefaultExpressionSAST(ctx *generated.DefaultExpressionSASTContext) interface{} {
	return v.VisitChildren(ctx)
}
