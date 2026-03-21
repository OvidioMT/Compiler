package encoder

import "errors"


type arreglos struct {
	nombreVariable string
	Numeros        []int
}

// Función para crear un nuevo objeto de la estructura
func nuevoArreglo(nombre string, numeros []int) arreglos {
	return arreglos{
		nombreVariable: nombre,
		Numeros:        numeros,
	}
}

// Función para agregar un número a la lista
func (m *arreglos) AgregarNumero(numero int) {
	m.Numeros = append(m.Numeros, numero)

}

// Función para quitar un número de la lista por posición
func (m *arreglos) QuitarNumero(posicion int) error {
	if posicion < 0 || posicion >= len(m.Numeros) {
		return errors.New("posición fuera de rango")
	}
	m.Numeros = append(m.Numeros[:posicion], m.Numeros[posicion+1:]...)
	return nil
}

// Función para obtener un número de la lista por posición
func (m *arreglos) ObtenerNumero(posicion int) int {
	if posicion < 0 || posicion >= len(m.Numeros) {
		return 0
	}
	return m.Numeros[posicion]
}

// Función para obtener el valor del string Texto
func (m *arreglos) ObtenerTexto() string {
	return m.nombreVariable
}

// Lista de arreglosList
type ListaArreglos struct {
	Elementos []arreglos
}

func NuevaListaArreglos() *ListaArreglos {
	return &ListaArreglos{Elementos: []arreglos{}}
}

// Función para agregar un nuevo elemento al inicio de la lista
func (l *ListaArreglos) AgregarAlInicio(nuevoElemento arreglos) {
	l.Elementos = append([]arreglos{nuevoElemento}, l.Elementos...)
}

// Función para agregar un número a la lista de un elemento específico usando su Texto
func (l *ListaArreglos) AgregarNumeroPorTexto(texto string, numero int) error {
	for i, elem := range l.Elementos {
		if elem.nombreVariable == texto {
			l.Elementos[i].AgregarNumero(numero)
			return nil
		}
	}
	return errors.New("elemento no encontrado")
}

func (l *ListaArreglos) ObtenerPorNombre(nombre string) *arreglos {
	for _, elem := range l.Elementos {
		if elem.nombreVariable == nombre {
			return &elem
		}
	}
	return nil
}