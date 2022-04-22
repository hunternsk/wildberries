package types

// MonthDetailSale Структура данных отчёта валберис "Ежемесячный отчет о продажах по реализации"
type MonthDetailSale struct {
	Realizationreport_id        uint64          `json:"realizationreport_id"`  // Номер отчёта
	Suppliercontract_code       string          `json:"suppliercontract_code"` // Договор
	Rrd_id                      uint64          `json:"rrd_id"`                // FK // Уникальный идентификатор номера строки отчёта
	Gi_id                       uint64          `json:"gi_id"`                 // Номер поставки
	Subject_name                string          `json:"subject_name"`          // Предмет или название товара
	Nm_id                       uint64          `json:"nm_id"`                 // Код валберис, он же номенклатура валберис, он же код 1С, он же артикул
	Brand_name                  string          `json:"brand_name"`            // Бренд
	Sa_name                     string          `json:"sa_name"`               // Артикул товара поставщика
	Ts_name                     string          `json:"ts_name"`               // Технический размер
	Barcode                     string          `json:"barcode"`               // Штрихкод
	Doc_type_name               string          `json:"doc_type_name"`         // Тип документа
	Quantity                    int             `json:"quantity"`
	Retail_price                float64         `json:"retail_price"`                // Цена розничная
	Retail_amount               float64         `json:"retail_amount"`               // Сумма продаж, возвратов
	Sale_percent                float64         `json:"sale_percent"`                // Согласованная скидка
	Commission_percent          float64         `json:"commission_percent"`          // Процент комиссии
	Office_name                 string          `json:"office_name"`                 // Название склада
	Supplier_oper_name          string          `json:"supplier_oper_name"`          // Обоснование для оплаты
	Order_dt                    WildberriesTime `json:"order_dt"`                    // Дата заказа
	Sale_dt                     WildberriesTime `json:"sale_dt"`                     // Дата продажи
	Rr_dt                       WildberriesTime `json:"Rr_dt"`                       // Дата заказа
	Shk_id                      int64           `json:"shk_id"`                      // ШК
	Retail_price_withdisc_rub   float64         `json:"retail_price_withdisc_rub"`   // Цена розничная с учётом согласованной скидки
	Delivery_amount             uint64          `json:"delivery_amount"`             // Количество доставок
	Return_amount               uint64          `json:"return_amount"`               // Количество возвратов
	Delivery_rub                float64         `json:"delivery_rub"`                // Стоимость логистики
	Gi_box_type_name            string          `json:"gi_box_type_name"`            // Тип коробов
	Product_discount_for_report float64         `json:"product_discount_for_report"` // Согласованный продуктовый дисконт
	Supplier_promo              float64         `json:"supplier_promo"`              // Промокод
	Rid                         uint64          `json:"rid"`                         // Rid
	Ppvz_spp_prc                float64         `json:"ppvz_spp_prc"`
	Ppvz_kvw_prc_base           float64         `json:"ppvz_kvw_prc_base"`
	Ppvz_kvw_prc                float64         `json:"ppvz_kvw_prc"`
	Ppvz_sales_commission       float64         `json:"ppvz_sales_commission"`
	Ppvz_for_pay                float64         `json:"ppvz_for_pay"`
	Ppvz_reward                 float64         `json:"ppvz_reward"`
	Ppvz_vw                     float64         `json:"ppvz_vw"`
	Ppvz_vw_nds                 float64         `json:"ppvz_vw_nds"`
	Ppvz_office_id              uint64          `json:"ppvz_office_id"`
	Ppvz_office_name            string          `json:"ppvz_office_name"`
	Ppvz_supplier_id            uint64          `json:"ppvz_supplier_id"`
	Declaration_number          string          `json:"declaration_number"`
	Sticker_id                  string          `json:"sticker_id"`

	/*
		Supplier_spp              float64         `json:"supplier_spp"`                // Скидка постоянного покупателя
		Quantity                 int64           `json:"quantity"`                    // Количество
		Nds                      float64         `json:"nds"`                         // Ставка НДС
		CostAmount               float64         `json:"cost_amount"`                 // Себестоимость, сумма
		RetailCommission         float64         `json:"retail_commission"`           // Сумма комиссии продаж
		CustomerReward           float64         `json:"customer_reward"`             // Вознаграждение покупателю
		SupplierReward           float64         `json:"supplier_reward"`             // Вознаграждение поставщику
		ForPay                   float64         `json:"for_pay"`                     // К перечислению поставщику
		ForPayNds                float64         `json:"for_pay_nds"`                 // К перечислению поставщику, НДС
	*/
}
