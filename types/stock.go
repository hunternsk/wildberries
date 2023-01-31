package types

// Stock Структура данных отчёта валберис "Склад"
type Stock struct {
	LastChangeDate  WildberriesTime `json:"lastChangeDate"`  // Дата и время последнего обновления информации отчёта в сервисе
	SupplierArticle string          `json:"supplierArticle"` // FK // Артикул товара поставщика
	TechSize        string          `json:"techSize"`        // Технический размер
	Barcode         string          `json:"barcode"`         // Штрихкод
	Quantity        int64           `json:"quantity"`        // Количество доступное для продажи - доступно на сайте, можно добавить в корзину
	IsSupply        bool            `json:"isSupply"`        // Договор поставки
	IsRealization   bool            `json:"isRealization"`   // Договор реализации
	QuantityFull    int64           `json:"quantityFull"`    // Количество полное - то, что не продано (числится на складе)
	WarehouseName   string          `json:"warehouseName"`   // FK // Название склада
	NmID            uint64          `json:"nmId"`            // Код валберис, он же номенклатура валберис, он же код 1С
	Subject         string          `json:"subject"`         // Предмет или название товара
	Category        string          `json:"category"`        // Категория
	DaysOnSite      uint64          `json:"daysOnSite"`      // Количество дней на сайте
	Brand           string          `json:"brand"`           // Бренд
	SCCode          string          `json:"SCCode"`          // FK // Код контракта
	Price           float64         `json:"Price"`           // Цена
	Discount        float64         `json:"Discount"`        // Скидка
}
