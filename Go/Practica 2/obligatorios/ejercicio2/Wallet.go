package main

type Wallet struct {
	ID        string
	FirstName string
	LastName  string	
}

func (wallet Wallet) EnviarTransaccion(amount float64, receiverID string) Transaction{
	return NewTransaction(amount, wallet.ID, receiverID)
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
