package checker

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"proyecto3/generated"
)

const (
	PACKAGE = 0
	INT     = 1
	RUNE    = 2 //char
	FLOAT   = 3
	STRING  = 4
	BOOLEAN = 5
)

// Checker struct represents the symbol table and error list.
type Checker struct {
	*generated.BaseMiniGoParserVisitor
	symbolTable *SymbolTable
	errorList   []string
}

// NewChecker initializes a new Checker instance.
func NewChecker() *Checker {
	return &Checker{
		BaseMiniGoParserVisitor: &generated.BaseMiniGoParserVisitor{},
		symbolTable:             NewSymbolTable(),
		errorList:               make([]string, 0),
	}
}

func (v *Checker) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *Checker) VisitChildren(tree antlr.RuleNode) any {
	n := tree.GetChildCount()
	for i := 0; i < n; i++ {
		c := tree.GetChild(i)
		val := c.(antlr.ParseTree)
		_ = v.Visit(val)
	}
	return 0
}

// HasErrors checks if there are any errors.
func (c *Checker) HasErrors() bool {
	return len(c.errorList) > 0
}

// ErrorList returns the list of errors.
func (c *Checker) ErrorList() string {
	if !c.HasErrors() {
		return "0 errors"
	}
	var builder string
	for _, s := range c.errorList {
		builder += fmt.Sprintf("%s\n", s)
	}
	return builder
}

