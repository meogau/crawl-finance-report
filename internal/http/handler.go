package http

import (
	"craw-finance-report/internal/file"
	"craw-finance-report/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"net/http"
)

type Handler struct {
	FileService  *file.Service
	CrawlService *service.CrawlService
}

func (h *Handler) getExcelType1(c *gin.Context) {
	f := excelize.NewFile()
	data := h.CrawlService.GetDataBitcoin()
	f, err := h.FileService.GenerateExcelType1(data.Data[0], f, "BTCUSDT")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate Excel file"})
		return
	}
	data = h.CrawlService.GetDataGold()
	f, err = h.FileService.GenerateExcelType1(data.Data[0], f, "XAUUSD")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate Excel file"})
		return
	}
	// Write file to browser
	c.Header("Content-Disposition", `attachment; filename="data.xlsx"`)
	c.Header("Content-Type", "application/octet-stream")
	err = f.Write(c.Writer)
	if err != nil {
		return
	}
}

func (h *Handler) getExcelType2(c *gin.Context) {
	f := excelize.NewFile()
	data := h.CrawlService.GetDailyData("https://www.barchart.com/proxies/core-api/v1/quotes/get?fields=symbol%2CcontractSymbol%2CdailyLastPrice%2CdailyPriceChange%2CdailyOpenPrice%2CdailyHighPrice%2CdailyLowPrice%2CdailyPreviousPrice%2CdailyVolume%2CdailyOpenInterest%2CdailyDate1dAgo%2CsymbolCode%2CsymbolType%2ChasOptions&lists=futures.contractInRoot&root=CL&meta=field.shortName%2Cfield.type%2Cfield.description%2Clists.lastUpdate&hasOptions=true&page=1&limit=100&raw=1")
	f, err := h.FileService.GenerateExcelType2(data.Data, f, "CLQ24")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate Excel file"})
		return
	}
	data = h.CrawlService.GetDailyData("https://www.barchart.com/proxies/core-api/v1/quotes/get?fields=symbol%2CcontractSymbol%2CdailyLastPrice%2CdailyPriceChange%2CdailyOpenPrice%2CdailyHighPrice%2CdailyLowPrice%2CdailyPreviousPrice%2CdailyVolume%2CdailyOpenInterest%2CdailyDate1dAgo%2CsymbolCode%2CsymbolType%2ChasOptions&lists=futures.contractInRoot&root=DX&meta=field.shortName%2Cfield.type%2Cfield.description%2Clists.lastUpdate&hasOptions=true&page=1&limit=100&raw=1")
	f, err = h.FileService.GenerateExcelType2(data.Data, f, "DXY00")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate Excel file"})
		return
	}
	data = h.CrawlService.GetDailyData("https://www.barchart.com/proxies/core-api/v1/quotes/get?fields=symbol%2CcontractSymbol%2CdailyLastPrice%2CdailyPriceChange%2CdailyOpenPrice%2CdailyHighPrice%2CdailyLowPrice%2CdailyPreviousPrice%2CdailyVolume%2CdailyOpenInterest%2CdailyDate1dAgo%2CsymbolCode%2CsymbolType%2ChasOptions&lists=futures.contractInRoot&root=ZC&meta=field.shortName%2Cfield.type%2Cfield.description%2Clists.lastUpdate&hasOptions=true&page=1&limit=100&raw=1")
	f, err = h.FileService.GenerateExcelType2(data.Data, f, "ZCZ24")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate Excel file"})
		return
	}
	data = h.CrawlService.GetDailyData("https://www.barchart.com/proxies/core-api/v1/quotes/get?fields=symbol%2CcontractSymbol%2CdailyLastPrice%2CdailyPriceChange%2CdailyOpenPrice%2CdailyHighPrice%2CdailyLowPrice%2CdailyPreviousPrice%2CdailyVolume%2CdailyOpenInterest%2CdailyDate1dAgo%2CsymbolCode%2CsymbolType%2ChasOptions&lists=futures.contractInRoot&root=ZM&meta=field.shortName%2Cfield.type%2Cfield.description%2Clists.lastUpdate&hasOptions=true&page=1&limit=100&raw=1")
	f, err = h.FileService.GenerateExcelType2(data.Data, f, "ZMZ24")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate Excel file"})
		return
	}
	data = h.CrawlService.GetDailyData("https://www.barchart.com/proxies/core-api/v1/quotes/get?fields=symbol%2CcontractSymbol%2CdailyLastPrice%2CdailyPriceChange%2CdailyOpenPrice%2CdailyHighPrice%2CdailyLowPrice%2CdailyPreviousPrice%2CdailyVolume%2CdailyOpenInterest%2CdailyDate1dAgo%2CsymbolCode%2CsymbolType%2ChasOptions&lists=futures.contractInRoot&root=ZL&meta=field.shortName%2Cfield.type%2Cfield.description%2Clists.lastUpdate&hasOptions=true&page=1&limit=100&raw=1")
	f, err = h.FileService.GenerateExcelType2(data.Data, f, "ZLZ24")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate Excel file"})
		return
	}
	data = h.CrawlService.GetDailyData("https://www.barchart.com/proxies/core-api/v1/quotes/get?fields=symbol%2CcontractSymbol%2CdailyLastPrice%2CdailyPriceChange%2CdailyOpenPrice%2CdailyHighPrice%2CdailyLowPrice%2CdailyPreviousPrice%2CdailyVolume%2CdailyOpenInterest%2CdailyDate1dAgo%2CsymbolCode%2CsymbolType%2ChasOptions&lists=futures.contractInRoot&root=ZS&meta=field.shortName%2Cfield.type%2Cfield.description%2Clists.lastUpdate&hasOptions=true&page=1&limit=100&raw=1")
	f, err = h.FileService.GenerateExcelType2(data.Data, f, "ZSX24")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate Excel file"})
		return
	}
	data = h.CrawlService.GetDailyData("https://www.barchart.com/proxies/core-api/v1/quotes/get?fields=symbol%2CcontractSymbol%2CdailyLastPrice%2CdailyPriceChange%2CdailyOpenPrice%2CdailyHighPrice%2CdailyLowPrice%2CdailyPreviousPrice%2CdailyVolume%2CdailyOpenInterest%2CdailyDate1dAgo%2CsymbolCode%2CsymbolType%2ChasOptions&lists=futures.contractInRoot&root=ZW&meta=field.shortName%2Cfield.type%2Cfield.description%2Clists.lastUpdate&hasOptions=true&page=1&limit=100&raw=1")
	f, err = h.FileService.GenerateExcelType2(data.Data, f, "ZWU24")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate Excel file"})
		return
	}
	data = h.CrawlService.GetDailyData("https://www.barchart.com/proxies/core-api/v1/quotes/get?fields=symbol%2CcontractSymbol%2CdailyLastPrice%2CdailyPriceChange%2CdailyOpenPrice%2CdailyHighPrice%2CdailyLowPrice%2CdailyPreviousPrice%2CdailyVolume%2CdailyOpenInterest%2CdailyDate1dAgo%2CsymbolCode%2CsymbolType%2ChasOptions&lists=futures.contractInRoot&root=KE&meta=field.shortName%2Cfield.type%2Cfield.description%2Clists.lastUpdate&hasOptions=true&page=1&limit=100&raw=1")
	f, err = h.FileService.GenerateExcelType2(data.Data, f, "KEU24")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate Excel file"})
		return
	}
	// Write file to browser
	c.Header("Content-Disposition", `attachment; filename="data.xlsx"`)
	c.Header("Content-Type", "application/octet-stream")
	err = f.Write(c.Writer)
	if err != nil {
		return
	}
}

