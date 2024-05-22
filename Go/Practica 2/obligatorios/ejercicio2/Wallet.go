package main

type Wallet struct {
	ID        string
	FirstName string
	LastName  string
}

func (wallet Wallet) EnviarTransaccion(amount float64, receiverID string) Transaction{
	return NewTransaction(amount, wallet.ID, receiverID)
}



func CrearWallet(id, firstName, lastName string) Wallet {
	return Wallet{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
	}
}
