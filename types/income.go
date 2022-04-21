package types

// Income Структура данных отчёта валберис "Поставки"
type Income struct {
	IncomeId        uint64          `json:"incomeId"`        // Уникальный идентификатор поставки
	Number          string          `json:"number"`          // Номер УПД
	Date            WildberriesTime `json:"date"`            // Дата поставки
	LastChangeDate  WildberriesTime `json:"lastChangeDate"`  // Дата и время последнего обновления информации отчёта в сервисе
	SupplierArticle string          `json:"supplierArticle"` // Артикул товара поставщика
	TechSize        string          `json:"techSize"`        // Технический размер
	Barcode         string          `json:"barcode"`         // Штрихкод
	Quantity        int64           `json:"quantity"`        // Количество
	TotalPrice      float64         `json:"totalPrice"`      // Цена товара из УПД
	DateClose       WildberriesTime `json:"dateClose"`       // Дата и время принятия (закрытия) в валберис
	WarehouseName   string          `json:"warehouseName"`   // Название склада
	NmID            uint64          `json:"nmId"`            // Код валберис, он же номенклатура валберис, он же код 1С
	Status          string          `json:"status"`          // Текущий статус поставки
}
