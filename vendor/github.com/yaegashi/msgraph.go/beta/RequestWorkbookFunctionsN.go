// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsNRequestBuilder struct{ BaseRequestBuilder }

// N action undocumented
func (b *WorkbookFunctionsRequestBuilder) N(reqObj *WorkbookFunctionsNRequestParameter) *WorkbookFunctionsNRequestBuilder {
	bb := &WorkbookFunctionsNRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/n"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsNRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsNRequestBuilder) Request() *WorkbookFunctionsNRequest {
	return &WorkbookFunctionsNRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsNRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
