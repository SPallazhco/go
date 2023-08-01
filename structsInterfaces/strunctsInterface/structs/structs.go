package structs

type Product struct {
	ID    uint     `json: "id"`
	Name  string   `json: "name"`
	Type  Type     `json: "type"`
	Count uint     `json: "count"`
	Price float64  `json: "price"`
	Tags  []string `json: "tags"`
}

type Type struct {
	ID          uint   `json: "id"`
	Code        string `json: "code"`
	Description string `json: "description"`
}

func (product Product) Total() float64 {
	return float64(product.Count) * product.Price
}

func (product *Product) SetName(name string) {
	product.Name = name
}

func (p *Product) AddTags(tags ...string) {
	p.Tags = append(p.Tags, tags...)
}

type Car struct {
	ID       uint
	UserID   uint
	Products []Product
}

func (c *Car) AddProducts(products ...Product) {
	c.Products = append(c.Products, products...)
}

func (c Car) Total() float64 {
	var total float64

	for _, c := range c.Products {
		total += c.Total()
	}

	return total
}

func NewCar(userID uint) Car {
	return Car{UserID: userID}
}
