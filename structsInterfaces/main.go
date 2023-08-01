package main

import (
	"encoding/json"
	"fmt"

	"github.com/SPallazhco/go/structsInterfaces/strunctsInterface/structs"
)

func main() {
	var p1 structs.Product
	fmt.Println(p1)
	p1.ID = 1
	p1.Name = "test"
	fmt.Println(p1)

	p2 := structs.Product{}

	fmt.Println(p2)

	p3 := structs.Product{2, "test_1", structs.Type{3, "test_0", "D"}, 1, 12.2, nil}

	fmt.Println(p3)

	p4 := structs.Product{
		ID:   3,
		Name: "test_2",
		// Type:  "B",
		Price: 1.9,
	}
	fmt.Println(p4)

	p5 := structs.Product{
		Name:  "Heladera",
		Price: 100,
		Type: structs.Type{
			Code:        "A",
			Description: "Electrodomesticos",
		},
		Tags:  []string{"helados", "sarasa", "freezer"},
		Count: 5,
	}
	v, err := json.Marshal(p5)
	fmt.Println(err)
	fmt.Println(v) // Marshal return the json in bytes format
	fmt.Println()
	fmt.Println(string(v))

	total := p5.Total()
	fmt.Println("Precio total: ", total)
	fmt.Println()

	fmt.Println(p5)
	p5.SetName("other_name")
	fmt.Println(p5)
	fmt.Println()

	p5.AddTags("tag1", "tag2", "tag3")
	fmt.Println(p5)

	p6 := structs.Product{
		Name:  "Naranja",
		Price: 10,
		Type: structs.Type{
			Code:        "B",
			Description: "Alimento",
		},
		Tags:  []string{"alimento", "fruta"},
		Count: 2,
	}
	p7 := structs.Product{
		Name:  "Mesa",
		Price: 4,
		Type: structs.Type{
			Code:        "C",
			Description: "Mueble",
		},
		Tags:  []string{"Arreglo", "Tabla", "Caoba"},
		Count: 8,
	}

	car := structs.NewCar(1111)
	car.AddProducts(p5, p6, p7)

	fmt.Println()
	fmt.Println("TOTAL PRODUCTS", len(car.Products))
	fmt.Printf("TOTAL CAR: $%.2f\n", car.Total())

}
