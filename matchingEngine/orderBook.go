package matchingengine

import ds "hammadExchange/datastructures"
import "container/list"
import "time"

type Order struct {
	id               int
	createdTimeStamp time.Time
	userId           int
	ticker           string
	isBid            bool
	quantity         float64
	priceLevel       float64
}
type OrderBook struct {
	bids *ds.RedBlackTree[float64, *list.List]
	asks *ds.RedBlackTree[float64, *list.List]
}

func (orderBook *OrderBook) placeBuyOrderAndGetMatches(order *Order) []Order {
	counterPartExists := orderBook.bids.IsKeyLessThanOrEqualExists(&order.priceLevel)
	if counterPartExists {
		orderBook.getMatchingOrdersForBid(order.quantity, order.priceLevel)
	}
	return nil
}
func (orderBook *OrderBook) getMatchingOrdersForBid(quantity float64, priceLevel float64) {
	var matchingOrders []Order
	var root = orderBook.asks.GetRoot()
	orderBook.traverseUntilPriceLimitIncreasing(root, &matchingOrders, &quantity, priceLevel)
}

func (orderBook *OrderBook) placeSellOrderAndGetMatches(order *Order) []Order {

}
func (orderBook *OrderBook) traverseUntilPriceLimitIncreasing(node *ds.TreeNode[float64, *list.List], matchingOrders *[]Order, quantityLeft *float64, priceLevel float64) {
	if node == nil {
		return
	}
	orderBook.traverseUntilPriceLimitIncreasing(node.GetLeft(), matchingOrders, quantityLeft, priceLevel)
	var allOrders = node.GetVal()
	for e := (*allOrders).Front(); e != nil; e = e.Next() {
		order, ok := e.Value.(Order)
		if !ok {
			continue
		}
		if(order.quantity<= *quantityLeft) {
			*quantityLeft = *quantityLeft- order.quantity
		} else {
			*quantityLeft = 0
		}
	}

	orderBook.traverseUntilPriceLimitIncreasing(node.GetRight(), matchingOrders, quantityLeft, priceLevel)

}
