package pair

import "github.com/thrasher-/gocryptotrader/currency/symbol"

// BitcoinSymbol defines the standard symbols for crypto coins, all in upper
const (
	BitcoinSymbol    = "BTC"
	AltBitcoinSymbol = "XBT"
	EthereumSymbol   = "ETH"
)

// FirstEqual compares FirstCurrencies of two currency pairs and returns whether or not they are equal
func (p CurrencyPair) FirstEqual(b CurrencyPair) bool {
	return p.FirstCurrency.StdCoinSymbol() == b.FirstCurrency.StdCoinSymbol()
}

// IsFiat checks whether provided currency is a Fiat currency
func (c CurrencyItem) IsFiat() bool {
	_, err := symbol.GetSymbolByCurrencyName(c.ToUpper())
	if err != nil {
		return false
	}
	return true
}

// PairToString converts a currencyPair to upper string, no delimiter
// for use as string in spreadKey
func (p CurrencyPair) String() string {
	return p.Display("", true).String()
}

// ToUpper converts a currencyPair to upper string
func (p CurrencyPair) ToUpper() string {
	return p.Display(p.Delimiter, true).ToUpper()
}

// ToLower converts a currencyPair to loweer string
func (p CurrencyPair) ToLower() string {
	return p.Display(p.Delimiter, false).ToLower()
}

// StdCoinSymbol returns the "standard Coin symbol, e.g. BTC for Bitcoin, all returned symbols are upper
func (c CurrencyItem) StdCoinSymbol() string {
	switch c.ToUpper() {
	case AltBitcoinSymbol:
		return BitcoinSymbol
	default:
		return c.ToUpper()
	}
}

// StdBitcoin return a standard bitcoin currency item for the provided currency item
func (c CurrencyItem) StdBitcoin() CurrencyItem {
	return CurrencyItem(c.StdCoinSymbol())
}

// StdBitcoinPair return a standard bitcoin currency pair for the provided currency pair
func (p CurrencyPair) StdBitcoinPair() CurrencyPair {
	return CurrencyPair{
		FirstCurrency:  CurrencyItem(p.FirstCurrency.StdCoinSymbol()),
		SecondCurrency: CurrencyItem(p.SecondCurrency.StdCoinSymbol()),
		Delimiter:      p.Delimiter,
	}
}

// ChangeBitcoin swaps between the std and alt BitcoinSymbol
func (c CurrencyItem) ChangeBitcoin(std bool) CurrencyItem {
	if std {
		if c.ToUpper() == AltBitcoinSymbol {
			return CurrencyItem(BitcoinSymbol)
		}
		return c
	}
	if c.ToUpper() == BitcoinSymbol {
		return CurrencyItem(AltBitcoinSymbol)
	}
	return c
}

// IsCrypto checks whether provided currency is a Crypto currency
func (c CurrencyItem) IsCrypto() bool {
	return !c.IsFiat()
}

// IsBitcoin tests whether provided currency item is a bitcoin or not
func (c CurrencyItem) IsBitcoin() bool {
	if c.IsStdBitcoin() || c.IsAltBitcoin() {
		return true
	}
	return false
}

// IsAltBitcoin tests whether currency item is a alt bitcoin symbol
func (c CurrencyItem) IsAltBitcoin() bool {
	if c.ToUpper() == AltBitcoinSymbol {
		return true
	}
	return false
}

// IsStdBitcoin tests whether currency item is BTC
func (c CurrencyItem) IsStdBitcoin() bool {
	if c.ToUpper() == BitcoinSymbol {
		return true
	}
	return false
}

// StdCryptoString converts a crypto CurrencyItem into a std formatted string
func (c CurrencyItem) StdCryptoString() string {
	return c.StdBitcoin().ToUpper()
}

// ToUpper converts a currencyItem into an Upper string
func (c CurrencyItem) ToUpper() string {
	return c.Upper().String()
}

// ToLower converts a currencyItem into an Upper string
func (c CurrencyItem) ToLower() string {
	return c.Lower().String()
}

// NewFromCurrencyItems creates a new currency pair from a pair of currency items
func NewFromCurrencyItems(cF, cS CurrencyItem, delimiter string) CurrencyPair {
	return CurrencyPair{
		FirstCurrency:  cF,
		SecondCurrency: cS,
		Delimiter:      delimiter,
	}
}

// Equal tests whether the provided currency items are equal
func (c CurrencyItem) Equal(b CurrencyItem) bool {
	return c.Upper() == b.Upper()
}

// EqualStdBitcoin tests whether the provided currency pairs are equal
func (p CurrencyPair) EqualStdBitcoin(b CurrencyPair, stdBTC bool) bool {
	if stdBTC {
		return p.StdBitcoinPair().ToUpper() == b.StdBitcoinPair().ToUpper()
	}
	return p.ToUpper() == b.ToUpper()
}
