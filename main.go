package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("[CRASH] ", r)
	// 	}
	// }()

	fmt.Printf("Name: %s, ID Student: %s\n", Name, IdStudent)
	fmt.Println("========================================")
	fmt.Println("Welcome to Sigmart Point of Sales")
	fmt.Println("Please input your command below")
	fmt.Println("========================================")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		line := scanner.Text()
		err := scanner.Err()
		if err != nil {
			fmt.Println("[CRASH] ", err.Error())
			os.Exit(1)
		}

		spl := strings.Split(line, " ")
		executeCommand(spl[0], spl[1:])
	}
}

func executeCommand(command string, data []string) {
	switch command {
	case "ADD_ITEM":
		if len(data) != 4 {
			PrintMessage("", errors.New("your input command is incorrect"))
		} else {
			if integer_price, err := strconv.Atoi(data[2]); err == nil {
				if integer_stockQty, err := strconv.Atoi(data[3]); err == nil {
					PrintMessage(AddItem(data[0], data[1], int32(integer_price), int32(integer_stockQty)))
				} else {
					PrintMessage("", errors.New("your input command is incorrect"))
				}
			} else {
				PrintMessage("", errors.New("your input command is incorrect"))
			}
		}
	case "DELETE_ITEM":
		if len(data) != 1 {
			PrintMessage("", errors.New("your input command is incorrect"))
		} else {
			PrintMessage(DeleteItem(data[0]))
		}
	case "ADD_MEMBER":
		if len(data) != 2 {
			PrintMessage("", errors.New("your input command is incorrect"))
		} else {
			PrintMessage(AddMember(data[0], data[1]))
		}
	case "DELETE_MEMBER":
		if len(data) != 1 {
			PrintMessage("", errors.New("your input command is incorrect"))
		} else {
			PrintMessage(DeleteMember(data[0]))
		}
	case "ADD_TRANSACTION":
		if len(data) != 2 || len(data) != 3 {
			PrintMessage("", errors.New("your input command is incorrect"))
		} else {
			if integer_qty, err := strconv.Atoi(data[0]); err == nil {
				PrintMessage(AddTransaction(int32(integer_qty), data[1:]...))
			} else {
				PrintMessage("", errors.New("your input command is incorrect"))
			}
		}
	case "RESTOCK_ITEM":
		if len(data) != 2 {
			PrintMessage("", errors.New("your input command is incorrect"))
		} else {
			if integer_qty, err := strconv.Atoi(data[1]); err == nil {
				PrintMessage(RestockItem(data[0], int32(integer_qty)))
			} else {
				PrintMessage("", errors.New("your input command is incorrect"))
			}
		}
	case "TRANSACTION_ITEM_RECAP":
		if len(data) != 1 {
			PrintMessage("", errors.New("your input command is incorrect"))
		} else {
			for i := 0; i < len(Items); i++ {
				if Items[i].GetData().(Item).SKU == data[0] {
					PrintTransactionRecap(Items[i].GetData().(Item).Transactions, nil)
				}
			}
		}
	case "TRANSACTION_MEMBER_RECAP":
		if len(data) != 1 {
			PrintMessage("", errors.New("your input command is incorrect"))
		} else {
			for i := 0; i < len(Members); i++ {
				if Members[i].GetData().(Item).SKU == data[0] {
					PrintTransactionRecap(Members[i].GetData().(Item).Transactions, nil)
				}
			}
		}
	case "EXIT":
		os.Exit(1)
	default:
		os.Exit(1)
	}
}

func PrintMessage(successMsg string, errMsg error) {
	if errMsg != nil {
		fmt.Println("[FAILED] " + errMsg.Error())
	} else {
		fmt.Println("[SUCCESS] " + successMsg)
	}
}

func PrintTransactionRecap(transactions []Transaction, errMsg error) {
	if len(transactions) == 0 {
		fmt.Println("[FAILED] " + errMsg.Error())
	}
	fmt.Println("-x-x-x-x-x-x-x-x-x-x-x-x-")
	for i := 0; i < len(transactions); i++ {
		fmt.Println("SKU: " + transactions[i].SKU + ", ID Member: " + *transactions[i].IdMember + ", Total Price: " + strconv.Itoa(int(transactions[i].Price*transactions[i].Qty)))
	}
	fmt.Println("-x-x-x-x-x-x-x-x-x-x-x-x-")
}
