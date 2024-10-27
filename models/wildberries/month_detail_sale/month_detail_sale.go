package monthsale

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/hunternsk/wildberries/modules/communication"
	wildberriesTypes "github.com/hunternsk/wildberries/types"

	"gopkg.in/webnice/transport.v2/request"
)

// New creates a new object and return interface
func New(com communication.Interface, serverURI string, apiKey string, fromAt time.Time) Interface {
	var mds = &impl{
		ctx:       context.Background(),
		apiKey:    apiKey,
		serverURI: serverURI,
		com:       com,
		fromAt:    fromAt,
	}
	return mds
}

// WithContext Using context to interrupt requests to service
func (mds *impl) WithContext(ctx context.Context) Interface {
	if ctx == nil {
		return mds
	}
	mds.ctx = ctx
	return mds
}

// From Set of date and time of the beginning of the period for data request
func (mds *impl) From(fromAt time.Time) Interface {
	if fromAt.IsZero() {
		return mds
	}
	mds.fromAt = fromAt

	return mds
}

// Выбор значения fromAt
func (mds *impl) getFrom(fromAt ...time.Time) (ret time.Time) {
	var n int

	// Переопределения даты и времени начала периода для запроса, если fromAt передан
	ret = mds.fromAt
	for n = range fromAt {
		if fromAt[n].IsZero() {
			continue
		}
		ret = fromAt[n]
		break
	}

	return
}

// UntilDone Configures repeated requests with a progressive timeout until a
// response is successfully received from the server, but not more than retryMax requests
func (mds *impl) UntilDone(retryTimeout time.Duration, retryMax uint) Interface {
	mds.retryTimeout, mds.retryMax = retryTimeout, retryMax
	return mds
}

// Выполнение запроса к серверу, получение и разбор результата
func (mds *impl) request(
	rowID uint64,
	limit uint64,
	fromAt ...time.Time,
) (statusCode int, ret []*wildberriesTypes.MonthDetailSale, err error) {
	const (
		urn            = `%s/reportDetailByPeriod`
		keyDate        = `dateFrom`
		keyDateTo      = `dateTo`
		keyApi         = `key`
		keyLimit       = `limit`
		keyRowID       = `rrdid`
		rawQueryFmt    = `%s=%s`
		rawQueryAddFmt = `&%s=%s`
	)
	var (
		req  request.Interface
		from time.Time
		to   time.Time
		uri  *url.URL
	)

	// Подготовка данных
	from = mds.getFrom(fromAt...)
	if len(fromAt) > 1 {
		to = mds.getFrom(fromAt[1])
	}
	if uri, err = url.Parse(fmt.Sprintf(urn, mds.serverURI)); err != nil {
		err = fmt.Errorf("can't create request URI, error: %s", err)
		return
	}
	uri.RawQuery = fmt.Sprintf(
		rawQueryFmt,
		keyDate, from.In(wildberriesTypes.WildberriesTimezoneLocal).Format(`2006-01-02`),
	)
	if rowID > 0 {
		uri.RawQuery += fmt.Sprintf(rawQueryAddFmt, keyRowID, strconv.FormatUint(rowID, 10))
	}
	if limit > 0 {
		uri.RawQuery += fmt.Sprintf(rawQueryAddFmt, keyLimit, strconv.FormatUint(limit, 10))
	}
	//uri.RawQuery += fmt.Sprintf(rawQueryAddFmt, keyDateTo, to.In(wildberriesTypes.WildberriesTimezoneLocal).Format(wildberriesNonRFC3339TimeFormat))
	uri.RawQuery += fmt.Sprintf(rawQueryAddFmt, keyDateTo, to.In(wildberriesTypes.WildberriesTimezoneLocal).Format(`2006-01-02`))
	// Создание запроса
	req = mds.com.RequestJSON(mds.com.NewRequest(uri.String(), mds.com.Transport().Method().Get(), mds.apiKey))
	defer mds.com.Transport().RequestPut(req)
	// Выполнение запроса
	if statusCode, err = mds.com.RequestResponseJSON(mds.ctx, req, &ret); err != nil {
		err = fmt.Errorf("service response error: %s", err)
		return
	}

	return
}

// Report Load report data from the service.
// If not set the fromAt parameter, then the data will be loaded for the current day
// or starting from the date and time set by the From function
func (mds *impl) Report(
	rowID uint64,
	limit uint64,
	fromAt ...time.Time,
) (ret []*wildberriesTypes.MonthDetailSale, err error) {
	var (
		statusCode int
		n          uint
	)

	for {
		n++
		statusCode, ret, err = mds.request(rowID, limit, fromAt...)
		// Успешный ответ
		if err == nil && (statusCode > 199 && statusCode < 300) {
			break
		}
		// Если выключены повторы или попытки кончились
		if mds.retryTimeout == 0 || mds.retryMax <= n {
			break
		}
		// Если было выполнено прерывание через контекст
		if err = mds.ctx.Err(); err != nil {
			break
		}
		// Ожидание прерывания или таймаута между повторами
		select {
		case <-time.After(mds.retryTimeout * time.Duration(n)):
		case <-mds.ctx.Done():
			err = mds.ctx.Err()
			break
		}
	}

	return
}
