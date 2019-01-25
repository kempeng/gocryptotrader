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
	var item CurrencyItem

	item = CurrencyItem("BTC")
	expected := "BTC"
	actual := item.StdCoinSymbol()
	if actual != expected {
		t.Errorf("Test failed. TestStdBitcoinSymbol(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	item = CurrencyItem("btc")
	expected = "BTC"
	actual = item.StdCoinSymbol()
	if actual != expected {
		t.Errorf("Test failed. TestStdBitcoinSymbol(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	item = CurrencyItem("XBT")
	expected = "BTC"
	actual = item.StdCoinSymbol()
	if actual != expected {
		t.Errorf("Test failed. TestStdBitcoinSymbol(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	item = CurrencyItem("xbt")
	expected = "BTC"
	actual = item.StdCoinSymbol()
	if actual != expected {
		t.Errorf("Test failed. TestStdBitcoinSymbol(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	item = CurrencyItem("ETH")
	expected = "ETH"
	actual = item.StdCoinSymbol()
	if actual != expected {
		t.Errorf("Test failed. TestStdBitcoinSymbol(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	item = CurrencyItem("eth")
	expected = "ETH"
	actual = item.StdCoinSymbol()
	if actual != expected {
		t.Errorf("Test failed. TestStdBitcoinSymbol(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}
}

func TestStdBitcoin(t *testing.T) {
	var item CurrencyItem

	item = CurrencyItem("BTC")
	expected := CurrencyItem("BTC")
	actual := item.StdBitcoin()
	if actual != expected {
		t.Errorf("Test failed. TestStdBitcoin(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	item = CurrencyItem("btc")
	expected = CurrencyItem("BTC")
	actual = item.StdBitcoin()
	if actual != expected {
		t.Errorf("Test failed. TestStdBitcoin(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	item = CurrencyItem("XBT")
	expected = CurrencyItem("BTC")
	actual = item.StdBitcoin()
	if actual != expected {
		t.Errorf("Test failed. TestStdBitcoin(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	item = CurrencyItem("xbt")
	expected = CurrencyItem("BTC")
	actual = item.StdBitcoin()
	if actual != expected {
		t.Errorf("Test failed. TestStdBitcoin(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	item = CurrencyItem("ETH")
	expected = CurrencyItem("ETH")
	actual = item.StdBitcoin()
	if actual != expected {
		t.Errorf("Test failed. TestStdBitcoin(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}
}

func TestStdCoinPair(t *testing.T) {

	pair := NewCurrencyPair("BTC", "EUR")
	expected := NewCurrencyPair("BTC", "EUR")
	actual := pair.StdBitcoinPair()
	if actual != expected {
		t.Errorf("Test failed. TestStdBitcoinPair(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	pair = NewCurrencyPair("XBT", "EUR")
	expected = NewCurrencyPair("BTC", "EUR")
	actual = pair.StdBitcoinPair()
	if actual != expected {
		t.Errorf("Test failed. TestStdBitcoinPair(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	pair = NewCurrencyPair("EUR", "BTC")
	expected = NewCurrencyPair("EUR", "BTC")
	actual = pair.StdBitcoinPair()
	if actual != expected {
		t.Errorf("Test failed. TestStdBitcoinPair(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	pair = NewCurrencyPair("EUR", "XBT")
	expected = NewCurrencyPair("EUR", "BTC")
	actual = pair.StdBitcoinPair()
	if actual != expected {
		t.Errorf("Test failed. TestStdBitcoinPair(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	pair = NewCurrencyPair("XBT", "XBT")
	expected = NewCurrencyPair("BTC", "BTC")
	actual = pair.StdBitcoinPair()
	if actual != expected {
		t.Errorf("Test failed. TestStdBitcoinPair(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}
}

func TestChangeBitcoin(t *testing.T) {

	item := CurrencyItem("BTC")
	expected := CurrencyItem("BTC")
	actual := item.ChangeBitcoin(true)
	if actual != expected {
		t.Errorf("Test failed. TestChangeBitcoin(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	item = CurrencyItem("btc")
	expected = CurrencyItem("btc")
	actual = item.ChangeBitcoin(true)
	if actual != expected {
		t.Errorf("Test failed. TestChangeBitcoin(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	item = CurrencyItem("XBT")
	expected = CurrencyItem("BTC")
	actual = item.ChangeBitcoin(true)
	if actual != expected {
		t.Errorf("Test failed. TestChangeBitcoin(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	item = CurrencyItem("xbt")
	expected = CurrencyItem("BTC")
	actual = item.ChangeBitcoin(true)
	if actual != expected {
		t.Errorf("Test failed. TestChangeBitcoin(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	item = CurrencyItem("ETH")
	expected = CurrencyItem("ETH")
	actual = item.ChangeBitcoin(true)
	if actual != expected {
		t.Errorf("Test failed. TestChangeBitcoin(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	item = CurrencyItem("BTC")
	expected = CurrencyItem("XBT")
	actual = item.ChangeBitcoin(false)
	if actual != expected {
		t.Errorf("Test failed. TestChangeBitcoin(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	item = CurrencyItem("btc")
	expected = CurrencyItem("XBT")
	actual = item.ChangeBitcoin(false)
	if actual != expected {
		t.Errorf("Test failed. TestChangeBitcoin(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	item = CurrencyItem("XBT")
	expected = CurrencyItem("XBT")
	actual = item.ChangeBitcoin(false)
	if actual != expected {
		t.Errorf("Test failed. TestChangeBitcoin(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	item = CurrencyItem("xbt")
	expected = CurrencyItem("xbt")
	actual = item.ChangeBitcoin(false)
	if actual != expected {
		t.Errorf("Test failed. TestChangeBitcoin(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	item = CurrencyItem("ETH")
	expected = CurrencyItem("ETH")
	actual = item.ChangeBitcoin(false)
	if actual != expected {
		t.Errorf("Test failed. TestChangeBitcoin(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

}

func TestIsCrypto(t *testing.T) {

	item := CurrencyItem("BTC")
	expected := true
	actual := item.IsCrypto()
	if actual != expected {
		t.Errorf("Test failed. TestIsCrypto(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	item = CurrencyItem("ETH")
	expected = true
	actual = item.IsCrypto()
	if actual != expected {
		t.Errorf("Test failed. TestIsCrypto(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	item = CurrencyItem("btc")
	expected = true
	actual = item.IsCrypto()
	if actual != expected {
		t.Errorf("Test failed. TestIsCrypto(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	item = CurrencyItem("USD")
	expected = false
	actual = item.IsCrypto()
	if actual != expected {
		t.Errorf("Test failed. TestIsCrypto(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	item = CurrencyItem("eur")
	expected = false
	actual = item.IsCrypto()
	if actual != expected {
		t.Errorf("Test failed. TestIsCrypto(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}
}

func TestIsBitcoin(t *testing.T) {

	item := CurrencyItem("BTC")
	expected := true
	actual := item.IsBitcoin()
	if actual != expected {
		t.Errorf("Test failed. TestIsBitcoin(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	item = CurrencyItem("XBT")
	expected = true
	actual = item.IsBitcoin()
	if actual != expected {
		t.Errorf("Test failed. TestIsBitcoin()(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	item = CurrencyItem("btc")
	expected = true
	actual = item.IsBitcoin()
	if actual != expected {
		t.Errorf("Test failed. TestIsBitcoin(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	item = CurrencyItem("USD")
	expected = false
	actual = item.IsBitcoin()
	if actual != expected {
		t.Errorf("Test failed. TestIsBitcoin(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	item = CurrencyItem("eur")
	expected = false
	actual = item.IsBitcoin()
	if actual != expected {
		t.Errorf("Test failed. TestIsBitcoin(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}
}

func TestIsAltBitcoin(t *testing.T) {

	item := CurrencyItem("BTC")
	expected := false
	actual := item.IsAltBitcoin()
	if actual != expected {
		t.Errorf("Test failed. TestIsAltBitcoin(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	item = CurrencyItem("XBT")
	expected = true
	actual = item.IsAltBitcoin()
	if actual != expected {
		t.Errorf("Test failed. TestIsAltBitcoin(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	item = CurrencyItem("xbt")
	expected = true
	actual = item.IsAltBitcoin()
	if actual != expected {
		t.Errorf("Test failed. TestIsAltBitcoin(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}

	item = CurrencyItem("USD")
	expected = false
	actual = item.IsAltBitcoin()
	if actual != expected {
		t.Errorf("Test failed. TestIsAltBitcoin(): %v was not equal to expected value: %v",
			actual, expected,
		)
	}
}


/* TODO
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

