package main
import "fmt"

type VendingMachine struct {
	insertedMoney int
	coins		  map[string]int
	items		  map[string]int
}

func (m VendingMachine) InsertedMoney() int {
	return m.insertedMoney
}

func (m *VendingMachine) InsertCoin(coin string) {
	m.insertedMoney += m.coins[coin]

}

func (m *VendingMachine) ClearCoin() {
	m.insertedMoney = 0
}

func (m VendingMachine) change(c int) string {
	var str string
	values := [...]int{10, 5, 2, 1}
	coins := [...]string{"T", "F", "TW", "O"}

	for i := 0; i < len(values); i++ {
		if c >= values[i] {
			str += ", " + coins[i]
			c -= values[i]
			i--
		}
	}
	
	return str
}

func (vm *VendingMachine) CoinReturn() string {
	coins := vm.change(vm.insertedMoney)
	vm.insertedMoney = 0
	return coins[2:len(coins)]
}

func (m *VendingMachine) SelectSD() string {

	price := m.items["SD"]
	change := m.insertedMoney - price
	return "SD" + m.change(change)
}

func (m *VendingMachine) SelectCC() string {
	
	price := m.items["CC"]
	change := m.insertedMoney - price
	return "CC" + m.change(change)
}

func (m *VendingMachine) SelectDW() string {
	
	price := m.items["DW"]
	change := m.insertedMoney - price
	return "DW" + m.change(change)
}

func main() {
	var coins = map[string]int{"T": 10, "F": 5, "TW": 2, "O": 1}
	var items = map[string]int{"SD": 18, "CC": 12, "DW": 7}
	vm := VendingMachine{coins: coins, items: items}
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
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// Inserted Money: 20
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