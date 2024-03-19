package lists

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

// main description of the Go function.
func main() {
	hobbies := []string{"Trading", "Reading", "Hiking"}
	highlighted := hobbies[1:]
	fmt.Println(hobbies[0])
	fmt.Println(highlighted)
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	courseGoals := []string{"Finish the course", "Get a job"}
	courseGoals = append(courseGoals, "Learn Go")
	fmt.Println(courseGoals)

	products := []Product{
		{"P1", "Product 1", 10.99},
		{"P2", "Product 2", 20.99},
		{"P3", "Product 3", 30.99},
	}
	newProduct := Product{"P4", "Product 4", 40.99}
	products = append(products, newProduct)
	fmt.Println(products)
}
