package main

import "fmt"

type PuntoCardinal string

const (
	N  PuntoCardinal = "N"
	S  PuntoCardinal = "S"
	E  PuntoCardinal = "E"
	O  PuntoCardinal = "O"
	NE PuntoCardinal = "NE"
	NO PuntoCardinal = "NO"
	SE PuntoCardinal = "SE"
	SO PuntoCardinal = "SO"
)

// Función para determinar hacia dónde se dirige el viento
func DireccionDelViento(cardinal PuntoCardinal) PuntoCardinal {

	if cardinal == "" {
		return ""
	}

	// Lógica para determinar hacia dónde se dirige el viento
	switch cardinal {
	case N:
		return S
	case S:
		return N
	case E:
		return O
	case O:
		return E
	case NO:
		return SE
	case SE:
		return NO
	case NE:
		return SO
	case SO:
		return NE
	default:
		return "" // Valor inválido
	}
}

func main() {
	var cardinal PuntoCardinal
	fmt.Print("Ingrese el punto cardinal del cual viene el viento (N, S, E, O, NO, SE, NE, SO): ")
	fmt.Scan(&cardinal)

	// Determinar hacia dónde se dirige el viento
	direction := DireccionDelViento(cardinal)

	// Imprimir resultado
	fmt.Printf("El viento se dirige hacia %v.\n", direction)
}
