package main

import "errors"

var (
	Name      string = "Arya Lesmana" // please insert your name here
	IdStudent string = "2206081603"   // please insert your id student here
	Items     []Tool                  // contain array of item pointer
	Members   []Tool                  // contain array of member pointer
)

func AddItem(SKU string, itemName string, price int32, stockQty int32) (string, error) {
	for i := 0; i < len(Items); i++ {
		if Items[i].GetData().(Item).SKU == SKU {
			return "", errors.New("item " + SKU + " is already in list of items")
		}
	}
	return "successfully added item " + SKU + " to list of items", nil
}

func DeleteItem(SKU string) (string, error) {
	for i := 0; i < len(Items); i++ {
		if Items[i].GetData().(Item).SKU == SKU {
			if len(Items[i].GetData().(Item).Transactions) == 0 {
				Items[i] = Items[len(Items)-1]
				Items = Items[:len(Items)-1]
				return "successfully deleted item " + SKU + " from list of items", nil
			} else {
				return "", errors.New("there is at least one transaction taking item " + SKU)
			}
		}
	}
	return "", errors.New("item " + SKU + " is not in list of items")
}

func AddMember(idMember string, memberName string) (string, error) {
	for i := 0; i < len(Members); i++ {
		if Members[i].GetData().(Member).IdMember == idMember {
			return "", errors.New("member " + idMember + " is already in list of members")
		}
	}
	return "successfully added member " + idMember + " to list of members", nil
}

func DeleteMember(idMember string) (string, error) {
	for i := 0; i < len(Members); i++ {
		if Members[i].GetData().(Member).IdMember == idMember {
			if len(Members[i].GetData().(Member).Transactions) == 0 {
				Members[i] = Members[len(Members)-1]
				Members = Members[:len(Members)-1]
				return "successfully deleted member " + idMember + " from list of members", nil
			} else {
				return "", errors.New("there is at least one transaction taking member " + idMember)
			}
		}
	}
	return "", errors.New("member " + idMember + " is not in list of members")
}

func AddTransaction(qty int32, data ...string) (string, error) {
	if len(data) == 2 {
		for i := 0; i < len(Members); i++ {
			if Members[i].GetData().(Member).IdMember == data[1] {
				for j := 0; j < len(Items); j++ {
					if Items[j].GetData().(Item).SKU == data[0] {
						if Items[j].GetData().(Item).StockQty >= qty {
							var transaksi Transaction
							transaksi.SKU = data[0]
							transaksi.IdMember = &data[1]
							transaksi.Price = Items[j].GetData().(Item).Price
							transaksi.Qty = qty

							item := Items[j].GetData().(Item)
							item.Transactions = append(item.Transactions, transaksi)
							Items[j] = &item

							member := Members[i].GetData().(Member)
							member.Transactions = append(member.Transactions, transaksi)
							Members[i] = &member

							return "successfully added transaction item " + Items[j].GetData().(Item).SKU + " for member " + Members[i].GetData().(Member).IdMember, nil
						} else {
							return "", errors.New("stock qty for item " + Items[j].GetData().(Item).SKU + " is not sufficient")	
						}
					}
				}
				return "", errors.New("item " + data[0] + " is not in list of items")
			}
		}
		return "", errors.New("member " + data[1] + " is not in list of members")
	} else {
		for i := 0; i < len(Items); i++ {
			if Items[i].GetData().(Item).SKU == data[0] {
				if Items[i].GetData().(Item).StockQty >= qty {
					var transaksi Transaction
					transaksi.SKU = data[0]
					transaksi.Price = Items[i].GetData().(Item).Price
					transaksi.Qty = qty

					item := Items[i].GetData().(Item)
					item.Transactions = append(item.Transactions, transaksi)
					Items[i] = &item

					return "successfully added transaction item " + Items[i].GetData().(Item).SKU, nil
				} else {
					return "", errors.New("stock qty for item " + Items[i].GetData().(Item).SKU + " is not sufficient")	
				}
			}
		}
		return "", errors.New("item " + data[0] + " is not in list of items")
	}
}

func RestockItem(SKU string, qty int32) (string, error) {
	for i := 0; i < len(Items); i++ {
		if Items[i].GetData().(Item).SKU == SKU {
			item := Items[i].GetData().(Item)
			item.StockQty += qty
			Items[i] = &item
			return "successfully restocked qty for item " + SKU, nil
		}
	}
	return "", errors.New("item " + SKU + " is not in list of items")
}

func GetTransactionItem(SKU string) ([]Transaction, error) {
	for i := 0; i < len(Items); i++ {
		if Items[i].GetData().(Item).SKU == SKU {
			return Items[i].GetData().(Item).Transactions, nil
		}
	}
	return nil, errors.New("item " + SKU + " is not in list of items")
}

func GetTransactionMember(idMember string) ([]Transaction, error) {
	for i := 0; i < len(Members); i++ {
		if Members[i].GetData().(Member).IdMember == idMember {
			return Members[i].GetData().(Member).Transactions, nil
		}
	}
	return nil, errors.New("member " + idMember + " is not in list of members")
}
