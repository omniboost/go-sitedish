package sitedish

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sitedish/utils"
)

// Manual Journal Voucher service used to query and create and post
// ManualJournalVoucher. (Norwegian: Manuelt bilag) Querying this service will
// only return the vouchers created by the integration itself, not all vouchers
// of type One.Domain.Entities.Accounting.Vouchers.VoucherType.ManualJournal.

func (c *Client) NewOrdersArchivedGet() OrdersArchivedGet {
	r := OrdersArchivedGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type OrdersArchivedGet struct {
	client      *Client
	queryParams *OrdersArchivedGetQueryParams
	pathParams  *OrdersArchivedGetPathParams
	method      string
	headers     http.Header
	requestBody OrdersArchivedGetBody
}

func (r OrdersArchivedGet) NewQueryParams() *OrdersArchivedGetQueryParams {
	return &OrdersArchivedGetQueryParams{}
}

type OrdersArchivedGetQueryParams struct {
	Start DateTime `schema:"start,omitempty"`
	End   DateTime `schema:"end,omitempty"`
}

func (p OrdersArchivedGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *OrdersArchivedGet) QueryParams() *OrdersArchivedGetQueryParams {
	return r.queryParams
}

func (r OrdersArchivedGet) NewPathParams() *OrdersArchivedGetPathParams {
	return &OrdersArchivedGetPathParams{}
}

type OrdersArchivedGetPathParams struct {
	Platform string
}

func (p *OrdersArchivedGetPathParams) Params() map[string]string {
	return map[string]string{
		"platform": p.Platform,
	}
}

func (r *OrdersArchivedGet) PathParams() *OrdersArchivedGetPathParams {
	return r.pathParams
}

func (r *OrdersArchivedGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *OrdersArchivedGet) SetMethod(method string) {
	r.method = method
}

func (r *OrdersArchivedGet) Method() string {
	return r.method
}

func (r OrdersArchivedGet) NewRequestBody() OrdersArchivedGetBody {
	return OrdersArchivedGetBody{}
}

type OrdersArchivedGetBody struct {
}

func (r *OrdersArchivedGet) RequestBody() *OrdersArchivedGetBody {
	return nil
}

func (r *OrdersArchivedGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *OrdersArchivedGet) SetRequestBody(body OrdersArchivedGetBody) {
	r.requestBody = body
}

func (r *OrdersArchivedGet) NewResponseBody() *OrdersArchivedGetResponseBody {
	return &OrdersArchivedGetResponseBody{}
}

type OrdersArchivedGetResponseBody Orders

func (r *OrdersArchivedGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/orders/archived/{{.platform}}", r.PathParams())
	return &u
}

func (r *OrdersArchivedGet) Do() (OrdersArchivedGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
