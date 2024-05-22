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
	transaction := Transaction{0, "Genesis", "Genesis", time.Now()}
	return Block{
		Hash:        "0000000000000000000000000000000000000000000000000000000000000000",
		Transaction: transaction,
		Timestamp:   time.Now(),
	}
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
	prevBlock := bc.Blocks[len(bc.Blocks)-1] // accedemos al ultimo bloque actual
	newBlock := CrearBloque(prevBlock, transaction) // creamos un bloque con los datos de la transaccion y su bloque anterior para mantener la integridad de la blockchain
	bc.Blocks = append(bc.Blocks, newBlock) // agregamos el bloque a la blockchain
}
