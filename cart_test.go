package demo

import (
	"fmt"
	"github.com/cucumber/godog"
)

type CartSuite struct {
	*godog.ScenarioContext
	cart *Cart
}

func (cs *CartSuite) cartEmpty() error {
	cs.cart.Clear()
	return nil
}

func (cs *CartSuite) cartWithItem(qtyItems int) error {
	_ = cs.cart.AddItems(qtyItems)
	return nil
}

func (cs *CartSuite) iAddItemToCart(qtyItemToAdd int) error {
	return cs.iAddQtyOfItemInCart(qtyItemToAdd)
}

func (cs *CartSuite) iAddQtyOfItemInCart(qtyItemsToAdd int) error {
	return cs.cart.AddItems(qtyItemsToAdd)
}

func (cs *CartSuite) iRemoveItemToCart(qtyItemsToRemove int) error {
	return cs.cart.RemoveItem(qtyItemsToRemove)
}

func (cs *CartSuite) theCartShouldHaveItem(expectedResult int) error {
	result := cs.cart.QtyItems()
	if result == expectedResult {
		return nil
	}
	return fmt.Errorf("%d doesn't match expectation: %d", result, expectedResult)
}

func (cs *CartSuite) theCartShouldHaveItems(expectedResult int) error {
	result := cs.cart.QtyItems()
	if result == expectedResult {
		return nil
	}
	return fmt.Errorf("%d doesn't match expectation: %d", result, expectedResult)
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	c := CartSuite{
		cart:  new(Cart),
		ScenarioContext: ctx,
	}
	ctx.Step(`^cart empty$`, c.cartEmpty)
	ctx.Step(`^cart with (\d+) item$`, c.cartWithItem)
	ctx.Step(`^i add (\d+) item to cart$`, c.iAddItemToCart)
	ctx.Step(`^i add (\d+) qty of item in cart$`, c.iAddQtyOfItemInCart)
	ctx.Step(`^i remove (\d+) item to cart$`, c.iRemoveItemToCart)
	ctx.Step(`^the cart should have (\d+) item$`, c.theCartShouldHaveItem)
	ctx.Step(`^the cart should have (\d+) items$`, c.theCartShouldHaveItems)
}


