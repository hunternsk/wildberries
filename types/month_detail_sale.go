package types

// MonthDetailSale Структура данных отчёта валберис "Ежемесячный отчет о продажах по реализации"
type MonthDetailSale struct {
	ReportID                 uint64          `json:"realizationreport_id"`        // Номер отчёта
	SupplierContractCode     string          `json:"suppliercontract_code"`       // Договор
	RowReportID              uint64          `json:"rrd_id"`                      // FK // Уникальный идентификатор номера строки отчёта
	IncomeID                 uint64          `json:"gi_id"`                       // Номер поставки
	Name                     string          `json:"subject_name"`                // Предмет или название товара
	WbID                     uint64          `json:"nm_id"`                       // Код валберис, он же номенклатура валберис, он же код 1С, он же артикул
	BrandName                string          `json:"brand_name"`                  // Бренд
	VendorCode               string          `json:"sa_name"`                     // Артикул товара поставщика
	TechSize                 string          `json:"ts_name"`                     // Технический размер
	Barcode                  string          `json:"barcode"`                     // Штрихкод
	DocTypeName              string          `json:"doc_type_name"`               // Тип документа
	Quantity                 int64           `json:"quantity"`                    // Количество
	Nds                      float64         `json:"nds"`                         // Ставка НДС
	CostAmount               float64         `json:"cost_amount"`                 // Себестоимость, сумма
	RetailPrice              float64         `json:"retail_price"`                // Цена розничная
	RetailAmount             float64         `json:"retail_amount"`               // Сумма продаж, возвратов
	RetailCommission         float64         `json:"retail_commission"`           // Сумма комиссии продаж
	SalePercent              float64         `json:"sale_percent"`                // Согласованная скидка
	CommissionPercent        float64         `json:"commission_percent"`          // Процент комиссии
	CustomerReward           float64         `json:"customer_reward"`             // Вознаграждение покупателю
	SupplierReward           float64         `json:"supplier_reward"`             // Вознаграждение поставщику
	WarehouseName            string          `json:"office_name"`                 // Название склада
	SupplierOperName         string          `json:"supplier_oper_name"`          // Обоснование для оплаты
	OrderAt                  WildberriesTime `json:"order_dt"`                    // Дата заказа
	SaleAt                   WildberriesTime `json:"sale_dt"`                     // Дата продажи
	ShkID                    int64           `json:"shk_id"`                      // ШК
	RetailPriceWithDiscRub   float64         `json:"retail_price_withdisc_rub"`   // Цена розничная с учётом согласованной скидки
	ForPay                   float64         `json:"for_pay"`                     // К перечислению поставщику
	ForPayNds                float64         `json:"for_pay_nds"`                 // К перечислению поставщику, НДС
	DeliveryAmount           uint64          `json:"delivery_amount"`             // Количество доставок
	ReturnAmount             uint64          `json:"return_amount"`               // Количество возвратов
	DeliveryRub              float64         `json:"delivery_rub"`                // Стоимость логистики
	GiBoxTypeName            string          `json:"gi_box_type_name"`            // Тип коробов
	ProductDiscountForReport float64         `json:"product_discount_for_report"` // Согласованный продуктовый дисконт
	SupplierPromo            float64         `json:"supplier_promo"`              // Промокод
	SupplierSpp              float64         `json:"supplier_spp"`                // Скидка постоянного покупателя
	Rid                      int64           `json:"rid"`                         // Rid
}