func (h *Handler) getExcelType3(c *gin.Context) {
	f := excelize.NewFile()
	data := h.CrawlService.GetTradingData("https://www.barchart.com/proxies/core-api/v1/quotes/get?fields=symbol%2CcontractSymbol%2ClastPrice%2CpriceChange%2CopenPrice%2ChighPrice%2ClowPrice%2CpreviousPrice%2Cvolume%2CopenInterest%2CtradeTime%2CsymbolCode%2CsymbolType%2ChasOptions&lists=futures.contractInRoot&root=CL&meta=field.shortName%2Cfield.type%2Cfield.description%2Clists.lastUpdate&hasOptions=true&page=1&limit=100&raw=1")
	f, err := h.FileService.GenerateExcelType3(data.Data, f, "CLQ24")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate Excel file"})
		return
	}
	data = h.CrawlService.GetTradingData("https://www.barchart.com/proxies/core-api/v1/quotes/get?fields=symbol%2CcontractSymbol%2ClastPrice%2CpriceChange%2CopenPrice%2ChighPrice%2ClowPrice%2CpreviousPrice%2Cvolume%2CopenInterest%2CtradeTime%2CsymbolCode%2CsymbolType%2ChasOptions&lists=futures.contractInRoot&root=DX&meta=field.shortName%2Cfield.type%2Cfield.description%2Clists.lastUpdate&hasOptions=true&page=1&limit=100&raw=1")
	f, err = h.FileService.GenerateExcelType3(data.Data, f, "DXY00")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate Excel file"})
		return
	}
	data = h.CrawlService.GetTradingData("https://www.barchart.com/proxies/core-api/v1/quotes/get?fields=symbol%2CcontractSymbol%2ClastPrice%2CpriceChange%2CopenPrice%2ChighPrice%2ClowPrice%2CpreviousPrice%2Cvolume%2CopenInterest%2CtradeTime%2CsymbolCode%2CsymbolType%2ChasOptions&lists=futures.contractInRoot&root=ZC&meta=field.shortName%2Cfield.type%2Cfield.description%2Clists.lastUpdate&hasOptions=true&page=1&limit=100&raw=1")
	f, err = h.FileService.GenerateExcelType3(data.Data, f, "ZCZ24")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate Excel file"})
		return
	}
	data = h.CrawlService.GetTradingData("https://www.barchart.com/proxies/core-api/v1/quotes/get?fields=symbol%2CcontractSymbol%2ClastPrice%2CpriceChange%2CopenPrice%2ChighPrice%2ClowPrice%2CpreviousPrice%2Cvolume%2CopenInterest%2CtradeTime%2CsymbolCode%2CsymbolType%2ChasOptions&lists=futures.contractInRoot&root=ZM&meta=field.shortName%2Cfield.type%2Cfield.description%2Clists.lastUpdate&hasOptions=true&page=1&limit=100&raw=1")
	f, err = h.FileService.GenerateExcelType3(data.Data, f, "ZMZ24")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate Excel file"})
		return
	}
	data = h.CrawlService.GetTradingData("https://www.barchart.com/proxies/core-api/v1/quotes/get?fields=symbol%2CcontractSymbol%2ClastPrice%2CpriceChange%2CopenPrice%2ChighPrice%2ClowPrice%2CpreviousPrice%2Cvolume%2CopenInterest%2CtradeTime%2CsymbolCode%2CsymbolType%2ChasOptions&lists=futures.contractInRoot&root=ZL&meta=field.shortName%2Cfield.type%2Cfield.description%2Clists.lastUpdate&hasOptions=true&page=1&limit=100&raw=1")
	f, err = h.FileService.GenerateExcelType3(data.Data, f, "ZLZ24")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate Excel file"})
		return
	}
	data = h.CrawlService.GetTradingData("https://www.barchart.com/proxies/core-api/v1/quotes/get?fields=symbol%2CcontractSymbol%2ClastPrice%2CpriceChange%2CopenPrice%2ChighPrice%2ClowPrice%2CpreviousPrice%2Cvolume%2CopenInterest%2CtradeTime%2CsymbolCode%2CsymbolType%2ChasOptions&lists=futures.contractInRoot&root=ZS&meta=field.shortName%2Cfield.type%2Cfield.description%2Clists.lastUpdate&hasOptions=true&page=1&limit=100&raw=1")
	f, err = h.FileService.GenerateExcelType3(data.Data, f, "ZSX24")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate Excel file"})
		return
	}
	data = h.CrawlService.GetTradingData("https://www.barchart.com/proxies/core-api/v1/quotes/get?fields=symbol%2CcontractSymbol%2ClastPrice%2CpriceChange%2CopenPrice%2ChighPrice%2ClowPrice%2CpreviousPrice%2Cvolume%2CopenInterest%2CtradeTime%2CsymbolCode%2CsymbolType%2ChasOptions&lists=futures.contractInRoot&root=ZW&meta=field.shortName%2Cfield.type%2Cfield.description%2Clists.lastUpdate&hasOptions=true&page=1&limit=100&raw=1")
	f, err = h.FileService.GenerateExcelType3(data.Data, f, "ZWU24")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate Excel file"})
		return
	}
	data = h.CrawlService.GetTradingData("https://www.barchart.com/proxies/core-api/v1/quotes/get?fields=symbol%2CcontractSymbol%2ClastPrice%2CpriceChange%2CopenPrice%2ChighPrice%2ClowPrice%2CpreviousPrice%2Cvolume%2CopenInterest%2CtradeTime%2CsymbolCode%2CsymbolType%2ChasOptions&lists=futures.contractInRoot&root=KE&meta=field.shortName%2Cfield.type%2Cfield.description%2Clists.lastUpdate&hasOptions=true&page=1&limit=100&raw=1")
	f, err = h.FileService.GenerateExcelType3(data.Data, f, "KEU24")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate Excel file"})
		return
	}
	// Write file to browser
	c.Header("Content-Disposition", `attachment; filename="data.xlsx"`)
	c.Header("Content-Type", "application/octet-stream")
	err = f.Write(c.Writer)
	if err != nil {
		return
	}
}
