Feature: Interact with cart

  Scenario: Add item to cart
   Given cart empty
   When i add 1 item to cart
   Then the cart should have 1 item

  Scenario: Remove item in cart
   Given cart with 1 item
   When i remove 1 item to cart
   Then the cart should have 0 items

  Scenario: Modify qty of item in cart
   Given cart with 1 item
   When i add 2 qty of item in cart
   Then the cart should have 3 items
