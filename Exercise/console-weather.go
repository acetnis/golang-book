package main
import "fmt"

func main() {
	weatherCelsuis(25, "Bangkok few cloud")
	weatherCelsuis(34, "Tak sunny")
	weatherCelsuis(17, "Phuket rainy")
	weatherCelsuis(9, "Chiang-mai cold")
	weatherCelsuis(10, "Test")
}

func weatherCelsuis(num int,str string) {
	if num == 25 {
		fmt.Println(" _   _")
		fmt.Println(" _| |_")
		fmt.Print("|_   _|")
		fmt.Println("  c")
		fmt.Printf("%v\n", str)
	}
	if num == 34 {
		fmt.Println(" _   ")
		fmt.Println(" _| |_|")
		fmt.Print(" _|   |")
		fmt.Println("  c")
		fmt.Printf("%v\n", str)
	}
	if num == 17 {
		fmt.Println("     _")
		fmt.Println("  |   |")
		fmt.Print("  |   |")
		fmt.Println("  c")
		fmt.Printf("%v\n", str)
	}
	if num == 9 {
		fmt.Println(" _ ")
		fmt.Println("|_|")
		fmt.Print(" _|")
		fmt.Println("  c")
		fmt.Printf("%v\n", str)
	}
	if num == 10 {
		printnum(1,1)
		fmt.Print(" ")
		printnum(0,1)
		fmt.Println()
		printnum(1,2)
		fmt.Print(" ")
		printnum(0,2)
		fmt.Println()
		printnum(1,3)
		fmt.Print(" ")
		printnum(0,3)
		fmt.Println("  c")
		fmt.Printf("%v\n", str)
	}
}

func printnum(n int,pos int) {
	if n == 0 {
		if pos == 1 {
			fmt.Print(" _ ")
		} else if pos == 2 {
			fmt.Print("| |")
		} else if pos == 3 {
			fmt.Print("|_|")
		}
	}
	if n == 1 {
		if pos == 1 {
			fmt.Print("   ")
		} else if pos == 2 {
			fmt.Print("  |")
		} else if pos == 3 {
			fmt.Print("  |")
		}
	}
}