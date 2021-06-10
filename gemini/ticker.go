package gemini

import (
	"encoding/json"
	"io/ioutil"
)

type Symbol string

type Ticker struct {
	Symbol  string    `json:"symbol"`
	Open    float64   `json:"open,string"`
	High    float64   `json:"high,string"`
	Low     float64   `json:"low,string"`
	Close   float64   `json:"close,string"`
	Bid     float64   `json:"bid,string"`
	Ask     float64   `json:"ask,string"`
	Changes []string  `json:"changes"`
}

func (c Client) GetTickerSymbols() []Symbol {
	c.logger.Debug("Retrieving all ticket symbols.")
	url := c.baseURLv1 + "/symbols"
	resp, err := c.httpClient.Get(url)
	if err != nil {
		c.logger.Fatalf("Failed to retrieve all symbols: %s", err)
	}
	defer resp.Body.Close() // Best effort.

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.logger.Fatalf("Failed to read body: %e", err)
	}

	var symbols []Symbol
	if err := json.Unmarshal(body, &symbols); err != nil {
		c.logger.Fatalf("Failed to unmarshal body into symbols: %s", err)
	}
	c.logger.Debugf("Retrieved ticker symbols: %#v", symbols)

	return symbols
}

func (c Client) GetTicker(symbol Symbol) Ticker {
	url := c.baseURLv2 + "/ticker/" + string(symbol)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		c.logger.Fatalf("Failed to retrieve ticker: %s", err)
	}
	defer resp.Body.Close() // Best effort.

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.logger.Fatalf("Failed to read body: %e", err)
	}

	var ticker Ticker
	if err := json.Unmarshal(body, &ticker); err != nil {
		c.logger.Fatalf("Failed to unmarshal body into ticker: %s", err)
	}
	c.logger.Debugf("Retrieved ticker: %#v", ticker)

	return ticker
}
