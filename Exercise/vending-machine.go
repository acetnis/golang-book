package main
import "fmt"

type VendingMachine struct {

}

func (m VendingMachine) InsertedMoney() int {
	return 10
}

func (m *VendingMachine) InsertCoin(coin string) {

}

func main() {
	vm := VendingMachine{}
	fmt.Println("inserted Money:", vm.InsertedMoney())
	// Inserted Money: 0
	vm.InsertCoin("T")
	fmt.Println("Inserted Money", vm.InsertedMoney())
	// Inserted Money: 10
}