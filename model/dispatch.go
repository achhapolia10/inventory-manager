package model

import (
	"log"

	"github.com/achhapolia10/anjaniv2/opdatabase"
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
		e = append(e, s)
	}
	return e
}

//NewDispatchEntry saves  a new Dispatch Entry
func NewDispatchEntry(se opdatabase.StockEntry) bool {
	res := DispatchAddStock(se)
	return res
}

//DeleteDispatchEntry deletes a dispatch Entry
func DeleteDispatchEntry(se opdatabase.StockEntry) bool {
	res := DispatchDeleteStock(se)
	return res
}
