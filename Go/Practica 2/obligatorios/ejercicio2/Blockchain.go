package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// structs

// representacion del bloque (de momento, solo de a una transaccion a la vez)
type Block struct {
	Hash         string
	PreviousHash string
	Transaction  Transaction
	Timestamp    time.Time
}

// representacion de la blockchain como un slice de bloques
type Blockchain struct {
	Blocks []Block
}

// calculo del hash de un bloque
func CalcularHash(block Block) string {
	// creamos un hash a partir de los datos del bloque para que sea unico
	record := block.PreviousHash + fmt.Sprintf("%f", block.Transaction.Amount) + block.Transaction.SenderID + block.Transaction.ReceiverID + block.Timestamp.String()
	// usamos la implementacion de sha256 de hashing criptografico para convertirlo en una cadena de longitud fija
	h := sha256.New()
	// usamos el record creado anteriormente para calcular el hash
	// se convierte el registro en un slice de bytes y se escribe en el hash h. esto le proporciona los datos
	// necesarios al sha256 que luego realiza el calculo del hash
	h.Write([]byte(record))
	// se llama a Sum para obtener el resultado del hash calculado en h. el parametro nil es para
	// indicarle que lo queremos en un slice de bytes
	hashed := h.Sum(nil)

	// pasamos el slice de bytes a string para retornarlo
	return hex.EncodeToString(hashed)
}

// creacion del bloque genesis

// Descripcion de un bloque genesis:
// El bloque génesis es la plataforma sobre la cual se asientan los subsiguientes bloques que conforman un registro de blockchain,
// permitiendo así el funcionamiento de una criptodivisa. En otras palabras, el bloque génesis hace posible la creación de una criptomoneda.
func CrearBloqueGenesis() Block {
	transaction := Transaction{100000000, "coin", "Genesis", time.Now()}
	block := Block{
		PreviousHash: "0000000000000000000000000000000000000000000000000000000000000000",
		Transaction:  transaction,
		Timestamp:    time.Now(),
	}

	block.Hash = CalcularHash(block)

	return block
}

// creacion de una blockchain a partir de su bloque genesis
func CrearBlockchain() Blockchain {
	genesisBlock := CrearBloqueGenesis()
	return Blockchain{[]Block{genesisBlock}}
}

// creacion de un bloque a partir de los datos del ultimo bloque de la blockchain de una transaccion
func CrearBloque(prevBlock Block, transaction Transaction) Block {
	block := Block{
		PreviousHash: prevBlock.Hash,
		Transaction:  transaction,
		Timestamp:    time.Now(),
	}
	block.Hash = CalcularHash(block)
	return block
}

// metodo para agregar un bloque a la blockchain a partir de una transaccion
func (bc *Blockchain) AgregarBloque(transaction Transaction) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]        // accedemos al ultimo bloque actual
	newBlock := CrearBloque(prevBlock, transaction) // creamos un bloque con los datos de la transaccion y su bloque anterior para mantener la integridad de la blockchain
	bc.Blocks = append(bc.Blocks, newBlock)         // agregamos el bloque a la blockchain
}

// Metodo que nos devuelve el saldo de nuestra billera dentro del blockchain
func (bc Blockchain) ObtenerSaldo(wallet Wallet) float64 {
	var saldo float64
	for i := 0; i < len(bc.Blocks); i++ {
		// si enviamos dinero en esta chain, restamos del saldo de nuestra wallet
		if bc.Blocks[i].Transaction.SenderID == wallet.ID {
			saldo -= bc.Blocks[i].Transaction.Amount
		}

		// si recibimos dinero en esta chain, sumamos saldo a nuestra wallet
		if bc.Blocks[i].Transaction.ReceiverID == wallet.ID {
			saldo += bc.Blocks[i].Transaction.Amount
		}
	}

	return saldo
}

// para validar una blockchain hay que asegurarse de que cada bloque tiene el hash correcto del bloque anterior
func (bc Blockchain) Validar() bool {

	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		previousBlock := bc.Blocks[i-1]

		// recalculamos el hash anterior
		// recalculamos para validar que el hash verdaderamente hace referencia al bloque referenciado
		// esto se hace para verificar que su contenido no ha sido alterado
		recalculatedHash := CalcularHash(previousBlock)

		// si el bloque actual tiene un hash diferente al recalculado, retornamos false
		if currentBlock.PreviousHash != recalculatedHash {
			return false
		}
	}

	return true

}
