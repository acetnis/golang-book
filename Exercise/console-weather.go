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
	
	p1 := num%10
	p2 := (num - p1) / 10
	if num < 100 {
		printnum(p2,1)
		fmt.Print(" ")
		printnum(p1,1)
		fmt.Println()
		printnum(p2,2)
		fmt.Print(" ")
		printnum(p1,2)
		fmt.Println()
		printnum(p2,3)
		fmt.Print(" ")
		printnum(p1,3)
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
	if n == 2 {
		if pos == 1 {
			fmt.Print(" _ ")
		} else if pos == 2 {
			fmt.Print(" _|")
		} else if pos == 3 {
			fmt.Print("|_ ")
		}
	}
	if n == 3 {
		if pos == 1 {
			fmt.Print(" _ ")
		} else if pos == 2 {
			fmt.Print(" _|")
		} else if pos == 3 {
			fmt.Print(" _|")
		}
	}
	if n == 4 {
		if pos == 1 {
			fmt.Print("   ")
		} else if pos == 2 {
			fmt.Print("|_|")
		} else if pos == 3 {
			fmt.Print("  |")
		}
	}
	if n == 5 {
		if pos == 1 {
			fmt.Print(" _ ")
		} else if pos == 2 {
			fmt.Print("|_ ")
		} else if pos == 3 {
			fmt.Print(" _|")
		}
	}
	if n == 6 {
		if pos == 1 {
			fmt.Print(" _ ")
		} else if pos == 2 {
			fmt.Print("|_ ")
		} else if pos == 3 {
			fmt.Print("|_|")
		}
	}
	if n == 7 {
		if pos == 1 {
			fmt.Print(" _ ")
		} else if pos == 2 {
			fmt.Print("  |")
		} else if pos == 3 {
			fmt.Print("  |")
		}
	}
	if n == 8 {
		if pos == 1 {
			fmt.Print(" _ ")
		} else if pos == 2 {
			fmt.Print("|_|")
		} else if pos == 3 {
			fmt.Print("|_|")
		}
	}
	if n == 9 {
		if pos == 1 {
			fmt.Print(" _ ")
		} else if pos == 2 {
			fmt.Print("|_|")
		} else if pos == 3 {
			fmt.Print(" _|")
		}
	}
}