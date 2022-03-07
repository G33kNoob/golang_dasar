package arrfun

import "fmt"

func Arrlistdata() {
	myMaps := map[string]string{
		"ruby":   "is beautiful",
		"js":     "is popular",
		"golang": "is the fastest language",
	}
	fmt.Println(myMaps)
	for _, myMap := range myMaps {
		fmt.Println(myMap)
	}
}
func Arrlistdecimal() {
	Price := [...]int{100, 80, 90, 35, 95, 10, 28, 900, 87, 12, 11, 8}
	lenPrice := len(Price)
	var totalPrice int
	_, _ = totalPrice, Price
	for i := 0; i < lenPrice; i++ {
		// fmt.Println("awal", totalPrice)
		// fmt.Println(i)
		totalPrice = totalPrice + Price[i]
		// fmt.Println("ahir", totalPrice)
	}
	fmt.Println("length price", lenPrice)
	fmt.Println("total price", totalPrice)
	ratarata := float64(totalPrice) / float64(lenPrice)
	fmt.Printf("rata-rata = %.2f\n", ratarata)
	fmt.Println(ratarata)

}
