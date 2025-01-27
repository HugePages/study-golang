package main

import "fmt"

func main() {
	months := []string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	fmt.Println(months)

	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)     // ["April" "May" "June"]
	fmt.Println(summer) // ["June" "July" "August"]

	fmt.Println("len(Q2)--->", len(Q2))

	fmt.Println("cap(Q2)--->", cap(Q2))

	fmt.Println("len(summer)--->", len(summer))
	fmt.Println("cap(summer)--->", cap(summer))

}
