package main
import "fmt"

type VendingMachine struct {

}

func (m VendingMachine) InsertedMoney() int {
	return 0
}

func main() {
	vm := VendingMachine{}
	fmt.Println("inserted Money:", vm.InsertedMoney())
}