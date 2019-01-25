package pair

import (
	"testing"
)

func TestFirstEqual(t *testing.T) {
	t.Parallel()
	pair := NewCurrencyPair("BTC", "USD")
	secondPair := NewCurrencyPair("BTC", "EUR")
	actual := pair.FirstEqual(secondPair)
	expected := true
	if actual != expected {
		t.Errorf(
			"Test failed. TestFirstEqual(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	secondPair = NewCurrencyPair("ETH", "EUR")
	actual = pair.FirstEqual(secondPair)
	expected = false
	if actual != expected {
		t.Errorf(
			"Test failed. TestFirstEqual(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}
}

func TestString(t *testing.T) {
	var pair CurrencyPair

	pair = NewCurrencyPair("BTC", "USD")
	expected := "BTCUSD"
	actual := pair.String()
	if actual != expected {
		t.Errorf("Test failed. TestString(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	pair = NewCurrencyPair("BTC", "USD")
	pair.Delimiter = ":"
	expected = "BTCUSD"
	actual = pair.String()
	if actual != expected {
		t.Errorf("Test failed. TestString(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}
}

func TestToUpper(t *testing.T) {
	var pair CurrencyPair
	pair = NewCurrencyPair("BTC", "USD")
	expected := "BTCUSD"
	actual := pair.String()
	if actual != expected {
		t.Errorf("Test failed. TestToUpper(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	pair = NewCurrencyPair("btc", "usd")
	pair.Delimiter = ":"
	expected = "BTC:USD"
	actual = pair.ToUpper()

	if actual != expected {
		t.Errorf("Test failed. TestToUpper(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	pair = NewCurrencyPair("btc", "usd")
	pair.Delimiter = ":"
	expected = "BTC:USD"
	actual = pair.ToUpper()

	if actual != expected {
		t.Errorf("Test failed. TestToUpper(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}
}

func TestToLower(t *testing.T) {
	var pair CurrencyPair
	pair = NewCurrencyPair("BTC", "USD")

	expected := "btcusd"
	actual := pair.ToLower()

	if actual != expected {
		t.Errorf("Test failed. TestToLower(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	pair = NewCurrencyPair("BTC", "USD")
	pair.Delimiter = ":"
	expected = "btc:usd"
	actual = pair.ToLower()
	if actual != expected {
		t.Errorf("Test failed. TestToLower(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	pair = NewCurrencyPair("BTC", "USD")
	pair.Delimiter = ":"
	expected = "btc:usd"
	actual = pair.ToLower()
	if actual != expected {
		t.Errorf("Test failed. TestToLower(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}
}

func TestStdCoinSymbol(t *testing.T) {
	var pair CurrencyItem

	pair = CurrencyItem("BTC")
	expected := "BTC"
	actual := pair.StdCoinSymbol()
	if actual != expected {
		t.Errorf("Test failed. TestStdBitcoinSymbol(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	pair = CurrencyItem("btc")
	expected = "BTC"
	actual = pair.StdCoinSymbol()
	if actual != expected {
		t.Errorf("Test failed. TestStdBitcoinSymbol(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	pair = CurrencyItem("XBT")
	expected = "BTC"
	actual = pair.StdCoinSymbol()
	if actual != expected {
		t.Errorf("Test failed. TestStdBitcoinSymbol(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	pair = CurrencyItem("xbt")
	expected = "BTC"
	actual = pair.StdCoinSymbol()
	if actual != expected {
		t.Errorf("Test failed. TestStdBitcoinSymbol(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	pair = CurrencyItem("ETH")
	expected = "ETH"
	actual = pair.StdCoinSymbol()
	if actual != expected {
		t.Errorf("Test failed. TestStdBitcoinSymbol(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	pair = CurrencyItem("eth")
	expected = "ETH"
	actual = pair.StdCoinSymbol()
	if actual != expected {
		t.Errorf("Test failed. TestStdBitcoinSymbol(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}
}

/*
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
*/
