package main

import "fmt"
import "github.com/acetnis/vendingMachine"

var coins = map[string]int{"T": 10, "F": 5, "TW": 2, "O": 1}
var items = map[string]int{"SD": 18, "CC": 12, "DW": 7}

func main() {

	//vm := VendingMachine{coins: coins, items: items}
	vm := vendingMachine.NewVendingMachine(coins, items)
	fmt.Println("inserted Money:", vm.InsertedMoney())
	// Inserted Money: 0
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	vm.InsertCoin("O")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// Inserted Money: 18
	can := vm.SelectSD()
	fmt.Println(can) // SD

	vm.ClearCoin()
	vm.InsertCoin("T")
	vm.InsertCoin("T")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// Inserted Money: 20
	can = vm.SelectCC()
	fmt.Println(can) // CC, F, TW, O

	vm.ClearCoin()
	vm.InsertCoin("T")
	vm.InsertCoin("T")
	vm.InsertCoin("T")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// Inserted Money: 30
	can = vm.SelectDW()
	fmt.Println(can) // DW, T, TW, O

	vm.ClearCoin()
	vm.InsertCoin("T")
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// Inserted Money: 25
	coin := vm.CoinReturn()
	fmt.Println("Returned Coin:", coin) // T, T, F
}
