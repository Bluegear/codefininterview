package main

// AssetManagerService Interface
type AssetManagerService interface {
	Deposit(uid UserID, amount float64) BuyOrders
	Withdraw(uid UserID, amount float64) SellOrders
	/*
	 Portpolio data changed after each successful trade between
	 Cryptodini and Binance so Robo Server should not implement their own GetPort.
	*/
	// GetPort(uid UserID) Portfolio
}

// CryptodiniServiceInf Interface
type CryptodiniServiceInf interface {
	Adjust(uid UserID, desirePort *Portfolio)
	GetPort(uid UserID) *Portfolio
}

// CryptodiniService struct
type CryptodiniService struct {
}

// GetPort Get current user's portfolio
func (c CryptodiniService) GetPort(uid UserID) *Portfolio {
	return &Portfolio{}
}

// UserID struct
type UserID struct {
}

// Portfolio struct
type Portfolio struct {
}

// Balance struct
type Balance struct {
	amount   float64
	currency string
}

// BuyOrders struct
type BuyOrders struct {
	orders []Order
}

// SellOrders struct
type SellOrders struct {
	orders []Order
}

// Order struct
type Order struct {
	amount         float64
	currency       string
	targetCurrency string
}
