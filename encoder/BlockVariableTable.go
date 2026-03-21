package encoder

import (
	"errors"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

// BlockVariableTable define una estructura para almacenar nombres de variables y valores por bloque
type BlockVariableTable struct {
	table map[*ir.Block]map[string]value.Value
}

// NewBlockVariableTable crea una nueva instancia de BlockVariableTable
func NewBlockVariableTable() *BlockVariableTable {
	return &BlockVariableTable{
		table: make(map[*ir.Block]map[string]value.Value),
	}
}

// AddVariable agrega una variable a un bloque específico
func (bvt *BlockVariableTable) AddVariable(block *ir.Block, name string, val value.Value) {
	if _, exists := bvt.table[block]; !exists {
		bvt.table[block] = make(map[string]value.Value)
	}
	bvt.table[block][name] = val
}

// GetVariable obtiene el valor de una variable en un bloque específico
func (bvt *BlockVariableTable) GetVariable(block *ir.Block, name string) (value.Value, bool) {
	if vars, exists := bvt.table[block]; exists {
		val, found := vars[name]
		return val, found
	}
	return nil, false
}

// FindValue busca una variable por nombre en todos los bloques
func (bvt *BlockVariableTable) FindValue(name string) (value.Value, bool) {
	for _, vars := range bvt.table {
		if val, found := vars[name]; found {
			return val, true
		}
	}
	return nil, false
}

// RemoveVariable elimina una variable de un bloque específico
func (bvt *BlockVariableTable) RemoveVariable(block *ir.Block, name string) {
	if vars, exists := bvt.table[block]; exists {
		delete(vars, name)
	}
}

// ListVariables lista todas las variables en un bloque específico
func (bvt *BlockVariableTable) ListVariables(block *ir.Block) map[string]value.Value {
	if vars, exists := bvt.table[block]; exists {
		return vars
	}
	return nil
}

// ClearBlock elimina todas las variables de un bloque específico
func (bvt *BlockVariableTable) ClearBlock(block *ir.Block) {
	delete(bvt.table, block)
}

// SetVariableValue busca una variable por nombre y le asigna un nuevo valor
func (bvt *BlockVariableTable) SetVariableValue(block *ir.Block, name string, newVal value.Value) error {
	// Verificar si el bloque existe en la tabla
	if vars, exists := bvt.table[block]; exists {
		// Verificar si la variable existe en el bloque
		if _, varExists := vars[name]; varExists {
			// Asignar el nuevo valor a la variable
			vars[name] = newVal
			return nil
		} else {
			return errors.New("variable no encontrada en el bloque")
		}
	} else {
		return errors.New("bloque no encontrado")
	}
}

// GetBlockNumberWithVariable retorna el número de bloque donde se encuentra una variable
func (bvt *BlockVariableTable) GetBlockNumberWithVariable(variableName string) (int, error) {
	// Inicializamos un contador para el número de bloque
	blockNumber := 0
	// Recorremos la tabla de variables por bloque
	for block := range bvt.table {
		// Incrementamos el contador de bloque
		blockNumber++
		// Verificamos si la variable existe en el bloque actual
		if _, exists := bvt.table[block][variableName]; exists {
			return blockNumber, nil
		}
	}
	// Si no se encuentra la variable en ningún bloque
	return -1, errors.New("variable no encontrada en ningún bloque")
}