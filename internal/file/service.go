package file

import (
	"craw-finance-report/pkg/entity"
	"github.com/xuri/excelize/v2"
	"strconv"
)

type Service struct {
}

func (s *Service) GenerateExcelType1(data entity.BasicData, f *excelize.File, sheet string) (*excelize.File, error) {
	f.NewSheet(sheet)
	f.SetColWidth(sheet, "A", "H", 20)

	headers := []string{"Date", "Open", "High", "Low", "Last", "Change", "%Chg", "Volume"}
	writeHeader(f, sheet, headers)

	f.SetCellValue(sheet, "A2", data.Date1dAgo)
	f.SetCellValue(sheet, "A3", data.Date2dAgo)
	f.SetCellValue(sheet, "A4", data.Date3dAgo)
	f.SetCellValue(sheet, "A5", data.Date4dAgo)
	f.SetCellValue(sheet, "A6", data.Date5dAgo)

	f.SetCellValue(sheet, "B2", data.Open1dAgo)
	f.SetCellValue(sheet, "B3", data.Open2dAgo)
	f.SetCellValue(sheet, "B4", data.Open3dAgo)
	f.SetCellValue(sheet, "B5", data.Open4dAgo)
	f.SetCellValue(sheet, "B6", data.Open5dAgo)

	f.SetCellValue(sheet, "C2", data.High1dAgo)
	f.SetCellValue(sheet, "C3", data.High2dAgo)
	f.SetCellValue(sheet, "C4", data.High3dAgo)
	f.SetCellValue(sheet, "C5", data.High4dAgo)
	f.SetCellValue(sheet, "C6", data.High5dAgo)

	f.SetCellValue(sheet, "D2", data.Low1dAgo)
	f.SetCellValue(sheet, "D3", data.Low2dAgo)
	f.SetCellValue(sheet, "D4", data.Low3dAgo)
	f.SetCellValue(sheet, "D5", data.Low4dAgo)
	f.SetCellValue(sheet, "D6", data.Low5dAgo)

	f.SetCellValue(sheet, "E2", data.LastPrice1dAgo)
	f.SetCellValue(sheet, "E3", data.LastPrice2dAgo)
	f.SetCellValue(sheet, "E4", data.LastPrice3dAgo)
	f.SetCellValue(sheet, "E5", data.LastPrice4dAgo)
	f.SetCellValue(sheet, "E6", data.LastPrice5dAgo)

	f.SetCellValue(sheet, "F2", data.PriceChange1dAgo)
	f.SetCellValue(sheet, "F3", data.PriceChange2dAgo)
	f.SetCellValue(sheet, "F4", data.PriceChange3dAgo)
	f.SetCellValue(sheet, "F5", data.PriceChange4dAgo)
	f.SetCellValue(sheet, "F6", data.PriceChange5dAgo)

	f.SetCellValue(sheet, "G2", data.PercentChange1dAgo)
	f.SetCellValue(sheet, "G3", data.PercentChange2dAgo)
	f.SetCellValue(sheet, "G4", data.PercentChange3dAgo)
	f.SetCellValue(sheet, "G5", data.PercentChange4dAgo)
	f.SetCellValue(sheet, "G6", data.PercentChange5dAgo)

	f.SetCellValue(sheet, "H2", data.Volume1dAgo)
	f.SetCellValue(sheet, "H3", data.Volume2dAgo)
	f.SetCellValue(sheet, "H4", data.Volume3dAgo)
	f.SetCellValue(sheet, "H5", data.Volume4dAgo)
	f.SetCellValue(sheet, "H6", data.Volume5dAgo)

	return f, nil
}

func (s *Service) GenerateExcelType2(datas []entity.DailyData, f *excelize.File, sheet string) (*excelize.File, error) {
	f.NewSheet(sheet)
	f.SetColWidth(sheet, "A", "H", 20)

	headers := []string{"Contract", "Last", "Change", "Open", "High", "Low", "Previous", "Volume", "Time"}
	writeHeader(f, sheet, headers)

	for i, data := range datas {
		f.SetCellValue(sheet, "A"+strconv.Itoa(i+2), data.ContractSymbol)
		f.SetCellValue(sheet, "B"+strconv.Itoa(i+2), data.DailyLastPrice)
		f.SetCellValue(sheet, "C"+strconv.Itoa(i+2), data.DailyPriceChange)
		f.SetCellValue(sheet, "D"+strconv.Itoa(i+2), data.DailyOpenPrice)
		f.SetCellValue(sheet, "E"+strconv.Itoa(i+2), data.DailyHighPrice)
		f.SetCellValue(sheet, "F"+strconv.Itoa(i+2), data.DailyLowPrice)
		f.SetCellValue(sheet, "G"+strconv.Itoa(i+2), data.DailyPreviousPrice)
		f.SetCellValue(sheet, "H"+strconv.Itoa(i+2), data.DailyVolume)
		f.SetCellValue(sheet, "I"+strconv.Itoa(i+2), data.DailyDate1dAgo)
	}

	return f, nil
}
func (s *Service) GenerateExcelType3(datas []entity.TradingData, f *excelize.File, sheet string) (*excelize.File, error) {
	f.NewSheet(sheet)
	f.SetColWidth(sheet, "A", "H", 20)
	headers := []string{"Contract", "Last", "Change", "Open", "High", "Low", "Previous", "Volume", "Time"}
	writeHeader(f, sheet, headers)

	for i, data := range datas {
		f.SetCellValue(sheet, "A"+strconv.Itoa(i+2), data.ContractSymbol)
		f.SetCellValue(sheet, "B"+strconv.Itoa(i+2), data.LastPrice)
		f.SetCellValue(sheet, "C"+strconv.Itoa(i+2), data.PriceChange)
		f.SetCellValue(sheet, "D"+strconv.Itoa(i+2), data.OpenPrice)
		f.SetCellValue(sheet, "E"+strconv.Itoa(i+2), data.HighPrice)
		f.SetCellValue(sheet, "F"+strconv.Itoa(i+2), data.LowPrice)
		f.SetCellValue(sheet, "G"+strconv.Itoa(i+2), data.PreviousPrice)
		f.SetCellValue(sheet, "H"+strconv.Itoa(i+2), data.Volume)
		f.SetCellValue(sheet, "I"+strconv.Itoa(i+2), data.TradeTime)
	}

	return f, nil
}

func writeHeader(f *excelize.File, sheet string, headers []string) {
	cells := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N"}
	headerStyle, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
		Border: []excelize.Border{
			{
				Type:  "left",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "top",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "bottom",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "right",
				Color: "#000000",
				Style: 1,
			},
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#C0C0C0"},
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
	})

	for i, header := range headers {
		cell := cells[i] + "1"
		f.SetCellValue(sheet, cell, header)
		f.SetCellStyle(sheet, cell, cell, headerStyle)
	}
}
