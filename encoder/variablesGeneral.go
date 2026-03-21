package encoder

import "fmt"

// Definición de la estructura variablesGeneral
type variablesGeneral struct {
	nombre string
}

func (v variablesGeneral) ObtenerNombre() string {
	return v.nombre
}

// Definición de la lista de variablesGeneral
type listaVariablesGeneral struct {
	elementos []variablesGeneral
}

// Función para obtener el primer elemento de la lista
func (lvg *listaVariablesGeneral) obtenerPrimerElemento() (variablesGeneral, error) {
	if len(lvg.elementos) == 0 {
		return variablesGeneral{}, fmt.Errorf("la lista está vacía")
	}
	return lvg.elementos[0], nil
}

// Función para agregar un elemento al inicio de la lista
func (l *listaVariablesGeneral) agregarAlInicio(nombre string) {
	nuevaVariable := variablesGeneral{nombre: nombre}
	l.elementos = append([]variablesGeneral{nuevaVariable}, l.elementos...)
}

// Función para obtener un elemento por nombre
func (l *listaVariablesGeneral) obtenerPorNombre(nombre string) *variablesGeneral {
	for _, variable := range l.elementos {
		if variable.nombre == nombre {
			return &variable
		}
	}
	return nil
}

// Función para eliminar un elemento por nombre
func (l *listaVariablesGeneral) eliminarPorNombre(nombre string) bool {
	for i, variable := range l.elementos {
		if variable.nombre == nombre {
			l.elementos = append(l.elementos[:i], l.elementos[i+1:]...)
			return true
		}
	}
	return false
}
