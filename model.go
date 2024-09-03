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
	panic("fix me")
}

func (m *Member) GetData() any {
	panic("fix me")
}

type Item struct {
	SKU          string
	ItemName     string
	StockQty     int32
	Transactions []Transaction
	Price        int32
}

func (it *Item) AddTransaction(data any) {
	panic("fix me")
}

func (it *Item) GetData() any {
	panic("fix me")
}
