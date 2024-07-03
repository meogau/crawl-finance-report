package entity

type DailyData struct {
	Symbol             string `json:"symbol"`
	ContractSymbol     string `json:"contractSymbol"`
	DailyLastPrice     string `json:"dailyLastPrice"`
	DailyPriceChange   string `json:"dailyPriceChange"`
	DailyOpenPrice     string `json:"dailyOpenPrice"`
	DailyHighPrice     string `json:"dailyHighPrice"`
	DailyLowPrice      string `json:"dailyLowPrice"`
	DailyPreviousPrice string `json:"dailyPreviousPrice"`
	DailyVolume        string `json:"dailyVolume"`
	DailyOpenInterest  string `json:"dailyOpenInterest"`
	DailyDate1dAgo     string `json:"dailyDate1dAgo"`
	SymbolCode         string `json:"symbolCode"`
	SymbolType         int    `json:"symbolType"`
	HasOptions         string `json:"hasOptions"`
}

type RawData struct {
	Date1dAgo          string  `json:"date1dAgo"`
	Open1dAgo          float64 `json:"open1dAgo"`
	High1dAgo          float64 `json:"high1dAgo"`
	Low1dAgo           float64 `json:"low1dAgo"`
	LastPrice1dAgo     float64 `json:"lastPrice1dAgo"`
	PriceChange1dAgo   float64 `json:"priceChange1dAgo"`
	PercentChange1dAgo float64 `json:"percentChange1dAgo"`
	Volume1dAgo        int     `json:"volume1dAgo"`
	Date2dAgo          string  `json:"date2dAgo"`
	Open2dAgo          float64 `json:"open2dAgo"`
	High2dAgo          float64 `json:"high2dAgo"`
	Low2dAgo           float64 `json:"low2dAgo"`
	LastPrice2dAgo     float64 `json:"lastPrice2dAgo"`
	PriceChange2dAgo   float64 `json:"priceChange2dAgo"`
	PercentChange2dAgo float64 `json:"percentChange2dAgo"`
	Volume2dAgo        int     `json:"volume2dAgo"`
	Date3dAgo          string  `json:"date3dAgo"`
	Open3dAgo          float64 `json:"open3dAgo"`
	High3dAgo          float64 `json:"high3dAgo"`
	Low3dAgo           float64 `json:"low3dAgo"`
	LastPrice3dAgo     float64 `json:"lastPrice3dAgo"`
	PriceChange3dAgo   float64 `json:"priceChange3dAgo"`
	PercentChange3dAgo float64 `json:"percentChange3dAgo"`
	Volume3dAgo        int     `json:"volume3dAgo"`
	Date4dAgo          string  `json:"date4dAgo"`
	Open4dAgo          float64 `json:"open4dAgo"`
	High4dAgo          float64 `json:"high4dAgo"`
	Low4dAgo           float64 `json:"low4dAgo"`
	LastPrice4dAgo     float64 `json:"lastPrice4dAgo"`
	PriceChange4dAgo   float64 `json:"priceChange4dAgo"`
	PercentChange4dAgo float64 `json:"percentChange4dAgo"`
	Volume4dAgo        int     `json:"volume4dAgo"`
	Date5dAgo          string  `json:"date5dAgo"`
	Open5dAgo          float64 `json:"open5dAgo"`
	High5dAgo          float64 `json:"high5dAgo"`
	Low5dAgo           float64 `json:"low5dAgo"`
	LastPrice5dAgo     float64 `json:"lastPrice5dAgo"`
	PriceChange5dAgo   float64 `json:"priceChange5dAgo"`
	PercentChange5dAgo float64 `json:"percentChange5dAgo"`
	Volume5dAgo        int     `json:"volume5dAgo"`
}

