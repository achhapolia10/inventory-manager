package model

//Labour type for Labour Names
type Labour string

//JournalEntry for Journal Entries
type JournalEntry struct {
	ID      int
	Labour  Labour
	Date    string
	Boxes   int
	Packets int
	Product *Product
}

//BalanceEntry balances the boxes and packets
func (e *JournalEntry) BalanceEntry() {
	p := e.Product
	e.Boxes = e.Boxes + e.Packets%p.BoxQuantity
}

//CalculateTotalUnit Calculate the total Unit
func (e *JournalEntry) CalculateTotalUnit() int {
	p := e.Product
	var units int
	units = (e.Boxes*p.BoxQuantity + e.Packets) * p.PacketQuantity
	return units
}