package main
import "fmt"

func dompet(n, hasil *int) {
	if *n >= 0 {
		*hasil = *n
	}
}

func pengurangan(hasil, num *int) {
	*hasil -= *num 
}

func main() {
	var num, hasil int
	fmt.Scan(&num)
	dompet(&num, &hasil)

	for {
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		if num > 0 { 
			pengurangan(&hasil, &num)
		}
	}

	fmt.Print(hasil)
}

