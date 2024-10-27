package sales

import (
	"context"
	"time"

	"github.com/hunternsk/wildberries/modules/communication"
	wildberriesTypes "github.com/hunternsk/wildberries/types"
)

const wildberriesNonRFC3339TimeFormat = `2006-01-02T15:04:05`

// Interface is an interface of package
type Interface interface {
	// WithContext Using context to interrupt requests to service
	WithContext(ctx context.Context) Interface

	// From Set of date and time of the beginning of the period for data request
	From(fromAt time.Time) Interface

	// UntilDone Configures repeated requests with a progressive timeout until a
	// response is successfully received from the server, but not more than retryMax requests
	UntilDone(retryTimeout time.Duration, retryMax uint) Interface

	// Report Load report data from the service.
	// The onThisDay parameter indicates that data for the selected day is requested.
	// If not set the fromAt parameter, then the data will be loaded for the current day
	// or starting from the date and time set by the From function.
	// PriceWithDisc calculation formula:
	//   Pricewithdisc = totalprice*((100 – discountPercent)/100 ) *((100 – promoCodeDiscount)/100 ) *((100 – spp)/100 )
	Report(onThisDay bool, fromAt ...time.Time) (ret []*wildberriesTypes.Sale, err error)
}

// impl is an implementation of package
type impl struct {
	fromAt       time.Time               // Дата и время начала периода для запроса данных
	com          communication.Interface // Интерфейс коммуникации с сервисом
	ctx          context.Context         // Интерфейс контекста
	apiKey       string                  // Ключ API
	serverURI    string                  // URI адрес сервиса
	retryTimeout time.Duration           // Начальный таймаут для повторных запросов
	retryMax     uint                    // Максимальное количество повторных запросов
}
