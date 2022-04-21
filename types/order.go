package types

// Order Структура данных отчёта валберис "Заказы".
// totalPrice*(100-discountPercent)/100= pricewithdisc (как и Методах сервиса. Продажи)
type Order struct {
	Number          uint64          `json:"Number"`          // Number
	Date            WildberriesTime `json:"date"`            // Дата заказа
	LastChangeDate  WildberriesTime `json:"lastChangeDate"`  // Дата и время последнего обновления информации отчёта в сервисе
	SupplierArticle string          `json:"supplierArticle"` // Артикул товара поставщика
	TechSize        string          `json:"techSize"`        // Технический размер
	Barcode         string          `json:"barcode"`         // Штрихкод
	Quantity        int64           `json:"quantity"`        // Количество
	TotalPrice      float64         `json:"totalPrice"`      // Цена товара из УПД
	DiscountPercent float64         `json:"discountPercent"` // Согласованная итоговая скидка в процентах
	WarehouseName   string          `json:"warehouseName"`   // Название склада отгрузки товара
	Oblast          string          `json:"oblast"`          // Область
	IncomeID        uint64          `json:"incomeID"`        // Уникальный идентификатор поставки
	Odid            int64           `json:"odid"`            // FP // Уникальный идентификатор позиции заказа
	NmId            uint64          `json:"nmId"`            // Код валберис, он же номенклатура валберис, он же код 1С
	Subject         string          `json:"subject"`         // Предмет или название товара
	Category        string          `json:"category"`        // Категория
	Brand           string          `json:"brand"`           // Бренд
	IsCancel        bool            `json:"isCancel"`        // Отменённый заказ
	CancelDt        WildberriesTime `json:"cancel_dt"`       // Дата отмены заказа
	GNumber         string          `json:"gNumber"`         // GNumber
}
