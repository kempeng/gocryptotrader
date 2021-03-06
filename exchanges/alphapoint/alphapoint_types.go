package alphapoint

import "github.com/thrasher-/gocryptotrader/decimal"

// Response contains general responses from the exchange
type Response struct {
	IsAccepted    bool            `json:"isAccepted"`
	RejectReason  string          `json:"rejectReason"`
	Fee           decimal.Decimal `json:"fee"`
	FeeProduct    string          `json:"feeProduct"`
	CancelOrderID int64           `json:"cancelOrderId"`
	ServerOrderID int64           `json:"serverOrderId"`
	DateTimeUTC   decimal.Decimal `json:"dateTimeUtc"`
	ModifyOrderID int64           `json:"modifyOrderId"`
	Addresses     []DepositAddresses
}

// Ticker holds ticker information
type Ticker struct {
	High               decimal.Decimal `json:"high"`
	Last               decimal.Decimal `json:"last"`
	Bid                decimal.Decimal `json:"bid"`
	Volume             decimal.Decimal `json:"volume"`
	Low                decimal.Decimal `json:"low"`
	Ask                decimal.Decimal `json:"ask"`
	Total24HrQtyTraded decimal.Decimal `json:"Total24HrQtyTraded"`
	Total24HrNumTrades decimal.Decimal `json:"Total24HrNumTrades"`
	SellOrderCount     decimal.Decimal `json:"sellOrderCount"`
	BuyOrderCount      decimal.Decimal `json:"buyOrderCount"`
	NumOfCreateOrders  decimal.Decimal `json:"numOfCreateOrders"`
	IsAccepted         bool            `json:"isAccepted"`
	RejectReason       string          `json:"rejectReason"`
}

// Trades holds trade information
type Trades struct {
	IsAccepted   bool    `json:"isAccepted"`
	RejectReason string  `json:"rejectReason"`
	DateTimeUTC  int64   `json:"dateTimeUtc"`
	Instrument   string  `json:"ins"`
	StartIndex   int     `json:"startIndex"`
	Count        int     `json:"count"`
	StartDate    int64   `json:"startDate"`
	EndDate      int64   `json:"endDate"`
	Trades       []Trade `json:"trades"`
}

// Trade is a sub-type which holds the singular trade that occurred in the past
type Trade struct {
	TID                   int64           `json:"tid"`
	Price                 decimal.Decimal `json:"px"`
	Quantity              decimal.Decimal `json:"qty"`
	Unixtime              int             `json:"unixtime"`
	UTCTicks              int64           `json:"utcticks"`
	IncomingOrderSide     int             `json:"incomingOrderSide"`
	IncomingServerOrderID int             `json:"incomingServerOrderId"`
	BookServerOrderID     int             `json:"bookServerOrderId"`
}

// Orderbook holds the total Bids and Asks on the exchange
type Orderbook struct {
	Bids         []OrderbookEntry `json:"bids"`
	Asks         []OrderbookEntry `json:"asks"`
	IsAccepted   bool             `json:"isAccepted"`
	RejectReason string           `json:"rejectReason"`
}

// OrderbookEntry is a sub-type that takes has the individual quantity and price
type OrderbookEntry struct {
	Quantity decimal.Decimal `json:"qty"`
	Price    decimal.Decimal `json:"px"`
}

// ProductPairs holds the full range of product pairs that the exchange can
// trade between
type ProductPairs struct {
	ProductPairs []ProductPair `json:"productPairs"`
	IsAccepted   bool          `json:"isAccepted"`
	RejectReason string        `json:"rejectReason"`
}

// ProductPair holds the individual product pairs that are currently traded
type ProductPair struct {
	Name                  string `json:"name"`
	Productpaircode       int    `json:"productPairCode"`
	Product1Label         string `json:"product1Label"`
	Product1Decimalplaces int    `json:"product1DecimalPlaces"`
	Product2Label         string `json:"product2Label"`
	Product2Decimalplaces int    `json:"product2DecimalPlaces"`
}

