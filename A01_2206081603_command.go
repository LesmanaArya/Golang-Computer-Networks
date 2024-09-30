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
		item, ok := Items[i].GetData().(*Item)
            if !ok {
                return "", errors.New("Item type assertion failed")
            }
		if item.SKU == SKU {
			return "", errors.New("item " + SKU + " is already in list of items")
		}
	}
	new_item := Item{SKU: SKU, ItemName: itemName, Price: price, StockQty: stockQty}
	Items = append(Items, &new_item)
	return "successfully added item " + SKU + " to list of items", nil
}

func DeleteItem(SKU string) (string, error) {
	for i := 0; i < len(Items); i++ {
		item, ok := Items[i].GetData().(*Item)
            if !ok {
                return "", errors.New("Item type assertion failed")
            }
		if item.SKU == SKU {
			if len(item.Transactions) == 0 {
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
		member, ok := Members[i].GetData().(*Member)
            if !ok {
                return "", errors.New("Member type assertion failed")
            }
		if member.IdMember == idMember {
			return "", errors.New("member " + idMember + " is already in list of members")
		}
	}
	new_member := Member{IdMember: idMember, MemberName: memberName}
	Members = append(Members, &new_member)
	return "successfully added member " + idMember + " to list of members", nil
}

func DeleteMember(idMember string) (string, error) {
	for i := 0; i < len(Members); i++ {
		member, ok := Members[i].GetData().(*Member)
            if !ok {
                return "", errors.New("Member type assertion failed")
            }
		if member.IdMember == idMember {
			if len(member.Transactions) == 0 {
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
            member, ok := Members[i].GetData().(*Member)
            if !ok {
                return "", errors.New("Member type assertion failed")
            }
            if member.IdMember == data[1] {
                for j := 0; j < len(Items); j++ {
                    item, ok := Items[j].GetData().(*Item)
                    if !ok {
                        return "", errors.New("Item type assertion failed")
                    }
                    if item.SKU == data[0] {
                        if item.StockQty >= qty {
                            var transaksi Transaction
                            transaksi.SKU = data[0]
                            transaksi.IdMember = &data[1]
                            transaksi.Price = item.Price
                            transaksi.Qty = qty

                            item.AddTransaction(transaksi)
                            member.AddTransaction(transaksi)

                            return "successfully added transaction item " + item.SKU + " for member " + member.IdMember, nil
                        } else {
                            return "", errors.New("stock qty for item " + item.SKU + " is not sufficient")
                        }
                    }
                }
                return "", errors.New("item " + data[0] + " is not in list of items")
            }
        }
        return "", errors.New("member " + data[1] + " is not in list of members")
    } else {
        for i := 0; i < len(Items); i++ {
            item, ok := Items[i].GetData().(*Item)
            if !ok {
                return "", errors.New("Item type assertion failed")
            }
            if item.SKU == data[0] {
                if item.StockQty >= qty {
                    var transaksi Transaction
                    transaksi.SKU = data[0]
                    transaksi.Price = item.Price
                    transaksi.Qty = qty

                    item.AddTransaction(transaksi)

                    return "successfully added transaction item " + item.SKU, nil
                } else {
                    return "", errors.New("stock qty for item " + item.SKU + " is not sufficient")
                }
            }
        }
        return "", errors.New("item " + data[0] + " is not in list of items")
    }
}

func RestockItem(SKU string, qty int32) (string, error) {
	for i := 0; i < len(Items); i++ {
		item, ok := Items[i].GetData().(*Item)
            if !ok {
                return "", errors.New("Item type assertion failed")
            }
		if item.SKU == SKU {
			item.StockQty += qty
			return "successfully restocked qty for item " + SKU, nil
		}
	}
	return "", errors.New("item " + SKU + " is not in list of items")
}

func GetTransactionItem(SKU string) ([]Transaction, error) {
	for i := 0; i < len(Items); i++ {
		item, ok := Items[i].GetData().(*Item)
            if !ok {
                return nil, errors.New("Item type assertion failed")
            }
		if item.SKU == SKU {
			return item.Transactions, nil
		}
	}
	return nil, errors.New("item " + SKU + " is not in list of items")
}

func GetTransactionMember(idMember string) ([]Transaction, error) {
	for i := 0; i < len(Members); i++ {
		member, ok := Members[i].GetData().(*Member)
            if !ok {
                return nil, errors.New("Member type assertion failed")
            }
		if member.IdMember == idMember {
			return member.Transactions, nil
		}
	}
	return nil, errors.New("member " + idMember + " is not in list of members")
}
