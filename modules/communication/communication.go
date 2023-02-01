package communication

import (
	"bytes"
	"context"
	"fmt"
	"runtime"
	"strings"

	"gopkg.in/webnice/transport.v2"
	"gopkg.in/webnice/transport.v2/content"
	"gopkg.in/webnice/transport.v2/methods"
	"gopkg.in/webnice/transport.v2/request"
	"gopkg.in/webnice/transport.v2/response"
	"gopkg.in/webnice/web.v1/header"
	"gopkg.in/webnice/web.v1/mime"
	"gopkg.in/webnice/web.v1/status"
)

// New creates a new object and return interface
func New() Interface {
	var com = &impl{
		singleton: newTransport(),
	}
	runtime.SetFinalizer(com, destructor)
	return com
}

// Деструктор
func destructor(item *impl) {
	// Остановка пула воркеров, ожидание завершения всех запросов, закрытие каналов
	item.singleton.Done()
	item.singleton = nil
}

func newTransport() transport.Interface {
	return transport.New().
		MaximumIdleConnections(defaultMaximumIdleConnections).               // Максимальное общее число бездействующих keepalive соединений
		MaximumIdleConnectionsPerHost(defaultMaximumIdleConnectionsPerHost). // Максимальное число бездействующих keepalive соединений для каждого хоста
		DialContextTimeout(defaultDialContextTimeout).                       // Таймаут установки соединения с хостом
		IdleConnectionTimeout(defaultIdleConnectionTimeout).                 // Таймаут keepalive соединения до обрыва связи
		TotalTimeout(defaultTotalTimeout).                                   // Общий таймаут на весь процесс связи, включает соединение, отправку данных, получение ответа
		RequestPoolSize(defaultRequestPoolSize).                             // Размер пула воркеров готовых для выполнения запросов к хостам
		TLSSkipVerify(defaultSkipTlsVerify)
}

// Errors Ошибки известного состояни, которые могут вернуть функции пакета
func (com *impl) Errors() *Error { return Errors() }

// Transport Готовый к использованию интерфейс коммуникации с сервером
func (com *impl) Transport() transport.Interface { return com.singleton }

// NewRequest Базовый метод создания объекта запроса
func (com *impl) NewRequest(uri string, mtd methods.Value, apiKey string) (ret request.Interface) {
	ret = com.Transport().RequestGet().
		AcceptLanguage(AcceptLanguage).
		UserAgent(UserAgent).
		Method(mtd).
		URL(uri).
		CustomHeader("Authorization", apiKey)

	return
}

// RequestJSON Подготовка запроса для получения JSON ответа
func (com *impl) RequestJSON(req request.Interface) (ret request.Interface) {
	ret = req.
		Accept(AcceptJSON).
		AcceptEncoding(AcceptEncoding)
	ret.Header().Add(header.CacheControl, CacheControl)
	req.Header().Add(header.ContentType, mime.ApplicationJSONCharsetUTF8)
	return
}

// RequestResponse Выполнение запроса, ожидание и получение результата
func (com *impl) RequestResponse(ctx context.Context, req request.Interface) (ret response.Interface, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	// DEBUG
	//req.DebugFunc(func(d []byte) { log.Debug(string(d)) })
	// DEBUG
	// Выполнение запроса
	com.Transport().Do(req)
	// Ожидание ответа
	if err = req.DoneWithContext(ctx).
		Error(); err != nil {
		err = fmt.Errorf("execute request error: %s", err)
		return
	}
	// Анализ результата
	ret = req.Response()
	switch ret.StatusCode() {
	case status.Unauthorized:
		err = Errors().Unauthorized()
		return
	case status.Forbidden:
		err = Errors().Forbidden()
		return
	case status.NotFound:
		err = Errors().NotFound()
		return
	}
	if ret.StatusCode() < 200 || ret.StatusCode() > 299 {
		err = fmt.Errorf("request %s %q error, HTTP code %d (%s)", ret.Response().Request.Method, ret.Response().Request.URL.String(), ret.StatusCode(), ret.Status())
		return
	}

	return
}

// RequestResponseStatusCode Выполнение запроса, ожидание и получение результата в виде HTTP статуса
func (com *impl) RequestResponseStatusCode(ctx context.Context, req request.Interface) (statusCode int, err error) {
	var rsp response.Interface

	if rsp, err = com.
		RequestResponse(ctx, req); err != nil {
		return
	}
	statusCode = rsp.StatusCode()
	// DEBUG
	//req.Response().Content().BackToBegin()
	//log.Debug(req.Response().Content().String())
	// DEBUG

	return
}

// RequestResponsePlainText Выполнение запроса, ожидание и получение результата в виде текста
func (com *impl) RequestResponsePlainText(ctx context.Context, req request.Interface) (ret *bytes.Buffer, statusCode int, err error) {
	var (
		rsp response.Interface
		cnt content.Interface
	)

	if rsp, err = com.
		RequestResponse(ctx, req); err != nil {
		return
	}
	ret, statusCode, cnt = &bytes.Buffer{}, rsp.StatusCode(), rsp.Content()
	if strings.EqualFold(rsp.Header().Get(header.ContentEncoding), EncodingGzip) {
		cnt = cnt.UnGzip()
	}
	if strings.EqualFold(rsp.Header().Get(header.ContentEncoding), EncodingDeflate) {
		cnt = cnt.UnFlate()
	}
	_, err = cnt.WriteTo(ret)
	// DEBUG
	//req.Response().Content().BackToBegin()
	//log.Debug(req.Response().Content().String())
	// DEBUG

	return
}

// RequestResponseJSON Выполнение запроса, ожидание и получение результата в виде JSON
func (com *impl) RequestResponseJSON(ctx context.Context, req request.Interface, data interface{}) (statusCode int, err error) {
	var (
		rsp response.Interface
		cnt content.Interface
	)

	if rsp, err = com.
		RequestResponse(ctx, req); err != nil {
		return
	}
	cnt, statusCode = rsp.Content(), rsp.StatusCode()
	if strings.EqualFold(rsp.Header().Get(header.ContentEncoding), EncodingGzip) {
		cnt = cnt.UnGzip()
	}
	if strings.EqualFold(rsp.Header().Get(header.ContentEncoding), EncodingDeflate) {
		cnt = cnt.UnFlate()
	}
	if err = cnt.UnmarshalJSON(data); err != nil {
		err = fmt.Errorf("json unmarshal error: %s", err)
		return
	}
	// DEBUG
	//req.Response().Content().BackToBegin()
	//log.Debug(req.Response().Content().String())
	// DEBUG

	return
}
