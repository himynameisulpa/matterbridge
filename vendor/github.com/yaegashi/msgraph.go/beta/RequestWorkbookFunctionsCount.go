// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsCountRequestBuilder struct{ BaseRequestBuilder }

// Count action undocumented
func (b *WorkbookFunctionsRequestBuilder) Count(reqObj *WorkbookFunctionsCountRequestParameter) *WorkbookFunctionsCountRequestBuilder {
	bb := &WorkbookFunctionsCountRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/count"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsCountRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsCountRequestBuilder) Request() *WorkbookFunctionsCountRequest {
	return &WorkbookFunctionsCountRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsCountRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookFunctionsCountARequestBuilder struct{ BaseRequestBuilder }

// CountA action undocumented
func (b *WorkbookFunctionsRequestBuilder) CountA(reqObj *WorkbookFunctionsCountARequestParameter) *WorkbookFunctionsCountARequestBuilder {
	bb := &WorkbookFunctionsCountARequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/countA"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsCountARequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsCountARequestBuilder) Request() *WorkbookFunctionsCountARequest {
	return &WorkbookFunctionsCountARequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsCountARequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookFunctionsCountBlankRequestBuilder struct{ BaseRequestBuilder }

// CountBlank action undocumented
func (b *WorkbookFunctionsRequestBuilder) CountBlank(reqObj *WorkbookFunctionsCountBlankRequestParameter) *WorkbookFunctionsCountBlankRequestBuilder {
	bb := &WorkbookFunctionsCountBlankRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/countBlank"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsCountBlankRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsCountBlankRequestBuilder) Request() *WorkbookFunctionsCountBlankRequest {
	return &WorkbookFunctionsCountBlankRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsCountBlankRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookFunctionsCountIfRequestBuilder struct{ BaseRequestBuilder }

// CountIf action undocumented
func (b *WorkbookFunctionsRequestBuilder) CountIf(reqObj *WorkbookFunctionsCountIfRequestParameter) *WorkbookFunctionsCountIfRequestBuilder {
	bb := &WorkbookFunctionsCountIfRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/countIf"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsCountIfRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsCountIfRequestBuilder) Request() *WorkbookFunctionsCountIfRequest {
	return &WorkbookFunctionsCountIfRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsCountIfRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookFunctionsCountIfsRequestBuilder struct{ BaseRequestBuilder }

// CountIfs action undocumented
func (b *WorkbookFunctionsRequestBuilder) CountIfs(reqObj *WorkbookFunctionsCountIfsRequestParameter) *WorkbookFunctionsCountIfsRequestBuilder {
	bb := &WorkbookFunctionsCountIfsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/countIfs"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsCountIfsRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsCountIfsRequestBuilder) Request() *WorkbookFunctionsCountIfsRequest {
	return &WorkbookFunctionsCountIfsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsCountIfsRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