// Products holds the full range of supported currency products
type Products struct {
	Products     []Product `json:"products"`
	IsAccepted   bool      `json:"isAccepted"`
	RejectReason string    `json:"rejectReason"`
}

// Product holds the a single currency product that is supported
type Product struct {
	Name          string `json:"name"`
	IsDigital     bool   `json:"isDigital"`
	ProductCode   int    `json:"productCode"`
	DecimalPlaces int    `json:"decimalPlaces"`
	FullName      string `json:"fullName"`
}

// UserInfo holds current user information associated with the apiKey details
type UserInfo struct {
	UserInforKVPs []UserInfoKVP `json:"userInfoKVP"`
	IsAccepted    bool          `json:"isAccepted"`
	RejectReason  string        `json:"rejectReason"`
}

// UserInfoKVP is a sub-type that holds key value pairs
type UserInfoKVP struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// UserInfoSet is the returned response from set user information request
type UserInfoSet struct {
	IsAccepted        string `json:"isAccepted"`
	RejectReason      string `json:"rejectReason"`
	RequireAuthy2FA   bool   `json:"requireAuthy2FA"`
	Val2FaRequestCode string `json:"val2FaRequestCode"`
}

// AccountInfo holds your current account information like balances, trade count
// and volume
type AccountInfo struct {
	Currencies []struct {
		Name    string `json:"name"`
		Balance int    `json:"balance"`
		Hold    int    `json:"hold"`
	} `json:"currencies"`
	ProductPairs []struct {
		ProductPairName string `json:"productPairName"`
		ProductPairCode int    `json:"productPairCode"`
		TradeCount      int    `json:"tradeCount"`
		TradeVolume     int    `json:"tradeVolume"`
	} `json:"productPairs"`
	IsAccepted   bool   `json:"isAccepted"`
	RejectReason string `json:"rejectReason"`
}

// Order is a generalised order type
type Order struct {
	Serverorderid int   `json:"ServerOrderId"`
	AccountID     int   `json:"AccountId"`
	Price         int   `json:"Price"`
	QtyTotal      int   `json:"QtyTotal"`
	QtyRemaining  int   `json:"QtyRemaining"`
	ReceiveTime   int64 `json:"ReceiveTime"`
	Side          int   `json:"Side"`
}

// OpenOrders holds the full range of orders by instrument
type OpenOrders struct {
	Instrument string  `json:"ins"`
	Openorders []Order `json:"openOrders"`
}

// OrderInfo holds all open orders across the entire range of all instruments
type OrderInfo struct {
	OpenOrders   []OpenOrders `json:"openOrdersInfo"`
	IsAccepted   bool         `json:"isAccepted"`
	DateTimeUTC  int64        `json:"dateTimeUtc"`
	RejectReason string       `json:"rejectReason"`
}

// DepositAddresses holds information about the generated deposit address for
// a specific currency
type DepositAddresses struct {
	Name           string `json:"name"`
	DepositAddress string `json:"depositAddress"`
}

// WebsocketTicker holds current up to date ticker information
type WebsocketTicker struct {
	MessageType             string          `json:"messageType"`
	ProductPair             string          `json:"prodPair"`
	High                    decimal.Decimal `json:"high"`
	Low                     decimal.Decimal `json:"low"`
	Last                    decimal.Decimal `json:"last"`
	Volume                  decimal.Decimal `json:"volume"`
	Volume24Hrs             decimal.Decimal `json:"volume24hrs"`
	Volume24HrsProduct2     decimal.Decimal `json:"volume24hrsProduct2"`
	Total24HrQtyTraded      decimal.Decimal `json:"Total24HrQtyTraded"`
	Total24HrProduct2Traded decimal.Decimal `json:"Total24HrProduct2Traded"`
	Total24HrNumTrades      decimal.Decimal `json:"Total24HrNumTrades"`
	Bid                     decimal.Decimal `json:"bid"`
	Ask                     decimal.Decimal `json:"ask"`
	BuyOrderCount           int             `json:"buyOrderCount"`
	SellOrderCount          int             `json:"sellOrderCount"`
}
