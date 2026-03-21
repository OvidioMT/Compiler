parser grammar MiniGoParser;


//import the lexer tokens
options {
    tokenVocab = MiniGoScanner;

}


root			: PACKAGE  ID PyCOMA topDeclarationList EOF                                                             #rootAST
            ;
topDeclarationList	: ( variableDecl | typeDecl | funcDecl)*                                                            #topDeclarationListAST
            ;
variableDecl		: VAR singleVarDecl PyCOMA                                                                          #variableDeclSingAST
			| VAR  PARIZQ innerVarDecls PARDER PyCOMA                                                                   #variableDeclPaRAST
			| VAR  PARIZQ PARDER PyCOMA                                                                                 #variableDeclNullAST
			;
innerVarDecls		: singleVarDecl PyCOMA (singleVarDecl PyCOMA)*                                                      #innerVarDeclsAST
            ;
singleVarDecl		: identifierList declType IGUAL expressionList                                                      #singleVarDeclIdenAST
			| identifierList IGUAL expressionList                                                                       #singleVarDeclIAST
			| singleVarDeclNoExps                                                                                       #singleVarDeclNoExpsSVDAST
			;
singleVarDeclNoExps	: identifierList declType                                                                           #singleVarDeclNoExpsAST
            ;
typeDecl		: TYPE singleTypeDecl PyCOMA                                                                            #typeDeclTypeAST
			| TYPE PARIZQ innerTypeDecls PARDER PyCOMA                                                                  #typeDeclPARAST
			| TYPE PARIZQ PARDER PyCOMA                                                                                 #typeDeclNULPARAST
			;
innerTypeDecls		: singleTypeDecl PyCOMA (singleTypeDecl PyCOMA)*                                                    #innerTypeDeclsAST
            ;
singleTypeDecl		: ID declType                                                                                       #singleTypeDeclAST
             ;
funcDecl		: funcFrontDecl block PyCOMA                                                                            #funcDeclAST
            ;
funcFrontDecl		: FUNC ID PARIZQ (funcArgDecls)? PARDER (declType)?                                                 #funcFrontDeclAST
            ;
funcArgDecls		: singleVarDeclNoExps (COMA singleVarDeclNoExps)*                                                   #funcArgDeclsAST
            ;
declType		: PARIZQ declType PARDER                                                                                #declTypeParAST
			| ID                                                                                                        #declTypeIDAST
			| sliceDeclType                                                                                             #sliceDeclTypeDeclAST
			| arrayDeclType                                                                                             #arrayDeclTypeDeclAST
			| structDeclType                                                                                            #structDeclTypeDeckAST
			;
sliceDeclType		: CORCHIZQ CORCHDER declType                                                                        #sliceDeclTypeAST
            ;
arrayDeclType		: CORCHIZQ INT CORCHDER declType                                                                    #arrayDeclTypeAST
            ;
structDeclType		: STRUCT LLAVEIZQ (structMemDecls)? LLAVEDER                                                        #structDeclTypeAST
            ;
structMemDecls	: singleVarDeclNoExps PyCOMA (singleVarDeclNoExps PyCOMA)*                                              #structMemDeclsAST
            ;
identifierList		: ID (COMA ID)*                                                                                     #identifierListAST
            ;
expression		: primaryExpression                                                                                     #expressionAST
			| expression  MULTI expression                                                                              #multExpresionAST
			| expression DIV expression                                                                                 #divExpresionAST
			| expression PORCENT expression                                                                             #porcentExpresionAST
			| expression MAYORM expression                                                                              #mayorMExpresionAST
			| expression MEMORM expression                                                                              #menorMExpresionAST
			| expression AND expression                                                                                 #andExpresionAST
			| expression BITCLE expression                                                                              #bitExpresionAST
			| expression SUMA expression                                                                                #sumaExpresionAST
			| expression RESTA expression                                                                               #restaExpresionAST
			| expression OR expression                                                                                  #orExpresionAST
			| expression XOR expression                                                                                 #xorExpresionAST
			| expression IDENTICO expression                                                                            #identicoExpresionAST
			| expression DIFERENTE expression                                                                           #diferenteExpresionAST
			| expression MAYOR expression                                                                               #mayorExpresionAST
			| expression MAYORIG expression                                                                             #mayorIgExpresionAST
			| expression MENOR expression                                                                               #menorExpresionAST
			| expression MENORIG expression                                                                             #menorIgExpresionAST
			| expression ANDD expression                                                                                #anddExpresionAST
			| expression ORD expression                                                                                 #ordExpresionAST
			| SUMA expression                                                                                           #sumaExExpresionAST
			| RESTA expression                                                                                          #restaExExpresionAST
			| EXCL expression                                                                                           #exclExpresionAST
			| XOR expression                                                                                            #xorExExpresionAST
			;
expressionList		: expression (COMA expression)*#expressionListAST
            ;
primaryExpression	: operand                                                                                           #primaryExpressionOpAST
			| primaryExpression selector                                                                                #primaryExpressionSeAST
			| primaryExpression index                                                                                   #primaryExpressionInAST
			| primaryExpression arguments                                                                               #primaryExpressionArAST
			| appendExpression                                                                                          #primaryExpressionAppendAST
			| lengthExpression                                                                                          #primaryExpressionLengAST
			| capExpression                                                                                             #primaryExpressionCapEAST
			;
