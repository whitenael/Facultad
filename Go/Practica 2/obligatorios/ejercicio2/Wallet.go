package main

import "fmt"

type Wallet struct {
	ID        string
	FirstName string
	LastName  string
}

// para enviar una transaccion, pedimos el monto a enviar, el ID de la persona que la va a recibir, y la blockchain a la que se va a enviar
func (wallet Wallet) EnviarTransaccion(amount float64, receiverID string, bc Blockchain) {

	// vemos que saldo tiene el usuario dentro de la blockchain
	// de no tener saldo suficiente, le devolvemos una excepcion informandole
	if bc.ObtenerSaldo(wallet) < amount {
		panic(fmt.Sprintf("No se puede ejecutar la transaccion: %s -> %s ", wallet.ID, receiverID))
	}

	// caso contrario, agregamos la transaccion a un bloque que luego formara parte de la blockchain
	bc.AgregarBloque(NewTransaction(amount, wallet.ID, receiverID))
}

func CrearWallet(id, firstName, lastName string) Wallet {
	return Wallet{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
	}
}
