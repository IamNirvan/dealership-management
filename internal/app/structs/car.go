package structs

import "fmt"

type Car struct {
	Manufacturer string
	Model        string
	Colour       string
	Price        string
}

func (c *Car) String() string {
	return fmt.Sprintf(`
		Manufacturer: %s,
		Model: %s,
		Colour: %s,
		Price: %s
	`, c.Manufacturer, c.Model, c.Colour, c.Price)
}
