package checker

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
)

type Ident interface {
}

type VarIdent struct {
	tok        antlr.Token // token
	typ        int         // tipo
	level      int         // nivel
	isConstant bool
}

type MethodIdent struct {
	tok    antlr.Token
	typ    int
	level  int
	params []int
}

type SymbolTable struct {
	table        []Ident // lista de identificadores
	currentLevel int     // nivel actual
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		table:        []Ident{},
		currentLevel: 0,
	}
}

func (t *SymbolTable) insert(id antlr.Token, tp int) {
	v := &VarIdent{
		tok:        id,
		typ:        tp,
		level:      t.currentLevel,
		isConstant: false,
	}
	t.table = append([]Ident{v}, t.table...) // se inserta al inicio
}

func (t *SymbolTable) insertWithParams(id antlr.Token, tipo int, p []int) {
	m := &MethodIdent{
		tok:    id,
		typ:    tipo,
		level:  t.currentLevel,
		params: p,
	}
	t.table = append([]Ident{m}, t.table...) // se inserta al inicio
}

func (t *SymbolTable) search(name string) Ident {
	for _, id := range t.table {
		if id.(*VarIdent).tok.GetText() == name { //|| id.(*MethodIdent).tok.GetText() == name {
			return id // se encontró
		}
	}
	return nil
}

func (t *SymbolTable) searchInCurrentLevel(name string) Ident {
	var temp Ident
	tempLevel := t.currentLevel // se guarda el nivel actual para comparar
	for _, id := range t.table {
		if tempLevel == id.(*VarIdent).level { //|| tempLevel == id.(*MethodIdent).level {
			if id.(*VarIdent).tok.GetText() == name { // || id.(*MethodIdent).tok.GetText() == name {
				temp = id // se encontró en el nivel actual
			}
		} else {

			break // se termina la búsqueda
		}
	}
	return temp
}

func (t *SymbolTable) openScope() { // abre un nuevo nivel en la tabla
	t.currentLevel++
}

func (t *SymbolTable) closeScope() { // cierra el nivel actual en la tabla
	var newTabla []Ident
	for _, ident := range t.table {
		if ident.(*VarIdent).level != t.currentLevel { //&& ident.(*MethodIdent).level != t.currentLevel {
			newTabla = append(newTabla, ident) // se copian los identificadores
		}
	}
	t.table = newTabla // se actualiza la tabla
	t.currentLevel--   // se disminuye el nivel
}

func (t *SymbolTable) printTable() {
	fmt.Println("---------- START TABLE ----------")
	for _, ident := range t.table {
		if v, ok := ident.(*VarIdent); ok {
			fmt.Printf("Nombre: %s , 	n| %d , 	t| %d\n", v.tok.GetText(), v.level, v.typ)
		} //else if m, ok := ident.(*MethodIdent); ok {
		//fmt.Printf("NAME: %s |-| %d |-| %d\n", m.tok.GetText(), m.level, m.typ)
		//}
	}
	fmt.Println("----------   End   ----------")
}
