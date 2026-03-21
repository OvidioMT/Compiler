// Code generated from C:/Users/Lenovo/Desktop/l Semestre 2024/Compiladores e interpretes/Semana 15/proyecto3/MiniGoParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package generated // MiniGoParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by MiniGoParser.
type MiniGoParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MiniGoParser#rootAST.
	VisitRootAST(ctx *RootASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#topDeclarationListAST.
	VisitTopDeclarationListAST(ctx *TopDeclarationListASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#variableDeclSingAST.
	VisitVariableDeclSingAST(ctx *VariableDeclSingASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#variableDeclPaRAST.
	VisitVariableDeclPaRAST(ctx *VariableDeclPaRASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#variableDeclNullAST.
	VisitVariableDeclNullAST(ctx *VariableDeclNullASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#innerVarDeclsAST.
	VisitInnerVarDeclsAST(ctx *InnerVarDeclsASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#singleVarDeclIdenAST.
	VisitSingleVarDeclIdenAST(ctx *SingleVarDeclIdenASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#singleVarDeclIAST.
	VisitSingleVarDeclIAST(ctx *SingleVarDeclIASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#singleVarDeclNoExpsSVDAST.
	VisitSingleVarDeclNoExpsSVDAST(ctx *SingleVarDeclNoExpsSVDASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#singleVarDeclNoExpsAST.
	VisitSingleVarDeclNoExpsAST(ctx *SingleVarDeclNoExpsASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#typeDeclTypeAST.
	VisitTypeDeclTypeAST(ctx *TypeDeclTypeASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#typeDeclPARAST.
	VisitTypeDeclPARAST(ctx *TypeDeclPARASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#typeDeclNULPARAST.
	VisitTypeDeclNULPARAST(ctx *TypeDeclNULPARASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#innerTypeDeclsAST.
	VisitInnerTypeDeclsAST(ctx *InnerTypeDeclsASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#singleTypeDeclAST.
	VisitSingleTypeDeclAST(ctx *SingleTypeDeclASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#funcDeclAST.
	VisitFuncDeclAST(ctx *FuncDeclASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#funcFrontDeclAST.
	VisitFuncFrontDeclAST(ctx *FuncFrontDeclASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#funcArgDeclsAST.
	VisitFuncArgDeclsAST(ctx *FuncArgDeclsASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#declTypeParAST.
	VisitDeclTypeParAST(ctx *DeclTypeParASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#declTypeIDAST.
	VisitDeclTypeIDAST(ctx *DeclTypeIDASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#sliceDeclTypeDeclAST.
	VisitSliceDeclTypeDeclAST(ctx *SliceDeclTypeDeclASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#arrayDeclTypeDeclAST.
	VisitArrayDeclTypeDeclAST(ctx *ArrayDeclTypeDeclASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#structDeclTypeDeckAST.
	VisitStructDeclTypeDeckAST(ctx *StructDeclTypeDeckASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#sliceDeclTypeAST.
	VisitSliceDeclTypeAST(ctx *SliceDeclTypeASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#arrayDeclTypeAST.
	VisitArrayDeclTypeAST(ctx *ArrayDeclTypeASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#structDeclTypeAST.
	VisitStructDeclTypeAST(ctx *StructDeclTypeASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#structMemDeclsAST.
	VisitStructMemDeclsAST(ctx *StructMemDeclsASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#identifierListAST.
	VisitIdentifierListAST(ctx *IdentifierListASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#sumaExpresionAST.
	VisitSumaExpresionAST(ctx *SumaExpresionASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionAST.
	VisitExpressionAST(ctx *ExpressionASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#mayorExpresionAST.
	VisitMayorExpresionAST(ctx *MayorExpresionASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#menorIgExpresionAST.
	VisitMenorIgExpresionAST(ctx *MenorIgExpresionASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#restaExExpresionAST.
	VisitRestaExExpresionAST(ctx *RestaExExpresionASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#anddExpresionAST.
	VisitAnddExpresionAST(ctx *AnddExpresionASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#mayorMExpresionAST.
	VisitMayorMExpresionAST(ctx *MayorMExpresionASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#menorExpresionAST.
	VisitMenorExpresionAST(ctx *MenorExpresionASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#sumaExExpresionAST.
	VisitSumaExExpresionAST(ctx *SumaExExpresionASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#bitExpresionAST.
	VisitBitExpresionAST(ctx *BitExpresionASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#diferenteExpresionAST.
	VisitDiferenteExpresionAST(ctx *DiferenteExpresionASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#multExpresionAST.
	VisitMultExpresionAST(ctx *MultExpresionASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#divExpresionAST.
	VisitDivExpresionAST(ctx *DivExpresionASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#orExpresionAST.
	VisitOrExpresionAST(ctx *OrExpresionASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#xorExpresionAST.
	VisitXorExpresionAST(ctx *XorExpresionASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#identicoExpresionAST.
	VisitIdenticoExpresionAST(ctx *IdenticoExpresionASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#mayorIgExpresionAST.
	VisitMayorIgExpresionAST(ctx *MayorIgExpresionASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#exclExpresionAST.
	VisitExclExpresionAST(ctx *ExclExpresionASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#andExpresionAST.
	VisitAndExpresionAST(ctx *AndExpresionASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#ordExpresionAST.
	VisitOrdExpresionAST(ctx *OrdExpresionASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#restaExpresionAST.
	VisitRestaExpresionAST(ctx *RestaExpresionASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#xorExExpresionAST.
	VisitXorExExpresionAST(ctx *XorExExpresionASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#porcentExpresionAST.
	VisitPorcentExpresionAST(ctx *PorcentExpresionASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#menorMExpresionAST.
	VisitMenorMExpresionAST(ctx *MenorMExpresionASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionListAST.
	VisitExpressionListAST(ctx *ExpressionListASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#primaryExpressionInAST.
	VisitPrimaryExpressionInAST(ctx *PrimaryExpressionInASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#primaryExpressionArAST.
	VisitPrimaryExpressionArAST(ctx *PrimaryExpressionArASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#primaryExpressionAppendAST.
	VisitPrimaryExpressionAppendAST(ctx *PrimaryExpressionAppendASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#primaryExpressionSeAST.
	VisitPrimaryExpressionSeAST(ctx *PrimaryExpressionSeASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#primaryExpressionCapEAST.
	VisitPrimaryExpressionCapEAST(ctx *PrimaryExpressionCapEASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#primaryExpressionOpAST.
	VisitPrimaryExpressionOpAST(ctx *PrimaryExpressionOpASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#primaryExpressionLengAST.
	VisitPrimaryExpressionLengAST(ctx *PrimaryExpressionLengASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#operandLitAST.
	VisitOperandLitAST(ctx *OperandLitASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#operandIDAST.
	VisitOperandIDAST(ctx *OperandIDASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#operandPARAST.
	VisitOperandPARAST(ctx *OperandPARASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#literalIntAST.
	VisitLiteralIntAST(ctx *LiteralIntASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#literalFloatAST.
	VisitLiteralFloatAST(ctx *LiteralFloatASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#literalRuneAST.
	VisitLiteralRuneAST(ctx *LiteralRuneASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#literalRawsAST.
	VisitLiteralRawsAST(ctx *LiteralRawsASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#literalInterAST.
	VisitLiteralInterAST(ctx *LiteralInterASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#indexAST.
	VisitIndexAST(ctx *IndexASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#argumentsAST.
	VisitArgumentsAST(ctx *ArgumentsASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#selectorAST.
	VisitSelectorAST(ctx *SelectorASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#appendExpressionAST.
	VisitAppendExpressionAST(ctx *AppendExpressionASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#lengthExpressionAST.
	VisitLengthExpressionAST(ctx *LengthExpressionASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#capExpressionAST.
	VisitCapExpressionAST(ctx *CapExpressionASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#statementListAST.
	VisitStatementListAST(ctx *StatementListASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#blockAST.
	VisitBlockAST(ctx *BlockASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#statementAST.
	VisitStatementAST(ctx *StatementASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#printlStAST.
	VisitPrintlStAST(ctx *PrintlStASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#returnStAST.
	VisitReturnStAST(ctx *ReturnStASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#breakStAST.
	VisitBreakStAST(ctx *BreakStASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#continueStAST.
	VisitContinueStAST(ctx *ContinueStASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#simpleStatementStAST.
	VisitSimpleStatementStAST(ctx *SimpleStatementStASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#blockStAST.
	VisitBlockStAST(ctx *BlockStASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#switchStAST.
	VisitSwitchStAST(ctx *SwitchStASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#ifStatementStAST.
	VisitIfStatementStAST(ctx *IfStatementStASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#loopStAST.
	VisitLoopStAST(ctx *LoopStASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#typeDeclStAST.
	VisitTypeDeclStAST(ctx *TypeDeclStASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#variableDeclStAST.
	VisitVariableDeclStAST(ctx *VariableDeclStASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#simpleStatementAST.
	VisitSimpleStatementAST(ctx *SimpleStatementASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionSimpStAST.
	VisitExpressionSimpStAST(ctx *ExpressionSimpStASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#assignmentStatementSimpStAST.
	VisitAssignmentStatementSimpStAST(ctx *AssignmentStatementSimpStASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#assignSimpStAST.
	VisitAssignSimpStAST(ctx *AssignSimpStASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#assignmentStatementAST.
	VisitAssignmentStatementAST(ctx *AssignmentStatementASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#incremAsSTAST.
	VisitIncremAsSTAST(ctx *IncremAsSTASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#andCompAsSTAST.
	VisitAndCompAsSTAST(ctx *AndCompAsSTASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#decremeAsSTAST.
	VisitDecremeAsSTAST(ctx *DecremeAsSTASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#assignNorAsSTAST.
	VisitAssignNorAsSTAST(ctx *AssignNorAsSTASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#assignMultAsSTAST.
	VisitAssignMultAsSTAST(ctx *AssignMultAsSTASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#assingnXorAsSTAST.
	VisitAssingnXorAsSTAST(ctx *AssingnXorAsSTASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#assignMaAsSTAST.
	VisitAssignMaAsSTAST(ctx *AssignMaAsSTASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#assignMeAsSTAST.
	VisitAssignMeAsSTAST(ctx *AssignMeAsSTASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#assignBitAsSTAST.
	VisitAssignBitAsSTAST(ctx *AssignBitAsSTASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#assignPorAsSTAST.
	VisitAssignPorAsSTAST(ctx *AssignPorAsSTASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#assignDivAsSTAST.
	VisitAssignDivAsSTAST(ctx *AssignDivAsSTASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#ifStatementAST.
	VisitIfStatementAST(ctx *IfStatementASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#ifElseStAST.
	VisitIfElseStAST(ctx *IfElseStASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#ifElseBlockStAST.
	VisitIfElseBlockStAST(ctx *IfElseBlockStASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#ifElseExpressionBlAST.
	VisitIfElseExpressionBlAST(ctx *IfElseExpressionBlASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#ifElseExpIfstAST.
	VisitIfElseExpIfstAST(ctx *IfElseExpIfstASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#ifElseExpBlockBAST.
	VisitIfElseExpBlockBAST(ctx *IfElseExpBlockBASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#forBlLoopAST.
	VisitForBlLoopAST(ctx *ForBlLoopASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#forExpBlLoopAST.
	VisitForExpBlLoopAST(ctx *ForExpBlLoopASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#forSimpStEXLoopAST.
	VisitForSimpStEXLoopAST(ctx *ForSimpStEXLoopASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#forSimpStSimpStLoopAST.
	VisitForSimpStSimpStLoopAST(ctx *ForSimpStSimpStLoopASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#switchAST.
	VisitSwitchAST(ctx *SwitchASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#switchExprAST.
	VisitSwitchExprAST(ctx *SwitchExprASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#switchSimpStAST.
	VisitSwitchSimpStAST(ctx *SwitchSimpStASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#switchExCaseAST.
	VisitSwitchExCaseAST(ctx *SwitchExCaseASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionCaseClauseListAST.
	VisitExpressionCaseClauseListAST(ctx *ExpressionCaseClauseListASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionCaseClauseExpAST.
	VisitExpressionCaseClauseExpAST(ctx *ExpressionCaseClauseExpASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionCaseClauseAST.
	VisitExpressionCaseClauseAST(ctx *ExpressionCaseClauseASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionSwitchCaseAST.
	VisitExpressionSwitchCaseAST(ctx *ExpressionSwitchCaseASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#defaultExpressionSAST.
	VisitDefaultExpressionSAST(ctx *DefaultExpressionSASTContext) interface{}
}
