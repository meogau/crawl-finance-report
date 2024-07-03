package service

import (
	"craw-finance-report/pkg/entity"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type CrawlService struct {
	AuthService *AuthService
}

func (c *CrawlService) GetDataBitcoin() entity.BasicResponse {
	url := "https://www.barchart.com/proxies/core-api/v1/quotes/get?symbols=%5EBTCUSDT&fields=date1dAgo%2Copen1dAgo%2Chigh1dAgo%2Clow1dAgo%2ClastPrice1dAgo%2CpriceChange1dAgo%2CpercentChange1dAgo%2Cvolume1dAgo%2Cdate2dAgo%2Copen2dAgo%2Chigh2dAgo%2Clow2dAgo%2ClastPrice2dAgo%2CpriceChange2dAgo%2CpercentChange2dAgo%2Cvolume2dAgo%2Cdate3dAgo%2Copen3dAgo%2Chigh3dAgo%2Clow3dAgo%2ClastPrice3dAgo%2CpriceChange3dAgo%2CpercentChange3dAgo%2Cvolume3dAgo%2Cdate4dAgo%2Copen4dAgo%2Chigh4dAgo%2Clow4dAgo%2ClastPrice4dAgo%2CpriceChange4dAgo%2CpercentChange4dAgo%2Cvolume4dAgo%2Cdate5dAgo%2Copen5dAgo%2Chigh5dAgo%2Clow5dAgo%2ClastPrice5dAgo%2CpriceChange5dAgo%2CpercentChange5dAgo%2Cvolume5dAgo&meta=field.shortName&raw=1"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}
	req.Header.Set("accept", "application/json")
	cookie, token := c.AuthService.GetCookieAndToken()
	req.Header.Set("cookie", cookie)
	req.Header.Set("x-xsrf-token", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
	}
	var responseData entity.BasicResponse
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON response: %v", err)
	}
	return responseData
}

func (c *CrawlService) GetDataGold() entity.BasicResponse {
	url := "https://www.barchart.com/proxies/core-api/v1/quotes/get?symbols=%5EXAUUSD&fields=date1dAgo%2Copen1dAgo%2Chigh1dAgo%2Clow1dAgo%2ClastPrice1dAgo%2CpriceChange1dAgo%2CpercentChange1dAgo%2Cvolume1dAgo%2Cdate2dAgo%2Copen2dAgo%2Chigh2dAgo%2Clow2dAgo%2ClastPrice2dAgo%2CpriceChange2dAgo%2CpercentChange2dAgo%2Cvolume2dAgo%2Cdate3dAgo%2Copen3dAgo%2Chigh3dAgo%2Clow3dAgo%2ClastPrice3dAgo%2CpriceChange3dAgo%2CpercentChange3dAgo%2Cvolume3dAgo%2Cdate4dAgo%2Copen4dAgo%2Chigh4dAgo%2Clow4dAgo%2ClastPrice4dAgo%2CpriceChange4dAgo%2CpercentChange4dAgo%2Cvolume4dAgo%2Cdate5dAgo%2Copen5dAgo%2Chigh5dAgo%2Clow5dAgo%2ClastPrice5dAgo%2CpriceChange5dAgo%2CpercentChange5dAgo%2Cvolume5dAgo&meta=field.shortName&raw=1"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}
	req.Header.Set("accept", "application/json")
	cookie, token := c.AuthService.GetCookieAndToken()
	req.Header.Set("cookie", cookie)
	req.Header.Set("x-xsrf-token", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
	}
	var responseData entity.BasicResponse
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON response: %v", err)
	}
	return responseData
}

func (c *CrawlService) GetDailyData(url string) entity.DailyDataResponse {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}
	// Set headers
	req.Header.Set("accept", "application/json")
	cookie, token := c.AuthService.GetCookieAndToken()
	req.Header.Set("cookie", cookie)
	req.Header.Set("x-xsrf-token", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
	}
	var responseData entity.DailyDataResponse
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON response: %v", err)
	}
	return responseData
}

func (c *CrawlService) GetTradingData(url string) entity.TradingResponse {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}
	// Set headers
	req.Header.Set("accept", "application/json")
	cookie, token := c.AuthService.GetCookieAndToken()
	req.Header.Set("cookie", cookie)
	req.Header.Set("x-xsrf-token", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
	}
	var responseData entity.TradingResponse
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON response: %v", err)
	}
	return responseData
}
