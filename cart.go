package demo

type Cart struct {
	qtyItems   int
}

func (c *Cart) Clear() {
	c.qtyItems = 0
}

func (c *Cart) QtyItems() int {
	return c.qtyItems
}

func (c *Cart) AddItems(qtyItem int) error{
	c.qtyItems = c.qtyItems + qtyItem
	return nil
}

func (c *Cart) RemoveItem(qtyItem int) error{
	c.qtyItems = c.qtyItems - qtyItem
	return nil
}

func main() {}
