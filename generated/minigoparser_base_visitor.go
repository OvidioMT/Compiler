// Code generated from C:/Users/Lenovo/Desktop/l Semestre 2024/Compiladores e interpretes/Semana 15/proyecto3/MiniGoParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package generated // MiniGoParser
import "github.com/antlr4-go/antlr/v4"

type BaseMiniGoParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMiniGoParserVisitor) VisitRootAST(ctx *RootASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitTopDeclarationListAST(ctx *TopDeclarationListASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitVariableDeclSingAST(ctx *VariableDeclSingASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitVariableDeclPaRAST(ctx *VariableDeclPaRASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitVariableDeclNullAST(ctx *VariableDeclNullASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitInnerVarDeclsAST(ctx *InnerVarDeclsASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitSingleVarDeclIdenAST(ctx *SingleVarDeclIdenASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitSingleVarDeclIAST(ctx *SingleVarDeclIASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitSingleVarDeclNoExpsSVDAST(ctx *SingleVarDeclNoExpsSVDASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitSingleVarDeclNoExpsAST(ctx *SingleVarDeclNoExpsASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitTypeDeclTypeAST(ctx *TypeDeclTypeASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitTypeDeclPARAST(ctx *TypeDeclPARASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitTypeDeclNULPARAST(ctx *TypeDeclNULPARASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitInnerTypeDeclsAST(ctx *InnerTypeDeclsASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitSingleTypeDeclAST(ctx *SingleTypeDeclASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitFuncDeclAST(ctx *FuncDeclASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitFuncFrontDeclAST(ctx *FuncFrontDeclASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitFuncArgDeclsAST(ctx *FuncArgDeclsASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitDeclTypeParAST(ctx *DeclTypeParASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitDeclTypeIDAST(ctx *DeclTypeIDASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitSliceDeclTypeDeclAST(ctx *SliceDeclTypeDeclASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitArrayDeclTypeDeclAST(ctx *ArrayDeclTypeDeclASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitStructDeclTypeDeckAST(ctx *StructDeclTypeDeckASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitSliceDeclTypeAST(ctx *SliceDeclTypeASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitArrayDeclTypeAST(ctx *ArrayDeclTypeASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitStructDeclTypeAST(ctx *StructDeclTypeASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitStructMemDeclsAST(ctx *StructMemDeclsASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitIdentifierListAST(ctx *IdentifierListASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitSumaExpresionAST(ctx *SumaExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitExpressionAST(ctx *ExpressionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitMayorExpresionAST(ctx *MayorExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitMenorIgExpresionAST(ctx *MenorIgExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitRestaExExpresionAST(ctx *RestaExExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitAnddExpresionAST(ctx *AnddExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitMayorMExpresionAST(ctx *MayorMExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitMenorExpresionAST(ctx *MenorExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitSumaExExpresionAST(ctx *SumaExExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitBitExpresionAST(ctx *BitExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitDiferenteExpresionAST(ctx *DiferenteExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitMultExpresionAST(ctx *MultExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitDivExpresionAST(ctx *DivExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitOrExpresionAST(ctx *OrExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitXorExpresionAST(ctx *XorExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitIdenticoExpresionAST(ctx *IdenticoExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitMayorIgExpresionAST(ctx *MayorIgExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitExclExpresionAST(ctx *ExclExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitAndExpresionAST(ctx *AndExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitOrdExpresionAST(ctx *OrdExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitRestaExpresionAST(ctx *RestaExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitXorExExpresionAST(ctx *XorExExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitPorcentExpresionAST(ctx *PorcentExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitMenorMExpresionAST(ctx *MenorMExpresionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitExpressionListAST(ctx *ExpressionListASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitPrimaryExpressionInAST(ctx *PrimaryExpressionInASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitPrimaryExpressionArAST(ctx *PrimaryExpressionArASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitPrimaryExpressionAppendAST(ctx *PrimaryExpressionAppendASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitPrimaryExpressionSeAST(ctx *PrimaryExpressionSeASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitPrimaryExpressionCapEAST(ctx *PrimaryExpressionCapEASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitPrimaryExpressionOpAST(ctx *PrimaryExpressionOpASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitPrimaryExpressionLengAST(ctx *PrimaryExpressionLengASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitOperandLitAST(ctx *OperandLitASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitOperandIDAST(ctx *OperandIDASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitOperandPARAST(ctx *OperandPARASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitLiteralIntAST(ctx *LiteralIntASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitLiteralFloatAST(ctx *LiteralFloatASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitLiteralRuneAST(ctx *LiteralRuneASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitLiteralRawsAST(ctx *LiteralRawsASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitLiteralInterAST(ctx *LiteralInterASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitIndexAST(ctx *IndexASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitArgumentsAST(ctx *ArgumentsASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitSelectorAST(ctx *SelectorASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitAppendExpressionAST(ctx *AppendExpressionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitLengthExpressionAST(ctx *LengthExpressionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitCapExpressionAST(ctx *CapExpressionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitStatementListAST(ctx *StatementListASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitBlockAST(ctx *BlockASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitStatementAST(ctx *StatementASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitPrintlStAST(ctx *PrintlStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitReturnStAST(ctx *ReturnStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitBreakStAST(ctx *BreakStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitContinueStAST(ctx *ContinueStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitSimpleStatementStAST(ctx *SimpleStatementStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitBlockStAST(ctx *BlockStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitSwitchStAST(ctx *SwitchStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitIfStatementStAST(ctx *IfStatementStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitLoopStAST(ctx *LoopStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitTypeDeclStAST(ctx *TypeDeclStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitVariableDeclStAST(ctx *VariableDeclStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitSimpleStatementAST(ctx *SimpleStatementASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitExpressionSimpStAST(ctx *ExpressionSimpStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitAssignmentStatementSimpStAST(ctx *AssignmentStatementSimpStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitAssignSimpStAST(ctx *AssignSimpStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitAssignmentStatementAST(ctx *AssignmentStatementASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitIncremAsSTAST(ctx *IncremAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitAndCompAsSTAST(ctx *AndCompAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitDecremeAsSTAST(ctx *DecremeAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitAssignNorAsSTAST(ctx *AssignNorAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitAssignMultAsSTAST(ctx *AssignMultAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitAssingnXorAsSTAST(ctx *AssingnXorAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitAssignMaAsSTAST(ctx *AssignMaAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitAssignMeAsSTAST(ctx *AssignMeAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitAssignBitAsSTAST(ctx *AssignBitAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitAssignPorAsSTAST(ctx *AssignPorAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitAssignDivAsSTAST(ctx *AssignDivAsSTASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitIfStatementAST(ctx *IfStatementASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitIfElseStAST(ctx *IfElseStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitIfElseBlockStAST(ctx *IfElseBlockStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitIfElseExpressionBlAST(ctx *IfElseExpressionBlASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitIfElseExpIfstAST(ctx *IfElseExpIfstASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitIfElseExpBlockBAST(ctx *IfElseExpBlockBASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitForBlLoopAST(ctx *ForBlLoopASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitForExpBlLoopAST(ctx *ForExpBlLoopASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitForSimpStEXLoopAST(ctx *ForSimpStEXLoopASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitForSimpStSimpStLoopAST(ctx *ForSimpStSimpStLoopASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitSwitchAST(ctx *SwitchASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitSwitchExprAST(ctx *SwitchExprASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitSwitchSimpStAST(ctx *SwitchSimpStASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitSwitchExCaseAST(ctx *SwitchExCaseASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitExpressionCaseClauseListAST(ctx *ExpressionCaseClauseListASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitExpressionCaseClauseExpAST(ctx *ExpressionCaseClauseExpASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitExpressionCaseClauseAST(ctx *ExpressionCaseClauseASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitExpressionSwitchCaseAST(ctx *ExpressionSwitchCaseASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniGoParserVisitor) VisitDefaultExpressionSAST(ctx *DefaultExpressionSASTContext) interface{} {
	return v.VisitChildren(ctx)
}
