package main

// AssertManagerService struct
type AssertManagerService struct {
}

// Deposit return BuyOrders which should sastify DesirePortfolio
func (ams AssertManagerService) Deposit(uid UserID, amount float64) BuyOrders {
	return invest(amount)
}

// Invest using Robo's algorithm
// Let assume to invest 50% with BTC and 50% with ETH
func invest(amount float64) BuyOrders {
	btcOrder := Order{amount: amount / 2, currency: "TUSD", targetCurrency: "BTC"}
	ethOrder := Order{amount: amount / 2, currency: "TUSD", targetCurrency: "ETH"}

	orders := make([]Order, 2)
	orders[0] = btcOrder
	orders[1] = ethOrder

	return BuyOrders{orders}
}

// Withdraw return SellOrders which should keep DesirePortfolio in the same ratio
func (ams AssertManagerService) Withdraw(uid UserID, amount float64) SellOrders {
	return redeem(amount)
}

func redeem(amount float64) SellOrders {
	btcOrder := Order{amount: amount / 2, currency: "BTC", targetCurrency: "TUSD"}
	ethOrder := Order{amount: amount / 2, currency: "ETH", targetCurrency: "TUSD"}

	orders := make([]Order, 2)
	orders[0] = btcOrder
	orders[1] = ethOrder

	return SellOrders{orders}
}
