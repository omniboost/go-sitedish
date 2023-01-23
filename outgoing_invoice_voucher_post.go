package poweroffice

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-poweroffice/omitempty"
	"github.com/omniboost/go-poweroffice/utils"
)

func (c *Client) NewOutgoingInvoiceVoucherPost() OutgoingInvoiceVoucherPost {
	r := OutgoingInvoiceVoucherPost{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type OutgoingInvoiceVoucherPost struct {
	client      *Client
	queryParams *OutgoingInvoiceVoucherPostQueryParams
	pathParams  *OutgoingInvoiceVoucherPostPathParams
	method      string
	headers     http.Header
	requestBody OutgoingInvoiceVoucherPostBody
}

func (r OutgoingInvoiceVoucherPost) NewQueryParams() *OutgoingInvoiceVoucherPostQueryParams {
	return &OutgoingInvoiceVoucherPostQueryParams{}
}

type OutgoingInvoiceVoucherPostQueryParams struct{}

func (p OutgoingInvoiceVoucherPostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *OutgoingInvoiceVoucherPost) QueryParams() *OutgoingInvoiceVoucherPostQueryParams {
	return r.queryParams
}

func (r OutgoingInvoiceVoucherPost) NewPathParams() *OutgoingInvoiceVoucherPostPathParams {
	return &OutgoingInvoiceVoucherPostPathParams{}
}

type OutgoingInvoiceVoucherPostPathParams struct {
}

func (p *OutgoingInvoiceVoucherPostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *OutgoingInvoiceVoucherPost) PathParams() *OutgoingInvoiceVoucherPostPathParams {
	return r.pathParams
}

func (r *OutgoingInvoiceVoucherPost) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *OutgoingInvoiceVoucherPost) SetMethod(method string) {
	r.method = method
}

func (r *OutgoingInvoiceVoucherPost) Method() string {
	return r.method
}

func (r OutgoingInvoiceVoucherPost) NewRequestBody() OutgoingInvoiceVoucherPostBody {
	return OutgoingInvoiceVoucherPostBody{}
}

type OutgoingInvoiceVoucherPostBody OutgoingInvoiceVoucher

func (v OutgoingInvoiceVoucherPostBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(v)
}

func (r *OutgoingInvoiceVoucherPost) RequestBody() *OutgoingInvoiceVoucherPostBody {
	return &r.requestBody
}

func (r *OutgoingInvoiceVoucherPost) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *OutgoingInvoiceVoucherPost) SetRequestBody(body OutgoingInvoiceVoucherPostBody) {
	r.requestBody = body
}

func (r *OutgoingInvoiceVoucherPost) NewResponseBody() *OutgoingInvoiceVoucherPostResponseBody {
	return &OutgoingInvoiceVoucherPostResponseBody{}
}

type OutgoingInvoiceVoucherPostResponseBody struct{}

func (r *OutgoingInvoiceVoucherPost) URL() *url.URL {
	u := r.client.GetEndpointURL("/OutgoingInvoiceVoucher/", r.PathParams())
	return &u
}

func (r *OutgoingInvoiceVoucherPost) Do() (OutgoingInvoiceVoucherPostResponseBody, error) {
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