func (v *Checker) VisitRootAST(ctx *generated.RootASTContext) interface{} {

	packageNameSymbol := ctx.ID()                                 // Obtiene el nombre del paquete del contexto
	if v.symbolTable.search(packageNameSymbol.GetText()) != nil { // Verifica si el paquete ya está definido
		v.errorList = append(v.errorList, "Error: Package '"+packageNameSymbol.GetText()+"' already defined")
	} else {
		v.symbolTable.insert(packageNameSymbol.GetSymbol(), PACKAGE) // Inserta el paquete en la tabla de símbolos
	}
	//v.symbolTable.printTable()
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitTopDeclarationListAST(ctx *generated.TopDeclarationListASTContext) interface{} {
	for _, child := range ctx.GetChildren() {
		// Verificar si el hijo es una instancia de una declaración conocida
		switch decl := child.(type) {
		case *generated.VariableDeclSingASTContext:
			v.VisitVariableDeclSingAST(decl)
		case *generated.TypeDeclTypeASTContext:
			v.VisitTypeDeclTypeAST(decl)
		case *generated.FuncDeclASTContext:
			v.VisitFuncDeclAST(decl)
			// Agregar casos para otros tipos de declaraciones si es necesario
		}
	}
	return nil
}

func (v *Checker) VisitVariableDeclSingAST(ctx *generated.VariableDeclSingASTContext) interface{} {
	//println("================================================================", ctx.SingleVarDecl().GetStart().GetText())

	//varName := ctx.SingleVarDecl().GetStart()
	/*
		varName := ctx.SingleVarDecl().GetStart() // Obtiene el nombre de la variable del contexto
		varType := v.Visit(ctx.SingleVarDecl())
		if v.symbolTable.search(varName.GetText()) != nil {

			v.errorList = append(v.errorList, fmt.Sprintf("Error: Variable '%s' already defined"+
				" in this scope in line: '%d'", varName.GetText(), varName.GetLine()))

		} else {

			v.symbolTable.insert(varName, varType.(int)) //typeResult)
		}

		// Visitar los nodos hijos del contexto
		v.symbolTable.printTable()

	*/
	return nil

}

func (v *Checker) VisitVariableDeclPaRAST(ctx *generated.VariableDeclPaRASTContext) interface{} {
	println("================================================================", ctx.InnerVarDecls().GetText())
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitVariableDeclNullAST(ctx *generated.VariableDeclNullASTContext) interface{} {
	println("================================================================", ctx.GetStart().GetText())
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitInnerVarDeclsAST(ctx *generated.InnerVarDeclsASTContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Checker) VisitSingleVarDeclIdenAST(ctx *generated.SingleVarDeclIdenASTContext) interface{} {
	//println("                        VisitSingleVarDeclIdenAST                                                      ", ctx.GetStart().GetText())
	//println(v.symbolTable.currentLevel)
	//v.symbolTable.printTable()
	ids := ctx.IdentifierList()
	println(ids)
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitSingleVarDeclIAST(ctx *generated.SingleVarDeclIASTContext) interface{} {
	//println("                        VisitSingleVarDeclIAST                                                      ", ctx.GetStart().GetText())

	return v.VisitChildren(ctx)
}

func (v *Checker) VisitSingleVarDeclNoExpsSVDAST(ctx *generated.SingleVarDeclNoExpsSVDASTContext) interface{} {
	//println("                        VisitSingleVarDeclNoExpsSVDAST                                                      ", ctx.GetStart().GetText())
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitSingleVarDeclNoExpsAST(ctx *generated.SingleVarDeclNoExpsASTContext) interface{} {
	//TODO: AQUI SE DEFINEN LAS VARIABLES
	//C:\Users\Lenovo\Desktop\proyecto3\resources
	varName := ctx.GetStart()
	println("varrrrrrrrrrrrrrrrrrr                           ", ctx.DeclType().GetStart().GetText())

	if v.symbolTable.search(ctx.GetStart().GetText()) != nil {

		v.symbolTable.insert(varName, 1)

	} else {
		//var1 := v.symbolTable.search(ctx.GetStart().GetText())

		//println("==============================================================", var1)
		//v.symbolTable.insert(varName, 1)
		//v.errorList = append(v.errorList, fmt.Sprintf("Error: Variable '%s' already defined"+
		//" in this scope in line: '%d'", varName.GetText(), varName.GetLine()))
	}
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitTypeDeclTypeAST(ctx *generated.TypeDeclTypeASTContext) interface{} {
	//println("==============================================================", ctx.TypeDeclContext.GetText())
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitTypeDeclPARAST(ctx *generated.TypeDeclPARASTContext) interface{} {
	//println("=======================    VisitTypeDeclPARAST     ", ctx.TYPE().GetText())
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitTypeDeclNULPARAST(ctx *generated.TypeDeclNULPARASTContext) interface{} {
	//println("====================== VisitTypeDeclNULPARAST   ", ctx.TypeDeclContext.GetText())
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitInnerTypeDeclsAST(ctx *generated.InnerTypeDeclsASTContext) interface{} {
	//println("====================== VisitInnerTypeDeclsAST   ", ctx.GetStart().GetText())
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitSingleTypeDeclAST(ctx *generated.SingleTypeDeclASTContext) interface{} {
	//println("====================== VisitSingleTypeDeclAST   ", ctx.ID().GetText())

	return v.VisitChildren(ctx)
}

func (v *Checker) VisitFuncDeclAST(ctx *generated.FuncDeclASTContext) interface{} {
	//println("======================== VisitFuncDeclAST  ", ctx.FuncDeclContext)
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitFuncFrontDeclAST(ctx *generated.FuncFrontDeclASTContext) interface{} {
	//TODO AQUI ES PARA LAS FUNCIONES
	if v.symbolTable.search(ctx.ID().GetText()) != nil {
		//println("si entro")
	} else {

		//println("no entro")
	}

	return v.VisitChildren(ctx)
}

func (v *Checker) VisitFuncArgDeclsAST(ctx *generated.FuncArgDeclsASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitDeclTypeParAST(ctx *generated.DeclTypeParASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitDeclTypeIDAST(ctx *generated.DeclTypeIDASTContext) interface{} {
	println("====================== VisitDeclTypeIDAST    ", ctx.ID().GetText())

	return v.VisitChildren(ctx)
}

func (v *Checker) VisitSliceDeclTypeDeclAST(ctx *generated.SliceDeclTypeDeclASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitArrayDeclTypeDeclAST(ctx *generated.ArrayDeclTypeDeclASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitStructDeclTypeDeckAST(ctx *generated.StructDeclTypeDeckASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitSliceDeclTypeAST(ctx *generated.SliceDeclTypeASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitArrayDeclTypeAST(ctx *generated.ArrayDeclTypeASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitStructDeclTypeAST(ctx *generated.StructDeclTypeASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitStructMemDeclsAST(ctx *generated.StructMemDeclsASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitIdentifierListAST(ctx *generated.IdentifierListASTContext) interface{} {
	//println("====================== VisitIdentifierListAST   ", ctx.GetStart().GetText())
	if v.symbolTable.searchInCurrentLevel(ctx.GetStart().GetText()) != nil {
		//v.symbolTable.insert(ctx.GetStart(), 122)
	}
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitSumaExpresionAST(ctx *generated.SumaExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitExpressionAST(ctx *generated.ExpressionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitMayorExpresionAST(ctx *generated.MayorExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitMenorIgExpresionAST(ctx *generated.MenorIgExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitRestaExExpresionAST(ctx *generated.RestaExExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitAnddExpresionAST(ctx *generated.AnddExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitMayorMExpresionAST(ctx *generated.MayorMExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitMenorExpresionAST(ctx *generated.MenorExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitSumaExExpresionAST(ctx *generated.SumaExExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitBitExpresionAST(ctx *generated.BitExpresionASTContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Checker) VisitDiferenteExpresionAST(ctx *generated.DiferenteExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitMultExpresionAST(ctx *generated.MultExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitDivExpresionAST(ctx *generated.DivExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitOrExpresionAST(ctx *generated.OrExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitXorExpresionAST(ctx *generated.XorExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitIdenticoExpresionAST(ctx *generated.IdenticoExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitMayorIgExpresionAST(ctx *generated.MayorIgExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitExclExpresionAST(ctx *generated.ExclExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitAndExpresionAST(ctx *generated.AndExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitOrdExpresionAST(ctx *generated.OrdExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitRestaExpresionAST(ctx *generated.RestaExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitXorExExpresionAST(ctx *generated.XorExExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitPorcentExpresionAST(ctx *generated.PorcentExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitMenorMExpresionAST(ctx *generated.MenorMExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitExpressionListAST(ctx *generated.ExpressionListASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitPrimaryExpressionInAST(ctx *generated.PrimaryExpressionInASTContext) interface{} {

	v.symbolTable.printTable()

	return v.VisitChildren(ctx)
}

func (v *Checker) VisitPrimaryExpressionArAST(ctx *generated.PrimaryExpressionArASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitPrimaryExpressionAppendAST(ctx *generated.PrimaryExpressionAppendASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitPrimaryExpressionSeAST(ctx *generated.PrimaryExpressionSeASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitPrimaryExpressionCapEAST(ctx *generated.PrimaryExpressionCapEASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitPrimaryExpressionOpAST(ctx *generated.PrimaryExpressionOpASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitPrimaryExpressionLengAST(ctx *generated.PrimaryExpressionLengASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitOperandLitAST(ctx *generated.OperandLitASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitOperandIDAST(ctx *generated.OperandIDASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitOperandPARAST(ctx *generated.OperandPARASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitLiteralIntAST(ctx *generated.LiteralIntASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitLiteralFloatAST(ctx *generated.LiteralFloatASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitLiteralRuneAST(ctx *generated.LiteralRuneASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitLiteralRawsAST(ctx *generated.LiteralRawsASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitLiteralInterAST(ctx *generated.LiteralInterASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitIndexAST(ctx *generated.IndexASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitArgumentsAST(ctx *generated.ArgumentsASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitSelectorAST(ctx *generated.SelectorASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitAppendExpressionAST(ctx *generated.AppendExpressionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitLengthExpressionAST(ctx *generated.LengthExpressionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitCapExpressionAST(ctx *generated.CapExpressionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitStatementListAST(ctx *generated.StatementListASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitBlockAST(ctx *generated.BlockASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitStatementAST(ctx *generated.StatementASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitPrintlStAST(ctx *generated.PrintlStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitReturnStAST(ctx *generated.ReturnStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitBreakStAST(ctx *generated.BreakStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitContinueStAST(ctx *generated.ContinueStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitSimpleStatementStAST(ctx *generated.SimpleStatementStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitBlockStAST(ctx *generated.BlockStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitSwitchStAST(ctx *generated.SwitchStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitIfStatementStAST(ctx *generated.IfStatementStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitLoopStAST(ctx *generated.LoopStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitTypeDeclStAST(ctx *generated.TypeDeclStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitVariableDeclStAST(ctx *generated.VariableDeclStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitSimpleStatementAST(ctx *generated.SimpleStatementASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitExpressionSimpStAST(ctx *generated.ExpressionSimpStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitAssignmentStatementSimpStAST(ctx *generated.AssignmentStatementSimpStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitAssignSimpStAST(ctx *generated.AssignSimpStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitAssignmentStatementAST(ctx *generated.AssignmentStatementASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitIncremAsSTAST(ctx *generated.IncremAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitAndCompAsSTAST(ctx *generated.AndCompAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitDecremeAsSTAST(ctx *generated.DecremeAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitAssignNorAsSTAST(ctx *generated.AssignNorAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitAssignMultAsSTAST(ctx *generated.AssignMultAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitAssingnXorAsSTAST(ctx *generated.AssingnXorAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitAssignMaAsSTAST(ctx *generated.AssignMaAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitAssignMeAsSTAST(ctx *generated.AssignMeAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitAssignBitAsSTAST(ctx *generated.AssignBitAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitAssignPorAsSTAST(ctx *generated.AssignPorAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitAssignDivAsSTAST(ctx *generated.AssignDivAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitIfStatementAST(ctx *generated.IfStatementASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitIfElseStAST(ctx *generated.IfElseStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitIfElseBlockStAST(ctx *generated.IfElseBlockStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitIfElseExpressionBlAST(ctx *generated.IfElseExpressionBlASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitIfElseExpIfstAST(ctx *generated.IfElseExpIfstASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitIfElseExpBlockBAST(ctx *generated.IfElseExpBlockBASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitForBlLoopAST(ctx *generated.ForBlLoopASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitForExpBlLoopAST(ctx *generated.ForExpBlLoopASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitForSimpStEXLoopAST(ctx *generated.ForSimpStEXLoopASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitForSimpStSimpStLoopAST(ctx *generated.ForSimpStSimpStLoopASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitSwitchAST(ctx *generated.SwitchASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitSwitchExprAST(ctx *generated.SwitchExprASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitSwitchSimpStAST(ctx *generated.SwitchSimpStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitSwitchExCaseAST(ctx *generated.SwitchExCaseASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitExpressionCaseClauseListAST(ctx *generated.ExpressionCaseClauseListASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitExpressionCaseClauseExpAST(ctx *generated.ExpressionCaseClauseExpASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitExpressionCaseClauseAST(ctx *generated.ExpressionCaseClauseASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitExpressionSwitchCaseAST(ctx *generated.ExpressionSwitchCaseASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitDefaultExpressionSAST(ctx *generated.DefaultExpressionSASTContext) interface{} {
	return v.VisitChildren(ctx)

}
