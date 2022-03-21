package main

import "fmt"

func main(){
// 	for i :=0; i < 3 ; i ++{
// 		fmt.Println ("Angka", i)
// 	}

// var i = 0
// for i < 3 {
// 	fmt.Println ("Angka", i)
// 	i++
// }

	// var i = 0
	// for i < 100 {
	// 	fmt.Println ("Angka", i)
	// 	i++

	// 	if i == 80 {
	// 		break
	// 	}
	// }

	// for i:=1; i<=10000 ; i++{
	// 	if i%2 == 0{
	// 		continue
	// 	}
	// 	if i>100 {
	// 		break
	// 	}
	// 	fmt.Println("Angka", i)
	// }
	// 	for i:=0; i<5; i++{
	// 		for j:=1; j<5; j++
	// 		fmt.Println(j, " ")
	// 	}
	// 	fmt.Println()
	// }
		outerloop :
		for i:=0; i < 3; i++ {
			fmt.Println("Perulangan ke-", i+ 1)
		for j:= 0; j<3;j++{
			if i ==2 {
				break outerloop
			}
			fmt.Println(j, " ")
		}
		fmt.Println("\n")

		}
		


	}