operand		: literal                                                                                                   #operandLitAST
			| ID                                                                                                        #operandIDAST
			| PARIZQ expression PARDER                                                                                  #operandPARAST
			;
literal			: INT                                                                                                   #literalIntAST
			| FLOAT                                                                                                     #literalFloatAST
			| RUNE                                                                                                      #literalRuneAST
			| RAWSTRING                                                                                                 #literalRawsAST
			| INTERPRETEDSTRING                                                                                         #literalInterAST
			;
index			: CORCHIZQ expression CORCHDER                                                                          #indexAST
            ;
arguments		: PARIZQ expressionList? PARDER                                                                         #argumentsAST
            ;
selector		: PUNTO ID                                                                                              #selectorAST
            ;
appendExpression	: APPEND PARIZQ expression COMA expression PARDER                                                   #appendExpressionAST
            ;
lengthExpression	: LEN PARIZQ expression PARDER                                                                      #lengthExpressionAST
            ;
capExpression		: CAP PARIZQ expression PARDER                                                                      #capExpressionAST
            ;
statementList 		: statement*                                                                                        #statementListAST
            ;
block 			: LLAVEIZQ statementList LLAVEDER                                                                       #blockAST
            ;
statement		: PRINT PARIZQ expressionList? PARDER PyCOMA                                                            #statementAST
			| PRINTLN PARIZQ expressionList? PARDER PyCOMA                                                              #printlStAST
			| RETURN (expression)? PyCOMA                                                                               #returnStAST
			| BREAK PyCOMA                                                                                              #breakStAST
			| CONTINUE PyCOMA                                                                                           #continueStAST
			| simpleStatement PyCOMA                                                                                    #simpleStatementStAST
			| block PyCOMA                                                                                              #blockStAST
			| switch PyCOMA                                                                                             #switchStAST
			| ifStatement PyCOMA                                                                                        #ifStatementStAST
			| loop PyCOMA                                                                                               #loopStAST
			| typeDecl                                                                                                  #typeDeclStAST
			| variableDecl                                                                                              #variableDeclStAST
			;
simpleStatement	:                                                                                                       #simpleStatementAST
			| expression (SUMAD | RESTAD)?                                                                              #expressionSimpStAST
			| assignmentStatement                                                                                       #assignmentStatementSimpStAST
			| expressionList ASSIGN expressionList                                                                      #assignSimpStAST
			;
assignmentStatement 	: expressionList IGUAL expressionList                                                           #assignmentStatementAST
			|expression INCREM expression                                                                               #incremAsSTAST
			|expression ANDCOMP expression                                                                              #andCompAsSTAST
            |expression DECREME expression                                                                              #decremeAsSTAST
			|expression ASSIGNOR expression                                                                             #assignNorAsSTAST
			|expression ASSIGNMULT expression                                                                           #assignMultAsSTAST
			|expression ASSIGNXOR expression                                                                            #assingnXorAsSTAST
			|expression ASSIGMA expression                                                                              #assignMaAsSTAST
			|expression ASSIGNME expression                                                                             #assignMeAsSTAST
			|expression ASSIGNBIT expression                                                                            #assignBitAsSTAST
			|expression ASSIGNPOR expression                                                                            #assignPorAsSTAST
			|expression ASSINGDIV expression                                                                            #assignDivAsSTAST
			;
ifStatement 		: IF expression block                                                                               #ifStatementAST
			| IF expression block ELSE ifStatement                                                                      #ifElseStAST
			| IF expression block ELSE block                                                                            #ifElseBlockStAST
			| IF simpleStatement  PyCOMA expression block                                                               #ifElseExpressionBlAST
			| IF simpleStatement PyCOMA  expression block ELSE ifStatement                                              #ifElseExpIfstAST
			| IF simpleStatement  PyCOMA expression block ELSE block                                                    #ifElseExpBlockBAST
			;
loop			: FOR block                                                                                             #forBlLoopAST
			| FOR expression block                                                                                      #forExpBlLoopAST
			| FOR simpleStatement PyCOMA expression PyCOMA simpleStatement block                                        #forSimpStEXLoopAST
			| FOR simpleStatement PyCOMA PyCOMA simpleStatement block                                                   #forSimpStSimpStLoopAST
			;
switch			: SWITCH simpleStatement PyCOMA expression LLAVEIZQ expressionCaseClauseList LLAVEDER                   #switchAST
			| SWITCH expression LLAVEIZQ expressionCaseClauseList LLAVEDER                                              #switchExprAST
			| SWITCH simpleStatement PyCOMA LLAVEIZQ expressionCaseClauseList LLAVEDER                                  #switchSimpStAST
			| SWITCH LLAVEIZQ expressionCaseClauseList LLAVEDER                                                         #switchExCaseAST
			;
expressionCaseClauseList :                                                                                              #expressionCaseClauseListAST
			| expressionCaseClause expressionCaseClauseList                                                             #expressionCaseClauseExpAST
			;
expressionCaseClause 	: expressionSwitchCase DOSP statementList                                                       #expressionCaseClauseAST
            ;
expressionSwitchCase : CASE expressionList                                                                              #expressionSwitchCaseAST
			| DEFAULT                                                                                                   #defaultExpressionSAST
			;
