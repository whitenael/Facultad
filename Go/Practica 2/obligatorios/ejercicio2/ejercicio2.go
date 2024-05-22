package main

func main() {

	bc := CrearBlockchain()

	wallet1 := CrearWallet(1, "Jesse", "Pinkman")
	wallet2 := CrearWallet(2, "Walter", "White")
	
	wallet1.EnviarTransaccion(23000.1, 2).SubmitTransaction(&bc) // creamos una transaccion y la agregamos al blockchain
	

}
