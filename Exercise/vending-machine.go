package main
import "fmt"

type VendingMachine struct {
	insertedMoney int
	coins		  map[string]int
}

func (m VendingMachine) InsertedMoney() int {
	return m.insertedMoney
}

func (m *VendingMachine) InsertCoin(coin string) {
	m.insertedMoney += m.coins[coin]

}

func (m VendingMachine) change(c int) (change string) {
	return ", F, TW, O"
}

func (m *VendingMachine) SelectSD() string {
	m.insertedMoney = 0
	return "SD"
}

func (m *VendingMachine) SelectCC() string {
	m.insertedMoney = 0
	price := 12
	change := m.insertedMoney - price
	return "CC" + m.change(change)
}

func main() {
	var coins = map[string]int{"T": 10, "F": 5, "TW": 2, "O": 1}
	vm := VendingMachine{coins: coins}
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

	vm.InsertCoin("T")
	vm.InsertCoin("TW")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// Inserted Money: 12
	can = vm.SelectCC()
	fmt.Println(can) // CC

	vm.InsertCoin("T")
	vm.InsertCoin("T")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// Inserted Money: 20
	can = vm.SelectCC()
	fmt.Println(can) // CC, F, TW, O
}