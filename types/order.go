package types

// Order Структура данных отчёта валберис "Заказы".
// totalPrice*(100-discountPercent)/100= pricewithdisc (как и Методах сервиса. Продажи)
type Order struct {
	Date            WildberriesTime `json:"date"`
	LastChangeDate  WildberriesTime `json:"lastChangeDate"`
	SupplierArticle string          `json:"supplierArticle"`
	TechSize        string          `json:"techSize"`
	Barcode         string          `json:"barcode"`
	TotalPrice      float64         `json:"totalPrice"`
	DiscountPercent float64         `json:"discountPercent"`
	WarehouseName   string          `json:"warehouseName"`
	Oblast          string          `json:"oblast"`
	IncomeID        uint64          `json:"incomeID"`
	Odid            int64           `json:"odid"`
	NmId            uint64          `json:"nmId"`
	Subject         string          `json:"subject"`
	Category        string          `json:"category"`
	Brand           string          `json:"brand"`
	IsCancel        bool            `json:"isCancel"`
	CancelDt        WildberriesTime `json:"cancel_dt"`
	GNumber         string          `json:"gNumber"`
	Sticker         string          `json:"sticker"`
}
