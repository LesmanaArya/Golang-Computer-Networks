package main

type Tool interface {
	AddTransaction(data any)
	GetData() any
}

type Transaction struct {
	IdMember *string
	SKU      string
	Qty      int32
	Price    int32
}

type Member struct {
	IdMember     string
	MemberName   string
	Transactions []Transaction
}

func (m *Member) AddTransaction(data any) {
	transaksi_baru := data.(Transaction)
	if m.IdMember == *transaksi_baru.IdMember {
		for i := 0; i < len(m.Transactions); i++ {
			if m.Transactions[i].SKU == transaksi_baru.SKU {
				m.Transactions[i].Qty += transaksi_baru.Qty
				return
			}
		}
		m.Transactions = append(m.Transactions, transaksi_baru)
		return
	} else {
		m.Transactions = append(m.Transactions, transaksi_baru)
		return
	}
}

func (m *Member) GetData() any {
	return any(m)
}

type Item struct {
	SKU          string
	ItemName     string
	StockQty     int32
	Transactions []Transaction
	Price        int32
}

func (it *Item) AddTransaction(data any) {
	transaksi_baru := data.(Transaction)
	if it.SKU == transaksi_baru.SKU {
		for i := 0; i < len(it.Transactions); i++ {
			if it.Transactions[i].IdMember == transaksi_baru.IdMember {
				it.Transactions[i].Qty += transaksi_baru.Qty
				return
			}
		}
		it.Transactions = append(it.Transactions, transaksi_baru)
		return
	} else {
		it.Transactions = append(it.Transactions, transaksi_baru)
		return
	}
}

func (it *Item) GetData() any {
	return any(it)
}
