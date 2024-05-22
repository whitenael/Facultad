package main

type Wallet struct {
	ID        string
	FirstName string
	LastName  string	
}

// para enviar una transaccion, pedimos el monto a enviar, el ID de la persona que la va a recibir, y la blockchain a la que se va a enviar
func (wallet Wallet) EnviarTransaccion(amount float64, receiverID string, bc Blockchain){

	// vemos que saldo tiene el usuario dentro de la blockchain
	// de no tener saldo suficiente, le devolvemos una excepcion informandole
	if (ObtenerSaldo(bc) < amount){
		panic("No se puede ejecutar la transacción. Saldo insuficiente")
	}
	
	// caso contrario, agregamos la transaccion a un bloque que luego formara parte de la blockchain
	bc.AgregarBloque(NewTransaction(amount, wallet.ID, receiverID))	
}

func (wallet Wallet) ObtenerSaldo(bc Blockchain) float64{
	var saldo float64
	for i := 0; i < len(bc); i++ {
		// si enviamos dinero en esta chain, restamos del saldo de nuestra wallet
		if (bc[i].Transaction.SenderID == wallet.ID){
			saldo -= bc[i].Transaction.Amount
		}

		// si recibimos dinero en esta chain, sumamos saldo a nuestra wallet
		if (bc[i].Transaction.ReceiverID == wallet.ID){
			saldo += bc[i].Transaction.Amount
		}
	}

	return saldo
}

func CrearWallet(id, firstName, lastName string) Wallet {
	return Wallet{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
	}
}
