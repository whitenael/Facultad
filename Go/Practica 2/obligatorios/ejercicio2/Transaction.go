package main

// representacion de una transaccion
type Transaction struct {
	Amount     float64
	SenderID   string
	ReceiverID string
	Timestamp  time.Time
}

func NewTransaction(amount float64, senderID, receiverID string) Transaction {
	return Transaction{amount, senderID, receiverID, time.Now()}
}

func (transaction Transaction) SubmitTransaction(bc *Blockchain){
	bc.AgregarBloque(transaction)
}