type BasicData struct {
	Date1dAgo          string  `json:"date1dAgo"`
	Open1dAgo          string  `json:"open1dAgo"`
	High1dAgo          string  `json:"high1dAgo"`
	Low1dAgo           string  `json:"low1dAgo"`
	LastPrice1dAgo     string  `json:"lastPrice1dAgo"`
	PriceChange1dAgo   string  `json:"priceChange1dAgo"`
	PercentChange1dAgo string  `json:"percentChange1dAgo"`
	Volume1dAgo        string  `json:"volume1dAgo"`
	Date2dAgo          string  `json:"date2dAgo"`
	Open2dAgo          string  `json:"open2dAgo"`
	High2dAgo          string  `json:"high2dAgo"`
	Low2dAgo           string  `json:"low2dAgo"`
	LastPrice2dAgo     string  `json:"lastPrice2dAgo"`
	PriceChange2dAgo   string  `json:"priceChange2dAgo"`
	PercentChange2dAgo string  `json:"percentChange2dAgo"`
	Volume2dAgo        string  `json:"volume2dAgo"`
	Date3dAgo          string  `json:"date3dAgo"`
	Open3dAgo          string  `json:"open3dAgo"`
	High3dAgo          string  `json:"high3dAgo"`
	Low3dAgo           string  `json:"low3dAgo"`
	LastPrice3dAgo     string  `json:"lastPrice3dAgo"`
	PriceChange3dAgo   string  `json:"priceChange3dAgo"`
	PercentChange3dAgo string  `json:"percentChange3dAgo"`
	Volume3dAgo        string  `json:"volume3dAgo"`
	Date4dAgo          string  `json:"date4dAgo"`
	Open4dAgo          string  `json:"open4dAgo"`
	High4dAgo          string  `json:"high4dAgo"`
	Low4dAgo           string  `json:"low4dAgo"`
	LastPrice4dAgo     string  `json:"lastPrice4dAgo"`
	PriceChange4dAgo   string  `json:"priceChange4dAgo"`
	PercentChange4dAgo string  `json:"percentChange4dAgo"`
	Volume4dAgo        string  `json:"volume4dAgo"`
	Date5dAgo          string  `json:"date5dAgo"`
	Open5dAgo          string  `json:"open5dAgo"`
	High5dAgo          string  `json:"high5dAgo"`
	Low5dAgo           string  `json:"low5dAgo"`
	LastPrice5dAgo     string  `json:"lastPrice5dAgo"`
	PriceChange5dAgo   string  `json:"priceChange5dAgo"`
	PercentChange5dAgo string  `json:"percentChange5dAgo"`
	Volume5dAgo        string  `json:"volume5dAgo"`
	Raw                RawData `json:"raw"`
}

type ShortName struct {
	Date1dAgo          string `json:"date1dAgo"`
	Open1dAgo          string `json:"open1dAgo"`
	High1dAgo          string `json:"high1dAgo"`
	Low1dAgo           string `json:"low1dAgo"`
	LastPrice1dAgo     string `json:"lastPrice1dAgo"`
	PriceChange1dAgo   string `json:"priceChange1dAgo"`
	PercentChange1dAgo string `json:"percentChange1dAgo"`
	Volume1dAgo        string `json:"volume1dAgo"`
	Date2dAgo          string `json:"date2dAgo"`
	Open2dAgo          string `json:"open2dAgo"`
	High2dAgo          string `json:"high2dAgo"`
	Low2dAgo           string `json:"low2dAgo"`
	LastPrice2dAgo     string `json:"lastPrice2dAgo"`
	PriceChange2dAgo   string `json:"priceChange2dAgo"`
	PercentChange2dAgo string `json:"percentChange2dAgo"`
	Volume2dAgo        string `json:"volume2dAgo"`
	Date3dAgo          string `json:"date3dAgo"`
	Open3dAgo          string `json:"open3dAgo"`
	High3dAgo          string `json:"high3dAgo"`
	Low3dAgo           string `json:"low3dAgo"`
	LastPrice3dAgo     string `json:"lastPrice3dAgo"`
	PriceChange3dAgo   string `json:"priceChange3dAgo"`
	PercentChange3dAgo string `json:"percentChange3dAgo"`
	Volume3dAgo        string `json:"volume3dAgo"`
	Date4dAgo          string `json:"date4dAgo"`
	Open4dAgo          string `json:"open4dAgo"`
	High4dAgo          string `json:"high4dAgo"`
	Low4dAgo           string `json:"low4dAgo"`
	LastPrice4dAgo     string `json:"lastPrice4dAgo"`
	PriceChange4dAgo   string `json:"priceChange4dAgo"`
	PercentChange4dAgo string `json:"percentChange4dAgo"`
	Volume4dAgo        string `json:"volume4dAgo"`
	Date5dAgo          string `json:"date5dAgo"`
	Open5dAgo          string `json:"open5dAgo"`
	High5dAgo          string `json:"high5dAgo"`
	Low5dAgo           string `json:"low5dAgo"`
	LastPrice5dAgo     string `json:"lastPrice5dAgo"`
	PriceChange5dAgo   string `json:"priceChange5dAgo"`
	PercentChange5dAgo string `json:"percentChange5dAgo"`
	Volume5dAgo        string `json:"volume5dAgo"`
}

type Field struct {
	ShortName ShortName `json:"shortName"`
}

type Meta struct {
	Field Field `json:"field"`
}

type TradingData struct {
	Symbol         string `json:"symbol"`
	ContractSymbol string `json:"contractSymbol,omitempty"`
	LastPrice      string `json:"lastPrice"`
	PriceChange    string `json:"priceChange"`
	OpenPrice      string `json:"openPrice"`
	HighPrice      string `json:"highPrice"`
	LowPrice       string `json:"lowPrice"`
	PreviousPrice  string `json:"previousPrice"`
	Volume         string `json:"volume"`
	OpenInterest   string `json:"openInterest"`
	TradeTime      string `json:"tradeTime"`
	SymbolCode     string `json:"symbolCode"`
	SymbolType     int    `json:"symbolType"`
	HasOptions     string `json:"hasOptions"`
}
