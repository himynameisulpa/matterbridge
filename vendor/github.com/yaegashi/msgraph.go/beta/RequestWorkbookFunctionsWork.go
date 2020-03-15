// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsWorkDayRequestBuilder struct{ BaseRequestBuilder }

// WorkDay action undocumented
func (b *WorkbookFunctionsRequestBuilder) WorkDay(reqObj *WorkbookFunctionsWorkDayRequestParameter) *WorkbookFunctionsWorkDayRequestBuilder {
	bb := &WorkbookFunctionsWorkDayRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/workDay"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsWorkDayRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsWorkDayRequestBuilder) Request() *WorkbookFunctionsWorkDayRequest {
	return &WorkbookFunctionsWorkDayRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsWorkDayRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookFunctionsWorkDay_IntlRequestBuilder struct{ BaseRequestBuilder }

// WorkDay_Intl action undocumented
func (b *WorkbookFunctionsRequestBuilder) WorkDay_Intl(reqObj *WorkbookFunctionsWorkDay_IntlRequestParameter) *WorkbookFunctionsWorkDay_IntlRequestBuilder {
	bb := &WorkbookFunctionsWorkDay_IntlRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/workDay_Intl"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsWorkDay_IntlRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsWorkDay_IntlRequestBuilder) Request() *WorkbookFunctionsWorkDay_IntlRequest {
	return &WorkbookFunctionsWorkDay_IntlRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsWorkDay_IntlRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
