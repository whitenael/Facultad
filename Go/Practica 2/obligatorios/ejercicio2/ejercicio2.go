package main

func main() {

	bc := CrearBlockchain()

	admin := CrearWallet("Genesis", "admin", "admin")
	//wallet1 := CrearWallet("1", "Jesse", "Pinkman")
	//wallet2 := CrearWallet("2", "Walter", "White")

	admin.EnviarTransaccion(10000, "1", bc)
	admin.EnviarTransaccion(10000, "2", bc)

	// wallet1.EnviarTransaccion(225.2, "2", bc)
	// wallet2.EnviarTransaccion(300.0, "1", bc)

}
