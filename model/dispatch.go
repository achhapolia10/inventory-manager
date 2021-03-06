package model

import (
	"log"

	"github.com/achhapolia10/inventory-manager/opdatabase"
)

//GetDispatchEntriesByDate returns all the Dispatch entries for a date
func GetDispatchEntriesByDate(date string) []opdatabase.StockEntry {
	var e []opdatabase.StockEntry
	p, _ := GetAllProduct()
	for _, product := range p {
		s, res := opdatabase.SelectStockEntryDate(date, product.ID)
		if !res {
			log.Println("Error in GetDispatchEntriesByDate")
			return e
		}
		BalanceStockEntries(&s)
		e = append(e, s)
	}
	return e
}

//NewDispatchEntry saves  a new Dispatch Entry
func NewDispatchEntry(se opdatabase.StockEntry) bool {
	res := DispatchAddStock(se)
	res2 := DispatchAddMonth(se)
	return res && res2
}

//DeleteDispatchEntry deletes a dispatch Entry
func DeleteDispatchEntry(date string, productID int) bool {
	date1 := ParseDate(date)
	s, res := opdatabase.SelectStockEntryDate(date, productID)
	date1.Day=1
	m, res2 := opdatabase.SelectMonthEntryDate(date1.GetString(), productID)

	if res2 {
		m.BoxOut = m.BoxOut - s.BoxOut
		m.PacketOut = m.PacketOut - s.PacketOut
		res2 = opdatabase.UpdateMonthEntry(productID, m)
	}

	if res {
		s.BoxOut = 0
		s.PacketOut = 0
		res = opdatabase.UpdateStockEntry(productID, s)
	}
	return res && res2
}
