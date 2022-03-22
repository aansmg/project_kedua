package main

import ("fmt"
//  "strings"
)
func main(){
	// var person map[string]string
	// person = map[string]string{}

	// person["name"] = "Hacktic8"
	// person["age"] = "6"
	// person["address"] = "jalan lorem"


	// fmt.Println("name", person["name"])
	// fmt.Println("age", person["age"])
	// fmt.Println("address", person["address"])

	var person = map[string]string{
		"name" : "Hacktic8",
		"age"  : "6",
		"address" : "jalan lorem",
	}
	// fmt.Println("name", person["name"])
	// fmt.Println("age", person["age"])
	// fmt.Println("address", person["address"])
// for key, value := range person{
// 	fmt.Println(key, ":", value)
// // }
// 		fmt.Println("Before Deleting:", person)
// 		delete(person, "age")
// 		fmt.Println("After Deleting:", person)

// 		//detect value
// 		if exist {
// 			fmt.Println(value)
// 			}else {
// 				fmt.Println("value not exist")
// 			}

// 			delete(person, "age)")
// 			value, exist = person["age"]

// 			if exist {
// 				fmt.Println(value)
// 				}else {
// 					fmt.Println("value has been deleted")
// 				}
			var people =[]map[string]string{
				{"name": "hacktiv8", "age": "6"},
				{"name": "hacktivkidz", "age": "2"},
				{"name": "KODE", "age": "5"},
			}
			for i person := range people{
				fmt.Println("index: %d, name : %s, age:%s\n, i, person["name", person["age"])
			} 

}

