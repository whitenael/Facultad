package main

type Wallet struct {
	ID        string
	FirstName string
	LastName  string
}

func CrearWallet(id, firstName, lastName string) Wallet {
	return Wallet{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
	}
}
